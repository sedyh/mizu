package component

type Pos struct {
	X, Y float64
}

func NewPosI(x, y int) Pos {
	return Pos{float64(x), float64(y)}
}
