package system

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

type Debug struct{}

func NewDebug() *Debug {
	return &Debug{}
}

func (d *Debug) Draw(_ engine.World, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"Use WASD to control the player, space to jump, - and + to zoom in and out.\nTPS: %.2f FPS: %.2f",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(),
	))
}
