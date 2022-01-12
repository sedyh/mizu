package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

type Debug struct{}

func (d *Debug) Draw(_ engine.World, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"Change the emitter settings in the game scene to achieve other effects\nTPS: %.2f FPS: %.2f",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(),
	))
}
