package gomud

import (
	"os"
	"reflect"
	"testing"
	"strings"
)

func TestNewCharacter(t *testing.T) {
	name := "duncan"
	class := "cleric"
	r := strings.NewReader(class)
	w, err := os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	c, err := NewCharacter(name, r, w)
	if err != nil {
		t.Error(err)
	}
	if c.Name != name {
		t.Errorf("NewCharacter name mismatch.")
	}
	if c.Level != 1 {
		t.Errorf("NewCharacter level mismatch.")
	}
	if c.Cursor != "What would you like to do?\n" {
		t.Errorf("NewCharacter cursor mismatch.")
	}
	if c.Class != "cleric" {
		t.Errorf("NewCharacter class mismatch.")
	}
	if c.Stats.Raw == nil {
		t.Errorf("NewCharacter Stats are nil.")
	}
	switch {
	case c.Stats.Str == 0:
		t.Errorf("NewCharacter Str is 0.")
	case c.Stats.Dex == 0:
		t.Errorf("NewCharacter Dex is 0.")
	case c.Stats.Con == 0:
		t.Errorf("NewCharacter Con is 0.")
	case c.Stats.Wis == 0:
		t.Errorf("NewCharacter Wis is 0.")
	case c.Stats.Int == 0:
		t.Errorf("NewCharacter Int is 0.")
	case c.Stats.Cha == 0:
		t.Errorf("NewCharacter Cha is 0.")
	}
}

func TestSetName(t *testing.T) {
	name := "alice"
	var c Character

	c.SetName(name)
	if c.Name != name {
		t.Errorf("SetName: name mismatch")
	}
}

func TestSetCursor(t *testing.T) {
	cursor := `-> `
	var c Character

	c.SetCursor(cursor)
	if c.Cursor != cursor {
		t.Errorf("SetCursor: cursor mismatch")
	}
}

func TestFancyCursor(t *testing.T) {
	expected := "drago the mage-> "
	c := Character{Name: "drago", Class: "mage"}
	c.FancyCursor()
	if c.Cursor != expected {
		t.Errorf("FancyCursor unexpected cursor: %s", c.Cursor)
	}
}

func TestSave(t *testing.T) {
	c := initializeTestCharacter()
	err := c.Save()
	if err != nil {
		t.Error(err)
	}
}

func TestLoadCharacter(t *testing.T) {
	b := initializeTestCharacter()
	c, err := LoadCharacter("tester")
	if err != nil {
		t.Error(err)
	}
	if reflect.DeepEqual(b, c) != true {
		t.Errorf("Load Character mismatch.")
	}
}

func TestCharacterExists(t *testing.T) {
	if CharacterExists("tester") != true {
		t.Errorf("expected to find tester Character")
	}
	if CharacterExists("ghewirehaerawefertaerw") != false {
		t.Errorf("did not expect to find bogus Character")
	}
}

func TestSummonObjectId(t *testing.T) {
	c := initializeTestCharacter()
	c.Realm = &TestRealm
	name, err := c.SummonObjectId(`#1`)
	if err != nil {
		t.Error(err)
	}
	if name != "a mug of coffee" {
		t.Errorf(`failed to summon object #1.`)
	}
}
