package main

import "fmt"

func main() {
	var a float32 = 125.55
	var b float32 = 3.4e+38

	var c float64 = 225.55
	var d float64 = 1.7e+308

	fmt.Printf("Value: %v, Tipe Data: %T\n", a, a)
	fmt.Printf("Value: %v, Tipe Data: %T\n", b, b)

	fmt.Printf("Value: %v, Tipe Data: %T\n", c, c)
	fmt.Printf("Value: %v, Tipe Data: %T\n", d, d)
}
