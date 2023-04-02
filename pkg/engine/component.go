package engine

type component struct {
	id     int
	values []any
	dead   bool
}

func newComponent(id, capacity int) *component {
	return &component{
		id:     id,
		values: make([]any, capacity),
	}
}

func (c *component) setValue(entityID int, v any) {
	if len(c.values) <= entityID {
		c.values = append(c.values, make([]any, nextPowerOf2(entityID))...)
	}

	c.values[entityID] = v
}

func (c *component) remValue(entityID int) {
	c.values[entityID] = nil
}

func (c *component) hasValues() bool {
	for _, v := range c.values {
		if v != nil {
			return true
		}
	}

	return false
}

func nextPowerOf2(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x++
	return x
}
