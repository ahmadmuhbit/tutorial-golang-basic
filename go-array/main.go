package main

import "fmt"

func main() {
	var a1 = [3]int{100, 200, 300}
	a2 := [4]int{400, 500, 600, 700}

	var a3 = [...]int{10, 20, 30}
	a4 := [...]int{40, 50, 60, 70}

	var languages = [3]string{"Go", "Java", "C"}

	prices := [3]int{1000, 2000, 3000}
	prices[2] = 5000

	a5 := [4]int{}           // Tidak diinisialisasi
	a6 := [4]int{1, 2}       // Diinisialisasi sebagian
	a7 := [4]int{1, 2, 3, 4} // Diinisialisasi semuanya

	fmt.Println(a1)
	fmt.Println(a2)

	fmt.Println(a3)
	fmt.Println(a4)

	fmt.Println(languages)

	fmt.Println(prices[0])
	fmt.Println(prices[2])
	fmt.Println(prices)

	fmt.Println(a5)
	fmt.Println(a6)
	fmt.Println(a7)

	fmt.Println(len(a2))
	fmt.Println(len(languages))
}
