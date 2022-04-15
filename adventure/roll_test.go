package gomud

import "testing"

func TestText(t *testing.T) {
	expected := "STR(18) DEX(17) CON(16) INT(15) WIS(14) CHA(13)"
	s := Stats{
		Str:	18,
		Dex:	17,
		Con:	16,
		Int:	15,
		Wis:	14,
		Cha:	13,
	}

	got := s.Text()

	if expected != got {
		t.Errorf("Text output mismatch. expected: %s got: %s", expected, got)
	}
}

func TestGetRandomStat(t *testing.T) {
	low := 1
	high := 20
	for i := 1; i <= 100; i++ {
		result := getRandomStat(low, high)
		if result < low {
			t.Errorf("result too low")
		}
		if result > high {
			t.Errorf("result too high")
		}
	}

	flipped := getRandomStat(high, low)
	if flipped < low {
		t.Errorf("flipped too low")
	}
	if flipped > high {
		t.Errorf("flipped too high")
	}
}

func TestRollStats(t *testing.T) {
	min := 5
	max := 18

	f := RollStats(min, max, "fighter")
	switch {
	case f.Str < f.Con:
		t.Errorf("fighter: str is not highest stat")
	case f.Str < f.Wis:
		t.Errorf("fighter: str is not highest stat")
	case f.Str < f.Dex:
		t.Errorf("fighter: str is not highest stat")
	case f.Str < f.Int:
		t.Errorf("fighter: str is not highest stat")
	case f.Str < f.Cha:
		t.Errorf("fighter: str is not highest stat")
	}
	m := RollStats(min, max, "mage")
	switch {
	case m.Int < m.Str:
		t.Errorf("mage: int is not highest stat")
	case m.Int < m.Con:
		t.Errorf("mage: int is not highest stat")
	case m.Int < m.Wis:
		t.Errorf("mage: int is not highest stat")
	case m.Int < m.Int:
		t.Errorf("mage: int is not highest stat")
	case m.Int < m.Cha:
		t.Errorf("mage: int is not highest stat")
	}
	r := RollStats(min, max, "rogue")
	switch {
	case r.Dex < r.Str:
		t.Errorf("rogue: dex is not highest stat")
	case r.Dex < r.Con:
		t.Errorf("rogue: dex is not highest stat")
	case r.Dex < r.Int:
		t.Errorf("rogue: dex is not highest stat")
	case r.Dex < r.Wis:
		t.Errorf("rogue: dex is not highest stat")
	case r.Dex < r.Cha:
		t.Errorf("rogue: dex is not highest stat")
	}
	c := RollStats(min, max, "cleric")
	switch {
	case c.Wis < c.Str:
		t.Errorf("cleric: wis is not highest stat")
	case c.Wis < c.Dex:
		t.Errorf("cleric: wis is not highest stat")
	case c.Wis < c.Con:
		t.Errorf("cleric: wis is not highest stat")
	case c.Wis < c.Wis:
		t.Errorf("cleric: wis is not highest stat")
	case c.Wis < c.Int:
		t.Errorf("cleric: wis is not highest stat")
	}
}
