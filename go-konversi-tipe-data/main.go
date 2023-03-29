package main

import "fmt"

func main() {
	var a int8 = 10
	var b int16 = int16(a) // Konversi int8 ke int16

	var c uint32 = 257
	var d uint64 = uint64(c) // Konversi uint32 ke uint64

	var e float32 = 20.5
	var f float64 = float64(e) // Konversi float32 ke float64

	var g uint8 = uint8(c) // Hati-hati integer overflow

	var h float64 = float64(a) // Konversi int8 ke float64
	var i int = int(e)         // Konversi float ke int

	fmt.Printf("Tipe data a: %T, nilai a: %v\n", a, a)
	fmt.Printf("Tipe data b: %T, nilai b: %v\n", b, b)

	fmt.Printf("Tipe data c: %T, nilai c: %v\n", c, c)
	fmt.Printf("Tipe data d: %T, nilai d: %v\n", d, d)

	fmt.Printf("Tipe data e: %T, nilai e: %v\n", e, e)
	fmt.Printf("Tipe data f: %T, nilai f: %v\n", f, f)

	fmt.Println(c)
	fmt.Println(g)

	fmt.Printf("Tipe data h: %T, nilai h: %v\n", h, h)
	fmt.Printf("Tipe data i: %T, nilai i: %v\n", i, i)

	// Konversi tipe data byte ke string
	var name string = "Ahmad Muhbit"
	var charA byte = name[0]
	var charAString string = string(charA)

	fmt.Println(name)
	fmt.Println(charA)
	fmt.Println(charAString)
}
