package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationURL = locations.Next
	cfg.PreviousLocationURL = locations.Previous

	return nil
}
