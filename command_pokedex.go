package main

import "fmt"

func commandPokedex(cfg *config, userInput string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("your Pokedex is empty, catch some Pokemon first")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
