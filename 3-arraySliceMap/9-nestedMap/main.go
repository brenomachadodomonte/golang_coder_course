package main

import "fmt"

func main() {
	
	wokersAndSalariesPerChar := map[string]map[string]float64 {
		"M": {
			"Mary": 1251.66,
			"Marco": 4321.99,
		},
		"B": {
			"Breno": 15123.40,
			"Bruno": 10000.00,
		},
		"V": {
			"Vanessa": 16550.10,
			"Vasco": 23500.10,
		},
	}

	delete(wokersAndSalariesPerChar, "M")

	for letter, worker := range wokersAndSalariesPerChar {
		fmt.Println(letter, worker)
	}

}
