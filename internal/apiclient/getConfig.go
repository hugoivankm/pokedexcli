package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hugoivankm/pokedexcli/internal/pokecache"
)

func Get(url string) (*Config, error) {

	client := NewClient(29*time.Second, 10*time.Second)

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

	var config Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall JSON: %w", err)
	}

	return &config, nil

}

type HttpCachedClient struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(httpTimeout, cacheInterval time.Duration) *HttpCachedClient {
	return &HttpCachedClient{
		httpClient: http.Client{Timeout: httpTimeout},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

func (c *HttpCachedClient) Get(url string) (*http.Response, error) {
	if data, ok := c.cache.Get(url); ok {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader(data)),
			Header:     make(http.Header),
		}, nil
	}

	res, err := c.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, err
	}
	c.cache.Add(url, body)
	res.Body = io.NopCloser((bytes.NewBuffer(body)))
	return res, nil
}
