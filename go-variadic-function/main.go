package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums)

	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	sumAll := sum(1, 2, 3, 4, 5)
	fmt.Println(sumAll)

	// Variadic function slice parameter
	nums := []int{10, 20, 30, 40, 50}
	total := sum(nums...)
	fmt.Println(total)
}
