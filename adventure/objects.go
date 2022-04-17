package gomud

type Object struct {
	Id		int
	Title		string
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

func (rl *Objlist) Add(p *Object) int {
	old := rl.Tail
	old.Next = p
	p.Previous = old
	p.Id = old.Id + 1
	rl.Count = p.Id
	rl.Tail = p
	return rl.Count
}

func NewObjlist(p *Object) Objlist {
	return Objlist{
		Count:		1,
		Head:		p,
		Tail:		p,
		Current:	nil,
	}
}
