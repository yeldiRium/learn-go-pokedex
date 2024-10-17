package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

var ErrAreaUnknown = errors.New("area unknown")
var ErrPokemonInAreaRequestInvalid = errors.New("failed to create request for pokemon in area")
var ErrPokemonInAreaRequestFailed = errors.New("failed to request pokemon in area")
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

func GetPokemonInArea(httpClient HttpClient, cache pokecache.Cache, areaName string) ([]string, error) {
	url := fmt.Sprintf("%s%s/", BasePokemonInAreaUrl, areaName)

	responseBody, ok := cache.GetEntry(url)
	if !ok {
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrPokemonInAreaRequestInvalid, err)
		}
		response, err := httpClient.Do(request)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrPokemonInAreaRequestFailed, err)
		}
		if response.StatusCode == http.StatusNotFound {
			return nil, ErrAreaDoesntExist
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrPokemonInAreaRequestFailed, err)
		}

		responseBody = body
	}

	var apiResponse pokemonInAreaApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrPokemonInAreaRequestFailed, err)
	}

	foundPokemon := make([]string, len(apiResponse.PokemonEncounters))
	for i, encounter := range apiResponse.PokemonEncounters {
		foundPokemon[i] = encounter.Pokemon.Name
	}

	return foundPokemon, nil
}
