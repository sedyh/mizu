package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/bunnymark/assets"
	"github.com/sedyh/mizu/pkg/engine"
)

type Background struct{}

func (b *Background) Draw(_ engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)
}
