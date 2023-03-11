package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{"name": name}
}

func main() {
	// Nil pada slice, map, interface, pointer & channel
	var s []int
	fmt.Println("Slice:", s)
	fmt.Println("Slice == nil?", s == nil)

	var m map[string]int
	fmt.Println("Map:", m)
	fmt.Println("Map == nil?", m == nil)

	var i interface{}
	fmt.Println("Interface:", i)

	var ptr *int
	fmt.Println("Pointer:", ptr)

	var ch chan int
	fmt.Println("Channel:", ch)

	// Contoh dengan nil function NewMap
	user := NewMap("")

	if user == nil {
		fmt.Println("Kosong")
	} else {
		fmt.Println(user)
	}

	// Contoh tanpa nil function NewMap
	if user["name"] == "" {
		fmt.Println("Kosong")
	} else {
		fmt.Println(user)
	}
}
