package main

type node struct {
	data int
	next *node
}
type linkedList struct {
	head    *node
	current *node
}

func newNode(d int) *node {
	n := &node{
		data: d,
		next: nil,
	}
	return n
}

func (list *linkedList) appendNode(n *node) {
	if list.head != nil {
		list.current.next = n
		list.current = list.current.next
	} else {
		list.head = n
		list.current = list.head
	}
}

func main() {
}
