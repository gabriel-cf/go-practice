package syncing

import (
	"sync"
)

type Counter struct {
	value int
	mutx  sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mutx.Lock()
	defer c.mutx.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
