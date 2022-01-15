package tilemap

import "github.com/sedyh/mizu/examples/platformer/helper/tileset"

func Index(x, y, w int) int {
	return y*w + x
}

func Solid(solids []bool, x, y, w, h int) bool {
	xOut := x < 0 || x >= w
	yOut := y < 0 || y >= h
	if xOut || yOut {
		return false
	}

	return solids[Index(x, y, w)]
}

func Area(solids []bool, x, y, w, h int) *tileset.Neighbors {
	return &tileset.Neighbors{
		TopLeft:     Solid(solids, x-1, y-1, w, h),
		TopRight:    Solid(solids, x+1, y-1, w, h),
		Top:         Solid(solids, x, y-1, w, h),
		Right:       Solid(solids, x+1, y, w, h),
		Left:        Solid(solids, x-1, y, w, h),
		BottomRight: Solid(solids, x+1, y+1, w, h),
		Bottom:      Solid(solids, x, y+1, w, h),
		BottomLeft:  Solid(solids, x-1, y+1, w, h),
	}
}
