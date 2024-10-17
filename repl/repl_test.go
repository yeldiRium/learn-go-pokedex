package repl

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/learning-go-pokedex/commands"
)

func TestStartRepl(t *testing.T) {
	t.Run("repl executes correct command handler.", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var handlerCalled bool
		cliCommands := map[string]commands.CliCommand{
			"test": {
				Handler: func(_ *commands.CliConfig, _ []string) error {
					handlerCalled = true
					return nil
				},
			},
		}

		input := bytes.NewBufferString("test\n")
		output := new(bytes.Buffer)
		go StartRepl(ctx, input, output, cliCommands)

		assert.Eventually(t, func() bool {
			return handlerCalled == true
		}, 500*time.Millisecond, 50*time.Millisecond, "Handler was not executed within timeout.")
	})

	t.Run("input does not match any command", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var handlerCalled bool
		cliCommands := map[string]commands.CliCommand{
			"test": {
				Handler: func(_ *commands.CliConfig, _ []string) error {
					handlerCalled = true
					return nil
				},
			},
		}

		input := bytes.NewBufferString("unknown\n")
		output := new(bytes.Buffer)
		go StartRepl(ctx, input, output, cliCommands)

		assert.Never(t, func() bool {
			return handlerCalled == true
		}, 100*time.Millisecond, 20*time.Millisecond, "Handler was executed, but should not have been.")
	})

	t.Run("displays error returned by handler", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		cliCommands := map[string]commands.CliCommand{
			"test": {
				Handler: func(_ *commands.CliConfig, _ []string) error {
					return errors.New("test error")
				},
			},
		}

		input := bytes.NewBufferString("test\n")
		output := new(bytes.Buffer)
		go StartRepl(ctx, input, output, cliCommands)

		time.Sleep(100 * time.Millisecond)
		text, _ := io.ReadAll(output)
		assert.Equal(t, "pokedex > error: test error\npokedex > ", string(text))
	})
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedOutput []string
	}{
		{
			name:           "Should ignore empty input",
			input:          "   ",
			expectedOutput: []string{},
		},
		{
			name:           "Should clean up whitespace",
			input:          "  test ",
			expectedOutput: []string{"test"},
		},
		{
			name:           "Should split input into words",
			input:          "test arg1 arg2",
			expectedOutput: []string{"test", "arg1", "arg2"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			output := cleanInput(testCase.input)

			assert.Equal(t, testCase.expectedOutput, output)
		})
	}
}
