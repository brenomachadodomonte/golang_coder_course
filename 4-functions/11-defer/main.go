package main

import "fmt"

func getValue(number int) int {
	defer fmt.Println("=== END ===")

	if number >= 5000 {
		fmt.Println("High value")
		return 5000
	} 
	fmt.Println("Low Value")
	return number
}

func main() {
	fmt.Println(getValue(3000))
}