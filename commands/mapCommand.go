package commands

import (
	"errors"
	"fmt"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
)

var ErrNextMapRequestFailed = errors.New("failed to request next map section")
var ErrEndOfAreasReached = errors.New("end of areas reached, can't go further")

func MapCommand(config *CliConfig) error {
	if config.nextMapUrl == nil {
		return ErrEndOfAreasReached
	}

	result, err := pokeapi.GetAreaList(config.httpClient, config.cache, *config.nextMapUrl)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNextMapRequestFailed, err)
	}

	config.nextMapUrl = result.NextAreaUrl
	config.previousMapUrl = result.PreviousAreaUrl

	for _, area := range result.Areas {
		fmt.Fprintln(config.output, area.Name)
	}

	return nil
}
