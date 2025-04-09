package main

import "fmt"

func commandInspect(conf *config, args ...string) error {
	pokemonName := args[0]

	if pokemon, ok := conf.pokedexPokemon[pokemonName]; ok {
		fmt.Printf("pokemon: %s\n", pokemon.Name)
		return nil
	}

	fmt.Printf("No Pokemon with the name %s is registered\n", pokemonName)
	return nil
}
