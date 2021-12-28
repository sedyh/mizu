package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sedyh/mizu/pkg/engine"
	"github.com/sedyh/mizu/test/helper"
)

type Position struct {
	X, Y int
}

type Velocity struct {
	X, Y int
}

type Gravity struct {
	Value int
}

type First struct {
	Value bool
}

type Ball struct {
	First
	Position
	Velocity
	Gravity
}

type Movement struct {
	*Position
	*Velocity
}

func (m *Movement) Update(_ engine.World) {
	m.Position.X += m.Velocity.X
	m.Position.Y += m.Velocity.Y
}

type Falling struct {
	*Velocity
	*Gravity
}

func (f *Falling) Update(_ engine.World) {
	f.Velocity.Y += f.Gravity.Value
}

var _ = Describe("Game world", func() {
	It("Should be able to create a system that traverses each entity", func() {
		var world engine.World

		// Run the world with two systems in two entities
		helper.RunSingleSceneGame(func(w engine.World) {
			w.AddComponents(First{}, Position{}, Velocity{}, Gravity{})
			w.AddEntities(
				&Ball{First{true}, Position{1, 5}, Velocity{-3, 4}, Gravity{2}},
				&Ball{First{}, Position{-5, 3}, Velocity{1, 8}, Gravity{3}},
			)
			w.AddSystems(&Movement{}, &Falling{})
			world = w
		})

		// Check that each system done it work right
		world.View(First{}, Position{}, Velocity{}).Each(func(e engine.Entity) {
			var first *First
			var position *Position
			var velocity *Velocity
			e.Get(&first, &position, &velocity)
			if first.Value {
				Expect(*position).To(Equal(Position{-2, 9}))
				Expect(*velocity).To(Equal(Velocity{-3, 6}))
			}
		})
		world.View(First{}, Position{}, Velocity{}).Each(func(e engine.Entity) {
			var first *First
			var position *Position
			var velocity *Velocity
			e.Get(&first, &position, &velocity)
			if !first.Value {
				Expect(*position).To(Equal(Position{-4, 11}))
				Expect(*velocity).To(Equal(Velocity{1, 11}))
			}
		})
	})
})
