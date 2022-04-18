package gomud

import "testing"

func TestTreeAdd(t *testing.T) {
	titleOne := []byte("test node 1")
	titleTwo := []byte("test node 2")
	l := NewTree(NewNode(titleOne))
	c := l.Add(NewNode(titleTwo))
	if c != 2 {
		t.Errorf("Node count mismatch: expected: 2 got: %d", c)
	}
	head := l.Head
	tail := l.Tail
	if string(*head.Data) != string(titleOne) {
		t.Errorf("title 1 mismatch")
	}
	if string(*tail.Data) != string(titleTwo) {
		t.Errorf("title 2 mismatch")
	}
}

func TestTreeDrop(t *testing.T) {
	titleOne := []byte("test node 1")
	titleTwo := []byte("test node 2")
	titleThree := []byte("test node 3")
	l := NewTree(NewNode(titleOne))
	l.Add(NewNode(titleTwo))
	l.Add(NewNode(titleThree))

	if l.Count != 3 {
		t.Errorf("TreeDrop pre-test count mismatch: expected: 3 got: %d", l.Count)
	}
	h := l.Head.Next
	if h.Id != 2 {
		t.Errorf("TreeDrop pre-test link to 2 mismatch: h.Id is: %d", h.Id)
	}

	dropped := l.Drop(2)
	if l.Count != 2 {
		t.Errorf("TreeDrop result count mismatch: expected: 2 got: %d", l.Count)
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
