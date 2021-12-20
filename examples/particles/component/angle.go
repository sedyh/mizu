package component

import (
	"github.com/sedyh/mizu/examples/particles/helper"
)

type Angle struct {
	Deg  int
	Init func() Angle
}

func NewAngle(min, max int) Angle {
	init := func() Angle {
		return Angle{Deg: helper.RangeInt(min, max)}
	}
	res := init()
	res.Init = init
	return res
}
