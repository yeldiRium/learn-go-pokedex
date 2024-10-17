package pokeapi_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeldiRium/learning-go-pokedex/model"
	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

func TestGetAreaList(t *testing.T) {
	t.Run("requests an area list section and returns the found areas and pagination", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte(`{"next": "http://test-next-url/", "results": [{"name": "area1"}, {"name": "area2"}]}`),
		}
		cache := pokecache.Cache{}
		result, err := pokeapi.GetAreaList(&client, cache, "https://whatever/")
		assert.NoError(t, err)
		nextAreaUrl := "http://test-next-url/"
		assert.Equal(t, &pokeapi.AreaListResult{
			NextAreaUrl:     &nextAreaUrl,
			PreviousAreaUrl: nil,
			Areas: model.Areas{
				{Name: "area1"},
				{Name: "area2"},
			},
		}, result)
	})
	t.Run("returns an error if the given URL is invalid", func(t *testing.T) {
		cache := pokecache.Cache{}
		_, err := pokeapi.GetAreaList(http.DefaultClient, cache, "::/-_([>&}invalid-urld>-_[}]")
		assert.ErrorIs(t, err, pokeapi.ErrGetAreaList)
		assert.ErrorIs(t, err, pokeapi.ErrRequestInvalid)
	})

	t.Run("returns an error if the request failed", func(t *testing.T) {
		client := mockHttpClient{
			shouldError: fmt.Errorf("test error"),
		}
		cache := pokecache.Cache{}

		_, err := pokeapi.GetAreaList(&client, cache, "http://test-url/")
		assert.ErrorIs(t, err, pokeapi.ErrGetAreaList)
		assert.ErrorIs(t, err, pokeapi.ErrRequestFailed)
	})

	t.Run("returns an error if the response can not be parsed", func(t *testing.T) {
		client := mockHttpClient{
			shouldReturn: []byte("invalid-json"),
		}
		cache := pokecache.Cache{}

		_, err := pokeapi.GetAreaList(&client, cache, "http://test-url/")
		assert.ErrorIs(t, err, pokeapi.ErrGetAreaList)
		assert.ErrorIs(t, err, pokeapi.ErrRequestFailed)
	})
}
