package commands

import (
	"errors"
	"fmt"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
)

var ErrPreviousMapRequestFailed = errors.New("failed to request previous map section")

func MapbCommand(config *CliConfig) error {
	var result *pokeapi.AreaListResult
	var err error
	if config.previousMapUrl != nil {
		result, err = pokeapi.GetAreaListWithUrl(config.httpClient, *config.previousMapUrl)
	} else {
		result, err = pokeapi.GetAreaList(config.httpClient)
	}
	if err != nil {
		return fmt.Errorf("%w: %w", ErrPreviousMapRequestFailed, err)
	}

	config.previousMapUrl = result.PreviousAreaUrl
	config.nextMapUrl = result.NextAreaUrl

	for _, area := range result.Areas {
		fmt.Fprintln(config.output, area.Name)
	}

	return nil
}
