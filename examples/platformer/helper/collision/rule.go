package collision

import (
	"fmt"
	"math"

	"github.com/sedyh/mizu/examples/platformer/component"
	"github.com/sedyh/mizu/examples/platformer/helper/enum"
)

// Reaction is the way to resolve the collision.
type Reaction func(e Event, c Contact)

// Rule is the condition under which collisions are allowed and how they are resolved.
type Rule struct {
	A, B enum.CollisionGroup
	Reaction
}

// Slide is a way to resolve collisions in which one object will crawl on the surface of another.
func Slide(a, b enum.CollisionGroup) Rule {
	return Rule{a, b, func(e Event, c Contact) {
		e.VelA.L += c.NormalX * math.Abs(e.VelA.L) * (1 - c.Time)
		e.VelA.M += c.NormalY * math.Abs(e.VelA.M) * (1 - c.Time)
	}}
}

// Event is all the additional data needed to resolve a particular collision.
type Event struct {
	PosA, PosB     *component.Pos
	VelA, VelB     *component.Vel
	SizeA, SizeB   *component.Size
	SolidA, SolidB *component.Solid
	Reaction       Reaction
	Time           float64
}

func NewEvent(
	aPos *component.Pos, aVel *component.Vel, aSize *component.Size, aSolid *component.Solid,
	bPos *component.Pos, bVel *component.Vel, bSize *component.Size, bSolid *component.Solid,
	reaction Reaction, time float64,
) Event {
	return Event{aPos, bPos, aVel, bVel, aSize, bSize, aSolid, bSolid, reaction, time}
}

func (e Event) String() string {
	return fmt.Sprintf(
		"collision %s, %s, %s -> %s, %s, %s at time: %.2f",
		e.PosA, e.VelA, e.SizeA, e.PosB, e.VelB, e.SizeB, e.Time,
	)
}

// FindRule finds a collision resolution rule in the list that is similar to or reversed from the given one.
func FindRule(rules []Rule, a, b enum.CollisionGroup) (rule Rule, inverted, ok bool) {
	for _, r := range rules {
		if a == r.A && b == r.B {
			return r, false, true
		}

		if a == r.B && b == r.A {
			return r, true, true
		}
	}

	return Rule{}, false, false
}
