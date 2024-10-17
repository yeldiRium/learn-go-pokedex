package pokeapi_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
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

func TestGetAreaList(t *testing.T) {
	t.Run("returns an error if the given URL is invalid", func(t *testing.T) {
		_, err := pokeapi.GetAreaList(http.DefaultClient, "::/-_([>&}invalid-urld>-_[}]")
		assert.ErrorIs(t, err, pokeapi.ErrAreaListRequestInvalid)
	})

	t.Run("returns an error if the request failed", func(t *testing.T) {
		client := mockHttpClient{
			shouldError: fmt.Errorf("test error"),
		}

		_, err := pokeapi.GetAreaList(&client, "http://test-url/")
		assert.ErrorIs(t, err, pokeapi.ErrAreaListRequestFailed)
	})

	t.Run("returns an error if the response can not be parsed", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte("invalid-json"),
		}

		_, err := pokeapi.GetAreaList(&client, "http://test-url/")
		assert.ErrorIs(t, err, pokeapi.ErrAreaListRequestFailed)
	})
}
