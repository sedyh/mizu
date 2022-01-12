package system

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/sedyh/mizu/examples/bunnymark/component"
	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Metrics struct {
	*component.Settings
}

func (m *Metrics) Update(w engine.World) {
	select {
	case <-m.Ticker.C:
		m.Objects.Update(float64(w.Entities() - 1))
		m.Tps.Update(ebiten.CurrentTPS())
		m.Fps.Update(ebiten.CurrentFPS())
	default:
	}
}

func (m *Metrics) Draw(w engine.World, screen *ebiten.Image) {
	str := fmt.Sprintf(
		"GPU: %s\nTPS: %.2f, FPS: %.2f, Objects: %.f\nBatching: %t, Amount: %d\nResolution: %dx%d",
		m.Gpu, m.Tps.Last(), m.Fps.Last(), m.Objects.Last(),
		!m.Colorful, m.Amount,
		w.Bounds().Dx(), w.Bounds().Dy(),
	)

	rect := text.BoundString(basicfont.Face7x13, str)
	width, height := float64(rect.Dx()), float64(rect.Dy())

	padding := 20.0
	rectW, rectH := width+padding, height+padding
	plotW, plotH := 100.0, 40.0

	ebitenutil.DrawRect(screen, 0, 0, rectW, rectH, color.RGBA{A: 128})
	text.Draw(screen, str, basicfont.Face7x13, int(padding)/2, 10+int(padding)/2, colornames.White)

	m.Tps.Draw(screen, 0, padding+rectH, plotW, plotH)
	m.Fps.Draw(screen, 0, padding+rectH*2, plotW, plotH)
	m.Objects.Draw(screen, 0, padding+rectH*3, plotW, plotH)
}
