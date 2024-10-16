package main

import (
	"os"

	"github.com/yeldiRium/learning-go-pokedex/commands"
	"github.com/yeldiRium/learning-go-pokedex/repl"
)

func main() {
	input := os.Stdin
	cliCommands := commands.GetCommands()
	repl.StartRepl(input, cliCommands)
}
