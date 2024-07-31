package main

import "fmt"

func oddEven() {
	oddCount := 0

	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			fmt.Println("Even:", i)
		} else {
			oddCount++
		}
	}
	fmt.Println("Odd:", oddCount)
	for i := 0; i < 15; i++ {
		if i&1 == 0 {
			fmt.Println("Even:", i)
		} else {
			oddCount++
		}
	}
	fmt.Println("Odd:", oddCount)
}
