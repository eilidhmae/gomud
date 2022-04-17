package gomud

import "testing"

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
