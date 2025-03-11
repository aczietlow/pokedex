package main

import (
	"time"

	"github.com/aczietlow/pokedex/pkg/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		apiClient: pokeClient,
	}
	startRepl(conf)
}
