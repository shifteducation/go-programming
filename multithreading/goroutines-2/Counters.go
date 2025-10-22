package main

import (
	"sync"
)

type Counters struct {
	mutex sync.RWMutex
	m     map[int]int
}

func (c *Counters) Load(key int) (int, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	v, ok := c.m[key]
	return v, ok
}

func (c *Counters) Store(key int, value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.m[key] = value
}
