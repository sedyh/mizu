package component

import "github.com/sedyh/mizu/examples/particles/helper"

type Spin struct {
	Speed int
	Init  func() Spin
}

func NewSpin(min, max int) Spin {
	init := func() Spin {
		return Spin{Speed: helper.RangeInt(min, max)}
	}
	res := init()
	res.Init = init
	return res
}
