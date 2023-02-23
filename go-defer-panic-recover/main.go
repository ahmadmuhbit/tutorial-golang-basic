package main

import "fmt"

// Defer
func txtNumber() {
	fmt.Println("Tiga")
}

// Panic
func division(num1, num2 int) {
	defer handlePanic()

	if num2 == 0 {
		panic("Tidak dapat membagi angka dengan nol")
	} else {
		result := num1 / num2
		fmt.Println("Result:", result)
	}
}

// Recover
func handlePanic() {
	err := recover()

	if err != nil {
		fmt.Println("RECOVER", err)
	}
}

func main() {
	defer txtNumber()

	fmt.Println("Satu")
	fmt.Println("Dua")

	division(3, 0)
}
