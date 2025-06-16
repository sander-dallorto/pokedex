package main

import (
	"fmt"
	"os"
)

type Pokedex struct {
}

func commandExit(c *config, userInput string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
