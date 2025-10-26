package commands

import apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error)
}
