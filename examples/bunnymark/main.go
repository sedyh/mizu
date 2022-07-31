package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/profile"
	"github.com/sedyh/mizu/examples/bunnymark/scene"
	. "github.com/sedyh/mizu/pkg/engine"
)

func main() {
	defer profile.Start().Stop()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowSizeLimits(300, 200, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	rand.Seed(time.Now().UTC().UnixNano())
	if err := ebiten.RunGame(NewGame(scene.Game)); err != nil {
		log.Fatal(err)
	}
}
