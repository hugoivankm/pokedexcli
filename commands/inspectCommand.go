package commands

import (
	"fmt"
	"strings"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func InspectCommand(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error) {
	if len(rest) == 0 {
		return nil, fmt.Errorf("no pokemon was provided")
	}

	pokemonName, ok := rest[0].(string)
	if !ok {
		return nil, fmt.Errorf("pokemon name must be a string type")
	}
	pokemonName = strings.ToLower(pokemonName)

	pokedexData, ok := rest[len(rest)-1].(apiClient.PokedexData)
	if !ok {
		return nil, fmt.Errorf("unable to access pokedex data")
	}

	details, ok := pokedexData[pokemonName]
	if !ok {
		return nil, fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", details.Name)
	fmt.Printf("Height: %v\n", details.Height)
	fmt.Printf("Weight: %v\n", details.Weight)
	fmt.Printf("Stats:\n")
	for _, v := range details.Stats {
		fmt.Printf("  -%v: %v\n", v.Stat.Name, v.BaseStat)

	}
	fmt.Printf("Types:\n")
	for _, w := range details.Types {
		fmt.Printf("  -%v\n", w.Type.Name)

	}

	return cfg, nil
}
