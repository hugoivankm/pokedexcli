package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/hugoivankm/pokedexcli/commands"
	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg *apiClient.CommandConfig
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		registry := commands.GetCommands()
		commandWord, ok := registry[input[0]]

		var err error

		if ok {
			cfg, err = commandWord.Callback(cfg)
			if err != nil {
				fmt.Println(fmt.Errorf("error acquiring config: %w", err))
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	result := strings.ToLower((text))
	words := strings.Fields(result)
	return words
}
