package commands

type CliCommand struct {
	Name        string
	Description string
	Handler     func(config *CliConfig, arguments []string) error
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
		"explore": {
			Name:        "explore <area>",
			Description: "Explore an area",
			Handler:     ExploreCommand,
		},
		"catch": {
			Name:        "catch <pokemon>",
			Description: "Try to catch a pokemon!",
			Handler:     CatchPokemonCommand,
		},
		"inspect": {
			Name:        "inspect <pokemon>",
			Description: "Look at your pokedex entry for a pokemon.",
			Handler:     InspectCommand,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show all caught pokemon.",
			Handler:     PokedexCommand,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex.",
			Handler:     exitCommand,
		},
	}
}
