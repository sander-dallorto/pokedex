package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(c *config, userInput string) error {
	data, found := cache.Get(c.Next)
	if found {
		var locations Locations
		if err := json.Unmarshal(data, &locations); err != nil {
			return fmt.Errorf("error unmarshalling cached JSON: %v", err)
		}

		c.Next = locations.Next
		c.Previous = locations.Previous

		for _, location := range locations.Results {
			fmt.Println(location.Name)
		}

		return nil
	}

	res, err := http.Get(c.Next)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	cache.Add(c.Next, body)

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return err
	}

	var locations Locations
	if err := json.Unmarshal(body, &locations); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config, exploreArea string) error {
	if c.Previous == nil {
		fmt.Println("No previous URL found.")
		return nil
	}

	data, found := cache.Get(*c.Previous)
	if found {
		var locations Locations
		if err := json.Unmarshal(data, &locations); err != nil {
			return fmt.Errorf("error unmarshalling cached JSON: %v", err)
		}
		c.Next = locations.Next
		c.Previous = locations.Previous
		for _, location := range locations.Results {
			fmt.Println(location.Name)
		}
		return nil
	}

	res, err := http.Get(*c.Previous)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	cache.Add(*c.Previous, body)

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return err
	}

	var locations Locations
	if err := json.Unmarshal(body, &locations); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
