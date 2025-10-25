package commands

import (
	"fmt"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func MapCommand(cfg *apiClient.CommandConfig) (*apiClient.CommandConfig, error) {
	var currentCfg *apiClient.CommandConfig
	var err error
	if cfg == nil {
		currentCfg, err = apiClient.Get(apiClient.LocationAreaEndPoint)
		if err != nil {
			return nil, fmt.Errorf("error acquiring config: %w", err)
		}

	} else {
		if cfg.Config.Next != nil {
			currentCfg, err = apiClient.Get(*cfg.Config.Next)
			if err != nil {
				return nil, fmt.Errorf("error acquiring next config: %w", err)
			}
		} else {
			currentCfg = cfg
			fmt.Println("you're on the last page")
			return currentCfg, nil
		}
	}

	if currentCfg == nil {
		return cfg, fmt.Errorf("no page data")
	}

	for _, loc := range currentCfg.Config.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return currentCfg, nil
}
