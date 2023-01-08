package main

import "fmt"

func main() {
	// Membuat slice langsung
	var iniSlice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(len(iniSlice1))
	fmt.Println(cap(iniSlice1))
	fmt.Println(iniSlice1)

	iniSlice2 := []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jum'at",
		"Sabtu",
		"Minggu",
	}
	fmt.Println(len(iniSlice2))
	fmt.Println(cap(iniSlice2))
	fmt.Println(iniSlice2)

	// Membuat slice dari array
	iniArray := [8]int{10, 11, 12, 13, 14, 15, 16, 17}
	iniSlice3 := iniArray[3:5]

	fmt.Printf("iniSlice3 = %v\n", iniSlice3)
	fmt.Printf("length = %d\n", len(iniSlice3))
	fmt.Printf("capacity = %d\n", cap(iniSlice3))

	// Contoh pass by reference
	iniSlice3[0] = 20
	fmt.Printf("iniSlice3 = %v\n", iniSlice3)
	fmt.Printf("iniArray = %v\n", iniArray)

	// Membuat slice dengan fungsi make()
	iniSlice4 := make([]int, 5, 5)
	iniSlice4[0] = 10
	iniSlice4[1] = 11

	fmt.Printf("iniSlice4 = %v\n", iniSlice4)
	fmt.Printf("length = %d\n", len(iniSlice4))
	fmt.Printf("capacity = %d\n", cap(iniSlice4))
}
