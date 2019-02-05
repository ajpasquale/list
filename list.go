// Package list implements a simple singly-linked list.
package list

// Node is an element in the list
type Node struct {
	value interface{}
	next  *Node
}

// List represents a singly-linked list. The zero value for List is an empty list ready to use.
type List struct {
	head   *Node
	tail   *Node
	length int
}

// NewNode returns a new node with it's next link pointing to nil
func NewNode(v interface{}) *Node {
	n := &Node{
		value: v,
		next:  nil,
	}
	return n
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
func (list *List) Add(n *Node) {
	if list.length == 0 {
		list.head = n
	} else {
		list.tail.next = n
	}
	list.tail = n
	list.length++
}

// Pop should make the second node the new head
func (list *List) Pop() interface{} {
	var tNode = list.head
	if list.length != 0 {
		list.head = list.head.next
		list.length--
		if list.length == 0 {
			list.tail = nil
		}
		return tNode.value
	}
	return nil
}

// Push should add a new node to the head
func (list *List) Push(n *Node) {
	n.next = list.head
	list.head = n
	if list.length == 0 {
		list.tail = n
	}
	list.length++
}

// Len returns the current size/length of the list
func (list *List) Len() int {
	return list.length
}

// First returns the current head of the list
func (list *List) First() *Node {
	return list.head
}

// Last returns the current tail of the list
func (list *List) Last() *Node {
	return list.tail
}
