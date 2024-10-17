package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/yeldiRium/learning-go-pokedex/model"
	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

var ErrGetPokemonInArea = errors.New("GetPokemonInArea")
var ErrAreaDoesntExist = errors.New("area does not exist")

const BasePokemonInAreaUrl = "https://pokeapi.co/api/v2/location-area/"

type pokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}
type pokemonInAreaApiResponse struct {
	Name              string             `json:"name"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

func GetPokemonInArea(httpClient HttpClient, cache pokecache.Cache, areaName string) (model.PokemonEncounters, error) {
	url := fmt.Sprintf("%s%s/", BasePokemonInAreaUrl, areaName)

	responseBody, ok := cache.GetEntry(url)
	if !ok {
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("%w: %w: %w", ErrGetPokemonInArea, ErrRequestInvalid, err)
		}
		response, err := httpClient.Do(request)
		if err != nil {
			return nil, fmt.Errorf("%w: %w: %w", ErrGetPokemonInArea, ErrRequestFailed, err)
		}
		if response.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("%w: %w: %s", ErrGetPokemonInArea, ErrAreaDoesntExist, areaName)
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("%w: %w: %w", ErrGetPokemonInArea, ErrRequestFailed, err)
		}

		responseBody = body
	}

	var apiResponse pokemonInAreaApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: %w: %w", ErrGetPokemonInArea, ErrRequestFailed, err)
	}

	foundPokemon := make(model.PokemonEncounters, len(apiResponse.PokemonEncounters))
	for i, encounter := range apiResponse.PokemonEncounters {
		foundPokemon[i] = model.PokemonEncounter{
			Name: encounter.Pokemon.Name,
		}
	}

	return foundPokemon, nil
}
