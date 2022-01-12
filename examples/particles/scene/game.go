package scene

import (
	"time"

	"github.com/sedyh/mizu/examples/particles/assets"
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/examples/particles/entity"
	"github.com/sedyh/mizu/examples/particles/system"
	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Root{}, component.Birthrate{}, component.Pivot{},
		component.Pos{}, component.Vel{}, component.Accel{},
		component.Angle{}, component.Spin{}, component.Scale{},
		component.Growth{}, component.Life{}, component.Gradient{},
		component.Sprite{},
	)

	w.AddSystems(
		&system.Emit{}, &system.Velocity{}, &system.Acceleration{},
		&system.Spin{}, &system.Growth{}, &system.Age{},
		&system.Death{}, &system.Render{}, &system.Debug{},
	)

	// Fire and waterfall emitters.
	w.AddEntities(
		&entity.Emitter{
			Root:      component.NewRoot(),
			Birthrate: component.NewBirthrate(2, 20*time.Millisecond),
			Pivot:     component.NewPivot(0.35, 0.5),
			Pos:       component.NewPos(-0.02, 0.02, -0.005, 0.005),
			Vel:       component.NewVel(0.01, 0.01, 80, 100),
			Accel:     component.NewAccel(-0.0001, -0.0001, 40, 90),
			Angle:     component.NewAngle(-90, 90),
			Spin:      component.NewSpin(1, 2),
			Scale:     component.NewScale(0.2, 0.2),
			Growth:    component.NewGrowth(-0.007, -0.003),
			Life:      component.NewLife(0, 10, 30, 40),
			Gradient:  component.NewGradient(assets.FireGradient...),
			Sprite:    component.NewSprite(assets.FireA, assets.FireB),
		},
		&entity.Emitter{
			Root:      component.NewRoot(),
			Birthrate: component.NewBirthrate(25, 20*time.Millisecond),
			Pivot:     component.NewPivot(0.65, 0.35),
			Pos:       component.NewPos(-0.08, 0.08, 0, 0),
			Vel:       component.NewVel(-0.0025, -0.0025, 80, 100),
			Accel:     component.NewAccel(-0.0002, -0.0002, 90, 90),
			Angle:     component.NewAngle(0, 360),
			Spin:      component.NewSpin(-10, 10),
			Scale:     component.NewScale(0.3, 0.3),
			Growth:    component.NewGrowth(0, 0),
			Life:      component.NewLife(0, 10, 50, 50),
			Gradient:  component.NewGradient(assets.WaterGradient...),
			Sprite:    component.NewSprite(assets.Water),
		},
	)
}
