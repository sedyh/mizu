package entity

import "github.com/sedyh/mizu/examples/platformer/component"

// The object that takes up space for the tile.

type Tile struct {
	component.Pos    // Index for the tile.
	component.Vel    // Vel for the tile.
	component.Size   // Current size in tiles
	component.Solid  // Solid or empty tile.
	component.Sprite // Current tile image.
}
