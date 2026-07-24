package main

import (
	"fmt"
	"os"
)

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
		return GetPokeData(c.Previous, c)
	}

	return nil
}

func commandMap(c *config) error {
	return GetPokeData(c.Next, c)
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
