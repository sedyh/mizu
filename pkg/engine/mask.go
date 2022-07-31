package engine

import "fmt"

type mask []uint64

func newMask(size int) mask {
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

func (m mask) String() string {
	str := ""
	for _, bits := range m {
		str += fmt.Sprintf("%064b", bits)
	}
	return str
}
