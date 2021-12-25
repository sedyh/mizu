package component

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/sedyh/mizu/examples/bunnymark/helper"
)

type Settings struct {
	Ticker   *time.Ticker
	Sprite   *ebiten.Image
	Colorful bool
	Amount   int
	Gpu      string
	Tps      *helper.Plot
	Fps      *helper.Plot
	Objects  *helper.Plot
}
