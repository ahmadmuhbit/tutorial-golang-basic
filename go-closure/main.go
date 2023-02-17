package main

import "fmt"

func main() {
	name := "Ahmad Muhbit"

	displayName := func() {
		fmt.Println(name)
	}

	displayName()
}
