package gomud

import "testing"

func TestMoblistAdd(t *testing.T) {
	l := NewMoblist(&Mobile{Title: "test mobile 1",Id: 1})
	c := l.Add(&Mobile{Title: "test mobile 2"})
	if c != 2 {
		t.Errorf("Mobile count mismatch: expected: 2 got: %d", c)
	}
	if l.Head.Title != "test mobile 1" {
		t.Errorf("title 1 mismatch")
	}
	if l.Tail.Title != "test mobile 2" {
		t.Errorf("title 2 mismatch")
	}
}
