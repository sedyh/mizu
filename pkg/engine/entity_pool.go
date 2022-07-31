package engine

type entityPool struct {
	entities []*entity
	pool     indexPool
	garbage  bool
}

func newEntityPool() *entityPool {
	return &entityPool{
		make([]*entity, 0, InitialCapacity),
		newIndexPool(InitialCapacity),
		false,
	}
}

func (p *entityPool) addEntity(w *world, components ...any) *entity {
	e := newEntity(w, p.pool.get(), components...)
	p.entities = append(p.entities, e)

	return e
}

func (p *entityPool) remEntity(e *entity) {
	p.garbage = true
	e.dead = true
}

func (p *entityPool) clearGarbage() {
	if !p.garbage {
		return
	}
	p.garbage = false

	i := 0
	for _, e := range p.entities {
		if !e.dead {
			p.entities[i] = e
			i++
		}
	}
	for j := i; j < len(p.entities); j++ {
		p.entities[j] = nil
	}
	p.entities = p.entities[:i]
}
