package engine

import "reflect"

type View interface {
	Each(consumer func(entity Entity))
	Filter() []Entity
}

type view struct {
	w    *world
	mask mask
}

func makeView(world *world, components ...interface{}) *view {
	m := makeMask(len(world.componentStores))

	for _, component := range components {
		componentType := reflect.TypeOf(component)
		componentId := world.componentIds[componentType]
		m.set(componentId)
	}

	return &view{w: world, mask: m}
}

func (v *view) Each(consumer func(entity Entity)) {
	for element := v.w.entities.Front(); element != nil; element = element.Next() {
		e := element.Value.(*entity)
		if e.mask.contains(v.mask) {
			consumer(e)
		}
	}
}

func (v *view) Filter() []Entity {
	entities := make([]Entity, 0, 2)
	for element := v.w.entities.Front(); element != nil; element = element.Next() {
		e := element.Value.(*entity)
		if e.mask.contains(v.mask) {
			entities = append(entities, e)
		}
	}
	return entities
}
