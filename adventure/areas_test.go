package gomud

import "testing"

func TestArealistAdd(t *testing.T) {
	al := NewArealist(&Area{Title: "test area 1",Id: 1})
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
	if al.Tail.Data == nil {
		t.Errorf("No Data for al.Tail")
	}
}

func TestBuild(t *testing.T) {
	areasPath := "../areas/"
	al, err := BuildAreaList(areasPath)
	if err != nil {
		t.Error(err)
	}
	cur := al.Head
	for cur.Data == nil {
		cur = cur.Next
	}
	if err := cur.Build(); err != nil {
		t.Errorf("Area.Build: %s", err)
	}
}

func TestAreaLookup(t *testing.T) {
	areasPath := "../areas/"
	al, err := BuildAreaList(areasPath)
	if err != nil {
		t.Error(err)
	}
	cur := al.Head
	for cur.Data == nil {
		cur = cur.Next
	}
	if err := cur.Build(); err != nil {
		t.Errorf("Area.Build: %s", err)
	}
	cur = al.Lookup(42)
	if cur.Id != 42 {
		t.Errorf("current id mismatch")
	}
	if cur.Previous.Id != 41 {
		t.Errorf("previous id mismatch")
	}
	if cur.Next.Id != 43 {
		t.Errorf("next id mismatch")
	}

	// test out of range: lower returns Head, higher returns Tail
	cur = al.Lookup(0)
	if cur.Id != al.Head.Id {
		t.Errorf("lower bounds unexpected id %d", cur.Id)
	}
	cur = al.Lookup(100)
	if cur.Id != al.Tail.Id {
		t.Errorf("upper bounds unexpected id %d", cur.Id)
	}
}
