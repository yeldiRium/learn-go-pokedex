package commands

type CliCommand struct {
	Name        string
	Description string
	Handler     func(*CliConfig) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Show help text.",
			Handler:     helpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Show the next area of the map.",
			Handler:     MapCommand,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows the previous area of the map.",
			Handler:     MapbCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex.",
			Handler:     exitCommand,
		},
	}
}
