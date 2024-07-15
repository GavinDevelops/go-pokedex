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

type AreaInfo struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetAreaInfo(area string, cache pokecache.Cache) (AreaInfo, error) {
	address := fmt.Sprint("https://pokeapi.co/api/v2/location-area/", area)
	if body, exists := cache.Get(address); exists {
		fmt.Println("--- Fetching from Cache ---")
		return getAreaInfoFromBody(body)
	}
	return getAreaInfoFromApi(address, cache)
}

func getAreaInfoFromApi(address string, cache pokecache.Cache) (AreaInfo, error) {
	resp, getErr := http.Get(address)
	if getErr != nil {
		return AreaInfo{}, errors.New("Error getting area info")
	}
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return AreaInfo{}, errors.New("Error reading area body")
	}
	cache.Add(address, body)
	fmt.Println("--- Fetching from API ---")
	time.Sleep(1 * time.Second)
	return getAreaInfoFromBody(body)
}

func getAreaInfoFromBody(body []byte) (AreaInfo, error) {
	info := AreaInfo{}
	unmarshalErr := json.Unmarshal(body, &info)
	if unmarshalErr != nil {
		return AreaInfo{}, errors.New("Error unmarshaling area body")
	}
	return info, nil
}
