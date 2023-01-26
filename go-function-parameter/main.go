package main

import "fmt"

func getName(nickName string) {
	fmt.Println("Halo", nickName)
}

// Function multiple parameter
func getPeople(myName string, myAge int) {
	fmt.Println("Nama saya", myName, "umur saya", myAge, "tahun")
}

func main() {
	getName("Ubit")

	getPeople("Ahmad Muhbit", 30)
}
