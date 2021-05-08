package sync

// A Counter can be incremented and holds a counter value.
type Counter struct {
	value int
}

// Inc takes no aruguments, and will increment a Counter's value.
func (c *Counter) Inc() {
	c.value++
}

// Value takes no aruguments, and returns a Counter's value.
func (c *Counter) Value() int {
	return c.value
}
