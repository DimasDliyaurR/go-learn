package main

import "fmt"

func random() interface{} {
	return "Oke"
}

func main() {
	result := random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "Is string")
	case int:
		fmt.Println("Value", value, "Is int")
	default:
		fmt.Println("Unknow Type")
	}

}
