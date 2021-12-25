package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/sedyh/mizu/pkg/engine"
)

type Background struct{}

func (b *Background) Draw(_ engine.World, screen *ebiten.Image) {
	screen.Fill(colornames.Whitesmoke)
}
