package commands

import "fmt"

func PokedexCommand(config *CliConfig, _ []string) error {
	fmt.Fprintln(config.output, "Your pokedex:")

	for _, pokemon := range config.pokedex {
		fmt.Fprintf(config.output, "  - %s\n", pokemon.Name)
	}
	return nil
}
