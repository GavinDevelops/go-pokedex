package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/GavinDevelops/pokecache"
)

func GetPokemon(name string, cache pokecache.Cache) (Pokemon, error) {
	address := fmt.Sprint("https://pokeapi.co/api/v2/pokemon/", name)
	if body, exists := cache.Get(address); exists {
		fmt.Println("Fetching Pokemon from Cache")
		fmt.Println("=================")
		return getPokemonFromBody(body)
	}
	fmt.Println("Fetching Pokemon from API")
	fmt.Println("=================")
	return getPokemonFromApi(address)
}

func getPokemonFromApi(address string) (Pokemon, error) {
	resp, getErr := http.Get(address)
	if getErr != nil {
		return Pokemon{}, errors.New("Error getting pokemon")
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return Pokemon{}, errors.New("Error reading pokemon response body")
	}
	return getPokemonFromBody(body)
}

func getPokemonFromBody(body []byte) (Pokemon, error) {
	pokemon := Pokemon{}
	unmarshalErr := json.Unmarshal(body, &pokemon)
	if unmarshalErr != nil {
		return Pokemon{}, errors.New("error unmarshaling pokemon body")
	}
	return pokemon, nil
}
