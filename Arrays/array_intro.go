package main

import (
	"fmt"
	"os"
)

var x int = 2

func main() {
	//arrayIntro()
	Array2()
	fmt.Println(Xyz)
}
func arrayIntro() {
	// different ways to declare array
	// fixed size // store same data types

	var a = [5]int{}
	fmt.Printf("%p \n", &a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	a[0] = 1
	a[0] = 2
	a[4] = 4
	x--
	if x == 0 {
		os.Exit(0)
	}
	arrayIntro()
}
