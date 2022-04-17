package gomud

type Room struct {
	Id		int
	Title		string
	Area		*Area
	Objects		*Objlist
	Mobiles		*Moblist
	Next		*Room
	Previous	*Room
}

type Roomlist struct {
	Count		int
	Data		*[]byte
	Head		*Room
	Tail		*Room
	Current		*Room
}

func (rl *Roomlist) Add(p *Room) int {
	old := rl.Tail
	old.Next = p
	p.Previous = old
	p.Id = old.Id + 1
	rl.Count = p.Id
	rl.Tail = p
	return rl.Count
}

func NewRoomlist(p *Room) Roomlist {
	return Roomlist{
		Count:		1,
		Head:		p,
		Tail:		p,
		Current:	nil,
	}
}
