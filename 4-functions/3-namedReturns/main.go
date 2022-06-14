package main

import "fmt"

func change(p1, p2 int) (second int, first int) {
	second = p2
	first = p1

	return //clean returned
}

func main() {
	
	r1, r2 := change(2, 3)
	fmt.Println(r1, r2)

	second, first := change(3, 2)
	fmt.Println(second, first)

}