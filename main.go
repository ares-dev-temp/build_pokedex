package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandSet = map[string]cliCommand{}

func cleanInput(text string) []string {
	outs := []string{}
	curr_word := ""

	for _, char := range text {
		if char != ' ' {
			curr_word += string(char)
		}
		if len(curr_word) > 0 && char == ' ' {
			outs = append(outs, strings.ToLower(curr_word))
			curr_word = ""
		}
	}

	if len(curr_word) > 0 {
		outs = append(outs, strings.ToLower(curr_word))
	}

	return outs
}

func commandHelp() error {
	response := "Welcome to the Pokedex!\nUsage:\n\n"

	for commandName := range commandSet {
		cli := commandSet[commandName]
		response += fmt.Sprintf("%s: %s\n", cli.name, cli.description)
	}
	response = response[0 : len(response)-2]

	fmt.Println(response)
	return nil
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	commandSet = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text_out := scanner.Text()

		command := cleanInput(text_out)[0]

		if cli, ok := commandSet[command]; ok {
			err := cli.callback()

			if err != nil {
				fmt.Printf("errors %v\n", err)
			}
		} else {
			fmt.Printf("Unknown Command\n")
		}

	}
}
