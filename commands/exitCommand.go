package commands

import (
	"fmt"
	"os"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func ExitCommand(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error) {
	fmt.Println("Closing the Pokedex... Goodbye! ")
	os.Exit(0)
	return nil, nil
}
