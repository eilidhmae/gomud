package gomud

import "testing"

func TestRoomlistAdd(t *testing.T) {
	rl := NewRoomlist(&Room{Title: "test room 1",Id: 1})
	c := rl.Add(&Room{Title: "test room 2"})
	if c != 2 {
		t.Errorf("Room count mismatch: expected: 2 got: %d", c)
	}
	if rl.Head.Title != "test room 1" {
		t.Errorf("title 1 mismatch")
	}
	if rl.Tail.Title != "test room 2" {
		t.Errorf("title 2 mismatch")
	}
}
