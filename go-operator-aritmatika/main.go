package main

import "fmt"

func main() {
	fmt.Println(10 + 4)
	fmt.Println(10 - 4)
	fmt.Println(10 * 4)
	fmt.Println(10 / 2)
	fmt.Println(10 % 3)

	var a, b, c = 10, 4, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / c)
	fmt.Println(a % b)

	var d = a + b + c
	fmt.Println(d)

	i := 5
	i++ // i = i + 1

	j := 5
	j-- // j = j - 1

	fmt.Println(i)
	fmt.Println(j)
}
