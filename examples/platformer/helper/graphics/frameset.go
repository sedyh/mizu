package graphics

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Frameset is a set of images with a ticker to switch them.
type Frameset struct {
	Images  []*ebiten.Image
	Ticker  *time.Ticker
	Current int
}

func NewFramesetSingle(frame *ebiten.Image) *Frameset {
	return &Frameset{Ticker: time.NewTicker(1), Images: []*ebiten.Image{frame}}
}

func NewFramesetMulti(speed time.Duration, frames ...*ebiten.Image) *Frameset {
	return &Frameset{Ticker: time.NewTicker(speed), Images: frames}
}

func (s *Frameset) Image() *ebiten.Image {
	return s.Images[s.Current]
}
