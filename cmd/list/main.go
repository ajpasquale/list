package main

import (
	"fmt"

	"github.com/ajpasquale/list"
)

func main() {
	l := list.New()

	for i := 0; i < 100; i++ {
		l.Add(i)
	}

	for {
		v, err := l.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(v)
	}
}
