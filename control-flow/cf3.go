package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("This is init function in action")
}

func modify_select() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	d1 := time.Duration(rand.Intn(325))
	d2 := time.Duration(rand.Intn(325))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 3
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 12
	}()
	//for concurrent code
	select {
	case val := <-ch1:
		fmt.Println("ch1: ", val)
	case val := <-ch2:
		fmt.Println("ch2: ", val)
	}
}
