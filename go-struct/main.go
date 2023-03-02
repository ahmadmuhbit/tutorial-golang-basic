package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	var user User

	user.Name = "Ubit"
	user.Email = "ahmadmuhbit@gmail.com"
	user.Age = 30

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Email)
	fmt.Println(user.Age)

	// Stuct literals, menyebutkan nama field
	var user1 = User{
		Name:  "User 1",
		Email: "user1@gmail.com",
		Age:   25,
	}

	fmt.Println(user1)
	fmt.Println(user1.Name)
	fmt.Println(user1.Email)
	fmt.Println(user1.Age)

	// Struct literals, tanpa menyebutkan nama field
	user2 := User{"User 2", "user2@gmail.com", 23}

	fmt.Println(user2)
	fmt.Println(user2.Name)
	fmt.Println(user2.Email)
	fmt.Println(user2.Age)
}
