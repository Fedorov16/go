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
	val := cacheItem{key, value}
	_, exist := c.items[key]
	if exist {
		c.queue.MoveToFront(&ListItem{val, nil, nil})
		c.items[key] = c.queue.Front()
		return exist
	}
	c.items[key] = c.queue.PushFront(val)
	if len(c.items) > c.capacity {
		i := c.queue.Back().Value.(cacheItem).key

		delete(c.items, i)
		elem := c.queue.Back()
		c.queue.Remove(elem)
	}

	return exist
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	i := item.Value
	c.queue.MoveToFront(&ListItem{i, nil, nil})
	return i.(cacheItem).value, true
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = new(list)
}
