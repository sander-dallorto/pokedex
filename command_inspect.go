package main

import (
	"fmt"
)

func commandInspect(cfg *config, userInput string) error {
	if userInput == "" {
		return fmt.Errorf("please specify a Pokemon to inspect")
	}

	pokemonInfo, exists := cfg.caughtPokemon[userInput]
	if !exists {
		return fmt.Errorf("no Pokemon named '%s' found in your collection", userInput)
	}

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)

	fmt.Println("Stats: ")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types: ")
	for _, t := range pokemonInfo.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
