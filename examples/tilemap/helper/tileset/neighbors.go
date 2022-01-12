package tileset

import "fmt"

// Neighbors stores all the status of the presence of the nearest tiles within a radius of 3x3
type Neighbors struct {
	TopLeft     bool
	Top         bool
	TopRight    bool
	Left        bool
	Right       bool
	BottomLeft  bool
	Bottom      bool
	BottomRight bool
}

func (n *Neighbors) String() string {
	return fmt.Sprintf(
		"%d, %d, %d\n%d, X, %d\n%d, %d, %d\n",
		b(n.TopLeft), b(n.Top), b(n.TopRight),
		b(n.Left), b(n.Right),
		b(n.BottomLeft), b(n.Bottom), b(n.BottomRight),
	)
}

func b(v bool) int {
	if v {
		return 1
	}
	return 0
}
