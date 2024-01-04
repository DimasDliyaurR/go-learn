package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// Function in Pointer

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}

	/* Pointer Pass by Value
	Object akan dibuatkan lagi tetapi dengan nilai sama
	Jika Object dimodifikasi tidak akan mengubah Object asal
	*/
	var address2 Address = address1

	/* Pointer pass by refrence
	Object akan direfence yang artinya object yang lama akan diinisialisai dengan variabel baru
	Variable ini bisa memodifikasi object asal
	*/

	var address3 *Address = &address1

	address2.City = "Bandung"
	address3.City = "Lembang"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	/* The Output
	{Lembang Jawa Barat Indonesia}
	{Bandung Jawa Barat Indonesia}
	&{Lembang Jawa Barat Indonesia}
	*/

	var alamat = Address{
		City:     "Gresik",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)

}
