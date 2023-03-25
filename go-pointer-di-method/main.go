package main

import "fmt"

type User struct {
	Email string
}

func (user *User) GetInfo() {
	user.Email = "Alamat Email: " + user.Email
}

func main() {
	var user User
	user.Email = "user@gmail.com"

	user.GetInfo()
	fmt.Println(user.Email)
}
