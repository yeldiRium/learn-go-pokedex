package formulas

import "github.com/yeldiRium/learning-go-pokedex/model"

func CatchPokemonProbability(species model.Species) float64 {
	return float64(species.CaptureRate) / 255.0
}
