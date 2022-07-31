package engine

import (
	"fmt"
	"reflect"
)

type componentPool struct {
	components map[reflect.Type]*component
	pool       indexPool
	garbage    bool
}

func newComponentPool() *componentPool {
	return &componentPool{
		make(map[reflect.Type]*component),
		newIndexPool(InitialCapacity),
		false,
	}
}

func (p *componentPool) getComponent(component any) (c *component, ok bool) {
	assertKind(component, reflect.Ptr, func() {
		panic(fmt.Sprintf("Entity component %T should be a pointer", component))
	})

	c, ok = p.components[reflect.TypeOf(component)]
	return
}

func (p *componentPool) addEntityComponent(entity *entity, component any) {
	c := p.addComponent(component)
	c.setValue(entity.id, component)
	entity.m.set(c.id)
}

func (p *componentPool) remEntityComponent(entity *entity, component any) {
	t := reflect.TypeOf(component)
	c, ok := p.components[t]
	if !ok {
		return
	}

	c.remValue(entity.id)
	entity.m.clear(c.id)

	if !c.hasValues() {
		p.garbage = true
		c.dead = true
	}
}

func (p *componentPool) addComponent(component any) *component {
	assertKind(component, reflect.Ptr, func() {
		panic(fmt.Sprintf("Entity component %T should be a pointer", component))
	})

	t := reflect.TypeOf(component)
	c, ok := p.components[t]
	if !ok {
		c = newComponent(p.pool.get(), InitialCapacity)
		p.components[t] = c
	}

	return c
}

func (p *componentPool) clearGarbage() {
	if !p.garbage {
		return
	}
	p.garbage = false

	for t, c := range p.components {
		if !c.dead {
			continue
		}

		p.pool.rem(c.id)
		delete(p.components, t)
	}
}
