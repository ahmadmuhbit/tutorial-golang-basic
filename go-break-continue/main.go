package main

import "fmt"

func main() {
	// Break statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Selesai")

	// Continue statement
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}
	fmt.Println("Selesai")
}
