package helper

import (
	"image"
	"image/color"
	"image/draw"
)

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
