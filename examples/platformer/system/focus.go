package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/platformer/assets"
	"github.com/sedyh/mizu/examples/platformer/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Focus struct{}

func NewFocus() *Focus {
	return &Focus{}
}

func (f *Focus) Update(w engine.World) {
	// Get auxiliary objects
	constructEntity, found := w.View(component.Construct{}).Get()
	if !found {
		return
	}
	cameraEntity, found := w.View(component.Pos{}, component.Zoom{}).Get()
	if !found {
		return
	}
	playerEntity, found := w.View(component.Pos{}, component.Vel{}, component.Sprite{}, component.Control{}).Get()
	if !found {
		return
	}
	var player *component.Pos
	var construct *component.Construct
	var camera *component.Pos
	var zoom *component.Zoom
	constructEntity.Get(&construct)
	cameraEntity.Get(&camera, &zoom)
	playerEntity.Get(&player)

	// Apply zoom.
	if ebiten.IsKeyPressed(ebiten.KeyEqual) {
		zoom.Value += 0.01
	} else if ebiten.IsKeyPressed(ebiten.KeyMinus) {
		zoom.Value -= 0.01
	}

	// Move the camera to follow the player.
	x, y := player.X*float64(assets.Tilesize)*zoom.Value, player.Y*float64(assets.Tilesize)*zoom.Value
	screenWidth, screenHeight := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())
	camera.X = x - screenWidth/2 + assets.PlayerW/2*zoom.Value
	camera.Y = y - screenHeight/2 + assets.PlayerH/2*zoom.Value
}
