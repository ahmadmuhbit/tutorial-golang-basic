package main

import (
	"fmt"
)

func main() {
	var data interface{} = "Hello World"

	dataString := data.(string)
	fmt.Println(dataString)

	// Handle panic type assertion
	var myData interface{} = true
	num, ok := myData.(int)
	if !ok {
		recoverMyData := recover()
		fmt.Println("Tipe data myData bukan int", recoverMyData)
	} else {
		fmt.Println("Tipe data myData adalah int", num)
	}

	// Type assertion switch
	switch myData.(type) {
	case int:
		fmt.Println("myData adalah integer")
	case string:
		fmt.Println("myData adalah string")
	default:
		fmt.Println("myData adalah tipe data lainnya")
	}
}
