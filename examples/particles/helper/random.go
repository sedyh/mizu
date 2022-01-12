package helper

import (
	"math/rand"
)

func RangeInt(min, max int) int {
	if min == max {
		return min
	}
	return min + rand.Intn(max-min)
}

func RangeFloat(min, max float64) float64 {
	if min == max {
		return min
	}
	return min + rand.Float64()*(max-min)
}
