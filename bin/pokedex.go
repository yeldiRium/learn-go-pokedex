package main

import (
	"context"
	"os"

	"github.com/yeldiRium/learning-go-pokedex/commands"
	"github.com/yeldiRium/learning-go-pokedex/repl"
)

func main() {
	ctx := context.Background()
	input := os.Stdin
	output := os.Stdout
	cliCommands := commands.GetCommands()
	repl.StartRepl(ctx, input, output, cliCommands)
}
