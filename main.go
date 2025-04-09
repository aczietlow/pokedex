package main

import (
	"time"

	"github.com/aczietlow/pokedex/pkg/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(30 * time.Second)
	conf := &config{
		pokedexPokemon: map[string]pokeapi.Pokemon{},
		apiClient:      pokeClient,
	}
	startRepl(conf)
}
