package main

import "fmt"

func printType(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}

func main() {
	var x interface{}

	x = 42
	fmt.Println(x)

	x = "Hello"
	fmt.Println(x)

	x = true
	fmt.Println(x)

	x = []int{1, 2, 3}
	fmt.Println(x)

	printType(10)
	printType("Hello World!")
	printType(true)
}
