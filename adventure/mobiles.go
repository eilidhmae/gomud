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

func (rl *Moblist) Add(p *Mobile) int {
	old := rl.Tail
	old.Next = p
	p.Previous = old
	p.Id = old.Id + 1
	rl.Count = p.Id
	rl.Tail = p
	return rl.Count
}

func NewMoblist(p *Mobile) Moblist {
	return Moblist{
		Count:		1,
		Head:		p,
		Tail:		p,
		Current:	nil,
	}
}
