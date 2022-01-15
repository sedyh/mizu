package collision

import (
	"fmt"
	"math"
)

// Contact is information about the collision of two rectangles.
type Contact struct {
	ContactX, ContactY float64 // Contact point [-n..n].
	NormalX, NormalY   float64 // Contact normal [-1..1].
	Time               float64 // Contact time [-1..1].
}

func (c Contact) String() string {
	return fmt.Sprintf(
		"point: %.2f:%.2f, normal: %.2f:%.2f, time: %.2f",
		c.ContactX, c.ContactY,
		c.NormalX, c.NormalY,
		c.Time,
	)
}

// DynamicRectInRect is an implementation of the "SweptAABB" collision detection algorithm.
// It checks if the moving rectangle intersects the static one and at what point.
func DynamicRectInRect(
	currentX, currentY, currentW, currentH float64,
	targetX, targetY, targetW, targetH float64,
	directionX, directionY float64,
) (c Contact, ok bool) {
	if directionX == 0 && directionY == 0 {
		return
	}

	// Take into account the size of the collision area
	// in order to react in advance, and not in the middle.
	targetX, targetY = targetX-currentW/2, targetY-currentH/2
	targetW, targetH = currentW+targetW, currentH+targetH
	currentX += currentW / 2
	currentY += currentH / 2

	contact, ok := RayInRect(
		currentX, currentY, directionX, directionY,
		targetX, targetY, targetW, targetH,
	)
	if ok && contact.Time >= 0.0 && contact.Time <= 1.0 {
		// Rectangle collides with rectangle.
		return contact, true
	}

	return Contact{}, false
}

// RayInRect casts a ray and checks its intersection with the rectangle.
func RayInRect(
	originX, originY float64,
	directionX, directionY float64,
	rectX, rectY, rectW, rectH float64,
) (c Contact, ok bool) {
	// At what point did we cross the rectangle for the first time.
	tNearX := (rectX - originX) / directionX
	tNearY := (rectY - originY) / directionY

	if math.IsNaN(tNearX) || math.IsNaN(tNearY) {
		// No collision.
		return Contact{}, false
	}

	// At what point did we cross the rectangle for the second time.
	tFarX := (rectX + rectW - originX) / directionX
	tFarY := (rectY + rectH - originY) / directionY

	if math.IsNaN(tFarX) || math.IsNaN(tFarY) {
		// No collision.
		return Contact{}, false
	}

	// Make near always first and far always second.
	if tNearX > tFarX {
		tNearX, tFarX = tFarX, tNearX
	}
	if tNearY > tFarY {
		tNearY, tFarY = tFarY, tNearY
	}
	if tNearX > tFarY || tNearY > tFarX {
		// No collision.
		return Contact{}, false
	}

	// Event times along direction
	tHitNear := math.Max(tNearX, tNearY)
	tHitFar := math.Min(tFarX, tFarY)
	if tHitFar < 0 {
		// Event in the opposite direction.
		return Contact{}, false
	}

	// The ray crossed the rectangle.
	ok = true
	c.Time = tHitNear
	c.ContactX = originX + tHitNear*directionX
	c.ContactY = originY + tHitNear*directionY
	if tNearX > tNearY {
		if directionX < 0 {
			c.NormalX, c.NormalY = 1, 0
		} else {
			c.NormalX, c.NormalY = -1, 0
		}
	} else if tNearX < tNearY {
		if directionY < 0 {
			c.NormalX, c.NormalY = 0, 1
		} else {
			c.NormalX, c.NormalY = 0, -1
		}
	}
	return
}
