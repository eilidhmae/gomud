package gomud

import "testing"

func TestMoblistAdd(t *testing.T) {
	rl := NewMoblist(&Mobile{Title: "test mobile 1",Id: 1})
	c := rl.Add(&Mobile{Title: "test mobile 2"})
	if c != 2 {
		t.Errorf("Mobile count mismatch: expected: 2 got: %d", c)
	}
	if rl.Head.Title != "test mobile 1" {
		t.Errorf("title 1 mismatch")
	}
	if rl.Tail.Title != "test mobile 2" {
		t.Errorf("title 2 mismatch")
	}
}
