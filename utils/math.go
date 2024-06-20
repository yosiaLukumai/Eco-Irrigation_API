package utils

import "math"

func RoundToTwoDecimalPlaces(f float32) float32 {
	return float32(math.Round(float64(f)*100) / 100)
}
