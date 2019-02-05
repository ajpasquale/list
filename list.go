// Package list implements a simple singly-linked list.
package list

import "errors"

// List represents a singly-linked list. The zero value for List is an empty list ready to use.
type List struct {
	head   *Element
	tail   *Element
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
	e := newElement(v)
	if l.length == 0 {
		l.head = e
	} else {
		l.tail.next = e
	}
	l.tail = e
	l.length++
}

// Pop should make the second node the new head
func (l *List) Pop() (int, error) {
	e := l.head
	if l.length != 0 {
		l.head = l.head.next
		l.length--
		if l.length == 0 {
			l.tail = nil
		}
		return e.value, nil
	}
	return 0, errors.New("list is empty")
}

// Push should add a new node to the head
func (l *List) Push(v int) {
	e := newElement(v)
	e.next = l.head
	l.head = e
	if l.length == 0 {
		l.tail = e
	}
	l.length++
}

// Len returns the current size/length of the list
func (l *List) Len() int {
	return l.length
}

// First returns the current head of the list
func (l *List) First() *Element {
	return l.head
}

// Last returns the current tail of the list
func (l *List) Last() *Element {
	return l.tail
}
