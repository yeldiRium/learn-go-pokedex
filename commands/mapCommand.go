package commands

import (
	"errors"
	"fmt"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
)

var ErrNextMapRequestFailed = errors.New("failed to request next map section")

func MapCommand(state *CliConfig) error {
	var result *pokeapi.AreaListResult
	var err error
	if state.nextMapUrl != nil {
		result, err = pokeapi.GetAreaListWithUrl(state.httpClient, *state.nextMapUrl)
	} else {
		result, err = pokeapi.GetAreaList(state.httpClient)
	}
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNextMapRequestFailed, err)
	}

	if result.NextAreaUrl != nil {
		state.nextMapUrl = result.NextAreaUrl
		state.previousMapUrl = result.PreviousAreaUrl
	}

	for _, area := range result.Areas {
		fmt.Fprintln(state.output, area.Name)
	}

	return nil
}
