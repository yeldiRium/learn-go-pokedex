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

var ErrGetAreaList = errors.New("GetAreaList")

const BaseAreaListUrl = "https://pokeapi.co/api/v2/location-area/?limit=20"

type AreaListResult struct {
	NextAreaUrl     *string
	PreviousAreaUrl *string
	Areas           model.Areas
}

type areaListApiResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func GetAreaList(httpClient HttpClient, cache pokecache.Cache, areaUrl string) (*AreaListResult, error) {
	responseBody, ok := cache.GetEntry(areaUrl)
	if !ok {
		request, err := http.NewRequest("GET", areaUrl, nil)
		if err != nil {
			return nil, fmt.Errorf("%w: %w: %w", ErrGetAreaList, ErrRequestInvalid, err)
		}
		response, err := httpClient.Do(request)
		if err != nil {
			return nil, fmt.Errorf("%w: %w: %w", ErrGetAreaList, ErrRequestFailed, err)
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("%w: %w: %w", ErrGetAreaList, ErrRequestFailed, err)
		}

		responseBody = body
	}

	var apiResponse areaListApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: %w: %w", ErrGetAreaList, ErrRequestFailed, err)
	}

	areas := make(model.Areas, len(apiResponse.Results))
	for i, area := range apiResponse.Results {
		areas[i] = model.Area{
			Name: area.Name,
		}
	}
	return &AreaListResult{
		NextAreaUrl:     apiResponse.Next,
		PreviousAreaUrl: apiResponse.Previous,
		Areas:           areas,
	}, nil
}
