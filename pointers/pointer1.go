package main

import "fmt"

func addT[T int | string](a, b T) {
	fmt.Println(a + b)
}
func main() {
	a := "aman"
	b := "kabab"
	addT(a, b)
}
