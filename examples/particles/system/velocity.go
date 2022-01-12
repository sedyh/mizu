package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Velocity struct {
	*component.Root
	*component.Pos
	*component.Vel
}

func (v *Velocity) Update(_ engine.World) {
	if v.Root.Enabled {
		return
	}

	// Increase position.
	v.Pos.X += v.Vel.L
	v.Pos.Y += v.Vel.M
}
