package main

import "fmt"

func sliced3() {

	x := []int{1, 2, 3, 5, 6, 7, 8, 9}

	fmt.Println(x[1:4])

	y := x[:3]

	fmt.Println(y)

	fmt.Println(x[1:6], x[2:6], x[6:8])

	x = append(x, 6, 7, 8, 9)
	x = append(x, y...)

	fmt.Println(x)
}
