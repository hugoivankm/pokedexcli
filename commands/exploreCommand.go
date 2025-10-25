package commands

import (
	"fmt"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func exploreCommand(cfg *apiClient.CommandConfig) (*apiClient.CommandConfig, error) {

	if cfg.LocationAreaName == nil || *cfg.LocationAreaName == "" {
		return nil, fmt.Errorf("no location provided")
	}

	currentCfg, err := apiClient.Get(apiClient.LocationAreaEndPoint)
	if err != nil {
		return nil, fmt.Errorf("error acquiring location: %w", err)
	}
	fmt.Print(currentCfg)

	return cfg, nil
}
