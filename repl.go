package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var registry map[string]cliCommand

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	registry = registerCommands()
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		cliInput := cleanInput(scanner.Text())

		commandName := cliInput[0]
		if command, ok := registry[commandName]; ok {
			command.callback()
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
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	return strings.Fields(strings.ToLower(text))
}
