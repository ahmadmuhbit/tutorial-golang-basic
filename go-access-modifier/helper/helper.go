package helper

import "fmt"

var Txt1 string = "Hello World"
var txt2 string = "Contoh Unexported"

func Greet(name string) {
	fmt.Println("Hi " + name)
}

func greet2(name string) {
	fmt.Println("Hello " + name)
}
