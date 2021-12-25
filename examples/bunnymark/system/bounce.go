package system

import (
	"github.com/sedyh/mizu/examples/bunnymark/component"
	"github.com/sedyh/mizu/examples/bunnymark/helper"
	"github.com/sedyh/mizu/pkg/engine"
)

type Bounce struct {
	*component.Position
	*component.Velocity
	*component.Sprite
}

func (b *Bounce) Update(w engine.World) {
	sw, sh := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())
	iw, ih := float64(b.Sprite.Image.Bounds().Dx()), float64(b.Sprite.Image.Bounds().Dy())
	relW, relH := iw/sw, ih/sh
	if b.Position.X+relW > 1 {
		b.Velocity.X *= -1
		b.Position.X = 1 - relW
	}
	if b.Position.X < 0 {
		b.Velocity.X *= -1
		b.Position.X = 0
	}
	if b.Position.Y+relH > 1 {
		b.Velocity.Y *= -0.85
		b.Position.Y = 1 - relH
		if helper.Chance(0.5) {
			b.Velocity.Y -= helper.RangeFloat(0, 0.009)
		}
	}
	if b.Position.Y < 0 {
		b.Velocity.Y = 0
		b.Position.Y = 0
	}
}
