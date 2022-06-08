package main

import (
	m "math" //alias for package
	"fmt"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //type float64 deduced by compiler
	
	//Reduced form
	area := PI * m.Pow(raio, 2)
	fmt.Println("The are is", area)

	//declare multiple consts
	const (
		a = 1
		b = 2
	)

	//declare multiple vars
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d);

	var e, f bool = true, false
	fmt.Println(e, f);

	g, h, i := 2, false, "hello"
	fmt.Println(g, h, i);
}