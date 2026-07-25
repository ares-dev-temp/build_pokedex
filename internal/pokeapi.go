package internal

//should be in internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type PokeData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`

	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
}

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon        `json:"pokemon"`
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

var pokeCache = Cache{}

func NewClient(interval time.Duration) {
	pokeCache = *NewCache(interval)
}

// func GetPokeData(urlPath string, c *config) error {
func GetPokeData(urlPath string) (error, string, string) {
	body, ok := pokeCache.Get(urlPath)

	if !ok {
		res, err := http.Get(urlPath)

		if err != nil {
			fmt.Printf("Error A: %v", err)
			return err, "", ""
		}

		body, err = io.ReadAll(res.Body)

		if err != nil {
			fmt.Printf("Error B: %v", err)
			return err, "", ""
		}

		pokeCache.Add(urlPath, body)
	}

	pokeData := PokeData{}
	err := json.Unmarshal(body, &pokeData)
	if err != nil {
		fmt.Println(err)
	}

	for idx := range pokeData.Results {
		location := pokeData.Results[idx]
		fmt.Printf("%s\n", location.Name)
	}

	parsedURL, _ := url.Parse(pokeData.Next)
	queryParams := parsedURL.Query()

	offset, _ := strconv.Atoi(queryParams.Get("offset"))

	Next := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%v&limit=%v", offset, 20)
	Previous := pokeData.Previous

	//fmt.Printf("Next: %s\n", Next)
	//fmt.Printf("Previous: %s\n\n", Previous)

	return nil, Next, Previous
}

func GetLocationData(urlPath string) error {
	body, ok := pokeCache.Get(urlPath)

	if !ok {
		res, err := http.Get(urlPath)

		if err != nil {
			fmt.Printf("Error A: %v", err)
			return err
		}

		body, err = io.ReadAll(res.Body)

		if err != nil {
			fmt.Printf("Error B: %v", err)
			return err
		}

		pokeCache.Add(urlPath, body)
	}

	locationData := LocationArea{}
	err := json.Unmarshal(body, &locationData)
	if err != nil {
		fmt.Println(err)
	}

	for idx := range locationData.PokemonEncounters {
		pokemon := locationData.PokemonEncounters[idx]
		fmt.Printf("%s\n", pokemon.Name)
	}

	return nil
}

func GetPokemonData(urlPath string) (Pokemon, error) {
	body, ok := pokeCache.Get(urlPath)

	if !ok {
		res, err := http.Get(urlPath)

		if err != nil {
			fmt.Printf("Error A: %v", err)
			return Pokemon{}, err
		}

		body, err = io.ReadAll(res.Body)

		if err != nil {
			fmt.Printf("Error B: %v", err)
			return Pokemon{}, err
		}

		pokeCache.Add(urlPath, body)
	}

	pokemonData := Pokemon{}
	err := json.Unmarshal(body, &pokemonData)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonData, nil
}
