package main

import "fmt"

type HasName interface {
	GetFullName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetFullName())
}

type Person struct {
	Name string
}

func (person Person) GetFullName() string {
	return person.Name
}

func main() {
	var dimas Person

	dimas.Name = "Dimas"

	var dani Person
	dani.Name = "Dani"

	sayHello(dimas)
	sayHello(dani)
}
