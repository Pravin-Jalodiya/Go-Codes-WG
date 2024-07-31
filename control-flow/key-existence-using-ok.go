package main

import "fmt"

func main() {
	x := map[string]int{"foo": 1, "bar": 2}
	if _, ok := x["foo"]; ok {
		fmt.Println("foo exists")
	}
	if _, ok := x["boo"]; !ok {
		fmt.Println("Boo does not exist")
	}
}
