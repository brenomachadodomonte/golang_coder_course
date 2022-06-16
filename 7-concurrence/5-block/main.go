package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //blocking operation
	fmt.Println("it shows only after read the channel")
}

func main() {
	c := make(chan int)
	go routine(c)

	fmt.Println(<-c) //blocking operation
	fmt.Println("Read")
	fmt.Println(<-c) //deadlock
}
