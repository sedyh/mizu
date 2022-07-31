package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/examples/bunnymark/component"
	"github.com/sedyh/mizu/examples/bunnymark/helper"
	. "github.com/sedyh/mizu/pkg/engine"
)

func Spawn(w World) {
	e, ok := w.Get(And[component.Settings]())
	if !ok {
		return
	}
	s := Get[component.Settings](e)

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		addBunnies(w, s)
	}

	if ids := ebiten.AppendTouchIDs(nil); len(ids) > 0 {
		addBunnies(w, s)
	}

	if _, offset := ebiten.Wheel(); offset != 0 {
		s.Amount += int(offset * 10)
		if s.Amount < 0 {
			s.Amount = 0
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.Colorful = !s.Colorful
	}
}

func addBunnies(w World, s *component.Settings) {
	for i := 0; i < s.Amount; i++ {
		w.AddEntity(
			component.NewPosition(float64(w.Entities()%2), 0),
			component.NewVelocity(helper.RangeFloat(0, 0.005), helper.RangeFloat(0.0025, 0.005)),
			component.NewSprite(s.Sprite),
			component.NewHue(&s.Colorful, helper.RangeFloat(0, 2*math.Pi)),
			component.NewGravity(0.00095),
		)
	}
}
