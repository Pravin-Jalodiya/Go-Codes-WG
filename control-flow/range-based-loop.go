package main

import "fmt"

func ranged() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, val := range x {
		fmt.Println("Value on index", i, "is", val)
	}

}
