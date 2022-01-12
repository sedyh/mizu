package system

import (
	"math"

	"github.com/sedyh/mizu/examples/tilemap/assets"
	"github.com/sedyh/mizu/examples/tilemap/component"
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
	var construct *component.Construct
	var camera *component.Pos
	var zoom *component.Zoom
	constructEntity.Get(&construct)
	cameraEntity.Get(&camera, &zoom)

	// Get the size of the map and screen in pixels
	tilemapWidth := float64(assets.Tilesize * construct.Width)
	tilemapHeight := float64(assets.Tilesize * construct.Height)
	screenWidth, screenHeight := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())
	minScreenSize := math.Min(screenWidth, screenHeight)

	// Change the camera zoom to fit the tile map
	okWidth := tilemapWidth * zoom.Value
	okHeight := tilemapHeight * zoom.Value
	maxOkSize := math.Max(okWidth, okHeight)
	for maxOkSize > minScreenSize {
		zoom.Value -= 0.0005
		okWidth = tilemapWidth * zoom.Value
		okHeight = tilemapHeight * zoom.Value
		maxOkSize = math.Max(okWidth, okHeight)
	}
	for maxOkSize < minScreenSize {
		zoom.Value += 0.0005
		okWidth = tilemapWidth * zoom.Value
		okHeight = tilemapHeight * zoom.Value
		maxOkSize = math.Max(okWidth, okHeight)
	}

	// Move the camera to fit the tile map
	scaledTilesize := float64(assets.Tilesize) * zoom.Value
	camera.X = (scaledTilesize*float64(construct.Width) - screenWidth) / 2
	camera.Y = (scaledTilesize*float64(construct.Height) - screenHeight) / 2
}
