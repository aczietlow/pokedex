package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("Please provide a name or id of a pokemon")
	}

	pokemon, err := conf.apiClient.FetchPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if attemptCatch(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.pokedexPokemon[pokemon.Name] = pokemon
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}

func attemptCatch(baseExperience int) bool {
	catchRate := rand.Intn(255)

	if (catchRate*100)/baseExperience > 60 {
		return true
	}

	return false
}
