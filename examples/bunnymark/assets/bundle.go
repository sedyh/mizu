package assets

import (
	"embed"
	"image/color"

	"github.com/sedyh/mizu/examples/bunnymark/helper"
)

// This is where the images, colors and gradients are loaded.

var (
	Background = color.RGBA{R: 41, G: 44, B: 45, A: 255}
	Bunny      = helper.Image(fs, "data/image/mizu-logo.png")
)

//go:embed data
var fs embed.FS
