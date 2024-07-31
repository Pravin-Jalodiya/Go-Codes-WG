package main

import "fmt"

type Person struct {
	Name string
	Age  float32
}

func main() {
	s := foo("Pravin")
	fmt.Println(s)
	//s1, RoxDogYears := dogYears("Rox", 70)
	//fmt.Println(s1, RoxDogYears)
	people := []Person{{"Kalu Langda", 23}, {"Bhalu Tagda", 34}, {"Roxie", 42}, {"Xorie", 12}}
	//for _, person := range people {
	//	fmt.Println(person.Name, " is ", person.Age, " years old")
	//}
	s3 := dogYears2(people...)
	for _, value := range s3 {
		fmt.Printf("%s %1.1f\n", value.Name, value.Age)
	}
	dogYears("Anish", 72)

	func2()

	fmt.Println(re1()()())

	counter1 := counter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	counter2 := counter()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter1())

}

func foo(s string) string {
	return fmt.Sprint("My name is ", s)
}

func dogYears(s string, age int) (string, int) {
	age /= 7
	dogs := func() {
		fmt.Printf("%s is %d years old\n", s, age)
	}
	dogs()
	return fmt.Sprint("Age of ", s, " in dog years is"), age
}

func dogYears2(people ...Person) []Person {
	var temp_age float32
	var temp_name string
	var temp []Person
	for _, val := range people {
		temp_name = fmt.Sprint("Age of ", val.Name, " in dog years is ")
		temp_age = val.Age / 7
		temp = append(temp, Person{temp_name, temp_age})
	}
	return temp

}

func re1() func() func() float32 {
	return func() func() float32 {
		return func() float32 {
			return 1.1
		}
	}
}

func counter() func() int {
	x := 0
	return func() int {
		x += 1
		return x
	}
}
