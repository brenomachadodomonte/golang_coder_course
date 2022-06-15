package main

import "fmt"

type car struct {
	name     string
	velocity int
}

type ferrari struct {
	car
	turboOn bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.velocity = 0
	f.turboOn = true

	fmt.Printf("Is the turbo on: %v\n", f.turboOn)
	fmt.Println(f)
}
