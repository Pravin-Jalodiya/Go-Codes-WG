package main

import (
	"encoding/json"
	"fmt"
)

type intern struct {
	Name               string
	Age                int
	Salary             int
	InternshipDuration int
	Department         string
	Batch              int
}

func main() {
	intern2 := new(intern)
	intern2.Name = "Alex"
	intern2.Age = 12
	intern2.Salary = 100
	intern2.Department = "Software Development"
	fmt.Println(*intern2)
	intern1 := &intern{
		Name:               "Pravin",
		Age:                21,
		Salary:             1000,
		InternshipDuration: 12,
		Batch:              4,
	}

	user, err := json.Marshal(intern1)

	if err != nil {
		fmt.Println(err)
		return
	}

	intern3 := intern{}

	fmt.Print(string(user))

	json1 := `{"Name":"Pravin","Age":21,"Salary":1000,"InternshipDuration":12,"Department":"","Batch":4}`

	//json.Unmarshal([]byte(json1), &intern3)

	unmarshal(json1, &intern3)
	fmt.Println(intern3)
}
