package main

import "fmt"

func main() {

	nama := "Dimas"
	// Switch Expression

	switch nama {
	case "Dimas":
		fmt.Println("Dimas")
	case "Jiwa":
		fmt.Println("Jiwa")
	default:
		fmt.Println("Lainnya")
	}

	// Switch Short Expression

	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("Nama Kepanjangan")
	default:
		fmt.Println("Nama Pas")
	}

	// Switch without Variable expression
	length := len(nama)

	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Pas")
	}
}
