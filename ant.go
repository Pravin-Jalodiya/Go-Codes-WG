package main

import (
	"fmt"
	"github.com/Pravin-Jalodiya/Version-Testing/v2/testing2"
)

func callMe(fun1 func){
	fun1()
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(testing2.Testing2())
	fmt.Println()
	//callMe(func{
	//	fmt.Println("Hello World")
	//})()


}
