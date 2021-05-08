package sync

import "sync"

// A Counter can be incremented and holds a counter value.
type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter initializes a new *Counter struct.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc takes no aruguments, and will increment a Counter's value.
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value takes no aruguments, and returns a Counter's value.
func (c *Counter) Value() int {
	return c.value
}
