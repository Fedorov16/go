package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mu:       sync.Mutex{},
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.items[key]; ok {
		val.Value = cacheItem{key, value}
		c.queue.MoveToFront(val)
		return true
	}

	if len(c.items) >= c.capacity {
		i := c.queue.Back()
		delete(c.items, i.Value.(cacheItem).key)
		c.queue.Remove(i)
	}
	c.items[key] = c.queue.PushFront(cacheItem{key, value})

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, ok := c.items[key]; ok {
		i := item.Value
		c.queue.MoveToFront(item)
		return i.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = new(list)
}
