package commands

import (
	"errors"
	"fmt"

	"github.com/yeldiRium/learning-go-pokedex/pokeapi"
)

var ErrNextMapRequestFailed = errors.New("failed to request next map section")

func MapCommand(config *CliConfig) error {
	var result *pokeapi.AreaListResult
	var err error
	if config.nextMapUrl != nil {
		result, err = pokeapi.GetAreaListWithUrl(config.httpClient, *config.nextMapUrl)
	} else {
		result, err = pokeapi.GetAreaList(config.httpClient)
	}
	if err != nil {
		return fmt.Errorf("%w: %w", ErrNextMapRequestFailed, err)
	}

	if result.NextAreaUrl != nil {
		config.nextMapUrl = result.NextAreaUrl
		config.previousMapUrl = result.PreviousAreaUrl
	}

	for _, area := range result.Areas {
		fmt.Fprintln(config.output, area.Name)
	}

	return nil
}
