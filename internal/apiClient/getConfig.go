package apiClient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hugoivankm/pokedexcli/internal/apiClient/models"
)

func GetConfig(url string) (*models.Config, error) {
	client := http.Client{
		Timeout: 29 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to make GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received a non-OK HTTP status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var config models.Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall JSON: %w", err)
	}

	return &config, nil

}
