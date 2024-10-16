package commands

type CliCommand struct {
	Name        string
	Description string
	Handler     func() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Show help text.",
			Handler:     helpCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex.",
			Handler:     exitCommand,
		},
	}
}
