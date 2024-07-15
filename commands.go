package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/GavinDevelops/pokecache"
	pokeapi "github.com/GavinDevelops/pokedexcli/commands"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
	config      *Config
}

type Config struct {
	next     string
	previous string
	cache    *pokecache.Cache
}

func commandHelp(config *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, value := range getCommands(nil) {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandExit(config *Config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *Config) error {
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

func commandMapB(config *Config) error {
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

func commandExplor(config *Config) error {
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
			name:        "map",
			description: "Get the next 20 locations",
			callback:    commandMapB,
			config:      config,
		},
	}
}
