package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type config struct {
	Next     string
	Previous string
}

type PokeData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`

	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandSet = map[string]cliCommand{}

func SetCommands(c *config) {
	commandSet = map[string]cliCommand{
		"map": {
			name:        "map",
			description: "list locations",
			callback:    func() error { return commandMap(c) },
		},
		"mapb": {
			name:        "mapb",
			description: "list previous locations",
			callback:    func() error { return commandMapb(c) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return commandExit(c) },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp(c) },
		},
	}
}

func commandMapb(c *config) error {
	if len(c.Previous) > 0 {
		return commandList(c.Previous, c)
	}

	return nil
}

func commandMap(c *config) error {
	return commandList(c.Next, c)
}

func commandList(urlPath string, c *config) error {
	//res, err := http.Get("https://pokeapi.co/api/v2/location?name")
	res, err := http.Get(urlPath)

	if err != nil {
		fmt.Printf("Error A: %v", err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed, status code: %v", res.StatusCode)
		return err
	}

	if err != nil {
		fmt.Printf("Error B: %v", err)
		return err
	}

	pokeData := PokeData{}
	err = json.Unmarshal(body, &pokeData)
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

func commandHelp(c *config) error {
	response := "Welcome to the Pokedex!\nUsage:\n\n"

	for commandName := range commandSet {
		cli := commandSet[commandName]
		response += fmt.Sprintf("%s: %s\n", cli.name, cli.description)
	}
	response = response[0 : len(response)-1]

	fmt.Println(response)
	return nil
}

func commandExit(c *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
