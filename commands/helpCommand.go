package commands

import (
	"fmt"

	"github.com/hugoivankm/pokedexcli/internal/apiClient/models"
)

func HelpCommand(*models.Config) (*models.Config, error) {
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
