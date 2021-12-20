package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/sedyh/mizu/examples/particles/scene"
	"github.com/sedyh/mizu/pkg/engine"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
		log.Fatal(err)
	}
}
