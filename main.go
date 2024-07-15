package main

import (
	"time"

	"github.com/GavinDevelops/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	config := Config{
		next:     "https://pokeapi.co/api/v2/location-area/",
		previous: "",
		cache:    &cache,
	}
	startRepl(&config)
}
