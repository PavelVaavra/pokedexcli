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
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas",
			callback:    pokeapi.CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas",
			callback:    pokeapi.CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all Pokemons located in the area",
			callback:    pokeapi.CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a pokemon",
			callback:    pokeapi.CommandCatch,
		},
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
		ExploreBasis: "https://pokeapi.co/api/v2/location-area/",
		ExploreArea: "",
		CatchBasis: "https://pokeapi.co/api/v2/pokemon/",
		CatchPokemon: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			commandName := cleanInput(scanner.Text())
			if len(commandName) == 1 || len(commandName) == 2 {
				command, ok := commands[commandName[0]]
				if ok {
					if command.name == "explore" {
						urls.ExploreArea = commandName[1]
					}
					if command.name == "catch" {
						urls.CatchPokemon = commandName[1]
					}
					err := command.callback(urls)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println()
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