package commands

import apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*apiClient.CommandConfig) (*apiClient.CommandConfig, error)
}
