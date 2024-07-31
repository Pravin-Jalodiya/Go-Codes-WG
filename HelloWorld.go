package main

import (
	"fmt"
	"package1"
)

func main() {
	package1.HelloWorld()
	var arr = [...]int{10, 2, 3, 12, 45, 12, 3, 21, 1, 2, 3, 434, 745, 74, 5345, 2, 424}
	fmt.Printf("%v \n", arr)
	fmt.Print(len(arr))
}
