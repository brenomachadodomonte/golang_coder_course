package main

import "fmt"

func main() {
	i := 1

	// you can't do arithmetics operations (+ - * /) in pointers

	var p* int = nil
	p = &i
	*p++
	i++

	fmt.Println(p, *p, i, &i);
}