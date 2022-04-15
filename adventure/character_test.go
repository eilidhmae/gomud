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
