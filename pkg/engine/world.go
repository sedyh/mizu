package engine

import (
	"fmt"
	"image"
	"reflect"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// World is the main interface of the framework, which implements work with
// entities, components and systems, allows you to change scenes during the game,
// as well as create your own queries for entities.
type World interface {
	ChangeScene(next Scene)                  // Switch scenes.
	Bounds() image.Rectangle                 // Returns current physical window size.
	View(components ...interface{}) View     // Creates a query to filter entities by their components.
	AddComponents(components ...interface{}) // Registers the used components of your object properties.
	AddSystems(systems ...interface{})       // Adds a system that will work with each entity by its components.
	AddEntities(entities ...interface{})     // Adds an entities that will represent your objects.
	RemoveEntity(entity Entity)              // Removes an entity.
	GetEntity(id int) (e Entity, ok bool)    // Returns an entity by id and a search status.
	Components() int                         // Get current amount of registered components.
	Systems() int                            // Get current amount of added systems.
	Entities() int                           // Get current amount of added entities.
}

// world is an internal struct, which implements both engine.World and ebiten.Game interfaces
type world struct {
	bounds       image.Rectangle
	componentIds map[reflect.Type]int
	entitiesIds  pool
	stores       []*store
	systems      []*system
	entities     []*entity
	first        Scene
	once         sync.Once
}

// NewGame creates world and returns ebiten.Game, which you can use right away,
// or embed in your own ebiten.Game implementation if you want to add your own
// behavior there (for example, change the logical resolution which is same
// as physical by default).
func NewGame(first Scene) ebiten.Game {
	w := &world{
		componentIds: make(map[reflect.Type]int, 256),
		stores:       make([]*store, 0, 256),
		entities:     make([]*entity, 0, 256),
		entitiesIds:  makePool(256),
		first:        first,
	}
	return w
}

// Update implements Update() error from ebiten.Game.
func (w *world) Update() error {
	w.once.Do(func() { w.first.Setup(w) })
	for _, s := range w.systems {
		for _, e := range w.entities {
			s.updateForEachEntity(w, e)
		}
		s.updateForEachTick(w)
	}
	return nil
}

// Draw represents Draw(screen *ebiten.Image) from ebiten.Game.
func (w *world) Draw(screen *ebiten.Image) {
	for _, s := range w.systems {
		for _, e := range w.entities {
			s.drawForEachEntity(w, screen, e)
		}
		s.drawForEachTick(w, screen)
	}
}

// Layout represents Layout(width, height int) (int, int) from ebiten.Game.
func (w *world) Layout(width, height int) (int, int) {
	w.bounds = image.Rect(0, 0, width, height)
	return width, height
}

// ChangeScene resets the world and switches the scene to a new one.
func (w *world) ChangeScene(next Scene) {
	w.componentIds = make(map[reflect.Type]int, 256)
	w.stores = make([]*store, 0, 256)
	w.entities = make([]*entity, 0, 256)
	w.entitiesIds = makePool(256)
	w.systems = make([]*system, 0, 2)
	next.Setup(w)
}

// Bounds returns current physical window size.
// If you want to change this behavior, then you can embed this implementation
// of the ebiten.Game in your own and call the Layout(width, height int) (int, int)
// of this implementation manually.
func (w *world) Bounds() image.Rectangle {
	return w.bounds
}

// View creates a query to filter entities by their components.
func (w *world) View(components ...interface{}) View {
	return makeView(w, components...)
}

// AddComponents registers the used components that will represent your object properties.
func (w *world) AddComponents(components ...interface{}) {
	for _, component := range components {
		componentType := reflect.TypeOf(component)
		if _, ok := w.componentIds[componentType]; ok {
			continue
		}
		w.componentIds[componentType] = len(w.componentIds)
		w.stores = append(w.stores, makeStore(componentType))
	}
}

// AddEntities adds entities that will represent your game objects
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
		w.entities = append(w.entities, makeEntity(w, components...))
	}
}

// AddSystems adds systems that will work with each entity by its components
func (w *world) AddSystems(systems ...interface{}) {
	for _, s := range systems {
		w.systems = append(w.systems, makeSystem(w, s))
	}
}

// GetEntity finds an entity by its id and returns status code
func (w *world) GetEntity(id int) (e Entity, ok bool) {
	for _, e := range w.entities {
		if e.id == id {
			return e, true
		}
	}

	return nil, false
}

// RemoveEntity removes an entity
func (w *world) RemoveEntity(e Entity) {
	if v, ok := e.(*entity); ok {
		for i, candidate := range w.entities {
			if candidate.id == v.id {
				copy(w.entities[i:], w.entities[i+1:])
				w.entities[len(w.entities)-1] = nil
				w.entities = w.entities[:len(w.entities)-1]

				for j := 0; j < len(w.stores); j++ {
					if v.mask.get(j) {
						w.stores[j].rem(v.id)
					}
				}
				w.entitiesIds.rem(v.id)
				v.w, v.id, v.mask = nil, -1, nil

				break
			}
		}
	}
}

// Components returns current amount of registered components.
func (w *world) Components() int {
	return len(w.stores)
}

// Systems returns current amount of added systems.
func (w *world) Systems() int {
	return len(w.systems)
}

// Entities returns current amount of added entities.
func (w *world) Entities() int {
	return len(w.entities)
}
