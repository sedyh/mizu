package component

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Image *ebiten.Image
}

func NewSprite(Image *ebiten.Image) Sprite {
	return Sprite{Image: Image}
}
