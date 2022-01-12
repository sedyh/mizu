package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/examples/particles/entity"
	"github.com/sedyh/mizu/pkg/engine"
)

type Emit struct {
	*component.Root
	*component.Birthrate
	*component.Pivot
	*component.Pos
	*component.Vel
	*component.Accel
	*component.Angle
	*component.Spin
	*component.Scale
	*component.Growth
	*component.Life
	*component.Gradient
	*component.Sprite
}

func (e *Emit) Update(w engine.World) {
	select {
	// Emit multiple particles per tick, as configured
	case <-e.Birthrate.Ticker.C:
		for i := 0; i < e.Birthrate.Amount; i++ {
			// Calls each setting's constructor with the original arguments
			// to create a newly generated copy and creates a particle from them.
			w.AddEntities(&entity.Particle{
				Pivot:    *e.Pivot,
				Pos:      e.Pos.Init(),
				Vel:      e.Vel.Init(),
				Accel:    e.Accel.Init(),
				Angle:    e.Angle.Init(),
				Spin:     e.Spin.Init(),
				Scale:    e.Scale.Init(),
				Growth:   e.Growth.Init(),
				Life:     e.Life.Init(),
				Gradient: e.Gradient.Init(),
				Sprite:   e.Sprite.Init(),
			})
		}
	default:
	}
}
