package utils

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"math"
)

func FloatSimilarity(a float64, b float64) float64 {
	if a < b {
		return math.Abs(a / b)
	} else {
		return math.Abs(b / a)
	}
}

func BoolToFloat64(a bool) float64 {
	if a {
		return 1.0
	} else {
		return 0.0
	}
}

func StringSimilarity(a string, b string) float64 {
	return levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
}

func OverallSimilarity(
	locationSimilarity float64,
	bedroomSimilarity float64,
	unitSizeSimilarity float64,
	unitNumberSimilarity float64) float64 {
	return (locationSimilarity +
		bedroomSimilarity +
		unitSizeSimilarity +
		unitNumberSimilarity) / 4
}
