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
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays location areas in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon in the current location area",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inpsect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display the Pokedex",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	loweredWord := strings.ToLower(text)
	words := strings.Fields(loweredWord)

	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{Next: "https://pokeapi.co/api/v2/location-area/", Previous: nil, caughtPokemon: make(map[string]PokemonInfo)}
	var userInput string

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if len(words) > 1 {
			userInput = words[1]
		} else {
			userInput = ""
		}

		command, ok := getCommands()[commandName]

		if ok {
			if err := command.callback(&c, userInput); err != nil {
				fmt.Printf("Error executing command '%s': %v\n", commandName, err)
			}
		} else {
			fmt.Printf("Unknown command: '%s'. Type 'help' for a list of commands.\n", commandName)
		}
	}
}
