package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"github.com/PavelVaavra/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    pokeapi.CommandMap,
		},
	}
}

func commandExit(urls *pokeapi.Urls) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(urls *pokeapi.Urls) error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

`)
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(urls *pokeapi.Urls) error
}

func replMgr() {
	commands := getCommands()
	urls := &pokeapi.Urls{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			commandName := cleanInput(scanner.Text())
			if len(commandName) == 1 {
				command, ok := commands[commandName[0]]
				if ok {
					err := command.callback(urls)
					if err != nil {
						fmt.Println(err)
					}
					continue
				}
			}
			fmt.Println("Unknown command")
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}