package main

import "fmt"

func commandMapBack(cfg *config, args ...string) error {
	if cfg.PreviousLocationURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.PreviousLocationURL)
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
