package test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sedyh/mizu/pkg/engine"
	"github.com/sedyh/mizu/test/helper"
)

var _ = Describe("Game world", func() {
	It("Should be able to create the correct View", func() {
		type ComponentA struct {
			Value int
		}
		type ComponentB struct {
			Value int
		}
		type ComponentC struct {
			Value int
		}
		type ComponentD struct {
			Value int
		}
		type EntityA struct {
			ComponentA
			ComponentB
			ComponentC
		}
		type EntityB struct {
			ComponentD
		}
		i := 0
		helper.RunSingleSceneGame(func(w engine.World) {
			w.AddComponents(ComponentA{}, ComponentB{}, ComponentC{}, ComponentD{})
			w.AddEntities(&EntityA{ComponentA{}, ComponentB{}, ComponentC{}}, &EntityB{ComponentD{}})
			for j := 0; j < 100; j++ {
				// Each view should fire only one time per iteration
				for _, e := range w.View(ComponentA{}).Filter() {
					var a *ComponentA
					e.Get(&a)
					i++
				}
				for _, e := range w.View(ComponentA{}, ComponentB{}).Filter() {
					var a *ComponentA
					var b *ComponentB
					e.Get(&a, &b)
					i++
				}
				// The order of the components is not important
				for _, e := range w.View(ComponentB{}, ComponentC{}).Filter() {
					var b *ComponentB
					var c *ComponentC
					e.Get(&b, &c)
					i++
				}
				for _, e := range w.View(ComponentC{}, ComponentB{}).Filter() {
					var c *ComponentC
					var b *ComponentB
					e.Get(&c, &b)
					i++
				}
				for _, e := range w.View(ComponentA{}, ComponentB{}, ComponentB{}).Filter() {
					var a *ComponentA
					var b *ComponentB
					var c *ComponentC
					e.Get(&c, &b, &a)
					i++
				}
			}
		})
		Expect(i).To(Equal(500))
	})
})
