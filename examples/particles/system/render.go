package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	*component.Root
	*component.Pivot
	*component.Pos
	*component.Angle
	*component.Scale
	*component.Life
	*component.Gradient
	*component.Sprite
}

func (r *Render) Draw(_ engine.World, screen *ebiten.Image) {
	if r.Root.Root {
		return
	}

	imageWidth, imageHeight := float64(r.Sprite.Image.Bounds().Dx()), float64(r.Sprite.Image.Bounds().Dy())
	screenWidth, screenHeight := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	screenX, screenY := (*r.Pivot.X+r.Pos.X)*screenWidth, (*r.Pivot.Y+r.Pos.Y)*screenHeight
	angle := float64(r.Angle.Deg) * math.Pi / 180
	age := float64(r.Life.Current) / float64(r.Life.Total)
	color := r.Gradient.Colors[int(age*float64(len(r.Gradient.Colors)-1))]
	cr, cg, cb, ca := color.RGBA()
	red, green, blue, alpha := float64(cr)/0xFFFF, float64(cg)/0xFFFF, float64(cb)/0xFFFF, float64(ca)/0xFFFF

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-imageWidth/2, -imageHeight/2)
	op.GeoM.Rotate(angle)
	op.GeoM.Scale(r.Scale.W, r.Scale.H)
	op.GeoM.Translate(imageWidth/2, imageHeight/2)
	op.GeoM.Translate(screenX, screenY)
	op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	op.ColorM.Translate(red, green, blue, alpha)
	screen.DrawImage(r.Sprite.Image, op)
}
