package engine

type pool struct {
	free []int
	next int
}

func makePool(capacity int) pool {
	return pool{free: make([]int, 0, capacity)}
}

func (p *pool) get() int {
	if len(p.free) == 0 {
		index := p.next
		p.next++
		return index
	}

	index := p.free[len(p.free)-1]
	p.free = p.free[:len(p.free)-1]
	return index
}

func (p *pool) rem(index int) {
	p.free = append(p.free, index)
}
