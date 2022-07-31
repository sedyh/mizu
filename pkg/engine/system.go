package engine

import "github.com/hajimehoshi/ebiten/v2"

type Updater = func(w World)

type Drawer = func(w World, screen *ebiten.Image)
