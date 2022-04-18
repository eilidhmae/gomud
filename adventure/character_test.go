package gomud

import (
	"reflect"
	"testing"
)

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
