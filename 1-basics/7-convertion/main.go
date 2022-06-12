package main

import (
	"fmt"
	"strconv"
)

func main(){

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	// be careful... string will be convert to ASC table
	fmt.Println("Test " + string(97))

	// int to string
	fmt.Println("Test " + strconv.Itoa(123))

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)

	b, _ := strconv.ParseBool("true") // 1 and true is true
	if b {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}