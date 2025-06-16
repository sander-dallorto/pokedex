package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func printEncounters(encounters Encounters, userInput string) {
	fmt.Printf("Exploring area '%s'...:\n", userInput)
	fmt.Println("Found Pokemon:")

	for _, enc := range encounters.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
}

func commandExplore(c *config, userInput string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", userInput)

	data, found := cache.Get(url)
	if found {
		var encounters Encounters
		if err := json.Unmarshal(data, &encounters); err != nil {
			return fmt.Errorf("error unmarshalling cached JSON: %v", err)
		}

		printEncounters(encounters, userInput)

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
			return fmt.Errorf("location area '%s' not found", userInput)
		}
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	if err != nil {
		return err
	}

	var encounters Encounters
	if err := json.Unmarshal(body, &encounters); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	printEncounters(encounters, userInput)

	return nil
}
