package assets

import (
	"embed"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/tilemap/helper/load"
	"github.com/sedyh/mizu/examples/tilemap/helper/split"
)

// This is where the images and colors are loaded.

const (
	Tilesize = 128
	BlobW    = 8
	BlobH    = 6
)

var (
	Background = color.RGBA{R: 41, G: 44, B: 45, A: 255}
	Images     = make(map[string]*ebiten.Image)
)

//go:embed data
var fs embed.FS

func Init() {
	// Slicing and loading wang blob tileset.
	roads := load.Image(fs, "data/image/tileset.png")
	for y := 0; y < BlobH; y++ {
		for x := 0; x < BlobW; x++ {
			Images["road_"+strconv.Itoa(x+y*BlobW)] = split.Single(
				roads, x*Tilesize, y*Tilesize, Tilesize, Tilesize,
			)
		}
	}
}
