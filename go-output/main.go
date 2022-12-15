package main

import "fmt"

func main() {
	var firstName, lastName string = "Ahmad", "Muhbit"
	var a, b = 100, 200

	var nickName, myAddress = "Ubit", "Jakarta"

	var x string = "Halo"
	var y int = 25
	var angka = 15.5
	var txt = "Hello World!"

	fmt.Print(firstName, " ", lastName, "\n")
	fmt.Print(a, b, "\n")

	fmt.Println(nickName, myAddress)

	fmt.Printf("value x adalah : %v dan tipe data x adalah : %T\n", x, x)
	fmt.Printf("value y adalah : %v dan tipe data y adalah : %T\n", y, y)
	fmt.Printf("%v%%\n", angka)
	fmt.Printf("%#v", txt)
}
