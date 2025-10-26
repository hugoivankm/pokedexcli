package commands

import (
	"fmt"

	"github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func MapCommand(cfg *apiclient.Config, rest ...any) (*apiclient.Config, error) {
	var currentCfg *apiclient.Config
	var err error

	if cfg == nil {
		currentCfg, err = apiclient.Get[apiclient.Config](apiclient.LocationAreaEndPoint)
		if err != nil {
			return nil, fmt.Errorf("error acquiring config: %w", err)
		}

	} else {
		if cfg.Next != nil {
			currentCfg, err = apiclient.Get[apiclient.Config](*cfg.Next)
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
