package mizu

import (
	"image"
	"reflect"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type World interface {
	ChangeScene(next Scene)
	Bounds() image.Rectangle
	AddSystems(systems ...interface{})
	AddEntities(entities ...interface{})
}

type world struct {
	bounds          image.Rectangle
	componentIds    map[reflect.Type]int
	componentStores []*store
	entities        []*entity
	systems         []*system
	first           Scene
	once            sync.Once
}

func NewGame(first Scene) ebiten.Game {
	w := &world{
		componentIds:    make(map[reflect.Type]int, 256),
		componentStores: make([]*store, 0, 256),
		entities:        make([]*entity, 0, 256),
		first:           first,
	}
	return w
}

func (w *world) Update() error {
	w.once.Do(func() {
		w.first.Setup(w)
		w.first = nil
	})
	for _, s := range w.systems {
		for _, e := range w.entities {
			s.update(w, e)
		}
	}
	return nil
}

func (w *world) Draw(screen *ebiten.Image) {
	for _, s := range w.systems {
		for _, e := range w.entities {
			s.draw(screen, e)
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
	w.entities = make([]*entity, 0, 256)
	w.systems = make([]*system, 0, 2)
	next.Setup(w)
}

func (w *world) Bounds() image.Rectangle {
	return w.bounds
}

func (w *world) AddSystems(systems ...interface{}) {
	for _, s := range systems {
		w.systems = append(w.systems, makeSystem(w, s))
	}
}

func (w *world) AddEntities(entities ...interface{}) {
	for _, e := range entities {
		probablePointer := reflect.ValueOf(e)
		if probablePointer.Kind() != reflect.Ptr {
			panic("entity should be a pointer")
		}
		probableStruct := probablePointer.Elem()
		if probableStruct.Kind() != reflect.Struct {
			panic("entity under pointer should be a struct")
		}
		components := structFieldTypes(probableStruct)
		w.addComponents(components...)
		w.addEntity(components...)
	}
}

func (w *world) addComponents(components ...interface{}) {
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

func (w *world) addEntity(components ...interface{}) {
	w.entities = append(w.entities, makeEntity(w, components...))
}
