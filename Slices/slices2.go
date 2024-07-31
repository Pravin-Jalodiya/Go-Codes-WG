package main

import (
	"fmt"
)

func sliced() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:6]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}
