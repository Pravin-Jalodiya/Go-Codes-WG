package main

import (
	"fmt"
	"math/rand"
)

func switchh() {
	x := rand.Intn(250)
	fmt.Println("Variable x has value:", x)
	switch {
	case x >= 0 && x <= 100:
		fmt.Println("x is between 0-100")
	case x >= 101 && x <= 200:
		fmt.Println("x is between 101-200")
	case x >= 201 && x <= 250:
		fmt.Println("x is between 201-250")
	}

	Select()
	//modified_select()
}
