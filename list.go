// Package list implements a simple singly-linked list.
package list

// List represents a singly-linked list. The zero value for List is an empty list ready to use.
type List struct {
	head *Node
	tail *Node
	size int
}

// New return a new list with head/tail pointing to nil
func New() *List {
	l := &List{
		head: nil,
		tail: nil,
		size: 0,
	}
	return l
}

// Add should append a node to the end of the list
func (l *List) Add(v int) {
	n := newNode(v)
	if l.size == 0 {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
	l.size++
}

// Pop should make the second node the new head and return the popped node
func (l *List) Pop() *Node {
	n := l.head
	if l.size != 0 {
		l.head = l.head.next
		l.size--
		if l.size == 0 {
			l.tail = nil
		}
		return n
	}
	return nil
}

// Push should add a new node to the head
func (l *List) Push(v int) {
	n := newNode(v)
	n.next = l.head
	l.head = n
	if l.size == 0 {
		l.tail = n
	}
	l.size++
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

// Values should return the value stored in all elements as a slice of ints
func (l List) Values() []int {
	var vs []int
	var n *Node
	vs = make([]int, 0, l.size)
	for {
		n = l.Pop()
		if n == nil {
			break
		}
		vs = append(vs, n.value)
	}
	return vs
}

// Size returns the current size of the list
func (l *List) Size() int {
	return l.size
}

// First returns the current head of the list
func (l *List) First() *Node {
	return l.head
}

// Last returns the current tail of the list
func (l *List) Last() *Node {
	return l.tail
}

// Sort implement a simple bubble sort and returns a new List
func (l List) Sort() List {
	var nl List
	vs := l.Values()
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(vs); i++ {
			if vs[i-1] > vs[i] {
				vs[i], vs[i-1] = vs[i-1], vs[i]
				swapped = true
			}
		}
	}
	for _, v := range vs {
		nl.Add(v)
	}
	return nl
}
