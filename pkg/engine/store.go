package engine

import "reflect"

// store is a slice for the values of components of the same type.
type store struct {
	zero     reflect.Value
	slice    reflect.Value
	capacity int
}

// makeStore creates new store of specific component type.
func makeStore(componentType reflect.Type) *store {
	capacity := 256
	return &store{
		zero:     reflect.Zero(componentType),
		slice:    reflect.MakeSlice(reflect.SliceOf(componentType), capacity, capacity),
		capacity: capacity,
	}
}

// add appends new component.
func (s *store) add(id int, value reflect.Value) {
	s.ensureCapacity(id + 1)
	s.set(id, value)
}

// get finds the value from the slice and sets this value to the passed pointer of the same type.
func (s *store) get(id int, value reflect.Value) {
	value.Set(s.slice.Index(id).Addr())
}

// set replaces the value of the specified element.
func (s *store) set(id int, value reflect.Value) {
	s.slice.Index(id).Set(value)
}

// rem zeroes the value of the specified element.
func (s *store) rem(id int) {
	s.slice.Index(id).Set(s.zero)
}

// ensureCapacity extends the slice if it needed.
func (s *store) ensureCapacity(capacity int) {
	if capacity < s.capacity {
		return
	}
	capacity = nextPowerOf2(capacity)
	slice := reflect.MakeSlice(s.slice.Type(), capacity, capacity)
	reflect.Copy(slice, s.slice)
	s.slice, s.capacity = slice, capacity
}

// nextPowerOf2 returns the next power of 2 to extend the slice.
func nextPowerOf2(x int) int {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}
