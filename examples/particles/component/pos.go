package component

import "github.com/sedyh/mizu/examples/particles/helper"

type Pos struct {
	X, Y float64
	Init func() Pos
}

func NewPos(xMin, xMax, yMin, yMax float64) Pos {
	init := func() Pos {
		return Pos{
			X: helper.RangeFloat(xMin, xMax),
			Y: helper.RangeFloat(yMin, yMax),
		}
	}
	res := init()
	res.Init = init
	return res
}
