package models

import "github.com/hugoivankm/pokedexcli/internal/apiClient/models"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*models.Config) (*models.Config, error)
}
