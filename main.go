package main

import (
	"bigzano/Go-kedex/internal/pokeapi"
	"bigzano/Go-kedex/internal/pokecache"
	"time"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	pokeClient := pokeapi.NewClient(5*time.Second, cache)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
