package commands

import "fmt"

func helpCommand(config *CliConfig) error {
	fmt.Fprintln(config.output, "Welcome to the pokedex!")
	fmt.Fprintln(config.output, "Usage:")
	fmt.Fprintln(config.output)
	for _, command := range GetCommands() {
		fmt.Fprintf(config.output, "  %s: %s\n", command.Name, command.Description)
	}
	return nil
}
