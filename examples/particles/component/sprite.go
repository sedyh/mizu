package component

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Image *ebiten.Image
	Init  func() Sprite
}

func NewSprite(image *ebiten.Image) Sprite {
	init := func() Sprite {
		return Sprite{Image: image}
	}
	res := init()
	res.Init = init
	return res
}
