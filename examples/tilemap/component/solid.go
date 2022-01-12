package component

type Solid struct {
	Value bool
}

func NewSolid(value bool) Solid {
	return Solid{value}
}
