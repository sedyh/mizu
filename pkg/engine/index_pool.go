package engine

type indexPool struct {
	free []int
	next int
}

func newIndexPool(capacity int) indexPool {
	return indexPool{free: make([]int, 0, capacity)}
}

func (p *indexPool) get() int {
	if len(p.free) == 0 {
		index := p.next
		p.next++
		return index
	}

	index := p.free[len(p.free)-1]
	p.free = p.free[:len(p.free)-1]
	return index
}

func (p *indexPool) rem(index int) {
	p.free = append(p.free, index)
}
