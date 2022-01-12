package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/particles/assets"
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct{}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)

	// Get all particles
	w.View().Each(func(e engine.Entity) {
		var root *component.Root
		var pivot *component.Pivot
		var pos *component.Pos
		var angle *component.Angle
		var scale *component.Scale
		var life *component.Life
		var gradient *component.Gradient
		var sprite *component.Sprite
		e.Get(
			&root, &pivot, &pos, &angle,
			&scale, &life, &gradient, &sprite,
		)
		if root.Enabled {
			return
		}

		// Calculate parameters for image transformation.
		iw, ih := float64(sprite.Image.Bounds().Dx())*scale.W, float64(sprite.Image.Bounds().Dy())*scale.H
		sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
		screenX, screenY := (pivot.X+pos.X)*sw, (pivot.Y+pos.Y)*sh
		theta := float64(angle.Deg) * math.Pi / 180
		age := float64(life.Current) / float64(life.Total)
		color := gradient.Colors[int(age*float64(len(gradient.Colors)-1))]
		cr, cg, cb, ca := color.RGBA()
		red, green, blue, alpha := float64(cr)/0xFFFF, float64(cg)/0xFFFF, float64(cb)/0xFFFF, float64(ca)/0xFFFF

		// Smooth fade in and out.
		startFadeIn := 0.1
		endFadeIn := 0.3
		startFadeOut := 0.5
		endFadeOut := 1.0
		if age <= endFadeIn {
			alpha = 0
			if age >= startFadeIn {
				alpha = age / (startFadeIn + endFadeIn)
			}
		} else if age >= startFadeOut {
			alpha = 1 - (age)/endFadeOut
		}

		// Draw the particle.
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale.W, scale.H)
		op.GeoM.Translate(-iw/2, -ih/2)
		op.GeoM.Rotate(theta)
		op.GeoM.Translate(screenX, screenY)
		op.ColorM.Scale(red, green, blue, alpha)
		screen.DrawImage(sprite.Image, op)
	})
}
