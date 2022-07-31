package scene

import (
	"time"

	"github.com/sedyh/mizu/examples/bunnymark/assets"
	"github.com/sedyh/mizu/examples/bunnymark/component"
	"github.com/sedyh/mizu/examples/bunnymark/helper"
	"github.com/sedyh/mizu/examples/bunnymark/system"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Game(w World) {
	w.AddEntity(component.NewSettings(
		time.NewTicker(500*time.Millisecond),
		assets.Bunny,
		false,
		100,
		helper.GpuInfo(),
		helper.NewPlot(20, 60),
		helper.NewPlot(20, 60),
		helper.NewPlot(20, 60000),
	))

	w.AddSystem(system.Background)
	w.AddSystem(system.Metrics)
	w.AddSystem(system.Velocity)
	w.AddSystem(system.Gravity)
	w.AddSystem(system.Bounce)
	w.AddSystem(system.Spawn)
	w.AddSystem(system.Render)
	w.AddSystem(system.Plot)
}
