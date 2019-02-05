package list

// Element is an element in the list
type Element struct {
	value int
	next  *Element
}

// newElement returns a new node with it's next link pointing to nil
func newElement(v int) *Element {
	n := &Element{
		value: v,
		next:  nil,
	}
	return n
}
