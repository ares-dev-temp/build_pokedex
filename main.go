package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	url := "https://pokeapi.co/api/v2/location-area"
	c := config{Next: url + "?offset=0&limit=20", Previous: url + "?offset=0&limit=20"}
	cPtr := &c

	SetCommands(cPtr)

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
