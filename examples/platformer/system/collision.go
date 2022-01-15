package system

import (
	"sort"

	"github.com/sedyh/mizu/examples/platformer/helper/enum"

	"github.com/sedyh/mizu/examples/platformer/component"
	"github.com/sedyh/mizu/examples/platformer/helper/collision"
	"github.com/sedyh/mizu/pkg/engine"
)

type Collision struct {
	rules      []collision.Rule
	candidates []engine.Entity
	events     []collision.Event
}

func NewCollision() *Collision {
	return &Collision{
		// Which collisions should be resolved and how.
		rules: []collision.Rule{
			collision.Slide(enum.CollisionGroupPlayer, enum.CollisionGroupTile),
			collision.Slide(enum.CollisionGroupPlayer, enum.CollisionGroupCrate),
			collision.Slide(enum.CollisionGroupCrate, enum.CollisionGroupTile),
			collision.Slide(enum.CollisionGroupCrate, enum.CollisionGroupCrate),
		},
	}
}

func (c *Collision) Update(w engine.World) {
	// Clear collision candidates and events from previous update
	c.candidates = c.candidates[:0]
	c.events = c.events[:0]

	// Select possible candidates.
	w.View(component.Pos{}, component.Vel{}, component.Size{}, component.Solid{}).Each(func(e engine.Entity) {
		var solid *component.Solid
		e.Get(&solid)

		if solid.Empty() {
			return
		}

		c.candidates = append(c.candidates, e)
	})

	// Walk each object over every object except itself to find collisions.
	// This list can be optimized through spatial partition, dirty rects, camera culling etc.
	for _, a := range c.candidates {
		for _, b := range c.candidates {
			if a.ID() == b.ID() {
				continue
			}

			var aPos *component.Pos
			var aVel *component.Vel
			var aSize *component.Size
			var aSolid *component.Solid
			a.Get(&aPos, &aVel, &aSize, &aSolid)

			var bPos *component.Pos
			var bVel *component.Vel
			var bSize *component.Size
			var bSolid *component.Solid
			b.Get(&bPos, &bVel, &bSize, &bSolid)

			// Check if the given collision is allowed in the list of rules.
			// If the opposite rule is found - swap objects.
			rule, inverted, ok := collision.FindRule(c.rules, aSolid.Group, bSolid.Group)
			if !ok {
				continue
			}
			if inverted {
				aPos, bPos = bPos, aPos
				aVel, bVel = bVel, aVel
				aSize, bSize = bSize, aSize
				aSolid, bSolid = bSolid, aSolid
			}

			// Do a collision check with Swept AABB to get the time until the event starts.
			contact, ok := collision.DynamicRectInRect(
				aPos.X, aPos.Y, aSize.W, aSize.H,
				bPos.X, bPos.Y, bSize.W, bSize.H,
				aVel.L, aVel.M,
			)
			if !ok {
				continue
			}

			// Add collision event.
			c.events = append(c.events, collision.NewEvent(
				aPos, aVel, aSize, aSolid,
				bPos, bVel, bSize, bSolid,
				rule.Reaction, contact.Time,
			))
		}
	}

	// Determine the order of collisions
	sort.Slice(c.events, func(i, j int) bool {
		return c.events[i].Time < c.events[j].Time
	})

	// Resolve collisions
	for _, c := range c.events {

		// Recheck if a collision exists since the last resolution and if so, start resolution.
		contact, ok := collision.DynamicRectInRect(
			c.PosA.X, c.PosA.Y, c.SizeA.W, c.SizeA.H,
			c.PosB.X, c.PosB.Y, c.SizeB.W, c.SizeB.H,
			c.VelA.L, c.VelA.M,
		)
		if !ok {
			continue
		}

		// Resolve collision.
		c.Reaction(c, contact)
	}
}
