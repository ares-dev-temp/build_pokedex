package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex_program/internal"
	"strings"
	"time"
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
	//c := config{Next: url + "?offset=0&limit=20", Previous: url + "?offset=0&limit=20"}
	c := config{Next: url + "?offset=0&limit=20", Previous: ""}
	cPtr := &c

	SetCommands(cPtr)

	const baseTime = 5 * time.Second
	internal.NewClient(baseTime)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text_out := scanner.Text()

		commands := cleanInput(text_out)

		run_output(commands)

	}
}
