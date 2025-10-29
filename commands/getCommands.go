package commands

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    ExitCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore location",
			Callback:    exploreCommand,
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

		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    CatchCommand,
		},
	}
}
