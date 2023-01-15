package main

import "fmt"

func main() {
	// If statement
	number := 10

	if number == 10 {
		fmt.Println("Number adalah 10")
	}

	// If else statement
	angka := 5

	if angka%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	// Else if statement
	nilai := 85
	absen := 90

	if nilai > 90 && absen > 95 {
		fmt.Println("Sangat Baik")
	} else if nilai > 80 && absen > 85 {
		fmt.Println("Baik")
	} else if nilai > 60 && absen > 70 {
		fmt.Println("Cukup")
	} else {
		fmt.Println("Buruk")
	}

	// Nested if
	num := 100

	if num >= 80 {
		fmt.Println("Num lebih dari 80")
		if num > 90 {
			fmt.Println("Num juga lebih dari 90")
		}
	} else {
		fmt.Println("Num kurang dari 80")
	}
}
