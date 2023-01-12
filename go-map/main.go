package main

import "fmt"

func main() {
	// Membuat map dengan keyword var dan tanda :=
	var person = map[string]string{
		"name":    "Ahmad Muhbit",
		"job":     "programmer",
		"address": "Jakarta",
	}

	stock := map[string]int{
		"book":      10,
		"handphone": 5,
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["job"])

	fmt.Println(stock)
	fmt.Println(stock["book"])
	fmt.Println(stock["handphone"])

	// Membuat map dengan function make()
	var laptop = make(map[string]string)
	laptop["merk"] = "Apple"
	laptop["model"] = "Macbook Pro"

	computer := make(map[string]int)
	computer["stock"] = 100
	computer["harga"] = 8000000

	fmt.Println(laptop)
	fmt.Println(laptop["merk"])
	fmt.Println(laptop["model"])

	fmt.Println(computer)
	fmt.Println(computer["stock"])
	fmt.Println(computer["harga"])

	// Mendapatkan panjang map
	fmt.Println(len(person))
	fmt.Println(len(stock))

	// Membuat map kosong
	var mapKosong1 = make(map[string]string)
	var mapKosong2 map[string]string

	fmt.Println(mapKosong1)
	fmt.Println(mapKosong2)

	// Menambah dan mengubah elemen map
	laptop["tahun"] = "2020"
	laptop["model"] = "Macbook Air"

	fmt.Println(laptop)
	fmt.Println(laptop["tahun"])
	fmt.Println(laptop["model"])

	// Menghapus elemen map
	delete(laptop, "merk")
	fmt.Println(laptop)

	// Map bersifat pass by reference
	person2 := person

	fmt.Println(person)
	fmt.Println(person2)

	person2["address"] = "Tangerang"

	fmt.Println(person)
	fmt.Println(person2)
}
