package engine

// pool is set of indexes, it can issue a recently freed index or create a new one.
type pool struct {
	free []int
	next int
}

// makePool creates new pool with specific capacity, after freeing the index can be reused.
func makePool(capacity int) pool {
	return pool{free: make([]int, 0, capacity)}
}

// get returns last free index
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

// rem frees the index for reuse
func (p *pool) rem(index int) {
	p.free = append(p.free, index)
}
