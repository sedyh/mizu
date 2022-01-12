package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/tilemap/assets"
	"github.com/sedyh/mizu/examples/tilemap/component"
	"github.com/sedyh/mizu/examples/tilemap/helper/enum"
	"github.com/sedyh/mizu/examples/tilemap/helper/tilemap"
	"github.com/sedyh/mizu/pkg/engine"
)

type Building struct {
	mode enum.BuildMode
}

func NewBuilding() *Building {
	return &Building{}
}

func (b *Building) Update(w engine.World) {
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

	// Calculate cursor position
	curX, curY := ebiten.CursorPosition()
	scaledTilesize := float64(assets.Tilesize) * zoom.Value
	x, y := int((float64(curX)+camera.X)/scaledTilesize), int((float64(curY)+camera.Y)/scaledTilesize)
	xOut := x < 0 || x >= construct.Width
	yOut := y < 0 || y >= construct.Height
	if xOut || yOut {
		return
	}

	// Determine what will do with the tile
	b.mode = enum.BuildModeNone
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		b.mode = enum.BuildModeBuild
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		b.mode = enum.BuildModeDestroy
	}
	if b.mode == enum.BuildModeNone {
		return
	}

	// Get tile to modify
	tiles := w.View(component.Pos{}, component.Solid{}, component.Sprite{}).Filter()
	if len(tiles) == 0 {
		return
	}
	var solid *component.Solid
	tiles[tilemap.Index(x, y, construct.Width)].Get(&solid)

	// Modify tile
	switch b.mode {
	case enum.BuildModeBuild:
		solid.Value = true
	case enum.BuildModeDestroy:
		solid.Value = false
	}
}
