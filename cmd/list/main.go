package main

import (
	"github.com/ajpasquale/list"
)

func main() {
	l := list.New()

	for i := 0; i < 10; i++ {
		nn := list.NewNode("ad")
		l.Add(nn)
	}
}
