package mizu

import (
	"reflect"
)

type entity struct {
	w    *world
	id   int
	mask mask
}

func makeEntity(w *world, components ...interface{}) *entity {
	e := &entity{
		w:    w,
		id:   len(w.entities),
		mask: makeMask(len(w.componentStores)),
	}
	e.Set(components...)
	return e
}

func (e *entity) Get(components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component).Elem()
		componentType := componentValue.Type().Elem()
		componentId := e.w.componentIds[componentType]
		if e.mask.get(componentId) {
			e.w.componentStores[componentId].get(e.id, componentValue)
		} else {
			componentValue.Set(reflect.Zero(reflect.PtrTo(componentType)))
		}
	}
}

func (e *entity) Set(components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component)
		componentType := componentValue.Type()
		componentId := e.w.componentIds[componentType]
		if e.mask.get(componentId) {
			e.w.componentStores[componentId].set(e.id, componentValue)
			return
		}
		e.w.componentStores[componentId].add(e.id, componentValue)
		e.mask.set(componentId)
	}
}

func (e *entity) Rem(components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component)
		componentType := componentValue.Type()
		componentId := e.w.componentIds[componentType]
		if e.mask.get(componentId) {
			e.w.componentStores[componentId].rem(e.id)
			e.mask.clear(componentId)
		}
	}
}
