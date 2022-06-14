package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	return total / float64(len(numbers))
}

func main(){

	fmt.Printf("Average: %.2f\n", average(7.7, 9.0, 10, 15))

}