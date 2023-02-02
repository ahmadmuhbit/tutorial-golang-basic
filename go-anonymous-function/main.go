package main

import "fmt"

func main() {
	// Anonymous function
	var getHello = func() {
		fmt.Println("Hello anonymous function")
	}

	getHello()

	//Anonymous function dengan parameter
	var getPeople = func(myName string, myAge int) {
		fmt.Println("Nama saya", myName, "umur saya", myAge, "tahun")
	}

	getPeople("Ahmad Muhbit", 30)

	// Anonymous function return value
	var addNumbers = func(x, y int) int {
		result := x + y
		return result
	}

	total := addNumbers(20, 30)
	fmt.Println(total)
}
