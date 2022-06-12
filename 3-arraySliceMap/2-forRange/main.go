package main

import "fmt"

func main(){
	numbers := [...]int{1, 2, 3, 4, 5} // the compiler counts the array's size

	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i+1, number)
	}
	
	for _, number := range numbers { // ignore index
		fmt.Printf("%d\n", number)
	}
	
}