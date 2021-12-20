package component

type Root struct {
	Root bool
}

func NewRoot() Root {
	return Root{true}
}

func NewNonRoot() Root {
	return Root{false}
}
