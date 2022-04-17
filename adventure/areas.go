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

type Area struct {
	Id		int
	Title		string
	Data		*[]byte
	Rooms		*Roomlist
	Objects		*Objlist
	Mobiles		*Moblist
	Next		*Area
	Previous	*Area
}

type Arealist struct {
	Count		int
	Head		*Area
	Tail		*Area
	Current		*Area
}

// assumes Arealist already has one area assigned
func (al *Arealist) Add(p *Area) int {
	old := al.Tail
	old.Next = p
	p.Previous = old
	p.Id = old.Id + 1
	al.Count = p.Id
	al.Tail = p
	return al.Count
}

func NewArealist(p *Area) Arealist {
	return Arealist{
		Count:		1,
		Head:		p,
		Tail:		p,
		Current:	nil,
	}
}

func BuildAreaList(areasPath string) (Arealist, error) {
	al := NewArealist(&Area{Id: 1,Title: "{ 1 35} Eilidh\tThe Coffeehouse~\n"})
	al.Tail = al.Head
	al.Current = al.Head
	fh, err := os.Open(areasPath + "area.lst")
	if err != nil {
		return al, err
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		filename := scanner.Text()
		if filename != "$" {
			content, err := ioutil.ReadFile(areasPath + filename)
			if err != nil {
				return al, err
			}
			r := bufio.NewReader(bytes.NewReader(content))
			t, err := r.ReadString('\n')
			if err != nil {
				return al, err
			}
			matched, err := regexp.Match(`^#AREA`, []byte(t))
			if err != nil {
				return al, err
			}
			if matched {
				title := strings.TrimLeft(t, "#AREA\t")
				a := Area{Title: title}
				a.Data = &content
				al.Add(&a)
			}
		}
	}
	return al, nil
}

func (a *Area) Build() error {
	if a.Data == nil {
		return fmt.Errorf("BuildRooms: Area %d has no Data.", a.Id)
	}
	s := bufio.NewScanner(bytes.NewReader(*a.Data))
	// #MOBILES #OBJECTS #ROOMS marks beginning block, #0 marks end of block
	var rooms []string
	var objects []string
	var mobiles []string

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
	a.Rooms = &Roomlist{Data: packageBytes(rooms)}
	a.Objects = &Objlist{Data: packageBytes(objects)}
	a.Mobiles = &Moblist{Data: packageBytes(mobiles)}
	return nil
}

func packageBytes(lines []string) *[]byte {
	var stream string
	for _, l := range lines {
		stream = stream + l + "\n"
	}
	b := []byte(stream)
	return &b
}
