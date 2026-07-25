package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"pokedex_program/internal"
)

// for paging around location areas
type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandSet = map[string]cliCommand{}
var location = ""
var pokemon = ""

var caught_pokemons = map[string]internal.Pokemon{}

func SetCommands(c *config) {
	commandSet = map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "attempts catching a pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "lists out all pokemons located in an area",
			callback:    commandExplore,
		},
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

func run_output(commands []string) {
	if cli, ok := commandSet[commands[0]]; ok {
		if commands[0] == "explore" {
			location = commands[1]
		}

		if commands[0] == "catch" {
			pokemon = commands[1]
		}

		cli.callback()

		/*if err != nil {
			fmt.Printf("errors %v\n", err)
		}*/
	} else {
		fmt.Printf("Unknown Command\n")
	}
}

func commandCatch() error {
	url_path := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemonData, err := internal.GetPokemonData(url_path)

	if err != nil {
		fmt.Printf("%s doesn't exist!\n", pokemon)
	} else {
		val := (float64(pokemonData.BaseExperience) / 208.0)

		if val < rand.Float64() {
			fmt.Printf("%s was caught!\n", pokemon)
			caught_pokemons[pokemon] = pokemonData
		} else {
			fmt.Printf("%s escaped!\n", pokemon)
		}
	}

	return err
}

func commandExplore() error {
	url_path := "https://pokeapi.co/api/v2/location-area/" + location
	return internal.GetLocationData(url_path)
}

func commandMapb(c *config) error {
	if len(c.Previous) > 0 {
		//temp := c.Previous
		err, Next, Previous := internal.GetPokeData(c.Previous)
		c.Next = Next
		c.Previous = Previous
		return err
	}

	return nil
}

func commandMap(c *config) error {
	//temp := c.Next
	err, Next, Previous := internal.GetPokeData(c.Next)
	c.Next = Next
	c.Previous = Previous
	return err
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
