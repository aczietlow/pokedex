package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aczietlow/pokedex/pkg/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var registry map[string]cliCommand

type config struct {
	apiClient pokeapi.Client
	mapPager  int
}

// var mapPager int

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	registry = registerCommands()
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		cliInput := cleanInput(scanner.Text())

		commandName := cliInput[0]
		if command, ok := registry[commandName]; ok {
			err := command.callback(conf)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
		fmt.Print("Pokedex > ")
	}
}

func registerCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": cliCommand{
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": cliCommand{
			name:        "map",
			description: "List map locations, 20 at a time",
			callback:    commandMap,
		},
		"mapb": cliCommand{
			name:        "mapb",
			description: "Fetch the previous 20 locations",
			callback:    commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	return strings.Fields(strings.ToLower(text))
}
