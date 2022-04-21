package gomud

import (
	"regexp"
	"strings"
)

type Node struct {
	Id		int		`json:"id"`
	Name		string		`json:"name"`
	Data		*[]byte		`json:"data"`
	Next		*Node		`json:"-"`
	Previous	*Node		`json:"-"`
}

type Tree struct {
	Count		int		`json:"int"`
	Data		*[]byte		`json:"-"`  // holds raw data from area files
	Head		*Node		`json:"-"`
	Tail		*Node		`json:"-"`
	Current		*Node		`json:"-"`
}

func NewTree(node *Node) *Tree {
	return &Tree{
		Count:		1,
		Head:		node,
		Tail:		node,
		Current:	nil,
	}
}

func NewNode(data []byte) *Node {
	return &Node{Id: 1, Data: &data}
}

func NewNodeByName(name string, data []byte) *Node {
	return &Node{Id: 1, Name: name, Data: &data}
}

func (l *Tree) LookupId(index int) *Node {
	if index > l.Count {
		return nil
	}
	if index < 1 {
		return nil
	}
	cur := l.Head
	for cur != nil {
		if cur.Id == index {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (l *Tree) LookupName(name string) *Node {
	cur := l.Head
	for cur != nil {
		if cur.Name == name {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func (l *Tree) Add(p *Node) int {
	if l.Head == nil {
		l.Head = p
		l.Tail = p
		l.Count = l.Count + 1
	} else {
		old := l.Tail
		old.Next = p		// old.Next is set to new Tail pointer
		p.Previous = old	// new Tail Previous is set to old Tail pointer
		p.Id = old.Id + 1	// update node id
		l.Count = p.Id		// update tree count
		l.Tail = p		// update tree tail
	}
	return l.Count
}

func (l *Tree) Drop(index int) *Node {
	t := l.LookupId(index)
	if t == nil {
		return nil
	}
	switch {
	case t == l.Head:
		l.Head = t.Next
		l.Head.Previous = nil
	case t == l.Tail:
		l.Tail = t.Previous
		l.Tail.Next = nil
	default:
		p := t.Previous
		n := t.Next
		p.Next = n
		n.Previous = p
	}
	l.Count = l.Count - 1
	return t
}

func (l *Tree) HasData(reg string) (int, bool) {
	q := regexp.MustCompile(reg)
	cur := l.Head
	for cur != nil {
		if q.Match(*cur.Data) {
			return cur.Id, true
		}
		cur = cur.Next
	}
	return 0, false
}

// returns object id, object short desc, object long desc
func (n *Node) GetObjectData() (string, string, string) {
	lines := strings.Split(string(*n.Data), "\n")
	if len(lines) == 3 {
		return lines[0], lines[1], lines[2]
	}
	return "", "", ""
}

func (l *Tree) FindObjectById(id string) *Node {
	cur := l.Head
	for cur != nil {
		objectId, _, _ := cur.GetObjectData()
		if objectId == id {
			return &Node{
				Id:	1,
				Name:	cur.Name,
				Data:	cur.Data,
			}
		}
		cur = cur.Next
	}
	return nil
}
