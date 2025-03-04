package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	return strings.Split(strings.ToLower(text), " ")
}
