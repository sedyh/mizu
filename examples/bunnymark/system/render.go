package system

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/sedyh/mizu/examples/bunnymark/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	*component.Position
	*component.Sprite
	*component.Hue
}

func (r *Render) Draw(_ engine.World, screen *ebiten.Image) {
	sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.Position.X*sw, r.Position.Y*sh)
	if *r.Hue.Colorful {
		op.ColorM.RotateHue(r.Hue.Value)
	}
	screen.DrawImage(r.Sprite.Image, op)
}
