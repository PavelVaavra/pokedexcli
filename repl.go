package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func replMgr() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			words := cleanInput(scanner.Text())
			word := ""
			if len(words) >= 1 {
				word = words[0]
			}
			fmt.Printf("Your command was: %v\n", word)
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}