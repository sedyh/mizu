package component

import "github.com/sedyh/mizu/examples/particles/helper"

type Life struct {
	Current, Total int
	Init           func() Life
}

func NewLife(currentMin, currentMax, totalMin, totalMax int) Life {
	init := func() Life {
		return Life{
			Current: helper.RangeInt(currentMin, currentMax),
			Total:   helper.RangeInt(totalMin, totalMax),
		}
	}
	res := init()
	res.Init = init
	return res
}
