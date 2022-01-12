package entity

import "github.com/sedyh/mizu/examples/tilemap/component"

// The object that takes up space for the tile.

type Tile struct {
	component.Pos    // Index for the tile.
	component.Solid  // Solid or empty tile.
	component.Sprite // Current tile image.
}
