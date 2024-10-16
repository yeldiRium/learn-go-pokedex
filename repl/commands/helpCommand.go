package commands

import "fmt"

func helpCommand() error {
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range GetCommands() {
		fmt.Printf("  %s: %s\n", command.Name, command.Description)
	}
	return nil
}
