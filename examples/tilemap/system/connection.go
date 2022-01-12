package system

import (
	"strconv"

	"github.com/sedyh/mizu/examples/tilemap/assets"
	"github.com/sedyh/mizu/examples/tilemap/component"
	"github.com/sedyh/mizu/examples/tilemap/helper/tilemap"
	"github.com/sedyh/mizu/examples/tilemap/helper/tileset"
	"github.com/sedyh/mizu/pkg/engine"
)

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Update(w engine.World) {
	// Get auxiliary objects
	constructEntity, found := w.View(component.Construct{}).Get()
	if !found {
		return
	}
	var construct *component.Construct
	constructEntity.Get(&construct)

	// Get tiles and solid states for calculating blob index
	tiles := w.View(component.Pos{}, component.Solid{}, component.Sprite{}).Filter()
	solids := make([]bool, len(tiles))
	for i, te := range tiles {
		var solid *component.Solid
		te.Get(&solid)

		solids[i] = solid.Value
	}

	// Calculate blob indexes and change sprites
	for y := 0; y < construct.Height; y++ {
		for x := 0; x < construct.Width; x++ {
			var sprite *component.Sprite

			neighbors := tilemap.Area(solids, x, y, construct.Width, construct.Height)
			tile := tiles[tilemap.Index(x, y, construct.Width)]
			tile.Get(&sprite)

			sprite.Image = assets.Images["road_"+strconv.Itoa(tileset.BlobIndex48(neighbors))]
		}
	}
}
