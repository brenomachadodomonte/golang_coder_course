package main

import "fmt"

type sport interface {
	turnTurboOn()
}

type luxurious interface {
	makeGoal()
}

type sportAndLuxurious interface {
	sport
	luxurious
}

type bmw7 struct{}

func (b bmw7) makeGoal() {
	fmt.Println("making goal...")
}

func (b bmw7) turnTurboOn() {
	fmt.Println("turning on the turbo...")
}

func main() {
	var b sportAndLuxurious = bmw7{}
	b.turnTurboOn()
	b.makeGoal()

}
