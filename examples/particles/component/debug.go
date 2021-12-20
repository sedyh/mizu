package component

type Debug struct {
	Enabled bool
}

func NewDebug() Debug {
	return Debug{true}
}
