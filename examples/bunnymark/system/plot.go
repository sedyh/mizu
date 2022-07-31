package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/sedyh/mizu/examples/bunnymark/component"
	. "github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func Plot(w World, screen *ebiten.Image) {
	e, ok := w.Get(And[component.Settings]())
	if !ok {
		return
	}
	s := Get[component.Settings](e)

	str := fmt.Sprintf(
		"GPU: %s\nTPS: %.2f, FPS: %.2f, Objects: %.f\nBatching: %t, Amount: %d\nResolution: %dx%d",
		s.Gpu, s.Tps.Last(), s.Fps.Last(), s.Objects.Last(),
		!s.Colorful, s.Amount,
		w.Bounds().Dx(), w.Bounds().Dy(),
	)

	rect := text.BoundString(basicfont.Face7x13, str)
	width, height := float64(rect.Dx()), float64(rect.Dy())

	padding := 20.0
	rectW, rectH := width+padding, height+padding
	plotW, plotH := 100.0, 40.0

	ebitenutil.DrawRect(screen, 0, 0, rectW, rectH, color.RGBA{A: 128})
	text.Draw(screen, str, basicfont.Face7x13, int(padding)/2, 10+int(padding)/2, colornames.White)

	s.Tps.Draw(screen, 0, padding+rectH, plotW, plotH)
	s.Fps.Draw(screen, 0, padding+rectH*2, plotW, plotH)
	s.Objects.Draw(screen, 0, padding+rectH*3, plotW, plotH)
}
