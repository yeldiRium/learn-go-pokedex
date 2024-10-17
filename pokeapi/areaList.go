package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/yeldiRium/learning-go-pokedex/pokecache"
)

var ErrAreaListRequestInvalid = errors.New("failed to create area list request")
var ErrAreaListRequestFailed = errors.New("failed to request area list section")

const BaseAreaListUrl = "https://pokeapi.co/api/v2/location-area/?limit=20"

type Area struct {
	Name string
}
type AreaListResult struct {
	NextAreaUrl     *string
	PreviousAreaUrl *string
	Areas           []Area
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
			return nil, fmt.Errorf("%w: %w", ErrAreaListRequestInvalid, err)
		}
		response, err := httpClient.Do(request)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrAreaListRequestFailed, err)
		}

		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrAreaListRequestFailed, err)
		}

		responseBody = body
	}

	var apiResponse areaListApiResponse
	err := json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrAreaListRequestFailed, err)
	}

	areas := make([]Area, len(apiResponse.Results))
	for i, area := range apiResponse.Results {
		areas[i] = Area{
			Name: area.Name,
		}
	}
	return &AreaListResult{
		NextAreaUrl:     apiResponse.Next,
		PreviousAreaUrl: apiResponse.Previous,
		Areas:           areas,
	}, nil
}
