package main

import "fmt"

func main() {
	// Single case switch statement
	warnaFavorit := "ungu"

	switch warnaFavorit {
	case "biru":
		fmt.Println("Warna favorit kamu adalah biru")
	case "merah":
		fmt.Println("Warna favorit kamu adalah merah")
	case "hijau":
		fmt.Println("Warna favorit kamu adalah hijau")
	default:
		fmt.Println("Warna favorit kamu bukan biru, merah atau hijau")
	}

	// Multi case switch statement
	hari := "januari"

	switch hari {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		fmt.Println("Weekday")
	case "sabtu", "minggu":
		fmt.Println("Weekend")
	default:
		fmt.Println("Hari tidak valid")
	}
}
