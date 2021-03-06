package gomud

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Realm struct {
	Areas		*Tree
	Rooms		*Tree
	Objects		*Tree
	Mobiles		*Tree
}

func BuildRealm(areasPath string) (Realm, error) {
	var realm Realm
	areas := NewTree(NewNodeByName("{ 1 35} Eilidh\tThe Coffeehouse~\n", nil))
	fh, err := os.Open(areasPath + "area.lst")
	if err != nil {
		return realm, err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		filename := scanner.Text()
		if filename != "$" {
			content, err := ioutil.ReadFile(areasPath + filename)
			if err != nil {
				return realm, err
			}
			r := bufio.NewReader(bytes.NewReader(content))
			t, err := r.ReadString('\n')
			if err != nil {
				return realm, err
			}
			matched, err := regexp.Match(`^#AREA`, []byte(t))
			if err != nil {
				return realm, err
			}
			if matched {
				title := strings.TrimLeft(t, "#AREA\t")
				areas.Add(NewNodeByName(title, content))
			}
		}
	}
	realm.Areas = areas
	if err := realm.ParseAreaData(); err != nil {
		return realm, err
	}
	return realm, nil
}

func (r *Realm) ParseAreaData() error {
	var rooms []string
	var objects []string
	var mobiles []string
	cur := r.Areas.Head
	if cur == nil {
		return fmt.Errorf("Realm.Areas.Head is nil")
	}
	for cur != nil {
		if cur.Data == nil {
			return fmt.Errorf("ParseAreaData: Area %d has no Data.", cur.Id)
		}
		s := bufio.NewScanner(bytes.NewReader(*cur.Data))
		// #MOBILES #OBJECTS #ROOMS marks beginning block, #0 marks end of block
		readingRooms := false
		readingObjects := false
		readingMobiles := false
		for s.Scan() {
			l := s.Text()
			matchedEnd, err := regexp.Match(`^#0$`, []byte(l))
			if err != nil {
				return err
			}
			matchedRooms, err := regexp.Match(`#ROOMS`, []byte(l))
			if err != nil {
				return err
			}
			matchedObjects, err := regexp.Match(`#OBJECTS`, []byte(l))
			if err != nil {
				return err
			}
			matchedMobiles, err := regexp.Match(`#MOBILES`, []byte(l))
			if err != nil {
				return err
			}
			switch {
			case matchedEnd:
				readingRooms = false
				readingObjects = false
				readingMobiles = false
			case readingRooms:
				rooms = append(rooms, l)
			case readingObjects:
				objects = append(objects, l)
			case readingMobiles:
				mobiles = append(mobiles, l)
			case matchedRooms:
				readingRooms = true
			case matchedObjects:
				readingObjects = true
			case matchedMobiles:
				readingMobiles = true
			default:
			}
		}
		cur = cur.Next
	}
	r.Rooms = &Tree{}
	r.Rooms.Data = packageBytes(rooms)
	r.ParseObjects(objects)
	r.Mobiles = &Tree{}
	r.Mobiles.Data = packageBytes(mobiles)
	return nil
}

func (r *Realm) ParseObjects(data []string) {
	mugName := "a mug of coffee"
	mugData := packageBytes([]string{`#1`,mugName,"A mug of coffee lies here."})
	objs := NewTree(NewNodeByName(mugName, *mugData))
	leafName := "a leaf"
	leafData := packageBytes([]string{`#3`,leafName,"A dewy leaf lies here."})
	objs.Add(NewNodeByName(leafName, *leafData))
	newObjectPattern := `^#[0-9]*$` // object index
	newObjectRegex := regexp.MustCompile(newObjectPattern)
	if len(data) > 0 {
		for index, l := range data {
			if newObjectRegex.MatchString(l) {
				// from location get line 0, 2, 3
				// create object Node{Data: []string{#<objnum>,<short desc>,<long desc>}, Name: <short desc>}
				id := data[index]
				short := strings.TrimSuffix(data[index+2], "~")
				long := strings.TrimSuffix(data[index+3], "~")
				d := packageBytes([]string{id,short,long})
				objs.Add(NewNodeByName(short, *d))
			}
		}
		objs.Data = packageBytes(data)
		r.Objects = objs
	}
}
