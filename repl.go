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
	var cfg *apiClient.Config
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		registry := commands.GetCommands()
		commandWord, ok := registry[input[0]]

		if ok {
			var err error
			params := input[1:]
			args := make([]any, len(params))

			for i, v := range params {
				args[i] = v
			}

			cfg, err = commandWord.Callback(cfg, args...)
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
