package engine

import (
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

// SystemUpdater is a system that updates something every game tick.
type SystemUpdater interface {
	Update(w World) // Updates specific set of entities.
}

// SystemDrawer is a system that draws something every game frame.
type SystemDrawer interface {
	Draw(w World, screen *ebiten.Image) // Draws specific set of entities.
}

// system is a struct, that wraps user-defined one and contains
// its own mask for filtering entities by their components and references
// to the component store to overwrite the fields specified in the
// system for each entity.
type system struct {
	inner  interface{}
	empty  bool
	mask   mask
	ids    []int
	stores []*store
	values []reflect.Value
	zeros  []reflect.Value
}

// makeSystem creates new system based on the one defined by the user.
func makeSystem(world *world, s interface{}) *system {
	systemValue := reflect.ValueOf(s).Elem()
	systemType := systemValue.Type()
	systemFieldsNum := systemType.NumField()

	res := &system{
		inner:  s,
		mask:   makeMask(len(world.stores)),
		ids:    make([]int, 0, systemFieldsNum),
		stores: make([]*store, 0, systemFieldsNum),
		values: make([]reflect.Value, 0, systemFieldsNum),
	}

	for i := 0; i < systemFieldsNum; i++ {
		componentValue := systemValue.Field(i)
		if !componentValue.CanInterface() {
			continue // unexported field
		}
		componentType := componentValue.Type()
		if componentType.Kind() != reflect.Ptr {
			continue // system value field is not component bound
		}
		componentId, found := world.componentIds[componentType.Elem()]
		if !found {
			continue // system pointer field is not component bound
		}
		res.ids = append(res.ids, componentId)
		res.stores = append(res.stores, world.stores[componentId])
		res.values = append(res.values, componentValue)
		res.mask.set(componentId)
	}
	res.empty = len(res.ids) == 0

	return res
}

// updateForEachTick handles game updateForEachTick for each entity if there is any components.
func (s *system) updateForEachEntity(world World, entity *entity) {
	if s.empty {
		return
	}

	if entity.mask.contains(s.mask) {
		if v, ok := s.inner.(SystemUpdater); ok {
			for i := range s.stores {
				s.stores[i].get(entity.id, s.values[i])
			}
			v.Update(world)
		}
	}
}

// drawForEachTick handles game drawForEachTick for each entity if there is any components.
func (s *system) drawForEachEntity(world World, screen *ebiten.Image, entity *entity) {
	if s.empty {
		return
	}

	if entity.mask.contains(s.mask) {
		if v, ok := s.inner.(SystemDrawer); ok {
			for i := range s.stores {
				s.stores[i].get(entity.id, s.values[i])
			}
			v.Draw(world, screen)
		}
	}
}

// updateForEachTick handles game updateForEachTick for each tick if there is no components.
func (s *system) updateForEachTick(world World) {
	if !s.empty {
		return
	}

	if v, ok := s.inner.(SystemUpdater); ok {
		v.Update(world)
	}
}

// drawForEachTick handles game drawForEachTick for each tick if there is no components.
func (s *system) drawForEachTick(world World, screen *ebiten.Image) {
	if !s.empty {
		return
	}

	if v, ok := s.inner.(SystemDrawer); ok {
		v.Draw(world, screen)
	}
}
