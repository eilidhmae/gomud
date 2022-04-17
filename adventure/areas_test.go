package gomud

import "testing"

func TestArealistAdd(t *testing.T) {
	var al Arealist
	al.Count = 1
	al.Head = &Area{Title: "test area 1",Id: 1}
	al.Tail = al.Head

	c := al.Add(&Area{Title: "test area 2"})
	if c != 2 {
		t.Errorf("Area count mismatch: expected: 2 got: %d", c)
	}
	if al.Head.Title != "test area 1" {
		t.Errorf("title 1 mismatch")
	}
	if al.Tail.Title != "test area 2" {
		t.Errorf("title 2 mismatch")
	}
}

func TestBuildAreaList(t *testing.T) {
	areasPath := "../areas/"
	expectedCount := 44

	al, err := BuildAreaList(areasPath)
	if err != nil {
		t.Error(err)
	}

	if expectedCount != al.Count {
		t.Errorf("index size mismatch: expected: %d got: %d", expectedCount, al.Count)
	}
}
