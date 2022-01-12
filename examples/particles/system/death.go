package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Death struct{}

func (d *Death) Update(w engine.World) {
	view := w.View(component.Root{}, component.Life{})
	for _, e := range view.Filter() {
		var root *component.Root
		var life *component.Life
		e.Get(&root, &life)

		// Remove dead particles.
		if !root.Enabled && life.Current >= life.Total {
			w.RemoveEntity(e)
		}
	}
}
