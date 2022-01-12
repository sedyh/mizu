package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/particles/helper"
)

type Sprite struct {
	Image *ebiten.Image
	Init  func() Sprite
}

func NewSprite(images ...*ebiten.Image) Sprite {
	init := func() Sprite {
		index := helper.RangeInt(0, len(images))
		return Sprite{Image: images[index]}
	}
	res := init()
	res.Init = init
	return res
}
