package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

// method: function with a receiver
func (p product) priceWithDescount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "Pencil",
		price:    1.79,
		discount: 0.05,
	}

	fmt.Println(product1, product1.priceWithDescount())

	product2 := product{"Laptop", 5450.00, 0.10}
	fmt.Println(product2, product2.priceWithDescount())
}
