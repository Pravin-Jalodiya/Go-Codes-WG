package main

import (
	"GoModule/tests/subtests"
	"example2"
	"fmt"
	"github.com/Pravin-Jalodiya/Puppy"
)

func main() {
	fmt.Println("Hi pravin")
	subtests.Testing3()
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Print("\n" + s1 + " \n" + s2)
	fmt.Print("\n" + example2.Ex2())
}
