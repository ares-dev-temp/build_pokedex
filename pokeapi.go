package main

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

var pokeCache = Cache{}

func NewClient(interval time.Duration) {
	pokeCache = *NewCache(interval)
}

func GetPokeData(urlPath string, c *config) error {
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

	c.Next = fmt.Sprintf("https://pokeapi.co/api/v2/location-area?offset=%v&limit=%v", offset+20, 20)
	c.Previous = pokeData.Previous

	return nil
}
