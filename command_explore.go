package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		fmt.Print("args required, use explore <area-name>")
		return nil
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	resp, err := cfg.pokeapiClient.GetLocationData(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, enc := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
