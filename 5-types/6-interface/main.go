package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastname string
}

type product struct {
	name  string
	price float64
}

//interfaces are implemented implicitly
func (p person) toString() string {
	return p.name + " " + p.lastname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = person{"Breno", "Machado"}
	print(thing)

	thing = product{"Pants", 82.50}
	print(thing)

}
