package repl

import (
	"bytes"
	"context"
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
				Handler: func() error {
					handlerCalled = true
					return nil
				},
			},
		}

		input := bytes.NewBufferString("test\n")
		go StartRepl(ctx, input, cliCommands)

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
				Handler: func() error {
					handlerCalled = true
					return nil
				},
			},
		}

		input := bytes.NewBufferString("unknown\n")
		go StartRepl(ctx, input, cliCommands)

		assert.Never(t, func() bool {
			return handlerCalled == true
		}, 100*time.Millisecond, 20*time.Millisecond, "Handler was executed, but should not have been.")
	})
}
