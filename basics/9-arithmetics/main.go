package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Minus =>", a-b)
	fmt.Println("Division =>", a/b)
	fmt.Println("Multiply =>", a*b)
	fmt.Println("Module =>", a%b)
	
	// bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("OR =>", a|b)
	fmt.Println("XOR =>", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Max ab =>", math.Max(float64(c), float64(d)))
	fmt.Println("Max cd =>", math.Max(c, d))
	fmt.Println("Min cd =>", math.Min(c, d))
	fmt.Println("Pow cd =>", math.Pow(c, d))

}

