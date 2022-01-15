package component

type Gravity struct {
	Value float64
}

func NewGravity(value float64) Gravity {
	return Gravity{value}
}
