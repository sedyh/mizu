package helper

import "math/rand"

func RangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RangeFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
