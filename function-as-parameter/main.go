package main

import (
	"fmt"
	"strings"
)

type Formatter func(string) string

func formatCase(txt string, formatter Formatter) {
	formatted := formatter(txt)
	fmt.Println(formatted)
}

func upperCase(txt string) string {
	return strings.ToUpper(txt)
}

func lowerCase(txt string) string {
	return strings.ToLower(txt)
}

func main() {
	formatCase("hello world", upperCase)
	formatCase("HELLO WORLD", lowerCase)

	toUpper := upperCase // Function as value
	formatCase("hello world", toUpper)
}
