package main

import (
	"fmt"
)

var pokedexField = []string{"Name", "Height", "Weight", "Stats", "Types"}

func commandInspect(conf *config, args ...string) error {
	pokemonName := args[0]

	if pokemon, ok := conf.pokedexPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats: \n")
		for _, v := range pokemon.Stats {
			fmt.Printf("\t- %s: %d\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, v := range pokemon.Types {
			fmt.Printf("\t- %s\n", v.Type.Name)
		}
		return nil
	}

	fmt.Printf("No Pokemon with the name %s is registered. Trying exploring and CATCHING one.\n", pokemonName)
	return nil
}
