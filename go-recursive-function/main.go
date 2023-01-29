package main

import "fmt"

// Contoh program recursive function
func countRecursive(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return countRecursive(x + 1)
}

// Factorial dengan for loop
func factorialLoop(y int) int {
	result := 1
	for i := y; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

// Factorial dengan recursive function
func factorialRecursive(z int) int {
	if z == 0 {
		return 1
	}
	return z * factorialRecursive(z-1)
}

func main() {
	countRecursive(1)

	loop := factorialLoop(10)
	fmt.Println(loop)

	fmt.Println(factorialRecursive(10))
}
