package list

// Element is an element in the list
type Element struct {
	value int
	next  *Element
}

// newElement returns a new node with it's next link pointing to nil
func newElement(v int) *Element {
	e := &Element{
		value: v,
		next:  nil,
	}
	return e
}

// Next will return a pointer to the
func (e Element) Next() *Element {
	return e.next
}

// Value will return the value of the element
func (e Element) Value() int {
	return e.value
}
