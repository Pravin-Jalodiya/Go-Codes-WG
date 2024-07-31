package main

import "fmt"

func func2() {
	add := func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(4, 5))

}
