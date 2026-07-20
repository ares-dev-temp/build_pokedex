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

	//return []string{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text_out := scanner.Text()

		command := cleanInput(text_out)
		fmt.Printf("Your command was: %s\n", command[0])
	}
}
