package main

import "fmt"

func main() {
	
	wokersAndSalaries := map[string]float64 {
		"Mary": 1200.00,
		"Breno": 9500.00,
		"Vanessa": 2000.00, //comma is required
	}

	wokersAndSalaries["John"] = 13350.00

	delete(wokersAndSalaries, "Bruno") // it does nothing

	for name, salary := range wokersAndSalaries {
		fmt.Println(name, salary)
	}

}
