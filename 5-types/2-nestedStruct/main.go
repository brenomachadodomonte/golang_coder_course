package main

import "fmt"

type item struct {
	productID int
	qty       int
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) total() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.qty)
	}
	return total
}

func main() {
	order1 := order{
		userID: 1,
		items: []item{
			{1, 2, 12.10},
			{2, 1, 10.50},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("The total of the order is R$ %.2f", order1.total())
}
