package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // writing data in the channel
	<-ch    // reading data in the channel

	ch <- 2
	fmt.Println(<-ch)

}
