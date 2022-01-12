package entity

import "github.com/sedyh/mizu/examples/particles/component"

// An object that spawns particles every tick.

type Emitter struct {
	component.Root      // Is an emitter.
	component.Birthrate // How many particles to spawn in a certain time.
	component.Pivot     // Setting for particle, spawn anchor point.
	component.Pos       // Setting for particle, spawn position.
	component.Vel       // Setting for particle, spawn velocity.
	component.Accel     // Setting for particle, spawn acceleration.
	component.Angle     // Setting for particle, spawn angle.
	component.Spin      // Setting for particle, spawn rotation per tick.
	component.Scale     // Setting for particle, spawn scale.
	component.Growth    // Setting for particle, spawn stretching per tick.
	component.Life      // Setting for particle, spawn lifetime.
	component.Gradient  // Setting for particle, spawn colors throughout lifetime.
	component.Sprite    // Setting for particle, spawn images.
}
