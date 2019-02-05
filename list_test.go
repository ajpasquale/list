package list

import "testing"

func TestNewNode(t *testing.T) {
	n := NewNode(1)
	if n == nil {
		t.Fatal("should create new node")
	}
	if n.data != 1 {
		t.Fatal("should set data in node")
	}
}

func TestNewList(t *testing.T) {
	l := New()
	if l == nil {
		t.Fatal("should create a new list")
	}
	if l.head != nil {
		t.Fatal("head should be empty")
	}
	if l.tail != nil {
		t.Fatal("tail should be empty")
	}
	if l.Len() != 0 {
		t.Fatal("list should have zero length")
	}
}

func TestListAdd(t *testing.T) {
	list := New()
	node := NewNode(19)

	list.Add(node)

	if list.head != node {
		t.Fatal("should reference the same object")
	}
	if list.tail != node {
		t.Fatal("should reference the same object")
	}
	if list.length != 1 {
		t.Fatal("should have a length of one")
	}

	node2 := NewNode(666)

	list.Add(node2)

	if list.head == node2 {
		t.Fatal("should not reference same object")
	}
	if list.tail != node2 {
		t.Fatal("should reference the same object")
	}
	if list.Len() != 2 {
		t.Fatal("should have a length of two")
	}
}

func TestListPop(t *testing.T) {
	list := New()
	node := NewNode(19)
	list.Add(node)
	list.Pop()
	list.Pop()
}

func TestListPush(t *testing.T) {
	list := New()

	nodes := []Node{
		{0, nil},
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
		{5, nil},
		{6, nil},
		{7, nil},
		{8, nil},
		{9, nil},
	}

	for i, n := range nodes {
		list.Push(&n)
		if list.head == nil {
			t.Fatal("should be set after push")
		} else if list.head != &n {
			t.Fatal("should reference the current object")
		}
		if list.tail == nil {
			t.Fatal("should be set after push")
		}
		if list.Len() == 1 {
			if list.head != list.tail {
				t.Fatal("should reference the same object")
			}
		}
		if list.Len() != i+1 {
			t.Fatal("should have the same length")
		}
	}
}

func TestListFirst(t *testing.T) {
	list := New()
	an := Node{
		data: 1,
		next: nil,
	}

	pn := Node{
		data: 2,
		next: nil,
	}

	if list.First() != nil {
		t.Fatal("should not be set yet")
	}

	list.Add(&an)

	if list.First() != &an {
		t.Fatal("should reference to the same object")
	}

	list.Push(&pn)

	if list.First() != &pn {
		t.Fatal("should reference the current object")
	}
}
