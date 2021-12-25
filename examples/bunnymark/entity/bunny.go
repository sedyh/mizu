package entity

import "github.com/sedyh/mizu/examples/bunnymark/component"

type Bunny struct {
	component.Position
	component.Velocity
	component.Gravity
	component.Sprite
	component.Hue
}
