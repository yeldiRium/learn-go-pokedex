package pokeapi_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

func TestGetPokemonInArea(t *testing.T) {
	t.Run("returns the pokemon found in the area", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"name": "some area", "pokemon_encounters": [{ "pokemon": {"name": "pokemon-1" }}, { "pokemon": {"name": "pokemon-2" }}]}`),
		}
		cache := pokecache.Cache{}
		result, err := pokeapi.GetPokemonInArea(&client, cache, "some area")
		assert.NoError(t, err)
		assert.Equal(t, []string{"pokemon-1", "pokemon-2"}, result)
	})

	t.Run("returns an error if the area does not exist", func(t *testing.T) {
		client := mockHttpClient{
			returnStatus: http.StatusNotFound,
		}
		cache := pokecache.Cache{}
		_, err := pokeapi.GetPokemonInArea(&client, cache, "invalid-area")
		assert.ErrorIs(t, err, pokeapi.ErrAreaDoesntExist)
	})

	t.Run("returns an error if the request failed", func(t *testing.T) {
		client := mockHttpClient{
			shouldError: fmt.Errorf("test error"),
		}
		cache := pokecache.Cache{}

		_, err := pokeapi.GetPokemonInArea(&client, cache, "some area")
		assert.ErrorIs(t, err, pokeapi.ErrPokemonInAreaRequestFailed)
	})

	t.Run("returns an error if the response can not be parsed", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte("invalid-json"),
		}
		cache := pokecache.Cache{}

		_, err := pokeapi.GetPokemonInArea(&client, cache, "some area")
		assert.ErrorIs(t, err, pokeapi.ErrPokemonInAreaRequestFailed)
	})
}
