package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func printCatch(pokemonInfo PokemonInfo, userInput string) bool {
	fmt.Printf("Throwing a Pokeball at %s...\n", userInput)

	maxBaseExp := 700.0
	catchChance := 1.0 - (float64(pokemonInfo.BaseExperience) / maxBaseExp)

	if rand.Float64() < catchChance {
		fmt.Printf("%s was caught!\n", pokemonInfo.Name)
		return true
	} else {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
		return false
	}
}

func commandCatch(c *config, userInput string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", userInput)

	data, found := cache.Get(url)
	if found {
		var pokemonInfo PokemonInfo
		if err := json.Unmarshal(data, &pokemonInfo); err != nil {
			return fmt.Errorf("error unmarshalling cached JSON: %v", err)
		}

		if printCatch(pokemonInfo, userInput) {
			c.caughtPokemon[userInput] = pokemonInfo
		}

		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	cache.Add(url, body)

	if res.StatusCode > 299 {
		if res.StatusCode == 404 {
			return fmt.Errorf("pokemon '%s' not found", userInput)
		}
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	if err != nil {
		return err
	}

	var pokemonInfo PokemonInfo
	if err := json.Unmarshal(body, &pokemonInfo); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	if printCatch(pokemonInfo, userInput) {
		c.caughtPokemon[userInput] = pokemonInfo
	}

	return nil
}
