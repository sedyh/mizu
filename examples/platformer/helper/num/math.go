package num

import "math"

// Equal roughly compares a and b width delta.
func Equal(a, b float64, delta float64) bool {
	return math.Abs(a-b) <= delta
}

// Lerp returns a value between a and b at a point on a linear scale.
func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}
