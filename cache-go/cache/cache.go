package cache

import (
	"sync"
)

//cache with data and mutex lock

type Cache struct {
	data  []Data
	mutex sync.RWMutex
}

type Data struct {
	Key   string
	Value int
}

func NewCache() *Cache {
	return &Cache{
		data:  []Data{},
		mutex: sync.RWMutex{},
	}
}

func (c *Cache) Set(key string, value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = append(c.data, Data{Key: key, Value: value})
}

func (c *Cache) Get(key string) (int, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for _, d := range c.data {
		if d.Key == key {
			return d.Value, true
		}
	}
	return 0, false
}

func (c *Cache) GetAll() []Data {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.data
}
