package component

type Vel struct {
	L, M float64
}

func NewVel(l, m float64) Vel {
	return Vel{l, m}
}
