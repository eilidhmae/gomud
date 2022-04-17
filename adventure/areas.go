package gomud

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type Area struct {
	Id		int
	Title		string
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

func BuildAreaList(areasPath string) (Arealist, error) {
	al := Arealist{
		Count:		1,
		Head:		&Area{
					Id:		1,
					Title:		"{ 1 35} Eilidh\tThe Coffeehouse~\n",
					Next:		nil,
					Previous:	nil,
				},
	}
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
			handler, err := os.Open(areasPath + filename)
			if err != nil {
				return al, err
			}
			defer handler.Close()
			r := bufio.NewReader(handler)
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
				al.Add(&a)
			}
		}
	}
	return al, nil
}
