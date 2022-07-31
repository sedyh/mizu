package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/bunnymark/assets"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Background(_ World, screen *ebiten.Image) {
	screen.Fill(assets.Background)
}
