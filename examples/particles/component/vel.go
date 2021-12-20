package component

import (
	"math"

	"github.com/sedyh/mizu/examples/particles/helper"
)

type Vel struct {
	L, M float64
	Init func() Vel
}

func NewVel(speedMin, speedMax, directionMin, directionMax float64) Vel {
	init := func() Vel {
		speed := helper.RangeFloat(speedMin, speedMax)
		direction := helper.RangeFloat(directionMin, directionMax)

		return Vel{
			L: math.Cos(-direction*math.Pi/180) * speed,
			M: math.Sin(-direction*math.Pi/180) * speed,
		}
	}
	res := init()
	res.Init = init
	return res
}
