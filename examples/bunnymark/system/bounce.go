package system

import (
	"github.com/sedyh/mizu/examples/bunnymark/component"
	"github.com/sedyh/mizu/examples/bunnymark/helper"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Bounce(w World) {
	w.Each(func(e Entity) {
		pos := Get[component.Position](e)
		vel := Get[component.Velocity](e)
		sprite := Get[component.Sprite](e)

		sw, sh := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())
		iw, ih := float64(sprite.Image.Bounds().Dx()), float64(sprite.Image.Bounds().Dy())
		relW, relH := iw/sw, ih/sh
		if pos.X+relW > 1 {
			vel.X *= -1
			pos.X = 1 - relW
		}
		if pos.X < 0 {
			vel.X *= -1
			pos.X = 0
		}
		if pos.Y+relH > 1 {
			vel.Y *= -0.85
			pos.Y = 1 - relH
			if helper.Chance(0.5) {
				vel.Y -= helper.RangeFloat(0, 0.009)
			}
		}
		if pos.Y < 0 {
			vel.Y = 0
			pos.Y = 0
		}

	}, And[component.Position](), And[component.Velocity](), And[component.Sprite]())
}
