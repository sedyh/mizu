package system

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/examples/platformer/helper/name"
	"github.com/sedyh/mizu/examples/platformer/helper/num"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/examples/platformer/assets"
	"github.com/sedyh/mizu/examples/platformer/component"
	"github.com/sedyh/mizu/pkg/engine"
)

type Player struct{}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Update(w engine.World) {
	// Get auxiliary objects
	player, ok := w.View(component.Pos{}, component.Vel{}, component.Sprite{}, component.Control{}).Get()
	if !ok {
		return
	}
	var pos *component.Pos
	var vel *component.Vel
	var gravity *component.Gravity
	var sprite *component.Sprite
	var control *component.Control
	player.Get(&pos, &vel, &gravity, &sprite, &control)

	// Almost a jump buffer.
	grounded := num.Equal(vel.M, 0, 0.002)

	// Jump with height control.
	if grounded && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		vel.M -= control.JumpSpeed
	}
	if vel.M > 0 {
		vel.M += gravity.Value * (control.FallSpeed - 1)
	} else if vel.M < 0 && !ebiten.IsKeyPressed(ebiten.KeySpace) {
		vel.M += gravity.Value * (control.LowSpeed - 1)
	}

	// Horizontal movement.
	moveDirection := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		moveDirection = -1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		moveDirection = 1.0
	}
	vel.L = num.Lerp(vel.L, control.MoveSpeed*moveDirection, control.MoveSpeed)

	// Player movement animation.
	// You can easily add sharp turns here if you make the horizontal movement smoother.
	if moveDirection != 0 {
		sprite.Frameset = assets.Images["player_"+name.PlayerMovement(grounded)+"_"+name.PlayerDirection(moveDirection)]
	} else {
		sprite.Frameset = assets.Images["player_idle"]
	}
}

func (p *Player) Draw(w engine.World, screen *ebiten.Image) {
	// Get auxiliary objects
	player, ok := w.View(component.Pos{}, component.Vel{}, component.Sprite{}).Get()
	if !ok {
		return
	}
	cameraEntity, found := w.View(component.Pos{}, component.Zoom{}).Get()
	if !found {
		return
	}
	var pos *component.Pos
	var vel *component.Vel
	var sprite *component.Sprite
	var camera *component.Pos
	var zoom *component.Zoom
	player.Get(&pos, &vel, &sprite)
	cameraEntity.Get(&camera, &zoom)

	// Apply camera zoom.
	scaledTilesize := float64(assets.Tilesize) * zoom.Value

	// Draw the player.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(zoom.Value, zoom.Value)
	op.GeoM.Translate(-camera.X, -camera.Y)
	op.GeoM.Translate(pos.X*scaledTilesize, pos.Y*scaledTilesize)
	screen.DrawImage(sprite.Frameset.Image(), op)
}
