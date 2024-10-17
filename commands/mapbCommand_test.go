package commands_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/learning-go-pokedex/commands"
)

func TestMapbCommand(t *testing.T) {
	t.Run("initially fails, since no previous url is set", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"next": "http://test-next-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		buf := new(bytes.Buffer)
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithOutput(buf)

		err := commands.MapbCommand(config, []string{})
		assert.ErrorIs(t, err, commands.ErrBeginningOfAreasReached)
	})

	t.Run("requests the previous page", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"previous": "http://test-next-previous-map-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithPreviousMapUrl("http://test-current-previous-map-url/")

		commands.MapbCommand(config, []string{})
		err := commands.MapbCommand(config, []string{})
		assert.NoErrorf(t, err, "mapb command should not have returned an error")

		assert.Equal(t, []string{"http://test-current-previous-map-url/", "http://test-next-previous-map-url/"}, client.wasCalledWithUrls)
	})

	t.Run("returns an error if the request failed", func(t *testing.T) {
		client := mockHttpClient{
			shouldError: fmt.Errorf("test error"),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithPreviousMapUrl("http://previous-map-url/")

		err := commands.MapbCommand(config, []string{})
		assert.ErrorIs(t, err, commands.ErrPreviousMapRequestFailed)
	})

	t.Run("returns an error if the response can not be parsed", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte("invalid-json"),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithPreviousMapUrl("http://previous-map-url/")

		err := commands.MapbCommand(config, []string{})
		assert.ErrorIs(t, err, commands.ErrPreviousMapRequestFailed)
	})

	t.Run("returns an error when it can't go further back", func(t *testing.T) {
		config := commands.NewCliConfig()

		err := commands.MapbCommand(config, []string{})
		assert.ErrorIs(t, err, commands.ErrBeginningOfAreasReached)
	})
}
