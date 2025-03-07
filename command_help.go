package main

import "fmt"

func commandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range registry {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()
	return nil
}
