package commands_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/learning-go-pokedex/commands"
	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
)

func TestMapbCommand(t *testing.T) {
	t.Run("initially requests the first page and prints all found areas", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"next": "http://test-next-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		buf := new(bytes.Buffer)
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithOutput(buf)

		err := commands.MapbCommand(config)
		assert.NoErrorf(t, err, "mapb command should not have returned an error")

		assert.Equal(t, []string{"https://pokeapi.co/api/v2/location-area/?limit=20"}, client.wasCalledWithUrls)
		assert.Equal(t, "area1\narea2\n", buf.String())
	})

	t.Run("requests the previous page", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"previous": "http://test-next-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client)

		commands.MapbCommand(config)
		err := commands.MapbCommand(config)
		assert.NoErrorf(t, err, "mapb command should not have returned an error")

		assert.Equal(t, []string{"https://pokeapi.co/api/v2/location-area/?limit=20", "http://test-next-url/"}, client.wasCalledWithUrls)
	})

	t.Run("returns an error if the request failed", func(t *testing.T) {
		client := mockHttpClient{
			shouldError: fmt.Errorf("test error"),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client)

		err := commands.MapbCommand(config)
		assert.ErrorIs(t, err, commands.ErrPreviousMapRequestFailed)
	})

	t.Run("returns an error if the response can not be parsed", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte("invalid-json"),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client)

		err := commands.MapbCommand(config)
		assert.ErrorIs(t, err, commands.ErrPreviousMapRequestFailed)
	})

	t.Run("stays at the beginning of the list, when it is reached", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		config := commands.NewCliConfig().
			WithHttpClient(&client)

		commands.MapbCommand(config)
		commands.MapbCommand(config)
		assert.Len(t, client.wasCalledWithUrls, 2)
		assert.Equal(t, []string{pokeapi.BaseAreaListUrl, pokeapi.BaseAreaListUrl}, client.wasCalledWithUrls)
	})

	t.Run("running mapb command without a previous map url makes the next map command request the second page", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"next": "http://test-next-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		buf := new(bytes.Buffer)
		config := commands.NewCliConfig().
			WithHttpClient(&client).
			WithOutput(buf)

		commands.MapbCommand(config)
		commands.MapCommand(config)

		assert.Equal(t, []string{"https://pokeapi.co/api/v2/location-area/?limit=20", "http://test-next-url/"}, client.wasCalledWithUrls)
	})
}
