package engine

type Entity interface {
	ID() int
	Rem()
	setComponent(component any)
	getComponent(component any) (any, bool)
	hasComponent(component any) bool
	remComponent(component any)
}

type entity struct {
	id   int
	m    mask
	w    *world
	dead bool
}

func newEntity(w *world, id int, components ...any) *entity {
	e := &entity{id: id, m: newMask(InitialCapacity), w: w}

	for _, c := range components {
		w.componentPool.addEntityComponent(e, c)
	}

	return e
}

func (e *entity) ID() int {
	return e.id
}

func (e *entity) Rem() {
	e.w.entityPool.remEntity(e)
}

func (e *entity) setComponent(component any) {
	e.w.componentPool.addEntityComponent(e, component)
}

func (e *entity) getComponent(component any) (any, bool) {
	if c, ok := e.w.componentPool.getComponent(component); ok {
		return c.values[e.id], true
	}

	return nil, false
}

func (e *entity) hasComponent(component any) bool {
	_, ok := e.w.componentPool.getComponent(component)

	return ok
}

func (e *entity) remComponent(component any) {
	e.w.componentPool.remEntityComponent(e, component)
}

func Set(e Entity, component any) {
	e.setComponent(component)
}

func Get[T any](e Entity) *T {
	var x *T

	if c, ok := e.getComponent(x); ok {
		return c.(*T)
	}

	return nil
}

func Has[T any](e Entity) bool {
	var x *T

	return e.hasComponent(x)
}

func Rem[T any](e Entity) {
	var x *T

	e.remComponent(x)
}
