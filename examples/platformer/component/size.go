package component

type Size struct {
	W, H float64
}

func NewSizeF(x, y float64) Size {
	return Size{x, y}
}

func NewSizeI(x, y int) Size {
	return Size{float64(x), float64(y)}
}
