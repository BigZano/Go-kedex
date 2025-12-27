package main

import (
	"fmt"
)

func commandList(cfg *config, args ...string) error {

	fmt.Println("Your Pokedex:")

	for _, p := range cfg.pokedex {
		fmt.Printf(" - %s\n", p.Name)
	}

	return nil
}
