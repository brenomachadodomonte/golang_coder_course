package main

import "fmt"

func main(){
	// arrays are same type and static
	var grades [3]float64
	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1
	fmt.Println(grades)

	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	avg := total / float64(len(grades))
	fmt.Printf("Total: %.2f\n", total)
	fmt.Printf("Average: %.2f", avg)
}