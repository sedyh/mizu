package component

type Construct struct {
	Width  int
	Height int
	Level  string
}

func NewConstruct(width, height int, level string) Construct {
	return Construct{width, height, level}
}
