package hw04lrucache

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
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if val, ok := c.items[key]; ok {
		val.Value = cacheItem{key, value}
		c.queue.MoveToFront(val)
		return true
	}

	if len(c.items) >= c.capacity {
		i := c.queue.Back().Value.(cacheItem).key

		delete(c.items, i)
		elem := c.queue.Back()
		c.queue.Remove(elem)
	}
	c.items[key] = c.queue.PushFront(cacheItem{key, value})

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		i := item.Value
		c.queue.MoveToFront(item)
		return i.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = new(list)
}
