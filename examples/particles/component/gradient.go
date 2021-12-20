package component

import "image/color"

type Gradient struct {
	Colors []color.Color
	Init   func() Gradient
}

func NewGradient(colors ...color.Color) Gradient {
	init := func() Gradient {
		return Gradient{Colors: colors}
	}
	res := init()
	res.Init = init
	return res
}
