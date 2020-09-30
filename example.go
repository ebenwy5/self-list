package main

import (
	"fmt"
	"strconv"

	"github.com/ebenwy5/self-list/list"
)

func Example_1() {
	l := list.NewList()
	testCase := 10

	// Add 10 input
	for i := 1; i <= testCase; i++ {
		l.Add("input_" + strconv.Itoa(i))
	}

	fmt.Println(l.NodeMap)
	fmt.Println(len(l.NodeMap))

	for i := 0; i < testCase; i++ {
		fmt.Println(i, l.Get(i))
	}

	fmt.Println("Size:", l.Size())
	fmt.Println("Cap:", l.Cap, "\n")

	// Remove 5
	fmt.Println("Remove:", 5, l.Remove(5))

	for i := 0; i < testCase; i++ {
		fmt.Println(i, l.Get(i))
	}

	fmt.Println("Size:", l.Size())
	fmt.Println("Cap:", l.Cap, "\n")

	// Remove 8
	fmt.Println("Remove:", 8, l.Remove(8))

	for i := 0; i < testCase; i++ {
		fmt.Println(i, l.Get(i))
	}

	fmt.Println("Size:", l.Size())
	fmt.Println("Cap:", l.Cap, "\n")

	// Remove 0
	fmt.Println("Remove:", 0, l.Remove(0))

	for i := 0; i < testCase; i++ {
		fmt.Println(i, l.Get(i))
	}

	fmt.Println("Size:", l.Size())
	fmt.Println("Cap:", l.Cap)
}
