package engine

import "reflect"

type store struct {
	zero     reflect.Value
	slice    reflect.Value
	capacity int
}

func makeStore(componentType reflect.Type) *store {
	capacity := 256
	return &store{
		zero:     reflect.Zero(componentType),
		slice:    reflect.MakeSlice(reflect.SliceOf(componentType), capacity, capacity),
		capacity: capacity,
	}
}

func (s *store) add(id int, value reflect.Value) {
	s.ensureCapacity(id + 1)
	s.set(id, value)
}

func (s *store) get(id int, value reflect.Value) {
	value.Set(s.slice.Index(id).Addr())
}

func (s *store) set(id int, value reflect.Value) {
	s.slice.Index(id).Set(value)
}

func (s *store) rem(id int) {
	s.slice.Index(id).Set(s.zero)
}

func (s *store) ensureCapacity(capacity int) {
	if capacity < s.capacity {
		return
	}
	capacity = nextPowerOf2(capacity)
	slice := reflect.MakeSlice(s.slice.Type(), capacity, capacity)
	reflect.Copy(slice, s.slice)
	s.slice, s.capacity = slice, capacity
}

func nextPowerOf2(x int) int {
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}
