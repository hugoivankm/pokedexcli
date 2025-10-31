package commands

import (
	"fmt"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func PokedexCommand(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error) {

	pokedexData, ok := rest[len(rest)-1].(apiClient.PokedexData)
	if !ok {
		return nil, fmt.Errorf("unable to access pokedex data")
	}

	if len(pokedexData) == 0 {
		return nil, fmt.Errorf("no pokemons in the pokedex")
	}

	for k := range pokedexData {
		fmt.Printf("  - %v\n", k)

	}
	return cfg, nil
}
