package repl

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/yeldiRium/learning-go-pokedex/commands"
)

func scanReader(input io.Reader, lines chan string) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}

func cleanInput(input string) (words []string) {
	output := strings.Split(input, " ")

	outputWithoutEmptyWords := []string{}
	for _, word := range output {
		if word == "" {
			continue
		}
		outputWithoutEmptyWords = append(outputWithoutEmptyWords, word)
	}

	return outputWithoutEmptyWords
}

func StartRepl(ctx context.Context, input io.Reader, cliCommands map[string]commands.CliCommand) {
	cliState := commands.NewCliConfig()
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

			cleanedInput := cleanInput(line)

			if len(cleanedInput) == 0 {
				continue
			}

			commandName := cleanedInput[0]

			cliCommand, exists := cliCommands[commandName]
			if !exists {
				continue
			}

			cliCommand.Handler(cliState)
		}
	}
}
