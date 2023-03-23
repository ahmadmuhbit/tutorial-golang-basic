package main

import "fmt"

func main() {
	var a string = "Ahmad"
	var b *string = &a // b menunjuk pada alamat memori dari variabel a

	*b = "Muhbit"

	fmt.Println(a)
	fmt.Println(*b)
}
