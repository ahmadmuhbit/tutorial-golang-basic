package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	total := sum

	result := total(1, 2)
	fmt.Println(result)
}
