package main

import "fmt"

// Declaration function parameters
type Blacklist func(string) bool

// ===================== Function returning multiple values =============

// func sayhello() (string, string, string) {
// 	return "Dimas", "Dliyaur", "Rahman"
// }

// ====================== Function as Parameter ========================

// func sayHello(name string, filter func(string) string) string {
// 	filterName := filter(name)

// 	return "hallo " + filterName
// }

// func filterSpam(name string) string {
// 	if name == "Anjing" {
// 		name := "..."
// 		return name
// 	} else {
// 		return name
// 	}
// }

// ==================== Function return named =========================
// func getAllName() (firstName, middleName, lastName string) { // Jika tipe data sama semua ,  deklarasi tidak usah tiap parameter
// 	firstName = "Dimas"
// 	middleName = "Dliyaur"
// 	lastName = "Rahman"

// 	return
// }

// ========================== Property for Function Anonymous =========================

// func registrationUser(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("Kamu di blacklist " + name)
// 	} else {
// 		fmt.Println("Selamat Datang " + name)
// 	}
// }

// Recursive function

// func recursiveFactorial(value int) int {
// 	if value == 1 {
// 		return 1
// 	} else {
// 		return value * recursiveFactorial(value-1)
// 	}
// }

// Defer , Panic , And Recover

func logging() {
	massage := recover()
	if massage != nil {
		fmt.Println("Error massage dengan", massage)
	}
	fmt.Println("aplikasi login")
}

func runApp(error bool) {
	defer logging()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi berhasil")
}

func main() {
	// ========================= Output Multiple values "_" untuk tidak memakai return ======================

	// _, b, c := sayhello()

	// fmt.Println(b)
	// fmt.Println(c)

	// ========================= Output Function as parameter =========================

	// fmt.Println(sayHello("Anjing", filterSpam))

	// =========================== Ouput Function return named ============================

	// firstName, middleName, lastName := getAllName()

	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)

	// ================== Output Anonimous Function ====================

	// blacklist := func(name string) bool {
	// 	return name == "admin" || name == "root"
	// }

	// registrationUser("admin", blacklist)
	// registrationUser("root", func(name string) bool {
	// 	return name == "admin" || name == "root"
	// })
	// registrationUser("dimas", blacklist)

	// Output Recursive Function

	// recursive := recursiveFactorial(10)
	// fmt.Println(recursive)

	// Output defer ,panic and recovery

	runApp(true)
}
