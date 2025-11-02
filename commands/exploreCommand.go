package commands

import (
	"fmt"
	"strings"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func exploreCommand(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error) {

	if len(rest) == 0 {
		return nil, fmt.Errorf("no location was provided")
	}

	LocationAreaName, ok := rest[0].(string)
	if !ok {
		return nil, fmt.Errorf("location must be a string type")
	}

	if strings.TrimSpace(LocationAreaName) == "" {
		return nil, fmt.Errorf("empty location provided")
	}

	data, err := apiClient.Get[apiClient.LocationArea](apiClient.LocationAreaEndPoint + LocationAreaName)
	if err != nil {
		return nil, fmt.Errorf("error acquiring location: %w", err)
	}

	fmt.Printf("Exploring %s...\n", data.Name)
	fmt.Print("Found Pokemon:\n")
	for _, v := range data.PokemonEncounters {
		fmt.Print("- " + v.Pokemon.Name + "\n")
	}

	return cfg, nil
}
