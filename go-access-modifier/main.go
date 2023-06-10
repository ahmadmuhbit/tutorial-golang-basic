package main

import (
	"fmt"
	"go-access-modifier/helper"
)

func main() {
	fmt.Println(helper.Txt1)
	//fmt.Println(helper.txt2) // Error Unexported

	helper.Greet("Muhbit")
	//helper.greet2("Ahmad") // Error Unexported
}
