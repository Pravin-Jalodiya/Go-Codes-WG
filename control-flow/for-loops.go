package main

import "fmt"

func loops() {
	var x int
	for x < 10 {
		fmt.Println("x is", x)
		x++
	}
	for {
		fmt.Println("x is", x)
		if x == 15 {
			break
		}
		x++
	}
}
