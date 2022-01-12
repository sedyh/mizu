package engine

import (
	"fmt"
	"reflect"
)

// Entity represents any game object with inner id.
type Entity interface {
	ID() int                       // Gets entity id for storage elsewhere.
	Get(components ...interface{}) // Gets entity components, takes a set of pointers to pointers.
}

// entity is inner struct that contains it own id and mask based on the components passed to the constructor.
type entity struct {
	w    *world
	id   int
	mask mask
}

// makeEntity creates new entity with id and mask based on the passed components.
func makeEntity(w *world, components ...interface{}) *entity {
	e := &entity{
		w:    w,
		id:   w.entitiesIds.get(),
		mask: makeMask(len(w.stores)),
	}
	e.set(components...)
	return e
}

// ID returns entity id
func (e *entity) ID() int {
	return e.id
}

// Get sets the values of entity components according to the passed pointers.
// Example:
//     var pos *Pos
//     var rad *Rad
//     entity.Get(&pos, &rad)
func (e *entity) Get(components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component)
		if componentValue.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("received entity component %s must be a pointer", typeName(componentValue.Type())))
		}
		componentValueElem := componentValue.Elem()
		if componentValueElem.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("received entity component %s must be a pointer to pointer", typeName(componentValue.Type())))
		}
		componentType := componentValueElem.Type().Elem()
		componentId := e.w.componentIds[componentType]
		if e.mask.get(componentId) {
			e.w.stores[componentId].get(e.id, componentValueElem)
		} else {
			componentValueElem.Set(reflect.Zero(reflect.PtrTo(componentType)))
		}
	}
}

// set replaces the values of the entity components with the passed ones.
func (e *entity) set(components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component)
		componentType := componentValue.Type()
		if componentType.Kind() == reflect.Ptr {
			panic(fmt.Sprintf("entity component %s should not be a pointer", componentType.Elem().Name()))
		}
		componentId, found := e.w.componentIds[componentType]
		if !found {
			panic(fmt.Sprintf("entity has unregistered component %s", componentType.Name()))
		}
		if e.mask.get(componentId) {
			e.w.stores[componentId].set(e.id, componentValue)
			return
		}
		e.w.stores[componentId].add(e.id, componentValue)
		e.mask.set(componentId)
	}
}

// rem zeroes the values of the entity components.
func (e *entity) rem(components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component)
		componentType := componentValue.Type()
		componentId := e.w.componentIds[componentType]
		if e.mask.get(componentId) {
			e.w.stores[componentId].rem(e.id)
			e.mask.clear(componentId)
		}
	}
}
