package memstore

import (
	"sync"
)

type Cache struct {
	data  map[string]interface{}
	mutex sync.RWMutex
}

func newCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) value(name string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	v, ok := c.data[name]
	return v, ok
}

func (c *Cache) setValue(name string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[name] = value
}

func (c *Cache) delete(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.data[name]; ok {
		delete(c.data, name)
	}
}
