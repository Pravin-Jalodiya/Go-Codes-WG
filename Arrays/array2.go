package main

import "fmt"

var Xyz = 19

func Array2() {
	//specific index initialization
	arr1 := [5]int{3: 21, 0: 40}
	fmt.Println(arr1)
	//2D arrays

	var twoD = [2][...]int{{1, 2, 3}, {3, 4, 5}}
	fmt.Println(twoD)
	fmt.Printf("%T\n", twoD)
}
