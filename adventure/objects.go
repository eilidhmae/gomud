package gomud

type Object struct {
	Id		int
	Desc		string
	Area		*Area
	Next		*Object
	Previous	*Object
}

type Objlist struct {
	Count		int
	Data		*[]byte
	Head		*Object
	Tail		*Object
	Current		*Object
}

func NewObjlist(p *Object) Objlist {
	return Objlist{
		Count:		1,
		Head:		p,
		Tail:		p,
		Current:	nil,
	}
}

func (l *Objlist) Lookup(index int) *Object {
	if index > l.Count {
		return l.Tail
	}
	if index < 1 {
		return l.Head
	}
	cur := l.Head
	for cur.Id < index {
		cur = cur.Next
	}
	return cur
}

func (l *Objlist) Add(p *Object) int {
	old := l.Tail
	old.Next = p
	p.Previous = old
	p.Id = old.Id + 1
	l.Count = p.Id
	l.Tail = p
	return l.Count
}

func (l *Objlist) Drop(index int) *Object {
	t := l.Lookup(index)
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
