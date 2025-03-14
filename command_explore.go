package main

import "fmt"

func commandExplore(conf *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("Please provide a name or id of an area")
	}

	// TODO: not processing 404 errors correctly
	location, err := conf.apiClient.FetchLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")

	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
