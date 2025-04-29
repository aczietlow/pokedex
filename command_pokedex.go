package main

import "fmt"

func commandPokedex(conf *config, arg ...string) error {

	if len(conf.pokedexPokemon) < 1 {
		fmt.Printf("Pokedex contains no entries, try getting out there and exploring more.\n")
		return nil
	}
	fmt.Printf("Your Pokedex:\n")
	for _, v := range conf.pokedexPokemon {
		fmt.Printf("  - %s\n", v.Name)
	}

	return nil
}
