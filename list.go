//Package list implements a simple singly-linked list.
package list

type Node struct {
	data int
	next *Node
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

//NewNode returns a new node with it's next link pointing to nil
func NewNode(d int) *Node {
	n := &Node{
		data: d,
		next: nil,
	}
	return n
}

//New return a new list with head/tail pointing to nil
func New() *List {
	l := &List{
		head:   nil,
		tail:   nil,
		length: 0,
	}
	return l
}

//add should append a node to the end of the list
func (list *List) add(n *Node) {
	if list.length == 0 {
		list.head = n
	} else {
		list.tail.next = n
	}
	list.tail = n
	list.length++
}

//pop should make the second node the new head
func (list *List) pop() {
	if list.length != 0 {
		list.head = list.head.next
		list.length--

		if list.length == 0 {
			list.tail = nil
		}
	}
}

//push should add a new node to the head
func (list *List) push(n *Node) {
	//the node has to be become the new head and point to the old head as it's
	//next node
	n.next = list.head
	list.head = n
	if list.length == 0 {
		list.tail = n
	}
	list.length++
}

//len should return the size of the list
func (list *List) len() int {
	return list.length
}

func (list *List) first() *Node {
	return list.head
}
