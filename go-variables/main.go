package main

import "fmt"

var x int = 1

// y := 2 // Error, tidak bisa digunakan di luar fungsi

func main() {
	var myName string
	myName = "Ahmad Muhbit"
	var nickName = "Ubit"
	myAge := 28

	var a string
	var b int
	var c bool

	var d, e, f, g = 1, 2, 3, "Hello World"

	var (
		firsName string = "Ahmad"
		lastName        = "Muhbit"
		height   int    = 172
	)

	fmt.Println(myName)
	fmt.Println(nickName)
	fmt.Println(myAge)

	fmt.Println(x)
	// fmt.Println(y)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(firsName)
	fmt.Println(lastName)
	fmt.Println(height)
}
