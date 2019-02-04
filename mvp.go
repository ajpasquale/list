package main

import "fmt"

type node struct {
	data []byte
	next *node
}
type linkedList struct {
	head    *node
	current *node
	count   int
}

func (list linkedList) addNode(d []byte) {
	var h node
	var nx node

	if list.head != nil {
		list.current.next = &nx
		list.current = list.current.next
	} else {
		list.head = &h
		list.current = list.head
	}

}
func main() {

	var list linkedList

	var h node

	var nx node

	h.data = []byte{1, 2, 3, 4}
	nx.data = []byte{6, 6, 6}

	if list.head != nil {
		list.current.next = &nx
		list.current = list.current.next
	} else {
		list.head = &h
		list.current = list.head
	}

	fmt.Println(list.head.data)
	fmt.Println(list.current.data)
}
