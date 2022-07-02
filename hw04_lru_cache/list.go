package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	FindElem(v ListItem) (*ListItem, bool)
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
	cur, exist := l.FindElem(*i)
	if !exist {
		return
	}
	cur.Value = i.Value

	_, curKey, tailKey := getKeys(*l, *cur)

	if tailKey == curKey {
		l.tail.Prev.Next = nil
		l.tail = nil
	} else {
		cur.Prev.Next, cur.Next.Prev = cur.Next, cur.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	cur, ok := l.FindElem(*i)
	if !ok {
		return
	}
	cur.Value = i.Value

	headKey, curKey, tailKey := getKeys(*l, *cur)

	if headKey == curKey {
		return
	}

	if tailKey == curKey {
		l.tail.Prev.Next = nil
		l.tail = l.tail.Prev
	} else {
		cur.Prev.Next, cur.Next.Prev = cur.Next, cur.Prev
	}

	l.len--
	l.PushFront(cur.Value)
}

func NewList() List {
	return new(list)
}

func (l *list) FindElem(v ListItem) (*ListItem, bool) {
	if l.len == 0 {
		return nil, false
	}

	switch l.head.Value.(type) {
	case cacheItem:
		cur := *l.head
		for cur.Value.(cacheItem).key != v.Value.(cacheItem).key && cur.Next != nil {
			cur = *cur.Next
		}

		if cur.Value.(cacheItem).key != v.Value.(cacheItem).key {
			return nil, false
		}
		return &cur, true
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

func getKeys(l list, cur ListItem) (interface{}, interface{}, interface{}) {
	var headKey, curKey, tailKey interface{}
	_, isCache := cur.Value.(cacheItem)
	if isCache {
		headKey = l.head.Value.(cacheItem).key
		curKey = cur.Value.(cacheItem).key
		tailKey = l.tail.Value.(cacheItem).key
	} else {
		headKey = l.head.Value
		tailKey = l.tail.Value
		curKey = cur.Value
	}
	return headKey, curKey, tailKey
}
