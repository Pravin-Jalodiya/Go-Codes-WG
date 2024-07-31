package main

import "fmt"

type Human interface {
	talk()
	walk()
}

type person struct {
	name string
	age  int
}

func (p person) talk() {
	fmt.Println("Hii I am ", p.name)
}
func (p person) walk() {
	fmt.Println(p.name, " starts walking...")
}

func activity(h Human) {
	h.talk()
	h.walk()
}

func main() {

	p1 := person{
		name: "Pravin",
		age:  22}
	p2 := person{
		name: "Bhalu",
		age:  46}

	activity(p1)
	activity(p2)

}
