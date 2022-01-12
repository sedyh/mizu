package component

type Pos struct {
	X, Y float64
}

func NewPosF(x, y float64) Pos {
	return Pos{x, y}
}

func NewPosI(x, y int) Pos {
	return Pos{float64(x), float64(y)}
}

func (p Pos) Add(n Pos) Pos {
	return Pos{X: p.X + n.X, Y: p.Y + n.Y}
}

func (p Pos) Sub(n Pos) Pos {
	return Pos{X: p.X - n.X, Y: p.Y - n.Y}
}

func (p Pos) Div(n Pos) Pos {
	return Pos{X: p.X / n.X, Y: p.Y / n.Y}
}

func (p Pos) Trunc() Pos {
	return Pos{X: float64(int(p.X)), Y: float64(int(p.Y))}
}
