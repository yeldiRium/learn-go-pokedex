package commands_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/learning-go-pokedex/commands"
)

func TestMapCommand(t *testing.T) {
	t.Run("initially requests the first page and prints all found areas", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"next": "http://test-next-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		buf := new(bytes.Buffer)
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithOutput(buf)

		err := commands.MapCommand(config)
		assert.NoErrorf(t, err, "map command should not have returned an error")

		assert.Equal(t, []string{"https://pokeapi.co/api/v2/location-area/?limit=20"}, client.wasCalledWithUrls)
		assert.Equal(t, "area1\narea2\n", buf.String())
	})

	t.Run("requests the next page", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"next": "http://test-next-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client)

		commands.MapCommand(config)
		err := commands.MapCommand(config)
		assert.NoErrorf(t, err, "map command should not have returned an error")

		assert.Equal(t, []string{"https://pokeapi.co/api/v2/location-area/?limit=20", "http://test-next-url/"}, client.wasCalledWithUrls)
	})

	t.Run("returns an error if the request failed", func(t *testing.T) {
		client := mockHttpClient{
			shouldError: fmt.Errorf("test error"),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client)

		err := commands.MapCommand(config)
		assert.ErrorIs(t, err, commands.ErrNextMapRequestFailed)
	})

	t.Run("returns an error if the response can not be parsed", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte("invalid-json"),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client)

		err := commands.MapCommand(config)
		assert.ErrorIs(t, err, commands.ErrNextMapRequestFailed)
	})

	t.Run("returns an error when it can't go further", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithNextMapUrl("https://current-map-url/")

		commands.MapCommand(config)
		err := commands.MapCommand(config)
		assert.ErrorIs(t, err, commands.ErrEndOfAreasReached)
	})
}
