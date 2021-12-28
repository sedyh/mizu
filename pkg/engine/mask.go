package engine

import "fmt"

// mask is an implementation of an expanding bitmask.
type mask []uint64

// makeMask creates new bitmask.
func makeMask(size int) mask {
	return make(mask, 1+((size-1)/64))
}

// get the requested bit.
func (m mask) get(index int) bool {
	return m[index/64]&(1<<(index%64)) != 0
}

// set a bit in the bitmask.
func (m mask) set(index int) {
	m[index/64] |= 1 << (index % 64)
}

// clear frees a bit in the bitmask.
func (m mask) clear(index int) {
	m[index/64] &^= 1 << (index % 64)
}

// contains checks if a bit in the bitmask is set to 1.
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
