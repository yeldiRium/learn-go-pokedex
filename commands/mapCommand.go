package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrNextMapRequestInvalid = errors.New("failed to create next map request")
var ErrNextMapRequestFailed = errors.New("failed to request next map section")

type MapResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func MapCommand(state *CliConfig) error {
	var nextUrl string
	if state.nextMapUrl != nil {
		nextUrl = *state.nextMapUrl
	} else {
		nextUrl = "https://pokeapi.co/api/v2/location-area/?limit=20"
	}

	request, err := http.NewRequest("GET", nextUrl, nil)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNextMapRequestInvalid, err)
	}
	response, err := state.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNextMapRequestFailed, err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNextMapRequestFailed, err)
	}

	var mapResponse MapResponse
	err = json.Unmarshal(body, &mapResponse)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNextMapRequestFailed, err)
	}

	state.nextMapUrl = mapResponse.Next
	state.previousMapUrl = mapResponse.Previous

	for _, area := range mapResponse.Results {
		fmt.Fprintln(state.output, area.Name)
	}

	return nil
}
