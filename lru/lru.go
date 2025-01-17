package lru

import (
	"container/list"
)

type Cache[T comparable] struct {
	capacity int
	cache    map[T]*list.Element
	list     *list.List
}
type Pair[T comparable] struct {
	key   T
	value any
}

func Constructor[T comparable](capacity int) Cache[T] {
	return Cache[T]{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[T]*list.Element),
	}
}

func (c *Cache[T]) Get(key T) any {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		if pair, ok := elem.Value.(Pair[T]); ok {
			return pair.value
		}
	}
	return nil
}

func (c *Cache[T]) Put(key T, value any) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value = Pair[T]{key, value}
	} else {
		if c.list.Len() >= c.capacity {
			if pair, ok := c.list.Back().Value.(Pair[T]); ok {
				delete(c.cache, pair.key)
			}
			c.list.Remove(c.list.Back())
		}
		c.list.PushFront(Pair[T]{key, value})
		c.cache[key] = c.list.Front()
	}
}
