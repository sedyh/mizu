package system

import (
	"github.com/sedyh/mizu/examples/particles/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Age struct {
	*component.Root
	*component.Life
}

func (v *Age) Update(_ engine.World) {
	if v.Root.Root {
		return
	}

	if v.Current >= v.Total {
		return
	}
	v.Current++
}
