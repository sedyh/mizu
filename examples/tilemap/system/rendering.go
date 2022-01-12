package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/tilemap/assets"
	"github.com/sedyh/mizu/examples/tilemap/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Rendering struct {
	offscreen *ebiten.Image
}

func NewRendering(w, h int) *Rendering {
	return &Rendering{
		offscreen: ebiten.NewImage(w*assets.Tilesize, h*assets.Tilesize),
	}
}

func (t *Rendering) Draw(w engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)

	// Get auxiliary objects
	cameraEntity, found := w.View(component.Pos{}, component.Zoom{}).Get()
	if !found {
		return
	}
	var camera *component.Pos
	var zoom *component.Zoom
	cameraEntity.Get(&camera, &zoom)

	// Draw tiles to the offscreen
	view := w.View(component.Solid{}, component.Pos{}, component.Sprite{})
	view.Each(func(e engine.Entity) {
		var solid *component.Solid
		var pos *component.Pos
		var sprite *component.Sprite
		e.Get(&solid, &pos, &sprite)

		if !solid.Value {
			return
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X*float64(assets.Tilesize), pos.Y*float64(assets.Tilesize))
		t.offscreen.DrawImage(sprite.Image, op)
	})

	// Draw the offscreen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(zoom.Value, zoom.Value)
	op.GeoM.Translate(-camera.X, -camera.Y)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(t.offscreen, op)
	t.offscreen.Clear()
}
