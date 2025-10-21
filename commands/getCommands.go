package commands

import (
	"github.com/hugoivankm/pokedexcli/commands/models"
)

func GetCommands() map[string]models.CliCommand {
	return map[string]models.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    ExitCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    HelpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 map locations",
			Callback:    MapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 map locations",
			Callback:    MapbCommand,
		},
	}
}
