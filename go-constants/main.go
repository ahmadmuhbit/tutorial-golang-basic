package main

import "fmt"

const LENGTH int = 10
const WIDTH = 5

const (
	X int = 10
	Y     = 3.14
	Z     = "Hello World!"
)

func main() {
	// const PI = 3.14
	// PI = 20 // Tidak dapat diubah nilainya

	fmt.Println(LENGTH)
	fmt.Println(WIDTH)

	// fmt.Println(PI)

	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
}
