package gomud

import (
	"reflect"
	"testing"
)

var titleOne []byte = []byte("test node 1")
var titleTwo []byte = []byte("test node 2")
var titleThree []byte = []byte("test node 3")

func TestNewTree(t *testing.T) {
	p := &Node{}
	tree := NewTree(p)
	if tree.Count != 1 {
		t.Errorf("NewTree count mismatch.")
	}
	if tree.Head != p {
		t.Errorf("NewTree head mismatch.")
	}
	if tree.Tail != p {
		t.Errorf("NewTree tail mismatch.")
	}
}

func TestNewNode(t *testing.T) {
	p := NewNode(titleOne)
	if p.Id != 1 {
		t.Errorf("NewNode id mismatch.")
	}
	if reflect.DeepEqual(*p.Data, titleOne) != true {
		t.Errorf("NewNode data mismatch.")
	}
}

func TestNewNodeByName(t *testing.T) {
	name := "newbie"
	p := NewNodeByName(name, titleTwo)
	if p.Name != name {
		t.Errorf("NewNodeByName name mismatch.")
	}
	if reflect.DeepEqual(*p.Data, titleTwo) != true {
		t.Errorf("NewNodeByName data mismatch.")
	}
}

func TestLookupId(t *testing.T) {
	l := NewTree(NewNode(titleOne))
	l.Add(NewNode(titleTwo))
	l.Add(NewNode(titleThree))
	p := l.LookupId(2)
	if p == nil {
		t.Errorf("LookupId returned nil.")
	}
	if p.Id != 2 {
		t.Errorf("LookupId id mismatch.")
	}
	if reflect.DeepEqual(*p.Data, titleTwo) != true {
		t.Errorf("LookupId data mismatch.")
	}
}

func TestLookupName(t *testing.T) {
	l := NewTree(NewNodeByName("first", titleOne))
	l.Add(NewNodeByName("second", titleTwo))
	l.Add(NewNodeByName("third", titleThree))
	p := l.LookupName("second")
	if p == nil {
		t.Errorf("LookupName returned nil.")
	}
	if p.Id != 2 {
		t.Errorf("LookupName id mismatch.")
	}
	if reflect.DeepEqual(*p.Data, titleTwo) != true {
		t.Errorf("LookupName data mismatch.")
	}
}

func TestTreeAdd(t *testing.T) {
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

func TestHasData(t *testing.T) {
	l := NewTree(NewNode(titleOne))
	l.Add(NewNode(titleTwo))
	l.Add(NewNode(titleThree))

	// test for success
	id, ok := l.HasData(string(titleTwo))
	if !ok {
		t.Errorf("HasData failed to match.")
	}
	if id != 2 {
		t.Errorf("HasData id mismatch.")
	}

	// test for failure
	_, ok = l.HasData("blonk")
	if ok {
		t.Errorf("HasData false match.")
	}
}
