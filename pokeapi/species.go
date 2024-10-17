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

var ErrGetSpecies = errors.New("GetSpecies")
var ErrSpeciesDoesntExist = errors.New("species does not exist")

const BaseSpeciesUrl = "https://pokeapi.co/api/v2/pokemon-species/"

func GetSpecies(httpClient HttpClient, cache pokecache.Cache, speciesName string) (model.Species, error) {
	url := fmt.Sprintf("%s%s/", BaseSpeciesUrl, speciesName)

	responseBody, ok := cache.GetEntry(url)
	if !ok {
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return model.Species{}, fmt.Errorf("%w: %w: %w", ErrGetSpecies, ErrRequestInvalid, err)
		}
		response, err := httpClient.Do(request)
		if err != nil {
			return model.Species{}, fmt.Errorf("%w: %w: %w", ErrGetSpecies, ErrRequestFailed, err)
		}
		if response.StatusCode == http.StatusNotFound {
			return model.Species{}, fmt.Errorf("%w: %w: %s", ErrGetSpecies, ErrPokemonDoesntExist, speciesName)
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return model.Species{}, fmt.Errorf("%w: %w: %w", ErrGetSpecies, ErrRequestFailed, err)
		}

		responseBody = body
	}

	var apiResponse speciesApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return model.Species{}, fmt.Errorf("%w: %w: %w", ErrGetSpecies, ErrRequestFailed, err)
	}

	return ParseSpecies(apiResponse), nil
}

type speciesApiResponse struct {
	Name        string `json:"name"`
	CaptureRate uint8  `json:"capture_rate"`
}

func ParseSpecies(rawSpecies speciesApiResponse) model.Species {
	return model.Species{
		Name:        rawSpecies.Name,
		CaptureRate: rawSpecies.CaptureRate,
	}
}
