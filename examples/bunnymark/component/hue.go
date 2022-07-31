package component

type Hue struct {
	Colorful *bool
	Value    float64
}

func NewHue(colorful *bool, value float64) *Hue {
	return &Hue{colorful, value}
}
