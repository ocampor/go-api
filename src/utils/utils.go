package utils

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"math"
)

func floatComparison(a float64, b float64) float64 {
	if a < b {
		return math.Abs(a / b)
	} else {
		return math.Abs(b / a)
	}
}

func FloatSimilarity(a *float64, b *float64) *float64 {
	if a == nil || b == nil {
		return nil
	}

	var result = floatComparison(*a, *b)
	return &result
}

func IntegerSimilarity(a *int, b *int) *float64 {
	if a == nil || b == nil {
		return nil
	}

	var result = BoolToFloat64(*a == *b)
	return &result
}

func BoolToFloat64(a bool) float64 {
	if a {
		return 1.0
	} else {
		return 0.0
	}
}

func StringSimilarity(a *string, b *string) *float64 {
	if a == nil || b == nil {
		return nil
	}

	var result = levenshtein.RatioForStrings([]rune(*a), []rune(*b), levenshtein.DefaultOptions)
	return &result
}

func OverallSimilarity(
	locationSimilarity *float64,
	bedroomSimilarity *float64,
	unitSizeSimilarity *float64,
	unitNumberSimilarity *float64) *float64 {

	if locationSimilarity == nil || locationSimilarity == nil || unitSizeSimilarity == nil || unitNumberSimilarity == nil {
		return nil
	}

	var result = (*locationSimilarity +
		*bedroomSimilarity +
		*unitSizeSimilarity +
		*unitNumberSimilarity) / 4

	return &result;
}
