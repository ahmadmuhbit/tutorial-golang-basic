package main

import (
	"fmt"
	_ "go-init-blankidentifier/config"
)

func init() {
	fmt.Println("Ini adalah fungsi init")
}

func main() {
	fmt.Println("Ini adalah fungsi main")
	//fmt.Println(config.GetPostgres())
}
