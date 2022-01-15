package assets

import (
	"embed"
	"image/color"
	"strconv"
	"time"

	"github.com/sedyh/mizu/examples/platformer/helper/graphics"

	"github.com/sedyh/mizu/examples/platformer/helper/load"
	"github.com/sedyh/mizu/examples/platformer/helper/split"
)

// This is where the images, colors and levels are loaded.

const (
	Tilesize  = 32
	PlayerW   = 25
	PlayerH   = 36
	Cratesize = 64

	Animation = 100 * time.Millisecond

	Level = `
		oooooooooooooooooooo
		o                  o
		o                  o
		o           ooooo  o
		o  oo              o
		o  ox          oo  o
		o  o          x o  o
		o  oooooo       o  o
		o          oooooo  o
		o                  o
		o                  o
		oooooooooooooooooooo
	`
)

var (
	Background = color.RGBA{R: 41, G: 44, B: 45, A: 255}
	Images     = make(map[string]*graphics.Frameset)
)

//go:embed data
var fs embed.FS

func Init() {
	// Slicing and loading wang blob tileset.
	road := load.Image(fs, "data/image/wall.png")
	blobW, blobH := 8, 6
	for y := 0; y < blobH; y++ {
		for x := 0; x < blobW; x++ {
			Images["wall_"+strconv.Itoa(x+y*blobW)] = split.Single(
				road, x*Tilesize, y*Tilesize,
				Tilesize, Tilesize, false,
			)
		}
	}

	// Slicing and loading player frames.
	player := load.Image(fs, "data/image/player.png")
	Images["player_idle"] = split.Single(player, 0, 0, PlayerW, PlayerH, false)
	for f := 0; f < 2; f++ {
		flipped := f == 1
		postfix := "right"
		if flipped {
			postfix = "left"
		}
		Images["player_look_"+postfix] = split.Single(player, PlayerW, 0, PlayerW, PlayerH, flipped)
		Images["player_jump_"+postfix] = split.Single(player, PlayerW*2, 0, PlayerW, PlayerH, flipped)
		Images["player_run_"+postfix] = split.Multi(player, PlayerW*3, 0, PlayerW, PlayerH, 4, flipped, Animation)

	}

	// Slicing and loading create image.
	crate := load.Image(fs, "data/image/crate.png")
	Images["crate"] = split.Single(crate, 0, 0, Cratesize, Cratesize, false)
}
