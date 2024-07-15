package main

import (
	"time"

	"github.com/GavinDevelops/pokecache"
	pokeapi "github.com/GavinDevelops/pokedexcli/commands"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	config := Config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
		cache:    &cache,
		pokedex:  make(map[string]pokeapi.Pokemon),
	}
	startRepl(&config)
}
