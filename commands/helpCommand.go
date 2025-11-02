package commands

import (
	"fmt"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func HelpCommand(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error) {
	fmt.Print(`Welcome to the Pokedex!
Usage:
`)
	registry := GetCommands()
	for command := range registry {
		c := registry[command]
		fmt.Print(c.Name + ": " + c.Description + "\n")
	}
	return nil, nil
}
