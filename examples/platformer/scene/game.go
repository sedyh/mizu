package scene

import (
	"github.com/sedyh/mizu/examples/platformer/assets"
	"github.com/sedyh/mizu/examples/platformer/component"
	"github.com/sedyh/mizu/examples/platformer/entity"
	"github.com/sedyh/mizu/examples/platformer/helper/enum"
	"github.com/sedyh/mizu/examples/platformer/helper/load"
	"github.com/sedyh/mizu/examples/platformer/helper/tilemap"
	"github.com/sedyh/mizu/examples/platformer/system"
	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Pos{}, component.Vel{}, component.Size{}, component.Zoom{},
		component.Solid{}, component.Construct{}, component.Sprite{}, component.Gravity{}, component.Control{},
	)

	// Fills the entire space with tiles, some of them are marked as empty.
	construct := component.NewConstruct(load.Level(assets.Level))
	w.AddSystems(
		system.NewRendering(construct.Width, construct.Height),
		system.NewPlayer(), system.NewGravity(), system.NewCollision(), system.NewVelocity(),
		system.NewFocus(), system.NewConnection(), system.NewDebug(), system.NewAnimation(),
	)
	for y := 0; y < construct.Height; y++ {
		for x := 0; x < construct.Width; x++ {
			switch string(construct.Level[tilemap.Index(x, y, construct.Width)]) {
			case "o":
				w.AddEntities(&entity.Tile{
					Pos:    component.NewPosI(x, y),
					Vel:    component.NewVel(0, 0),
					Size:   component.NewSizeI(1, 1),
					Solid:  component.NewSolid(enum.CollisionGroupTile),
					Sprite: component.NewSprite(assets.Images["wall_0"]),
				})
			default:
				w.AddEntities(&entity.Tile{
					Pos:    component.NewPosI(x, y),
					Vel:    component.NewVel(0, 0),
					Size:   component.NewSizeI(1, 1),
					Solid:  component.NewSolid(enum.CollisionGroupNone),
					Sprite: component.NewSprite(assets.Images["wall_0"]),
				})
			}

		}
	}

	// Adding crates.
	for y := 0; y < construct.Height; y++ {
		for x := 0; x < construct.Width; x++ {
			switch string(construct.Level[tilemap.Index(x, y, construct.Width)]) {
			case "x":
				w.AddEntities(&entity.Crate{
					Pos: component.NewPosI(x, y),
					Size: component.NewSizeF(
						float64(assets.Cratesize)/float64(assets.Tilesize),
						float64(assets.Cratesize)/float64(assets.Tilesize),
					),
					Vel:     component.NewVel(0, 0),
					Solid:   component.NewSolid(enum.CollisionGroupCrate),
					Sprite:  component.NewSprite(assets.Images["crate"]),
					Gravity: component.NewGravity(0.04),
				})
			}
		}
	}

	// Adding a player.
	w.AddEntities(
		&entity.Player{
			Pos: component.NewPosI(construct.Width/2, construct.Height/2),
			Size: component.NewSizeF(
				float64(assets.PlayerW)/float64(assets.Tilesize),
				float64(assets.PlayerH)/float64(assets.Tilesize),
			),
			Vel:     component.NewVel(0, 0),
			Solid:   component.NewSolid(enum.CollisionGroupPlayer),
			Sprite:  component.NewSprite(assets.Images["player_run_right"]),
			Gravity: component.NewGravity(0.04),
			Control: component.NewControl(0.7, 2.5, 2.0, 0.2),
		},
		&entity.Construct{Construct: construct}, &entity.Camera{Zoom: component.NewZoom(1.0)},
	)
}
