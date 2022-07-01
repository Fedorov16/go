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

	cur.Prev.Next, cur.Next.Prev = cur.Next, cur.Prev

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.head == i {
		return
	}

	if l.tail == i {
		i.Prev.Next = nil
		l.PushFront(i.Value)
		return
	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	l.PushFront(i.Value)
}

func (l *list) CheckElem(v ListItem) (*ListItem, bool) {
	if l.len == 0 {
		return nil, false
	}
	cur := *l.head
	for cur != v && cur.Next != nil {
		cur = *l.head.Next
	}

	if cur != v {
		return nil, false
	}

	return &cur, true
}

func NewList() List {
	return new(list)
}
