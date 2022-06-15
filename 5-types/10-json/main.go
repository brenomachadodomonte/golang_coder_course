package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := product{1, "Laptop", 1899.99, []string{"OnSale", "Eletronic", "Laptop"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	// JSON to Product
	var p2 product
	jsonString := `{"id":2,"name":"smartphone","price":2500.00,"tags":["Eletronic","phones"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tags[1])

}
