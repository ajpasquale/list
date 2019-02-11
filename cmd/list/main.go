package main

import (
	"fmt"

	"github.com/ajpasquale/list"
)

func main() {
	l := list.New()
	l.Add(1)
	for e := l.First(); e != nil; e = e.Next() {
		fmt.Println(e.Value())
	}
}
