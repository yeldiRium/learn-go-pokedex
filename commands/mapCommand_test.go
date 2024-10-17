package commands_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/learning-go-pokedex/commands"
)

type mockHttpClient struct {
	wasCalledWithUrls []string
	shouldReturn      []byte
	shouldError       error
}

func (c *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	c.wasCalledWithUrls = append(c.wasCalledWithUrls, req.URL.String())
	return &http.Response{
		Body: io.NopCloser(bytes.NewReader(c.shouldReturn)),
	}, c.shouldError
}

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

		assert.Equal(t, client.wasCalledWithUrls, []string{"https://pokeapi.co/api/v2/location-area/?limit=20"})
		assert.Equal(t, buf.String(), "area1\narea2\n")
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

		assert.Equal(t, client.wasCalledWithUrls, []string{"https://pokeapi.co/api/v2/location-area/?limit=20", "http://test-next-url/"})
	})

	t.Run("returns an error if the next URL is invalid", func(t *testing.T) {
		config := commands.NewCliConfig().
			WithNextMapUrl("::/-_([>&}invalid-urld>-_[}]")

		err := commands.MapCommand(config)
		assert.ErrorIs(t, err, commands.ErrNextMapRequestInvalid)
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
}
