package gomud

type Room struct {
	Id		int
	Title		string
	Area		*Area
	Next		*Room
	Previous	*Room
}

type Roomlist struct {
	Count		int
	Head		*Room
	Tail		*Room
	Current		*Room
}

func (al *Roomlist) Add(p *Room) int {
	old := al.Tail
	old.Next = p
	p.Previous = old
	p.Id = old.Id + 1
	al.Count = p.Id
	al.Tail = p
	return al.Count
}

func NewRoomlist(p *Room) Roomlist {
	return Roomlist{
		Count:		1,
		Head:		p,
		Tail:		p,
		Current:	nil,
	}
}
