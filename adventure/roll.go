package gomud

import (
	"fmt"
	"sort"
)

type Stats struct {
	Raw	[]int
	Str	int
	Dex	int
	Wis	int
	Int	int
	Con	int
	Cha	int
}

func (c *Character) RollStats(min,max int) {
	var s Stats

	for i := 0; i < 6; i++ {
		roll := getRandomStat(min, max)
		s.Raw = append(s.Raw, roll)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(s.Raw)))

	switch c.Class {
	case "fighter":
		s.Str = s.Raw[0]
		s.Con = s.Raw[1]
		s.Dex = s.Raw[2]
		s.Cha = s.Raw[3]
		s.Wis = s.Raw[4]
		s.Int = s.Raw[5]
	case "mage":
		s.Int = s.Raw[0]
		s.Con = s.Raw[1]
		s.Wis = s.Raw[2]
		s.Dex = s.Raw[3]
		s.Cha = s.Raw[4]
		s.Str = s.Raw[5]
	case "cleric":
		s.Wis = s.Raw[0]
		s.Con = s.Raw[1]
		s.Int = s.Raw[2]
		s.Cha = s.Raw[3]
		s.Dex = s.Raw[4]
		s.Str = s.Raw[5]
	case "rogue":
		s.Dex = s.Raw[0]
		s.Con = s.Raw[1]
		s.Str = s.Raw[2]
		s.Cha = s.Raw[3]
		s.Wis = s.Raw[4]
		s.Int = s.Raw[5]
	}

	c.Stats = s
}

func (s Stats) Text() string {
	return fmt.Sprintf("STR(%2d) DEX(%2d) CON(%2d) INT(%2d) WIS(%2d) CHA(%2d)", s.Str, s.Dex, s.Con, s.Int, s.Wis, s.Cha)
}

