package scene

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/examples/particles/entity"
	"github.com/sedyh/mizu/examples/particles/helper"
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
		component.Sprite{}, component.Debug{},
	)

	w.AddSystems(
		&system.Emit{}, &system.Velocity{}, &system.Acceleration{},
		&system.Spin{}, &system.Growth{}, &system.Age{},
		&system.Death{}, &system.Render{}, &system.Debug{},
	)

	w.AddEntities(
		&entity.Emitter{
			Root:      component.NewRoot(),
			Birthrate: component.NewBirthrate(2, 20*time.Millisecond),
			Pivot:     component.NewPivot(0.5, 0.5),
			Pos:       component.NewPos(-0.001, 0.001, -0.001, 0.001),
			Vel:       component.NewVel(0.005, 0.005, 80, 100),
			Accel:     component.NewAccel(-0.0001, -0.0001, 40, 90),
			Angle:     component.NewAngle(-90, 90),
			Spin:      component.NewSpin(1, 2),
			Scale:     component.NewScale(0.8, 1.2),
			Growth:    component.NewGrowth(-0.01, -0.02),
			Life:      component.NewLife(0, 10, 50, 60),
			Gradient:  component.NewGradient(helper.Gradient()...),
			Sprite:    component.NewSprite(ebiten.NewImageFromImage(helper.Checkerboard(16, 16, 2))),
		},
		&entity.Debugger{
			Debug: component.NewDebug(),
		},
	)
}
