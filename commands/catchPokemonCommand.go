package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yeldiRium/learning-go-pokedex/model/formulas"
	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
	"github.com/yeldiRium/learning-go-pokedex/utilities"
)

var ErrCatchPokemonArguments = errors.New("exactly one argument required: pokemon name")
var ErrCatchFailed = errors.New("failed to catch pokemon")

func CatchPokemonCommand(config *CliConfig, arguments []string) error {
	if len(arguments) != 1 {
		return ErrCatchPokemonArguments
	}

	pokemonName := strings.ToLower(arguments[0])

	pokemon, err := pokeapi.GetPokemon(config.httpClient, config.cache, pokemonName)
	if err != nil {
		if errors.Is(err, pokeapi.ErrPokemonDoesntExist) {
			fmt.Fprintf(config.output, "Pokemon %s does not exist. Try to find one with the command 'expore <area>'.\n", pokemonName)
			return nil
		}
		return fmt.Errorf("%w: %w", ErrCatchFailed, err)
	}

	species, err := pokeapi.GetSpecies(config.httpClient, config.cache, pokemon.Species)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCatchFailed, err)
	}

	fmt.Fprintf(config.output, "Throwing a pokeball at %s...\n", pokemon.Name)

	catchResult := utilities.RandomWithProbability(formulas.CatchPokemonProbability(species))

	if !catchResult {
		fmt.Fprintf(config.output, "%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Fprintf(config.output, "%s was caught!\n", pokemon.Name)
	config.pokedex[pokemon.Name] = pokemon
	return nil
}
