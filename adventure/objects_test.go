package gomud

import "testing"

func TestObjlistAdd(t *testing.T) {
	rl := NewObjlist(&Object{Title: "test object 1",Id: 1})
	c := rl.Add(&Object{Title: "test object 2"})
	if c != 2 {
		t.Errorf("Object count mismatch: expected: 2 got: %d", c)
	}
	if rl.Head.Title != "test object 1" {
		t.Errorf("title 1 mismatch")
	}
	if rl.Tail.Title != "test object 2" {
		t.Errorf("title 2 mismatch")
	}
}
