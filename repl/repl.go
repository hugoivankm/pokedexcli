package repl

import (
	"fmt"
	"strings"

	commands "github.com/hugoivankm/pokedexcli/commands"
	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func StartRepl() {
	var cfg *apiClient.Config
	pdx := apiClient.PokedexData{}
	history := NewCommandHistory()

	for {
		prompt := ("Pokedex > ")
		text, err := readLineWithHistory(prompt, history)

		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("\nGoodbye!")
				break
			}
		}

		input, err := cleanInput(text, nil)

		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if len(input) == 0 {
			continue
		}
		history.Add(strings.Join(input, " "))
		registry := commands.GetCommands()
		commandWord, ok := registry[input[0]]

		if ok {
			var err error
			params := input[1:]
			args := make([]any, len(params))

			for i, v := range params {
				args[i] = v
			}

			args = append(args, pdx)

			cfg, err = commandWord.Callback(cfg, args...)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string, err error) ([]string, error) {
	if err != nil {
		return nil, err
	}
	result := strings.ToLower((text))
	words := strings.Fields(result)
	return words, nil
}
