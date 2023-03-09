package main

import "fmt"

type Greeting interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return p.Name
}

func SayHello(g Greeting) {
	fmt.Println("Hello my name is", g.Greet())
}

func main() {
	person := Person{Name: "Ubit"}

	SayHello(person)
}
