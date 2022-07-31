package system

import (
	"github.com/sedyh/mizu/examples/bunnymark/component"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Velocity(w World) {
	w.Each(func(e Entity) {
		pos := Get[component.Position](e)
		vel := Get[component.Velocity](e)

		pos.X += vel.X
		pos.Y += vel.Y
	}, And[component.Position](), And[component.Velocity]())
}
