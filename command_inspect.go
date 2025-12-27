package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		fmt.Print("args required. use inspect <pokemon-name>")
		return nil
	}

	pokemonName := args[0]

	fmt.Printf("Inspecting %s\n", pokemonName)

	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught this pokemon yet.")
		return nil
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Println(" -", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Println(" -", t.Type.Name)
		}
	}

	return nil
}
