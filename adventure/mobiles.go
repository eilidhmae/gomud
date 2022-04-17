package gomud

type Mobile struct {
	Id		int
	Title		string
	Area		*Area
	Next		*Mobile
	Previous	*Mobile
}

type Moblist struct {
	Count		int
	Data		*[]byte
	Head		*Mobile
	Tail		*Mobile
	Current		*Mobile
}

func (l *Moblist) Add(p *Mobile) int {
	old := l.Tail
	old.Next = p
	p.Previous = old
	p.Id = old.Id + 1
	l.Count = p.Id
	l.Tail = p
	return l.Count
}

func NewMoblist(p *Mobile) Moblist {
	return Moblist{
		Count:		1,
		Head:		p,
		Tail:		p,
		Current:	nil,
	}
}
