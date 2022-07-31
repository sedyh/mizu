package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/bunnymark/component"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Metrics(w World) {
	e, ok := w.Get(And[component.Settings]())
	if !ok {
		return
	}
	s := Get[component.Settings](e)

	select {
	case <-s.Ticker.C:
		s.Objects.Update(float64(w.Entities() - 1))
		s.Tps.Update(ebiten.CurrentTPS())
		s.Fps.Update(ebiten.CurrentFPS())
	default:
	}
}
