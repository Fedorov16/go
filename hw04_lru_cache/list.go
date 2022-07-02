package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	CheckElem(v ListItem) (*ListItem, bool)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head *ListItem
	tail *ListItem
	len  int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.head
}

func (l list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

func (l *list) PushBack(v interface{}) *ListItem {
	t := &ListItem{v, nil, nil}

	if l.head == nil && l.tail == nil {
		l.head = t
		l.tail = t
	} else {
		t.Prev = l.tail

		l.tail.Next = t
		l.tail = t
	}

	l.len++
	return l.head
}

func (l *list) PushFront(v interface{}) *ListItem {
	t := &ListItem{v, nil, nil}

	if l.head == nil && l.tail == nil {
		l.head = t
		l.tail = t
	} else {
		t.Next = l.head

		l.head.Prev = t
		l.head = t
	}

	l.len++
	return l.head
}

func (l *list) Remove(i *ListItem) {
	cur, ok := l.CheckElem(*i)
	if !ok {
		return
	}

	switch l.head.Value.(type) {
	case cacheItem:
		if l.tail.Value.(cacheItem).key == cur.Value.(cacheItem).key {
			l.tail.Prev.Next = nil
			l.tail = nil
			l.len--
			return
		}
	default:
		if l.tail.Value == cur.Value {
			l.tail.Prev.Next = nil
			l.tail = nil
			l.len--
			return
		}
	}

	cur.Prev.Next, cur.Next.Prev = cur.Next, cur.Prev

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	cur, ok := l.CheckElem(*i)
	if !ok {
		return
	}

	var key, headKey, tailKey interface{}

	switch l.head.Value.(type) {
	case cacheItem:
		key = cur.Value.(cacheItem).key
		headKey = l.head.Value.(cacheItem).key
		tailKey = l.tail.Value.(cacheItem).key
	default:
		key = cur.Value
		headKey = l.head.Value
		tailKey = l.tail.Value
	}

	if headKey == key {
		return
	}

	if tailKey == key {
		l.tail.Prev.Next = nil
		l.tail = l.tail.Prev
		l.len--
		l.PushFront(cur.Value)
		return
	}

	cur.Prev.Next = cur.Next
	cur.Next.Prev = cur.Prev
	l.len--
	l.PushFront(cur.Value)
}

func (l *list) CheckElem(v ListItem) (*ListItem, bool) {
	if l.len == 0 {
		return nil, false
	}

	switch l.head.Value.(type) {
	case cacheItem:
		return checkCacheItem(*l, v)
	default:
		cur := *l.head
		for cur.Value != v.Value && cur.Next != nil {
			cur = *cur.Next
		}

		if cur.Value != v.Value {
			return nil, false
		}

		return &cur, true
	}
}

func checkCacheItem(l list, v ListItem) (*ListItem, bool) {
	cur := l.head
	for cur.Value.(cacheItem).key != v.Value.(cacheItem).key && cur.Next != nil {
		cur = cur.Next
	}

	if cur.Value.(cacheItem).key != v.Value.(cacheItem).key {
		return nil, false
	}
	cur.Value = v.Value
	return cur, true
}

func NewList() List {
	return new(list)
}
