// Package list implements a simple singly-linked list.
package list

// List represents a singly-linked list. The zero value for List is an empty list ready to use.
type List struct {
	head   *Node
	tail   *Node
	length int
}

// New return a new list with head/tail pointing to nil
func New() *List {
	l := &List{
		head:   nil,
		tail:   nil,
		length: 0,
	}
	return l
}

// Add should append a node to the end of the list
func (l *List) Add(v int) {
	n := newNode(v)
	if l.length == 0 {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
	l.length++
}

// Push should add a new node to the head
func (l *List) Push(v int) {
	n := newNode(v)
	n.next = l.head
	l.head = n
	if l.length == 0 {
		l.tail = n
	}
	l.length++
}

// Pop should make the second node the new head and return the popped node
func (l *List) Pop() *Node {
	n := l.head
	if l.length != 0 {
		l.head = l.head.next
		l.length--
		if l.length == 0 {
			l.tail = nil
		}
		return n
	}
	return nil
}

// Traverse will return the node at a given index
func (l List) Traverse(index int) *Node {
	if index > l.length || index < 0 {
		return nil
	}
	var n *Node
	for i := 0; i <= index; i++ {
		n = l.Pop()
	}
	return n
}

// Insert will insert a value at a given index
func (l *List) Insert(v int, index int) *Node {
	if index < 0 {
		return nil
	}
	if index > l.length {
		l.Add(v)
		return l.Last()
	}
	n := newNode(v)
	leading := l.Traverse(index - 1)
	trailing := leading.next
	leading.next = n
	n.next = trailing
	l.length++
	return n
}

// Find will return a node in the list based on it's value
func (l List) Find(v int) *Node {
	var n *Node
	for {
		n = l.Pop()
		if n == nil {
			break
		} else if n.value == v {
			return n
		}
	}
	return nil
}

// ToSlice will return the value stored in all elements as a slice of ints
func (l List) ToSlice() []int {
	var vs []int
	var n *Node
	vs = make([]int, 0, l.length)
	for {
		n = l.Pop()
		if n == nil {
			break
		}
		vs = append(vs, n.value)
	}
	return vs
}

// Length returns the current length of the list
func (l *List) Length() int {
	return l.length
}

// First returns the current head of the list
func (l *List) First() *Node {
	return l.head
}

// Last returns the current tail of the list
func (l *List) Last() *Node {
	return l.tail
}
