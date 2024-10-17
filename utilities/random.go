package utilities

import "math/rand"

func RandomWithProbability(probability float64) bool {
	if probability <= 0 {
		return false
	}
	if probability >= 1 {
		return true
	}

	return probability > rand.Float64()
}
