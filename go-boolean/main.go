package main

import "fmt"

func main() {
	var a bool = true // Deklarasi dengan inisialisasi tipe data
	var b = true      // Deklarasi tanpa inisialisasi tipe data
	var c bool        // Deklarasi tanpa value awal
	d := true         // Deklarasi dengan tanda :=

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%v, %T\n", b, b)
}
