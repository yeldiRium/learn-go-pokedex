package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yeldiRium/learning-go-pokedex/repl/commands"
)

func StartRepl(input io.Reader) {
	scanner := bufio.NewScanner(input)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		command, exists := commands.GetCommands()[input]
		if !exists {
			continue
		}

		command.Handler()
	}
}
