package engine

type matcher interface {
	match(w *componentPool) int
}

type andMatcher[T any] struct{}

func And[T any]() *andMatcher[T] {
	return &andMatcher[T]{}
}

func (m *andMatcher[T]) match(p *componentPool) int {
	var x *T

	c, ok := p.getComponent(x)
	if !ok {
		c = p.addComponent(x)
	}

	return c.id
}
