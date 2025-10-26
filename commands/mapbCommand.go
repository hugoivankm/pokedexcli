package commands

import (
	"fmt"

	"github.com/hugoivankm/pokedexcli/internal/apiclient"
	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func MapbCommand(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error) {
	var currentCfg *apiClient.Config
	var err error
	if cfg == nil {
		currentCfg, err = apiClient.Get[apiclient.Config](apiClient.LocationAreaEndPoint)
		if err != nil {
			return nil, fmt.Errorf("error acquiring config: %w", err)
		}

	} else {
		if cfg.Previous != nil {
			currentCfg, err = apiClient.Get[apiclient.Config](*cfg.Previous)
			if err != nil {
				return nil, fmt.Errorf("error acquiring previous config: %w", err)
			}
		} else {
			currentCfg = cfg
			fmt.Println("you're on the first page")
			return currentCfg, nil
		}
	}

	if currentCfg == nil {
		return cfg, fmt.Errorf("no page data")
	}

	for _, loc := range currentCfg.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return currentCfg, nil
}
