package main

import "fmt"

func main() {
	var a = 15

	fmt.Println(a > 10 && a < 5)
	fmt.Println(a > 10 || a < 5)
	fmt.Println(!(a > 10 || a < 5))

	// var nilaiUjian = 85
	// var nilaiAbsen = 75

	// var lulusUjian = nilaiUjian > 80
	// var lulusAbsen = nilaiAbsen > 80
	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsen)

	// var lulus = lulusUjian && lulusAbsen

	// fmt.Println(nilaiUjian > 80 && nilaiAbsen > 80)
	// fmt.Println(lulus)
}
