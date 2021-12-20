package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sedyh/mizu/pkg/engine"
	"github.com/sedyh/mizu/test/helper"
)

var _ = Describe("Game world", func() {
	It("Should be able to add entity embedded components", func() {
		// These are components
		type Rad int                    // Can be any custom type
		type Pos struct{ X, Y float64 } // Can be any custom type
		type Vel struct{ X, Y float64 } // Can be any custom type

		// This is an entity
		type Ball struct {
			Rad // Should not be a pointer
			Pos // Should not be a pointer
			Vel // Should not be a pointer
		}

		// Create Game and init Scene
		helper.RunSingleSceneGame(func(w engine.World) {
			w.AddComponents(Rad(0), Pos{}, Vel{})
			w.AddEntities(&Ball{Rad(5), Pos{1, 1}, Vel{2, 2}}) // Should be a pointer
		})
	})

	It("Should panic if entity is not pointer", func() {
		defer func() {
			Expect(recover()).To(Equal("entity Ball should be a pointer"))
		}()
		type Pos struct {
			X, Y float64
		}
		type Ball struct {
			Pos
		}
		helper.RunSingleSceneGame(func(w engine.World) {
			w.AddComponents(Pos{})
			w.AddEntities(Ball{Pos{1, 1}}) // Not a pointer
		})
	})

	It("Should panic if entity under pointer is not a struct", func() {
		defer func() {
			Expect(recover()).To(Equal("entity Ball under pointer should be a struct"))
		}()
		type Ball int
		b := Ball(1)
		helper.RunSingleSceneGame(func(w engine.World) {
			w.AddEntities(&b) // Not a struct
		})
	})

	It("Should panic if entity components are pointers", func() {
		defer func() {
			Expect(recover()).To(Equal("entity component Pos should not be a pointer"))
		}()
		type Pos struct {
			X, Y float64
		}
		type Ball struct {
			*Pos // Should not be a pointer
		}
		helper.RunSingleSceneGame(func(w engine.World) {
			w.AddComponents(&Pos{})
			w.AddEntities(&Ball{&Pos{1, 1}})
		})
	})
})
