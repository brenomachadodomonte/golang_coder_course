package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" Line.")

	fmt.Println(" New")
	fmt.Println("Line.")

	x := 3.141516
	//fmt.Println("The x value is " + x) // doesn't work

	xs := fmt.Sprint(x)
	fmt.Println("The x value is " + xs)
	fmt.Println("The x value is", xs)

	fmt.Printf("The x value is %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "hello"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}