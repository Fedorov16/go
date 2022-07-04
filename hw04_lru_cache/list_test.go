package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("only one Front elem", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
	})

	t.Run("There is no elem for removing", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.PushFront(9)
		require.Equal(t, 2, l.Len())
		l.Remove(&ListItem{7, nil, nil})
		require.Equal(t, 2, l.Len())
	})

	t.Run("only one Back elem", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
	})

	t.Run("There is no element for moving to front", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.PushFront(9)
		require.Equal(t, 9, l.Front().Value)
		l.MoveToFront(&ListItem{7, nil, nil})
		require.Equal(t, 9, l.Front().Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestFindElem(t *testing.T) {
	l := NewList()
	_, found := l.FindElem(ListItem{7, nil, nil})
	require.Equal(t, false, found)

	l.PushFront(43)
	l.PushFront(55)
	_, found = l.FindElem(ListItem{7, nil, nil})
	require.Equal(t, false, found)

	elem, found := l.FindElem(ListItem{55, nil, nil})
	require.Equal(t, 55, elem.Value)
	require.Equal(t, true, found)

	// cache
	c := NewCache(3)
	c.Set("a1", 4)
	l = c.(*lruCache).queue
	_, found = l.FindElem(ListItem{cacheItem{"a2", 5}, nil, nil})
	require.Equal(t, false, found)

	r, found := l.FindElem(ListItem{cacheItem{"a1", 10}, nil, nil})
	require.Equal(t, "a1", string(r.Value.(cacheItem).key))
	require.Equal(t, 4, r.Value.(cacheItem).value)
	require.Equal(t, true, found)
}
