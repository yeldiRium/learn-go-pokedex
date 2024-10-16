package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yeldiRium/learning-go-pokedex/commands"
)

func StartRepl(input io.Reader, cliCommands map[string]commands.CliCommand) {
	scanner := bufio.NewScanner(input)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		command, exists := cliCommands[input]
		if !exists {
			continue
		}

		command.Handler()
	}
}
