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
	b := l.outOfBounds(index)
	if b {
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
	b := l.outOfBounds(index)
	if b {
		return nil
	}
	n := newNode(v)
	leading := l.Traverse(index - 1)
	trailing := leading.next
	leading.next = n
	n.next = trailing
	l.length++
	return n
}

// Remove will remove a node at a given index and returns the unwanted node
func (l *List) Remove(index int) *Node {
	b := l.outOfBounds(index)
	if b {
		return nil
	}
	if index == 0 {
		return l.Pop()
	}
	leading := l.Traverse(index - 1)
	unwanted := leading.next
	leading.next = unwanted.next
	l.length--
	return unwanted
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

// Reverse will reverse the list
func (l *List) Reverse() {
	if l.head.next == nil || l.head == nil {
		return
	}
	first := l.head
	l.tail = l.head
	second := first.next
	for second != nil {
		var temp = second.next
		second.next = first
		first = second
		second = temp
	}
	l.head.next = nil
	l.head = first
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

// outOfBounds simply ensures that an index passed to it is within the list length
func (l List) outOfBounds(index int) bool {
	if index < 0 || index >= l.length {
		return true
	}
	return false
}

// Length returns the current length of the list
func (l *List) Length() int {
	return l.length
}

// Head returns the current head of the list
func (l *List) Head() *Node {
	return l.head
}

// Tail returns the current tail of the list
func (l *List) Tail() *Node {
	return l.tail
}
