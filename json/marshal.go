package main

import (
	"encoding/json"
	"fmt"
)

func marshal(a intern) {
	x, err := json.Marshal(a)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(x))
}
func unmarshal(s string, v *intern) {
	u := []byte(s)
	err := json.Unmarshal(u, v)
	if err != nil {
		panic(err)
		return
	}
}
