package commands

import (
	"errors"
	"fmt"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
)

var ErrPreviousMapRequestFailed = errors.New("failed to request previous map section")
var ErrBeginningOfAreasReached = errors.New("beginning of areas reached, can't go further back")

func MapbCommand(config *CliConfig) error {
	if config.previousMapUrl == nil {
		return ErrBeginningOfAreasReached
	}

	result, err := pokeapi.GetAreaList(config.httpClient, config.cache, *config.previousMapUrl)
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
