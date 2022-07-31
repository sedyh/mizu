package component

type Velocity struct {
	X, Y float64
}

func NewVelocity(x, y float64) *Velocity {
	return &Velocity{x, y}
}
