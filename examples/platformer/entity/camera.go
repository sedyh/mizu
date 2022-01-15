package entity

import "github.com/sedyh/mizu/examples/platformer/component"

// An object that stores a point and distance for focusing on other objects.

type Camera struct {
	component.Pos  // Offset for rendering.
	component.Zoom // Scale for rendering.
}
