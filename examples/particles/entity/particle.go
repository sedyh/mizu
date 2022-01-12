package entity

import "github.com/sedyh/mizu/examples/particles/component"

// A particle that moves by itself after settings are passed to it from the emitter.

type Particle struct {
	component.Root     // Is a particle.
	component.Pivot    // Spawn anchor point.
	component.Pos      // Current position.
	component.Vel      // Current velocity.
	component.Accel    // Current acceleration
	component.Angle    // Current angle.
	component.Spin     // Current rotation per tick.
	component.Scale    // Current size.
	component.Growth   // Current stretching per tick.
	component.Life     // Current lifetime.
	component.Gradient // Current colors.
	component.Sprite   // Current sprite.
}
