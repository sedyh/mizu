package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Acceleration struct {
	*component.Root
	*component.Vel
	*component.Accel
}

func (a *Acceleration) Update(_ engine.World) {
	if a.Root.Root {
		return
	}

	a.Vel.L += a.Accel.O
	a.Vel.M += a.Accel.P
}
