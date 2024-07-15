package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/GavinDevelops/pokecache"
	pokeapi "github.com/GavinDevelops/pokedexcli/commands"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, parameters string) error
	config      *Config
}

type Config struct {
	next     string
	previous string
	cache    *pokecache.Cache
	pokedex  map[string]pokeapi.Pokemon
}

func commandHelp(config *Config, args string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range getCommands(nil) {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandExit(config *Config, args string) error {
	os.Exit(0)
	return nil
}

func commandMap(config *Config, args string) error {
	loc, err := pokeapi.GetLocations(config.next, *config.cache)
	if err != nil {
		return err
	}
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
	config.previous = config.next
	config.next = loc.Next
	return nil
}

func commandMapB(config *Config, args string) error {
	if config.previous == "" {
		return errors.New("Error going back")
	}
	loc, err := pokeapi.GetLocations(config.previous, *config.cache)
	if err != nil {
		return err
	}
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
	config.next = config.previous
	if loc.Previous == nil {
		config.previous = ""
	} else {
		config.previous = strings.Clone(*loc.Previous)
	}
	return nil
}

func commandExplore(config *Config, area string) error {
	areaInfo, err := pokeapi.GetAreaInfo(area, *config.cache)
	if err != nil {
		return err
	}
	fmt.Println("=================")
	fmt.Printf("Exploring %v...\n", area)
	fmt.Println("=================")
	fmt.Println("Found Pokemon:")
	for _, pokemon := range areaInfo.PokemonEncounters {
		fmt.Println("  -", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *Config, name string) error {
	pokemon, err := pokeapi.GetPokemon(name, *config.cache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum < pokemon.BaseExperience/2 {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%v was caught!\n", pokemon.Name)
	config.pokedex[pokemon.Name] = pokemon
	return nil
}

func commandInspect(config *Config, name string) error {
	pokemon, exists := config.pokedex[name]
	if !exists {
		fmt.Println("You have not caught that Pokemon!")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(config *Config, arg string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range config.pokedex {
		fmt.Printf("  - %v\n", key)
	}
	return nil
}

func getCommands(config *Config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      config,
		},
		"exit": {

			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			config:      config,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 locations",
			callback:    commandMap,
			config:      config,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 locations",
			callback:    commandMapB,
			config:      config,
		},
		"explore": {
			name:        "explore",
			description: "Explore area usage: explore [area]",
			callback:    commandExplore,
			config:      config,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon usage: catch [pokemon name]",
			callback:    commandCatch,
			config:      config,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon usage: catch [pokemon name]",
			callback:    commandInspect,
			config:      config,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemon in your pokedex",
			callback:    commandPokedex,
			config:      config,
		},
	}
}
