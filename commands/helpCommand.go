package commands

import (
	"fmt"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func HelpCommand(*apiClient.Config) (*apiClient.Config, error) {
	fmt.Print(`Welcome to the Pokedex!
Usage:
`)
	registry := GetCommands()
	for command := range registry {
		c := registry[command]
		fmt.Println(c.Name + ": " + c.Description)
	}
	return nil, nil
}
