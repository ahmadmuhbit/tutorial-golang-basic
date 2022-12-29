package main

import "fmt"

func main() {
	var x uint8 = 10
	var y uint8 = 5

	var txt1 = "Hello"
	var txt2 = "Hello"

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

	var result bool = txt1 == txt2
	fmt.Println(result)
}
