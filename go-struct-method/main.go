package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func (user User) GetName() string {
	return user.Name
}

func (user User) GetEmail() string {
	return user.Email
}

func main() {
	var user User
	user.Name = "Ubit"
	user.Email = "ahmadmuhbit@gmail.com"
	user.Age = 30

	fmt.Println(user.GetName())
	fmt.Println(user.GetEmail())
}
