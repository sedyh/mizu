package entity

import "github.com/sedyh/mizu/examples/bunnymark/component"

// Bunny object.

type Bunny struct {
	component.Position // Current bunny position.
	component.Velocity // Current bunny velocity.
	component.Gravity  // Current bunny gravity.
	component.Sprite   // Current bunny image.
	component.Hue      // Current bunny hue.
}
