package engine

import (
	"image"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type World interface {
	ChangeScene(next Scene)
	Bounds() image.Rectangle
	AddEntity(components ...any)
	AddSystem(system any)
	Get(matchers ...matcher) (entity Entity, ok bool)
	Append(entities []Entity, matchers ...matcher)
	Each(f func(e Entity), matchers ...matcher)
	Components() int
	Entities() int
	Systems() int
}

type world struct {
	bounds        image.Rectangle
	componentPool *componentPool
	entityPool    *entityPool
	updaters      []Updater
	drawers       []Drawer
	first         Scene
	once          sync.Once
	reload        bool
}

func NewGame(first Scene) ebiten.Game {
	return &world{
		componentPool: newComponentPool(),
		entityPool:    newEntityPool(),
		updaters:      make([]Updater, 0, InitialCapacity),
		drawers:       make([]Drawer, 0, InitialCapacity),
		first:         first,
	}
}

func (w *world) Update() error {
	w.once.Do(func() { w.first(w) })

	for _, system := range w.updaters {
		system(w)
	}

	w.reload = false
	w.componentPool.clearGarbage()
	w.entityPool.clearGarbage()

	return nil
}

func (w *world) Draw(screen *ebiten.Image) {
	for _, system := range w.drawers {
		system(w, screen)
	}
}

func (w *world) Layout(width, height int) (int, int) {
	w.bounds = image.Rect(0, 0, width, height)

	return width, height
}

func (w *world) ChangeScene(next Scene) {
	w.reload = true
	w.componentPool = newComponentPool()
	w.entityPool = newEntityPool()
	w.updaters = make([]Updater, 0, InitialCapacity)
	w.drawers = make([]Drawer, 0, InitialCapacity)

	next(w)
}

func (w *world) Bounds() image.Rectangle {
	return w.bounds
}

func (w *world) AddEntity(components ...any) {
	w.entityPool.addEntity(w, components...)
}

func (w *world) AddSystem(system any) {
	if u, ok := system.(Updater); ok {
		w.updaters = append(w.updaters, u)
	}
	if d, ok := system.(Drawer); ok {
		w.drawers = append(w.drawers, d)
	}
}

func (w *world) Get(matchers ...matcher) (entity Entity, ok bool) {
	entity, ok = nil, false

	w.Each(func(e Entity) {
		entity, ok = e, true
	}, matchers...)

	return
}

func (w *world) Append(entities []Entity, matchers ...matcher) {
	w.Each(func(e Entity) {
		entities = append(entities, e)
	}, matchers...)
}

func (w *world) Each(f func(e Entity), matchers ...matcher) {
	query := newMask(InitialCapacity)
	for _, m := range matchers {
		query.set(m.match(w.componentPool))
	}
	for _, e := range w.entityPool.entities {
		if w.reload {
			return
		}
		if e.m.contains(query) {
			f(e)
		}
	}
}

func (w *world) Components() int {
	return len(w.componentPool.components)
}

func (w *world) Systems() int {
	return len(w.updaters) + len(w.drawers)
}

func (w *world) Entities() int {
	return len(w.entityPool.entities)
}
