package engine

import (
	"container/list"
	"fmt"
	"reflect"
)

type Entity interface {
	Get(components ...interface{})
}

type entity struct {
	w       *world
	id      int
	mask    mask
	element *list.Element
}

func makeEntity(w *world, components ...interface{}) *entity {
	e := &entity{
		w:    w,
		id:   w.entitiesIndexes.get(),
		mask: makeMask(len(w.componentStores)),
	}
	e.set(components...)
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
			e.w.componentStores[componentId].set(e.id, componentValue)
			return
		}
		e.w.componentStores[componentId].add(e.id, componentValue)
		e.mask.set(componentId)
	}
}

func (e *entity) rem(components ...interface{}) {
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
