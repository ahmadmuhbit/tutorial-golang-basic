package main

import "fmt"

type Person struct {
	Name string
	Job  string
}

func GetInfo(person *Person) {
	person.Job = "Programmer"
}

func main() {
	var person = Person{
		Name: "Ahmad Muhbit",
		Job:  "",
	}

	GetInfo(&person)
	fmt.Println(person)
}
