package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/bunnymark/component"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Render(w World, screen *ebiten.Image) {
	w.Each(func(e Entity) {
		pos := Get[component.Position](e)
		sprite := Get[component.Sprite](e)
		hue := Get[component.Hue](e)

		sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X*sw, pos.Y*sh)
		if *hue.Colorful {
			op.ColorM.RotateHue(hue.Value)
		}
		screen.DrawImage(sprite.Image, op)
	}, And[component.Position](), And[component.Sprite](), And[component.Hue]())
}
