package main

import "fmt"

func main() {
	var nama = [3]string{
		"Dimas",
		"Dliyaur",
		"Rahman",
	}

	fmt.Println(len(nama))
	fmt.Println(cap(nama[1:3]))

	var month = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
	}

	fmt.Println(month)
	var (
		slice = month[4:]
	)

	var slice1 = append(slice, "Desember")
	fmt.Println(month)
	fmt.Println(slice)
	fmt.Println(slice1)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Senin"
	newSlice[1] = "Selasa"

	fmt.Println(newSlice)

	copySlice := make([]string, 1, cap(newSlice))

	copy(copySlice, newSlice)

	fmt.Println(copySlice)

}
