package commands

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInspectPokemonArguments = errors.New("exactly one argument required: pokemon name")

func InspectCommand(config *CliConfig, arguments []string) error {
	if len(arguments) != 1 {
		return ErrInspectPokemonArguments
	}

	pokemonName := strings.ToLower(arguments[0])

	pokemon, ok := config.pokedex[pokemonName]
	if !ok {
		fmt.Fprintln(config.output, "You have not caught that pokemon yet.")
		return nil
	}

	fmt.Fprintf(config.output, "Name: %s\n", pokemon.Name)
	fmt.Fprintf(config.output, "Height: %d decimeters\n", pokemon.Weight)
	fmt.Fprintf(config.output, "Weight: %d hectograms\n", pokemon.Weight)
	fmt.Fprintf(config.output, "Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Fprintf(config.output, "  - %s\n", pokemonType)
	}
	fmt.Fprintf(config.output, "Stats:\n")
	fmt.Fprintf(config.output, "  - HP             : %3d\n", pokemon.BaseStats.Hp)
	fmt.Fprintf(config.output, "  - Attack         : %3d\n", pokemon.BaseStats.Attack)
	fmt.Fprintf(config.output, "  - Defense        : %3d\n", pokemon.BaseStats.Defense)
	fmt.Fprintf(config.output, "  - Special Attack : %3d\n", pokemon.BaseStats.SpecialAttack)
	fmt.Fprintf(config.output, "  - Special Defense: %3d\n", pokemon.BaseStats.SpecialDefense)
	fmt.Fprintf(config.output, "  - Speed          : %3d\n", pokemon.BaseStats.Speed)
	return nil
}
