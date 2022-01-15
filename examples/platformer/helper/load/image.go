package load

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/platformer/helper/geo"
)

// Image loads the image from filepath and puts it on the GPU.
func Image(fs embed.FS, filepath string) *ebiten.Image {
	data, err := fs.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return ebiten.NewImageFromImage(Checkerboard(25, 32, 4))
	}

	m, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return ebiten.NewImageFromImage(Checkerboard(25, 32, 4))
	}

	return ebiten.NewImageFromImage(m)
}

// Checkerboard creates a placeholder image.
func Checkerboard(w, h, cells int) image.Image {
	m := image.NewRGBA(geo.Rect(0, 0, w, h))
	cellW, cellH := w/cells, h/cells
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			c := color.RGBA{R: 255, B: 254, A: 255}
			if (i+j)%2 == 0 {
				c = color.RGBA{A: 255}
			}
			draw.Draw(m, geo.Rect(i*cellW, j*cellH, cellW, cellH), &image.Uniform{C: c}, image.Point{}, draw.Src)
		}
	}
	return m
}
