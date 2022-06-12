package main

import "fmt"

func main() {
	//maps should be initialized
	//var people map[int] string // doesn't work
	people := make(map[int]string)

	people[1] = "Mary"
	people[2] = "Breno"
	people[3] = "Vanessa"

	fmt.Println(people)

	for id, name := range people {
		fmt.Printf("%s (ID: %d)\n", name, id)	
	}

	//remove an element
	delete(people, 1) 
	fmt.Println(people)
}