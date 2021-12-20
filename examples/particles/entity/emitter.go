package entity

import "github.com/sedyh/mizu/examples/particles/component"

type Emitter struct {
	component.Root
	component.Birthrate
	component.Pivot
	component.Pos
	component.Vel
	component.Accel
	component.Angle
	component.Spin
	component.Scale
	component.Growth
	component.Life
	component.Gradient
	component.Sprite
}
