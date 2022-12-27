package main

import "fmt"

func main() {
	var a uint8 = 3
	var b, c, d, e, f uint8 = 5, 10, 15, 20, 25

	b += 5 // b = b + 5
	c -= 5 // c = c - 5
	d *= 5 // d = d * 5
	e /= 5 // e = e / 5
	f %= 5 // f = f % 5

	fmt.Println("nilai a :", a)
	fmt.Println("nilai b :", b)
	fmt.Println("nilai c :", c)
	fmt.Println("nilai c :", d)
	fmt.Println("nilai e :", e)
	fmt.Println("nilai f :", f)
}
