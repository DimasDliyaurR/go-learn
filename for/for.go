package main

import "fmt"

func main() {
	// Perulangan map
	book := make(map[string]string)
	book["author"] = "Dimas"
	book["title"] = "Programer"

	for key, value := range book {
		fmt.Println(key, value)
	}

	// Perulangan continues

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println("Perulangan ke ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			break
		}

		fmt.Println("Perulangan ke ", i)
	}

	// For / Perulangan

	for i := 0; i < 5; i++ {
		fmt.Println("Baris ke ", i)
	}

	// perulangan range slice

	names := []string{"Dimas", "Dliyaur", "Rahman"}
	for i, name := range names {
		fmt.Printf("Hallo saya %v dan itu adalah nama ke %v \n", name, i+1)
	}
}
