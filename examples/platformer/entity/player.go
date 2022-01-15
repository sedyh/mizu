package entity

import "github.com/sedyh/mizu/examples/platformer/component"

// Movable object for player.

type Player struct {
	component.Pos     // Current pos in tiles.
	component.Vel     // Current velocity in tiles.
	component.Size    // Current size in tiles.
	component.Solid   // Collision group.
	component.Sprite  // Current player image.
	component.Gravity // Current player gravity.
	component.Control // Controllable by keyboard.
}
