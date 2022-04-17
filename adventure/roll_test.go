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

func TestRollStats(t *testing.T) {
	min := 5
	max := 18
	var testFighter Character
	var testMage Character
	var testCleric Character
	var testRogue Character

	testFighter.Class = "fighter"
	testMage.Class = "mage"
	testCleric.Class = "cleric"
	testRogue.Class = "rogue"

	testFighter.RollStats(min, max)
	switch {
	case testFighter.Stats.Str < testFighter.Stats.Con:
		t.Errorf("fighter: str is not highest stat")
	case testFighter.Stats.Str < testFighter.Stats.Wis:
		t.Errorf("fighter: str is not highest stat")
	case testFighter.Stats.Str < testFighter.Stats.Dex:
		t.Errorf("fighter: str is not highest stat")
	case testFighter.Stats.Str < testFighter.Stats.Int:
		t.Errorf("fighter: str is not highest stat")
	case testFighter.Stats.Str < testFighter.Stats.Cha:
		t.Errorf("fighter: str is not highest stat")
	}

	testMage.RollStats(min, max)
	switch {
	case testMage.Stats.Int < testMage.Stats.Str:
		t.Errorf("mage: int is not highest stat")
	case testMage.Stats.Int < testMage.Stats.Con:
		t.Errorf("mage: int is not highest stat")
	case testMage.Stats.Int < testMage.Stats.Wis:
		t.Errorf("mage: int is not highest stat")
	case testMage.Stats.Int < testMage.Stats.Int:
		t.Errorf("mage: int is not highest stat")
	case testMage.Stats.Int < testMage.Stats.Cha:
		t.Errorf("mage: int is not highest stat")
	}

	testRogue.RollStats(min, max)
	switch {
	case testRogue.Stats.Dex < testRogue.Stats.Str:
		t.Errorf("rogue: dex is not highest stat")
	case testRogue.Stats.Dex < testRogue.Stats.Con:
		t.Errorf("rogue: dex is not highest stat")
	case testRogue.Stats.Dex < testRogue.Stats.Int:
		t.Errorf("rogue: dex is not highest stat")
	case testRogue.Stats.Dex < testRogue.Stats.Wis:
		t.Errorf("rogue: dex is not highest stat")
	case testRogue.Stats.Dex < testRogue.Stats.Cha:
		t.Errorf("rogue: dex is not highest stat")
	}

	testCleric.RollStats(min, max)
	switch {
	case testCleric.Stats.Wis < testCleric.Stats.Str:
		t.Errorf("cleric: wis is not highest stat")
	case testCleric.Stats.Wis < testCleric.Stats.Dex:
		t.Errorf("cleric: wis is not highest stat")
	case testCleric.Stats.Wis < testCleric.Stats.Con:
		t.Errorf("cleric: wis is not highest stat")
	case testCleric.Stats.Wis < testCleric.Stats.Wis:
		t.Errorf("cleric: wis is not highest stat")
	case testCleric.Stats.Wis < testCleric.Stats.Int:
		t.Errorf("cleric: wis is not highest stat")
	}
}
