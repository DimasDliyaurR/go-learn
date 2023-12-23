package main

import "fmt"

const (
	umur   int32   = 12
	hobi   string  = "Makan"
	income float32 = 20000
)

func main() {
	// type health int32
	// type damage int32
	// type lari bool
	// type defense bool

	// // Warrior
	// var (
	// 	WarriorHealth  health  = -12000
	// 	WarriorDamage  damage  = 1300
	// 	WarriorLari    lari    = false
	// 	WarriorDefense defense = false
	// )

	// // Enemy
	// var (
	// 	EnemyHealth  health  = 300
	// 	EnemyDamage  damage  = 100
	// 	EnemyLari    lari    = false
	// 	EnemyDefense defense = false
	// )

	// /*
	// 	Warrior State
	// */
	// fmt.Printf("Health Warrior = %d\n", WarriorHealth)
	// fmt.Printf("Attach Warrior = %d\n", WarriorDamage)
	// fmt.Printf("Warrior escape = %t\n", WarriorLari)
	// fmt.Printf("Warrior Defense = %t\n", WarriorDefense)

	// /*
	// 	Enemy State
	// */
	// fmt.Printf("Health Enemy = %d\n", EnemyHealth)
	// fmt.Printf("Attach Enemy = %d\n", EnemyDamage)
	// fmt.Printf("Enemy escape = %t\n", EnemyLari)
	// fmt.Printf("Enemy Defense = %t\n", EnemyDefense)

	// Slice And Arrays
	// var nama = [3]string{
	// 	"Dimas",
	// 	"Dliyaur",
	// 	"Rahman",
	// }

	// fmt.Println(len(nama))
	// fmt.Println(cap(nama[1:3]))

	// var month = [...]string{
	// 	"januari",
	// 	"februari",
	// 	"maret",
	// 	"april",
	// 	"mei",
	// 	"juni",
	// 	"juli",
	// 	"agustus",
	// 	"september",
	// 	"oktober",
	// 	"november",
	// }

	// fmt.Println(month)
	// var (
	// 	slice = month[4:]
	// )

	// var slice1 = append(slice, "Desember")
	// fmt.Println(month)
	// fmt.Println(slice)
	// fmt.Println(slice1)

	// newSlice := make([]string, 2, 5)

	// newSlice[0] = "Senin"
	// newSlice[1] = "Selasa"

	// fmt.Println(newSlice)

	// copySlice := make([]string, 1, cap(newSlice))

	// copy(copySlice, newSlice)

	// fmt.Println(copySlice)

	// Tipe Data Map

	// person := map[string]string{
	// 	"name":    "dimas",
	// 	"address": "Gresik",
	// }

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])

	// var book map[string]string = make(map[string]string)

	// book["title"] = "Belajar GO"
	// book["author"] = "Dimas"
	// book["Ups"] = "Salah"

	// fmt.Println(book)

	// delete(book, "Ups")

	// fmt.Println(book)

	nama := "Dimas"

	// If Expression

	// if nama == "Dimas" {
	// 	fmt.Println("Hallo Dimas")
	// } else if nama == "Dliyaur" {
	// 	fmt.Println("Hallo Dliyaur !")
	// } else {
	// 	fmt.Println("Hallo")
	// }

	// // If Short Statement
	// if length := len(nama); length > 5 {
	// 	fmt.Printf("Nama Kepanjangan")
	// } else {
	// 	fmt.Println("Nama Pas 😭")
	// }

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
