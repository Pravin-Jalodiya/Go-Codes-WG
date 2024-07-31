package main

import "fmt"

func main1() {
	var slc1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slc1)
	fmt.Printf("%T\n", slc1)
	slc2 := make([]int, 10)
	// slc3 := make([]int,10,20)
	slc2[0] = 100
	fmt.Println("Length of slice : ", len(slc2))
	fmt.Println("Capacity of slice : ", cap(slc2))

	slc2 = append(slc2, 101) // Append one more element
	fmt.Println("After appending one element:")
	fmt.Println("Slice:", slc2)
	fmt.Println("Length:", len(slc2))
	fmt.Println("Capacity:", cap(slc2))
	sliced()
	sliced2()
	del()
	var slc3 [][]int
	slc3 = append(slc3, [][]int{{1}, {112}}...)
	fmt.Println(slc3)

	sliced3()
	//D2Slice()

}
