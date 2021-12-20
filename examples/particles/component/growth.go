package component

import "github.com/sedyh/mizu/examples/particles/helper"

type Growth struct {
	Speed float64
	Init  func() Growth
}

func NewGrowth(min, max float64) Growth {
	init := func() Growth {
		return Growth{Speed: helper.RangeFloat(min, max)}
	}
	res := init()
	res.Init = init
	return res
}
