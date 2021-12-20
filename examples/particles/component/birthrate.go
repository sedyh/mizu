package component

import "time"

type Birthrate struct {
	Amount int
	Ticker *time.Ticker
}

func NewBirthrate(particles int, each time.Duration) Birthrate {
	return Birthrate{Amount: particles, Ticker: time.NewTicker(each)}
}

func NewZeroBirthrate() Birthrate {
	return Birthrate{}
}
