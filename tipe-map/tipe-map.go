package main

import "fmt"

func main() {
	// Tipe Data Map

	person := map[string]string{
		"name":    "dimas",
		"address": "Gresik",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar GO"
	book["author"] = "Dimas"
	book["Ups"] = "Salah"

	fmt.Println(book)

	delete(book, "Ups")

	fmt.Println(book)

}
