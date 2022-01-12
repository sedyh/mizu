package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Growth struct {
	*component.Root
	*component.Scale
	*component.Growth
}

func (g *Growth) Update(_ engine.World) {
	if g.Root.Enabled {
		return
	}

	// Increases the size until it goes to zero
	old := component.Scale{W: g.W, H: g.H}
	now := component.Scale{W: g.W + g.Speed, H: g.H + g.Speed}
	if scaleShouldBeStopped(old, now) {
		g.Growth.Speed = 0
		return
	}
	g.W, g.H = now.W, now.H
}

func scaleShouldBeStopped(old, now component.Scale) bool {
	return scaleIsInverted(old.W, now.W) || scaleIsInverted(old.H, now.H)
}

func scaleIsInverted(old, now float64) bool {
	return (old < 0 && now > 0) || (old > 0 && now < 0)
}
