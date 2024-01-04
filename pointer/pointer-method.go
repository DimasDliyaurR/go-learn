package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	dimas := Man{"Dimas"}
	dimas.Married()

	fmt.Println(dimas.Name)
}
