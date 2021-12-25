package system

import (
	"github.com/sedyh/mizu/examples/bunnymark/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Gravity struct {
	*component.Velocity
	*component.Gravity
}

func (g *Gravity) Update(_ engine.World) {
	g.Velocity.Y += g.Gravity.Value
}
