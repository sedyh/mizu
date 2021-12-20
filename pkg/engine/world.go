package engine

import (
	"container/list"
	"fmt"
	"image"
	"reflect"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type World interface {
	ChangeScene(next Scene)
	Bounds() image.Rectangle
	View(components ...interface{}) View
	AddComponents(components ...interface{})
	AddSystems(systems ...interface{})
	AddEntities(entities ...interface{})
	RemoveEntity(entity Entity)
}

type world struct {
	bounds          image.Rectangle
	componentIds    map[reflect.Type]int
	componentStores []*store
	systems         []*system
	entities        *list.List
	entitiesIndexes pool
	first           Scene
	once            sync.Once
}

func NewGame(first Scene) ebiten.Game {
	w := &world{
		componentIds:    make(map[reflect.Type]int, 256),
		componentStores: make([]*store, 0, 256),
		entities:        list.New(),
		entitiesIndexes: makePool(256),
		first:           first,
	}
	return w
}

func (w *world) Update() error {
	w.once.Do(func() { w.first.Setup(w) })
	for _, s := range w.systems {
		for element := w.entities.Front(); element != nil; element = element.Next() {
			s.update(w, element.Value.(*entity))
		}
	}
	return nil
}

func (w *world) Draw(screen *ebiten.Image) {
	for _, s := range w.systems {
		for element := w.entities.Front(); element != nil; element = element.Next() {
			s.draw(w, screen, element.Value.(*entity))
		}
	}
}

func (w *world) Layout(width, height int) (int, int) {
	w.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func (w *world) ChangeScene(next Scene) {
	w.componentIds = make(map[reflect.Type]int, 256)
	w.componentStores = make([]*store, 0, 256)
	w.entities = list.New()
	w.entitiesIndexes = makePool(256)
	w.systems = make([]*system, 0, 2)
	next.Setup(w)
}

func (w *world) Bounds() image.Rectangle {
	return w.bounds
}

func (w *world) View(components ...interface{}) View {
	return makeView(w, components)
}

func (w *world) AddComponents(components ...interface{}) {
	for _, component := range components {
		componentValue := reflect.ValueOf(component)
		componentType := componentValue.Type()
		if _, ok := w.componentIds[componentType]; ok {
			continue
		}
		w.componentIds[componentType] = len(w.componentIds)
		w.componentStores = append(w.componentStores, makeStore(componentType))
	}
}

func (w *world) AddEntities(entities ...interface{}) {
	for _, e := range entities {
		probablePointer := reflect.ValueOf(e)
		if probablePointer.Kind() != reflect.Ptr {
			panic(fmt.Sprintf("entity %s should be a pointer", typeName(probablePointer.Type())))
		}
		probableStruct := probablePointer.Elem()
		if probableStruct.Kind() != reflect.Struct {
			panic(fmt.Sprintf("entity %s under pointer should be a struct", typeName(probableStruct.Type())))
		}
		components := structFieldTypes(probableStruct)
		en := makeEntity(w, components...)
		en.element = w.entities.PushBack(en)
	}
}

func (w *world) AddSystems(systems ...interface{}) {
	for _, s := range systems {
		w.systems = append(w.systems, makeSystem(w, s))
	}
}

func (w *world) RemoveEntity(e Entity) {
	if v, ok := e.(*entity); ok {
		w.entities.Remove(v.element)
		w.entitiesIndexes.rem(v.id)

		for i := 0; i < len(w.componentStores); i++ {
			if v.mask.get(i) {
				w.componentStores[i].rem(v.id)
			}
		}

		v.w, v.element, v.id, v.mask = nil, nil, -1, nil
	}
}
