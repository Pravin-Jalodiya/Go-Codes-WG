package main

import "fmt"

func nested() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Outer loop", i)
		for j := 1; j <= 10; j++ {
			fmt.Println("Inner loop ", j)
		}
	}
}
