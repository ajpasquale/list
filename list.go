// Package list implements a simple singly-linked list.
package list

import "errors"

// List represents a singly-linked list. The zero value for List is an empty list ready to use.
type List struct {
	head *Element
	tail *Element
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
	e := newElement(v)
	if l.size == 0 {
		l.head = e
	} else {
		l.tail.next = e
	}
	l.tail = e
	l.size++
}

// Pop should make the second node the new head and return the popped element
func (l *List) Pop() *Element {
	e := l.head
	if l.size != 0 {
		l.head = l.head.next
		l.size--
		if l.size == 0 {
			l.tail = nil
		}
		return e
	}
	return nil
}

// Push should add a new node to the head
func (l *List) Push(v int) {
	e := newElement(v)
	e.next = l.head
	l.head = e
	if l.size == 0 {
		l.tail = e
	}
	l.size++
}

// Find will return an element in the list based on it's value
func (l *List) Find(v int) (*Element, error) {
	var e *Element
	e = l.head
	for {
		if e == nil {
			return nil, errors.New("the value is not in this list")
		}
		if e.value == v {
			break
		}
		e = e.Next()
	}
	return e, nil
}

// Values should return the value stored in all elements as a slice of ints
func (l List) Values() []int {
	var vs []int
	var e *Element
	vs = make([]int, 0, l.size)
	for {
		e = l.Pop()
		if e == nil {
			break
		}
		vs = append(vs, e.Value())
	}
	return vs
}

// Size returns the current size of the list
func (l *List) Size() int {
	return l.size
}

// First returns the current head of the list
func (l *List) First() *Element {
	return l.head
}

// Last returns the current tail of the list
func (l *List) Last() *Element {
	return l.tail
}
