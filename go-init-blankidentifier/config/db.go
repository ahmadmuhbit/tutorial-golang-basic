package config

import "fmt"

var createConnetction string

func init() {
	fmt.Println("Init dipanggil")
	createConnetction = "PostgreSQL"
}

func GetPostgres() string {
	return createConnetction
}
