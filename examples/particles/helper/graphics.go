package helper

import (
	"image"
	"image/color"
	"image/draw"
)

func Gradient() []color.Color {
	return []color.Color{
		color.RGBA{R: 128, G: 128, B: 128, A: 150},
		color.RGBA{R: 128, G: 128, B: 128, A: 100},
		color.RGBA{R: 128, G: 128, B: 128, A: 50},
		color.RGBA{R: 128, G: 128, B: 128, A: 25},
		color.RGBA{R: 100, G: 100, B: 100, A: 20},
		color.RGBA{R: 80, G: 80, B: 80, A: 15},
		color.RGBA{R: 40, G: 40, B: 40, A: 10},
		color.RGBA{R: 20, G: 20, B: 20, A: 5},
		color.RGBA{R: 1, G: 1, B: 1, A: 1},
		color.RGBA{R: 20, G: 20, B: 20, A: 5},
		color.RGBA{R: 40, G: 40, B: 40, A: 10},
		color.RGBA{R: 80, G: 80, B: 80, A: 15},
		color.RGBA{R: 100, G: 100, B: 100, A: 20},
		color.RGBA{R: 100, G: 100, B: 100, A: 20},
		color.RGBA{R: 100, G: 100, B: 100, A: 20},
		color.RGBA{R: 80, G: 80, B: 80, A: 15},
		color.RGBA{R: 40, G: 40, B: 40, A: 10},
		color.RGBA{R: 20, G: 20, B: 20, A: 5},
		color.RGBA{R: 1, G: 1, B: 1, A: 1},
	}
}

func Checkerboard(w, h, cells int) image.Image {
	m := image.NewRGBA(Rect(0, 0, w, h))
	cellW, cellH := w/cells, h/cells
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			c := color.RGBA{R: 255, B: 254, A: 255}
			if (i+j)%2 == 0 {
				c = color.RGBA{A: 255}
			}
			draw.Draw(m, Rect(i*cellW, j*cellH, cellW, cellH), &image.Uniform{C: c}, image.Point{}, draw.Src)
		}
	}
	return m
}

func Rect(x, y, w, h int) image.Rectangle {
	return image.Rect(x, y, x+w, y+h)
}
