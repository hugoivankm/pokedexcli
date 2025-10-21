package commands

import (
	"fmt"

	"github.com/hugoivankm/pokedexcli/internal/apiClient"
	"github.com/hugoivankm/pokedexcli/internal/apiClient/models"
)

func MapCommand(cfg *models.Config) (*models.Config, error) {
	var currentCfg *models.Config
	var err error
	if cfg == nil {
		currentCfg, err = apiClient.GetConfig(apiClient.LocationAreaEndPoint)
		if err != nil {
			return nil, fmt.Errorf("error acquiring config: %w", err)
		}

	} else {
		if cfg.Next != nil {
			currentCfg, err = apiClient.GetConfig(*cfg.Next)
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

	for _, loc := range currentCfg.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	return currentCfg, nil
}
