package component

type Pivot struct {
	X, Y float64
}

func NewPivot(x, y float64) Pivot {
	return Pivot{x, y}
}
