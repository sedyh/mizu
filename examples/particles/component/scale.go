package component

import (
	"github.com/sedyh/mizu/examples/particles/helper"
)

type Scale struct {
	W, H float64
	Init func() Scale
}

func NewScale(min, max float64) Scale {
	init := func() Scale {
		wh := helper.RangeFloat(min, max)

		return Scale{
			W: wh,
			H: wh,
		}
	}
	res := init()
	res.Init = init
	return res
}
