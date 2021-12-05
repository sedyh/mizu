package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"mizu"
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

		// This is a game init
		entity := &Ball{Rad(5), Pos{1, 1}, Vel{2, 2}} // Should be a pointer
		mizu.NewGame(&SingleEntityScene{entity})
	})

	It("Should panic if entity is not pointer", func() {
		defer func() {
			Expect(recover()).To(Equal("entity should be a pointer"))
		}()
		type Pos struct {
			X, Y float64
		}
		type Ball struct {
			Pos
		}
		mizu.NewGame(&SingleEntityScene{Ball{Pos{1, 1}}}) // Not a pointer
	})

	It("Should panic if entity under pointer is not a struct", func() {
		defer func() {
			Expect(recover()).To(Equal("entity under pointer should be a struct"))
		}()
		type Ball int // Not a struct
		b := Ball(1)
		mizu.NewGame(&SingleEntityScene{&b})
	})

	It("Should panic if entity components are pointers", func() {
		defer func() {
			Expect(recover()).To(Equal("entity should be a pointer"))
		}()
		type Pos struct {
			X, Y float64
		}
		type Ball struct {
			*Pos // A pointer
		}
		mizu.NewGame(&SingleEntityScene{&Ball{&Pos{1, 1}}})
	})
})
