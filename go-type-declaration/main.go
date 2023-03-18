package main

import "fmt"

func main() {
	// Tidak menggunakan type declarations
	var txt1 string = "Hello"
	var num1 int = 20
	var isMarried1 bool = true

	fmt.Println(txt1)
	fmt.Println(num1)
	fmt.Println(isMarried1)

	// Menggunakan type declarations
	type MyString string
	type MyInt int
	type Married bool

	var txt MyString = "Hello World!"
	var num MyInt = 22
	var isMarried Married = false

	fmt.Println(txt)
	fmt.Println(num)
	fmt.Println(isMarried)
}
