package engine

import (
	"reflect"
)

// View is a request to the game world to independently filter entities
// by their components instead of specifying them in the system.
// This is useful when you need to make requests to different sets
// of components in the same system.
type View interface {
	Each(consumer func(entity Entity))
	Filter() []Entity
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

// Each iterates all entities with the previously selected components.
// This method is for iteration only.
func (v *view) Each(consumer func(entity Entity)) {
	for _, e := range v.w.entities {
		if e.mask.contains(v.mask) {
			consumer(e)
		}
	}
}

// Filter returns a list of entities with the previously selected components for separate sorting and iteration.
// It is safe to delete entities from here, you also can use this to sort your entities.
func (v *view) Filter() []Entity {
	entities := make([]Entity, 0, 2)
	for _, e := range v.w.entities {
		if e.mask.contains(v.mask) {
			entities = append(entities, e)
		}
	}
	return entities
}
