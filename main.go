package main

import (
	"time"
)

func TimeNow() time.Time {
	return time.Now()
}

type cacheValue struct {
	timeOfDeath time.Time
	value       any
}

type InmemoryCache struct {
	cache   map[string]cacheValue
	getTime func() time.Time
}

func New_cache() InmemoryCache {
	return InmemoryCache{
		cache:   map[string]cacheValue{},
		getTime: TimeNow,
	}
}
func New_cache_with_time(getTime func() time.Time) InmemoryCache {
	return InmemoryCache{
		cache:   map[string]cacheValue{},
		getTime: getTime,
	}
}

func (c *InmemoryCache) Set(key string, value any, ttl_seconds int64) {
	timeOfDeath := time.Now().Add(time.Duration(ttl_seconds) * time.Second)
	c.cache[key] = cacheValue{value: value, timeOfDeath: timeOfDeath}
}
func (c *InmemoryCache) Get(key string) any {
	cacheValue := c.cache[key]
	timeOfDeath := cacheValue.timeOfDeath
	if timeOfDeath.Before(c.getTime()) {
		return nil
	}
	return cacheValue.value
}
func (c *InmemoryCache) Delete(key_to_delete string) {
	delete(c.cache, key_to_delete)
}
func (c *InmemoryCache) Clear() {}

func main() {
	// cache := new_cache()
}
