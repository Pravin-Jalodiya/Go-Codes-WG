package main

import "fmt"

func main() {
	x := make([][]string, 3)
	for i, _ := range x {
		x[i] = make([]string, 1)
		for j, _ := range x[i] {
			x[i][j] = fmt.Sprintf("Hey %d", i)
		}
	}

	for _, val1 := range x {
		for _, val2 := range val1 {
			fmt.Printf("%s\n", val2)
		}
	}

}
