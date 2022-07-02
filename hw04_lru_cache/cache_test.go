package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("check values exists", func(t *testing.T) {
		c := NewCache(10)
		for i := 0; i < 10; i++ {
			exist := c.Set(Key("key"+strconv.Itoa(i)), i*11)
			require.False(t, exist)
		}

		for i := 0; i < 10; i++ {
			exist := c.Set(Key("key"+strconv.Itoa(i)), i*11)
			require.True(t, exist)
		}
	})

	t.Run("set the last added value first in queue", func(t *testing.T) {
		c := createSimpleCacheList(5, 10)

		lastItem := c.(*lruCache).queue.Back().Value.(cacheItem).value
		firstItem := c.(*lruCache).queue.Front().Value.(cacheItem).value
		require.Equal(t, 0, lastItem)
		require.Equal(t, 44, firstItem)
	})

	t.Run("clear cache", func(t *testing.T) {
		c := NewCache(10)

		c.Set("a1", 1)
		c.Set("a2", 2)

		v, ok := c.Get("a1")
		require.True(t, ok)
		require.Equal(t, 1, v)

		v, ok = c.Get("a2")
		require.True(t, ok)
		require.Equal(t, 2, v)

		c.Clear()

		_, ok = c.Get("a1")
		require.False(t, ok)

		_, ok = c.Get("a2")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := createSimpleCacheList(10, 10)
		firstItem := c.(*lruCache).queue.Front().Value.(cacheItem).value
		require.Equal(t, 99, firstItem)

		// first Elem After Set
		c.Set("key5", 123)
		firstItem = c.(*lruCache).queue.Front().Value.(cacheItem).value
		require.Equal(t, 123, firstItem)

		// check Lengh After Set
		require.Equal(t, 10, c.(*lruCache).queue.Len())
		require.Equal(t, 10, len(c.(*lruCache).items))

		// first Elem After Get
		c.Get("key4")
		firstItem = c.(*lruCache).queue.Front().Value.(cacheItem).value
		require.Equal(t, 44, firstItem)

		// check Lengh After Get
		require.Equal(t, 10, c.(*lruCache).queue.Len())
		require.Equal(t, 10, len(c.(*lruCache).items))

		// check length and value after new one
		c.Set("over10", 666)
		require.Equal(t, 10, c.(*lruCache).queue.Len())
		require.Equal(t, 10, len(c.(*lruCache).items))
	})
}

func createSimpleCacheList(count, capping int) Cache {
	c := NewCache(capping)
	for i := 0; i < count; i++ {
		c.Set(Key("key"+strconv.Itoa(i)), i*11)
	}
	return c
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
