package main

import "fmt"

func main() {
	// Mengakses elemen slice
	numbers := []int{1, 2, 3}

	fmt.Println(numbers[0])
	fmt.Println(numbers[2])

	// Mengubah elemen slice
	numbers2 := []int{4, 5, 6}
	numbers2[2] = 7

	fmt.Println(numbers2)

	// Menambah elemen slice
	iniSlice := []int{10, 20, 30, 40, 50}
	fmt.Printf("iniSlice = %v\n", iniSlice)
	fmt.Printf("length = %d\n", len(iniSlice))
	fmt.Printf("capacity = %d\n", cap(iniSlice))

	iniSlice = append(iniSlice, 60, 70)
	fmt.Printf("iniSlice = %v\n", iniSlice)
	fmt.Printf("length = %d\n", len(iniSlice))
	fmt.Printf("capacity = %d\n", cap(iniSlice))

	// Mengubah panjang slice
	iniArray := [6]int{9, 10, 11, 12, 13, 14} // Array
	iniSliceArray := iniArray[1:5]            // Slice array
	fmt.Println("iniSLiceArray =", iniSliceArray)
	fmt.Println("length =", len(iniSliceArray))
	fmt.Println("capacity =", cap(iniSliceArray))

	iniSliceArray = iniArray[1:3] // Mengubah length dengan re-slicing sebuah array
	fmt.Println("iniSLiceArray =", iniSliceArray)
	fmt.Println("length =", len(iniSliceArray))
	fmt.Println("capacity =", cap(iniSliceArray))

	iniSliceArray = append(iniSliceArray, 20, 21, 22, 23) // Mengubah panjang slice dengan menambahkan elemen
	fmt.Println("iniSLiceArray =", iniSliceArray)
	fmt.Println("length =", len(iniSliceArray))
	fmt.Println("capacity =", cap(iniSliceArray))

	// Copy slice
	angka := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Angka =", angka)
	fmt.Println("Length =", len(angka))
	fmt.Println("Capacity =", cap(angka))

	copyAngka := make([]int, len(angka), cap(angka))
	copy(copyAngka, angka)
	fmt.Println("Copy angka =", angka)
	fmt.Println("Length =", len(angka))
	fmt.Println("Capacity =", cap(angka))
}
