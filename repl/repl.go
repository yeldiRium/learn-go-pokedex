package repl

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/yeldiRium/learning-go-pokedex/commands"
)

func scanReader(input io.Reader, lines chan string) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}

func StartRepl(ctx context.Context, input io.Reader, cliCommands map[string]commands.CliCommand) {
	lines := make(chan string)
	go scanReader(input, lines)

	for {
		fmt.Printf("pokedex > ")

		select {
		case <-ctx.Done():
			return
		case line, ok := <-lines:
			if !ok {
				return
			}

			command, exists := cliCommands[line]
			if !exists {
				continue
			}

			command.Handler()
		}
	}
}
