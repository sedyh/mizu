package entity

import "github.com/sedyh/mizu/examples/particles/component"

type Particle struct {
	component.Root
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
