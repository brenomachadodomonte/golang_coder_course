package main

import (
	"fmt"
)

func main() {
	
	x := 1
	y := 2

	x++ // x += 1 --> x = x + 1
	fmt.Println(x)

	y--
	fmt.Println(y)

	// x == y-- doesn't work
}

