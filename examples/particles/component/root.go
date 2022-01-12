package component

type Root struct {
	Enabled bool
}

func NewRoot() Root {
	return Root{true}
}

func NewNonRoot() Root {
	return Root{false}
}
