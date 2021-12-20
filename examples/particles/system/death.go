package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Death struct{}

func (d *Death) Update(w engine.World) {
	view := w.View(component.Root{}, component.Life{})
	view.Each(func(e engine.Entity) {
		var root *component.Root
		var life *component.Life
		e.Get(&root, &life)

		if !root.Root && life.Current >= life.Total {
			w.RemoveEntity(e)
		}
	})
}
