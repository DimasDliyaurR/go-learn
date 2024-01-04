package main

import "fmt"

// struct
type Customer struct {
	Name, Address string
	Age           int8
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "I am", customer.Name)
}

func main() {
	// Output struct
	var dimas Customer

	dimas.Name = "Dimas Dliyaur Rahman"
	dimas.Address = "Gresik"
	dimas.Age = 20

	fmt.Println(dimas)
	fmt.Println(dimas.Name)
	fmt.Println(dimas.Address)
	fmt.Println(dimas.Age)
	dimas.sayHello("Joko")
}
