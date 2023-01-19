package main

import "fmt"

func main() {
	// For loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Nested loop
	a := [2]string{"buah", "jus"}
	b := [3]string{"jambu", "mangga", "alpukat"}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Println(a[i], b[j])
		}
	}

	// For range array
	cars := [3]string{"Toyota", "Honda", "Nissan"}
	for _, car := range cars {
		fmt.Println(car)
	}

	// For range map
	var person = map[string]string{
		"name": "Ahmad Muhbit",
		"job":  "Programmer",
	}
	for key, value := range person {
		fmt.Println(key, value)
	}
}
