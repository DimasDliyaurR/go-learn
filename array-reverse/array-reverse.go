package main

import "fmt"

func main() {
	var list1 = [...]int{7, 6, 5, 4, 3, 2, 1}
	/*
		Ascending Program untuk mengurutkan array

		[7,6,5,4,3,2,1]
		 0,1,2,3,4,5,6

		Jika 7 bertemu yang lebih kecil maka index kecil akan disimpan terlebih dahulu seterusnya sampai no 7 berurut

		[6,5,4,3,2,1,7]
		[5,4,3,2,1,6,7]
		[4,3,2,1,5,6,7]
		[3,2,1,4,5,6,7]
		[2,1,3,4,5,6,7]
		[1,2,3,4,5,6,7]
		[1,2,3,4,5,6,7]
	*/

	for i := 0; i < len(list1)-1; i++ {

		fmt.Println(list1[i], list1[i+1])
		flex := 0
		newIndex := i + 1
		temp := list1[i]

		if list1[i] > list1[newIndex] {
			flex = list1[newIndex]
		}
		fmt.Println(list1, list1[newIndex], flex)

		list1[newIndex] = temp
		list1[i] = flex

	}
}
