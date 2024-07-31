package package1

import (
	"fmt"
	puppy "github.com/Pravin-Jalodiya/Puppy"
)

func HelloWorld() {
	fmt.Println("Hello World")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Version())
	x := 100
	if x != 10 {
		fmt.Println("x is not 10")
	} else if x == 10 {
		fmt.Println("x is 10")
	}
}
