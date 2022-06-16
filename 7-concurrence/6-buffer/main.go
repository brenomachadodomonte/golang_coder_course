package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executed") // not execute because buffer is full
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)

	time.Sleep(time.Second)
	go routine(ch)
	fmt.Println(<-ch)
}
