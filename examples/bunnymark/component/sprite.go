package component

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Image *ebiten.Image
}

func NewSprite(image *ebiten.Image) *Sprite {
	return &Sprite{image}
}
