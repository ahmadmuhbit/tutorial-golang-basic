package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Tidak bisa melakukan pembagian dengan nol")
	} else {
		result := x / y
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("Contoh Error")
	fmt.Println(contohError.Error())

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
