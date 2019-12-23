package list

// Node is an node in the list
type Node struct {
	value int
	next  *Node
}

// newNode returns a new node with it's next link pointing to nil
func newNode(v int) *Node {
	n := &Node{
		value: v,
		next:  nil,
	}
	return n
}

// Next will return a pointer to the
func (n Node) Next() *Node {
	return n.next
}

// Value will return the value of the element
func (n Node) Value() int {
	return n.value
}
