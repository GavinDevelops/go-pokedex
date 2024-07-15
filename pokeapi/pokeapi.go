package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

func GetLocations(address string) (Locations, error) {
	resp, getErr := http.Get(address)
	if getErr != nil {
		return Locations{}, errors.New("Error getting location")
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return Locations{}, errors.New("Error reading response body")
	}
	location := Locations{}
	unmarshalErr := json.Unmarshal(body, &location)
	if unmarshalErr != nil {
		return Locations{}, errors.New("Error unmarshaling body")
	}
	return location, nil
}
