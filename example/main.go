package main

import (
	"fmt"
	"math/rand"
	"time"

	"mizu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type Pos struct {
	X, Y float64
}
type Vel struct {
	L, M float64
}
type Rad struct {
	Value float64
}

type Ball struct {
	Pos
	Vel
	Rad
}

type Velocity struct {
	*Pos
	*Vel
}

func (v *Velocity) Update(w mizu.World) {
	v.Pos.X += v.Vel.L
	v.Pos.Y += v.Vel.M
	if v.Pos.X < 0 ||
		v.Pos.Y < 0 ||
		v.Pos.X >= float64(w.Bounds().Dx()) ||
		v.Pos.Y >= float64(w.Bounds().Dy()) {
		w.ChangeScene(&Menu{})
	}
}

type Render struct {
	*Pos
	*Rad
}

func (r *Render) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %.2f FPS: %.2f",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(),
	))
	ebitenutil.DrawRect(
		screen,
		r.Pos.X-r.Rad.Value/2, r.Pos.Y-r.Rad.Value/2,
		r.Rad.Value, r.Rad.Value,
		colornames.Aliceblue,
	)
}

type Menu struct{}

func (m *Menu) Setup(w mizu.World) {
	//w.AddComponents(Pos{}, Vel{}, Rad{})
	radius := 15.0
	width, height := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())
	cx, cy := (width-radius)/2, (height-radius)/2
	min, max := -10.0, 10.0
	vx, vy := min+rand.Float64()*(max-min), min+rand.Float64()*(max-min)
	w.AddEntities(
		&Ball{Pos{cx, cy}, Vel{vx, vy}, Rad{radius}},
		//&Ball{Pos{cx, cy}, Vel{1, -1}, Rad{radius}},
		//&Ball{Pos{cx, cy}, Vel{-1, 1}, Rad{radius}},
		//&Ball{Pos{cx, cy}, Vel{1, 1}, Rad{radius}},
	)
	w.AddSystems(
		&Velocity{},
		&Render{},
	)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.RunGame(mizu.NewGame(&Menu{}))
}
