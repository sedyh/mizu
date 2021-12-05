package engine

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

type SystemUpdater interface {
	Update(w World)
}

type SystemDrawer interface {
	Draw(screen *ebiten.Image)
}

type system struct {
	inner  interface{}
	mask   mask
	ids    []int
	stores []*store
	values []reflect.Value
	zeros  []reflect.Value
}

func makeSystem(world *world, s interface{}) *system {
	systemValue := reflect.ValueOf(s).Elem()
	systemType := systemValue.Type()
	systemFieldsNum := systemType.NumField()
	res := &system{
		inner:  s,
		mask:   makeMask(len(world.componentStores)),
		ids:    make([]int, 0, systemFieldsNum),
		stores: make([]*store, 0, systemFieldsNum),
		values: make([]reflect.Value, 0, systemFieldsNum),
	}
	for i := 0; i < systemFieldsNum; i++ {
		componentValue := systemValue.Field(i)
		if !componentValue.CanInterface() {
			continue
		}
		componentType := componentValue.Type().Elem()
		componentId := world.componentIds[componentType]
		res.ids = append(res.ids, componentId)
		res.stores = append(res.stores, world.componentStores[componentId])
		res.values = append(res.values, componentValue)
		res.mask.set(componentId)
	}
	return res
}

func (s *system) update(world World, entity *entity) {
	if entity.mask.contains(s.mask) {
		for i := range s.stores {
			s.stores[i].get(entity.id, s.values[i])
		}
		if v, ok := s.inner.(SystemUpdater); ok {
			v.Update(world)
		}
	}
}

func (s *system) draw(screen *ebiten.Image, entity *entity) {
	if entity.mask.contains(s.mask) {
		for i := range s.stores {
			s.stores[i].get(entity.id, s.values[i])
		}
		if v, ok := s.inner.(SystemDrawer); ok {
			v.Draw(screen)
		}
	}
}
