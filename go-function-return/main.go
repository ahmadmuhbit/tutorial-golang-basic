package main

import "fmt"

// Function return value
func getSum(a int, b int) int {
	return a + b
}

// Named return value
func addNumbers(x int, y int) (result int) {
	result = x + y
	return
}

// Multiple return value
func getCombine(i int, j string) (sum int, txt string) {
	sum = i + i
	txt = j + " Muhbit"
	return
}

func main() {
	fmt.Println(getSum(10, 20))

	total := addNumbers(20, 30)
	fmt.Println(total)

	a, _ := getCombine(5, "Ahmad")
	fmt.Println(a)
}
