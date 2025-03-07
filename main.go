package main

import "github.com/aczietlow/pokedex/pkg/pokeapi"

func main() {
	pokeClient := pokeapi.NewClient()
	conf := &config{
		apiClient: pokeClient,
	}
	startRepl(conf)
}
