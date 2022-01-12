package component

type Construct struct {
	Width  int
	Height int
}

func NewConstruct(totalW, totalH int) Construct {
	return Construct{totalW, totalH}
}
