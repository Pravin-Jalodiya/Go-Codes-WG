package main

import "fmt"

func del() {
	sl := []int{1, 2, 3}
	sl = append(sl[:1], sl[2:]...) //delete range here is : [1,2) half open interval
	fmt.Println(sl)
}
