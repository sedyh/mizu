package system

import (
	"strconv"

	"github.com/sedyh/mizu/examples/platformer/assets"
	"github.com/sedyh/mizu/examples/platformer/component"
	"github.com/sedyh/mizu/examples/platformer/helper/tilemap"
	"github.com/sedyh/mizu/examples/platformer/helper/tileset"
	"github.com/sedyh/mizu/pkg/engine"
)

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Update(w engine.World) {
	// Get auxiliary objects.
	constructEntity, found := w.View(component.Construct{}).Get()
	if !found {
		return
	}
	var construct *component.Construct
	constructEntity.Get(&construct)

	// Get tiles and solid states for calculating blob index.
	tiles := w.View(component.Pos{}, component.Solid{}, component.Sprite{}).Filter()
	solids := make([]bool, len(tiles))
	for i, te := range tiles {
		var solid *component.Solid
		te.Get(&solid)

		solids[i] = !solid.Empty()
	}

	// Calculate blob indexes and change sprites.
	// These indices can also be calculated in advance during the construction of tiles.
	for y := 0; y < construct.Height; y++ {
		for x := 0; x < construct.Width; x++ {
			var sprite *component.Sprite

			neighbors := tilemap.Area(solids, x, y, construct.Width, construct.Height)
			tile := tiles[tilemap.Index(x, y, construct.Width)]
			tile.Get(&sprite)

			sprite.Frameset = assets.Images["wall_"+strconv.Itoa(tileset.BlobIndex48(neighbors))]
		}
	}
}
