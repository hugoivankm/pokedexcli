package commands

import (
	"fmt"
	"os"

	"github.com/hugoivankm/pokedexcli/internal/apiClient/models"
)

func ExitCommand(*models.Config) (*models.Config, error) {
	fmt.Println("Closing the Pokedex... Goodbye! ")
	os.Exit(0)
	return nil, nil
}
