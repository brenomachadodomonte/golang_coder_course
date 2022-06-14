package main

import "fmt"

func printApproveds(approveds ...string) {
	fmt.Println("List of Approveds")
	for i, approved := range approveds {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main(){

	// variatic functions only works with slices not arrays
	approveds := []string{ "Mary", "Breno", "Vanessa", "Bruno" }
	printApproveds(approveds...)
	
}