package component

type Zoom struct {
	Value float64
}

func NewZoom(value float64) Zoom {
	return Zoom{Value: value}
}
