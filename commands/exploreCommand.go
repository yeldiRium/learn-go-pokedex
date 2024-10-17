package commands

import (
	"errors"
	"fmt"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
)

var ErrExploreArguments = errors.New("exactly one argument required: area name")
var ErrExploreFailed = errors.New("failed to explore")

func ExploreCommand(config *CliConfig, arguments []string) error {
	if len(arguments) != 1 {
		return ErrExploreArguments
	}

	areaName := arguments[0]
	fmt.Fprintf(config.output, "Exploring %s...\n", areaName)

	foundPokemon, err := pokeapi.GetPokemonInArea(config.httpClient, config.cache, areaName)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrExploreFailed, err)
	}

	fmt.Fprintln(config.output, "Found pokemon:")
	for _, pokemon := range foundPokemon {
		fmt.Fprintln(config.output, pokemon)
	}

	return nil
}
