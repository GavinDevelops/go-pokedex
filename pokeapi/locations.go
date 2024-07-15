package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/GavinDevelops/pokecache"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(address string, cache pokecache.Cache) (Locations, error) {
	if body, exists := cache.Get(address); exists {
		fmt.Println("--- Fetching from Cache ---")
		return getLocationsFromBody(body)
	}
	return getLocationsFromApi(address, cache)
}

func getLocationsFromApi(address string, cache pokecache.Cache) (Locations, error) {
	resp, getErr := http.Get(address)
	if getErr != nil {
		return Locations{}, errors.New("Error getting location")
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return Locations{}, errors.New("Error reading response body")
	}
	cache.Add(address, body)
	fmt.Println("--- Fetching from API ---")
	time.Sleep(1 * time.Second)
	return getLocationsFromBody(body)
}

func getLocationsFromBody(body []byte) (Locations, error) {
	location := Locations{}
	unmarshalErr := json.Unmarshal(body, &location)
	if unmarshalErr != nil {
		return Locations{}, errors.New("Error unmarshaling body")
	}
	return location, nil
}
