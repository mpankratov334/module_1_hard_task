package cache

import "sync"

// Interface - реализуйте этот интерфейс
type Interface interface {
	Set(k string, v string)
	Get(k string) (v string, ok bool)
}

// Не меняйте названия структуры и название метода создания экземпляра Cache, иначе не будут проходить тесты

type Cache struct {
	_map map[string]string
	sync.RWMutex
}

// NewCache создаёт и возвращает новый экземпляр Cache.
func NewCache() Interface {
	return &Cache{_map: make(map[string]string)}
}

func (c *Cache) Set(k, v string) {
	c.RWMutex.Lock()
	c._map[k] = v
	c.RWMutex.Unlock()
}

func (c *Cache) Get(k string) (v string, ok bool) {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	v, ok = c._map[k]
	return v, ok
}
