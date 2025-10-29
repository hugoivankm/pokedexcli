package commands

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	apiClient "github.com/hugoivankm/pokedexcli/internal/apiclient"
)

func CatchCommand(cfg *apiClient.Config, rest ...any) (*apiClient.Config, error) {
	if len(rest) == 0 {
		return nil, fmt.Errorf("no pokemon was provided")
	}

	pokemonName, ok := rest[0].(string)
	if !ok {
		return nil, fmt.Errorf("pokemon name must be a string type")
	}
	pokemonName = strings.ToLower(pokemonName)

	pokedexData, ok := rest[len(rest)-1].(apiClient.PokedexData)
	if !ok {
		return nil, fmt.Errorf("unable to access pokedex data")
	}

	if strings.TrimSpace(pokemonName) == "" {
		return nil, fmt.Errorf("no pokemon name was provided")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonDetails, err := apiClient.Get[apiClient.PokemonDetails](apiClient.PokemonEndPoint + pokemonName)

	if err != nil {
		return nil, fmt.Errorf("no such a pokemon, please check the name: %w", err)
	}

	baseProbability := calculateCatchProbability(*pokemonDetails)
	shakeProbability := 1.0
	escaped := false

	// There are 4 chances for the pokemon to shake the pokeball
	for range 4 {
		shakeProbability = rand.Float64() * 0.6
		if shakeProbability >= baseProbability {
			escaped = true
			break
		}
	}

	if escaped {
		fmt.Printf("%s escaped!...\n", pokemonName)
	} else {
		pokedexData[pokemonName] = *pokemonDetails
		fmt.Printf("%s was caught!\n", pokemonName)
	}

	return cfg, nil
}

type catchOptions struct {
	BallModifier float64
}

type catchOption func(*catchOptions)

func newCatchOptions() catchOptions {
	return catchOptions{
		BallModifier: 1.0,
	}
}

// func withBallModifier(mod float64) catchOption {
// 	return func(co *catchOptions) {
// 		co.BallModifier = mod
// 	}
// }

func calculateCatchProbability(pokemonDetails apiClient.PokemonDetails, opts ...catchOption) float64 {
	options := newCatchOptions()

	for _, opt := range opts {
		opt(&options)
	}

	difficulty := math.Min(float64(pokemonDetails.BaseExperience)/300.0, 1.0)
	probability := 1.0 - (difficulty * difficulty)
	baseCatchProbability := math.Max(0.03, probability)
	finalProbability := baseCatchProbability * options.BallModifier
	return finalProbability
}
