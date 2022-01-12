package engine

import (
	"reflect"
)

// View is a request to the game world to independently filter entities
// by their components instead of specifying them in the system.
// This is useful when you need to make requests to different sets
// of components in the same system.
type View interface {
	Each(consumer func(e Entity)) // Iterates all entities containing the given components.
	Filter() []Entity             // Returns a list of entities containing the given components.
	Get() (e Entity, ok bool)     // Returns the first entity containing the given components and a search status.
}

// view is inner struct that contains it own mask based on the components passed to the constructor.
type view struct {
	w    *world
	mask mask
}

// makeView creates new view based on the components passed to the constructor.
func makeView(world *world, components ...interface{}) *view {
	m := makeMask(len(world.stores))

	for _, component := range components {
		componentType := reflect.TypeOf(component)
		componentId, ok := world.componentIds[componentType]
		if !ok {
			continue
		}
		m.set(componentId)
	}

	return &view{w: world, mask: m}
}

// Each iterates all entities containing the given components.
// This method is for iteration only.
func (v *view) Each(consumer func(e Entity)) {
	for _, en := range v.w.entities {
		if en.mask.contains(v.mask) {
			consumer(en)
		}
	}
}

// Filter returns a list of entities containing the given components for separate sorting and iteration.
// It is safe to delete entities from here, you also can use this to sort your entities.
func (v *view) Filter() []Entity {
	entities := make([]Entity, 0, 2)
	for _, en := range v.w.entities {
		if en.mask.contains(v.mask) {
			entities = append(entities, en)
		}
	}
	return entities
}

// Get returns the first entity containing the given components and a search status.
func (v *view) Get() (e Entity, ok bool) {
	for _, en := range v.w.entities {
		if en.mask.contains(v.mask) {
			return en, true
		}
	}

	return nil, false
}
