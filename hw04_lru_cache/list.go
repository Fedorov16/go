package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
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
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushBack(v interface{}) *ListItem {
	t := &ListItem{v, nil, nil}

	if l.len == 0 {
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

	if l.len == 0 {
		l.tail = t
	} else {
		t.Next = l.head
		l.head.Prev = t
	}

	l.head = t
	l.len++
	return t
}

func (l *list) Remove(cur *ListItem) {
	switch {
	case l.len == 1:
		l.tail = nil
		l.head = nil
	case cur.Next == nil:
		l.tail = cur.Prev
		cur.Prev.Next = nil
	case cur.Prev == nil:
		cur.Next.Prev = nil
	default:
		cur.Prev.Next, cur.Next.Prev = cur.Next, cur.Prev
	}

	l.len--
}

func (l *list) MoveToFront(cur *ListItem) {
	if cur.Prev == nil {
		return
	}

	if cur.Next == nil {
		cur.Prev.Next = nil
		l.tail = cur.Prev
	} else {
		cur.Prev.Next, cur.Next.Prev = cur.Next, cur.Prev
	}

	l.len--
	l.PushFront(cur.Value)
}

func NewList() List {
	return new(list)
}
