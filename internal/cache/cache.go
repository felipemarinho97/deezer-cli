package cache

import (
	"encoding/json"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

type Cache struct {
	store *gocache.Cache
}

func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	return &Cache{
		store: gocache.New(defaultExpiration, cleanupInterval),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	data, err := json.Marshal(value)
	if err == nil {
		c.store.SetDefault(key, data)
	}
}

func (c *Cache) Get(key string, target interface{}) bool {
	if data, found := c.store.Get(key); found {
		if bytes, ok := data.([]byte); ok {
			return json.Unmarshal(bytes, target) == nil
		}
	}
	return false
}

func (c *Cache) Delete(key string) {
	c.store.Delete(key)
}

func (c *Cache) Flush() {
	c.store.Flush()
}