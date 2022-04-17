package gomud

import "testing"

func TestObjlistAdd(t *testing.T) {
	l := NewObjlist(&Object{Desc: "test object 1",Id: 1})
	c := l.Add(&Object{Desc: "test object 2"})
	if c != 2 {
		t.Errorf("Object count mismatch: expected: 2 got: %d", c)
	}
	if l.Head.Desc != "test object 1" {
		t.Errorf("title 1 mismatch")
	}
	if l.Tail.Desc != "test object 2" {
		t.Errorf("title 2 mismatch")
	}
}

func TestObjlistDrop(t *testing.T) {
	l := NewObjlist(&Object{Desc: "test object 1", Id: 1})
	l.Add(&Object{Desc: "test object 2", Id: 2})
	l.Add(&Object{Desc: "test object 3", Id: 3})

	if l.Count != 3 {
		t.Errorf("ObjlistDrop pre-test count mismatch: expected: 3 got: %d", l.Count)
	}
	h := l.Head.Next
	if h.Id != 2 {
		t.Errorf("ObjlistDrop pre-test link mismatch: Head.Next.Id is: %d", h.Id)
	}

	dropped := l.Drop(2)
	if l.Count != 2 {
		t.Errorf("ObjlistDrop result count mismatch: expected: 2 got: %d", l.Count)
	}
	if l.Head.Next != l.Tail {
		t.Errorf("Head.Next does not point to Tail after Drop")
	}
	if l.Tail.Previous != l.Head {
		t.Errorf("Tail.Previous does not point to Head after Drop")
	}
	if dropped.Id != 2 {
		t.Errorf("Drop returned different Id: expected: 2 got: %d", dropped.Id)
	}
}
