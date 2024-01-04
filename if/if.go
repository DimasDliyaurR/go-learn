package main

import "fmt"

func main() {

	nama := "Dimas"

	// If Expression

	if nama == "Dimas" {
		fmt.Println("Hallo Dimas")
	} else if nama == "Dliyaur" {
		fmt.Println("Hallo Dliyaur !")
	} else {
		fmt.Println("Hallo")
	}

	// If Short Statement
	if length := len(nama); length > 5 {
		fmt.Printf("Nama Kepanjangan")
	} else {
		fmt.Println("Nama Pas 😭")
	}
}
