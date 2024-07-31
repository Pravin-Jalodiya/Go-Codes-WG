package main

import (
	"fmt"
)

func maps2() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // defining the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}

	newmap := make(map[string]int, 10)
	fmt.Println(newmap)
	fmt.Println(len(newmap))
}
