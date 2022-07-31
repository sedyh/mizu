package system

import (
	"github.com/sedyh/mizu/examples/bunnymark/component"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Gravity(w World) {
	w.Each(func(e Entity) {
		vel := Get[component.Velocity](e)
		gravity := Get[component.Gravity](e)

		vel.Y += gravity.Value
	}, And[component.Velocity](), And[component.Gravity]())
}
