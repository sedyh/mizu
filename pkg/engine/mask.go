package engine

type mask []uint64

func makeMask(size int) mask {
	return make(mask, 1+((size-1)/64))
}

func (m mask) get(index int) bool {
	return m[index/64]&(1<<(index%64)) != 0
}

func (m mask) set(index int) {
	m[index/64] |= 1 << (index % 64)
}

func (m mask) clear(index int) {
	m[index/64] &^= 1 << (index % 64)
}

func (m mask) contains(mask mask) bool {
	for i, bits := range m {
		if bits&mask[i] != mask[i] {
			return false
		}
	}
	return true
}
