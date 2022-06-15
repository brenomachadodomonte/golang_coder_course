package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastname string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastname
}

func (p *person) setFullName(fullname string) {
	parts := strings.Split(fullname, " ")
	p.name = parts[0]
	p.lastname = parts[1]
}

func main() {
	p1 := person{"Breno", "Machado"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Vanessa Oliveira")
	fmt.Println(p1.getFullName())
}
