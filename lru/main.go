package main

import (
	"container/list"
	"fmt"
)

func main() {
	LRUconstructor := Constructor(3)
	LRUconstructor.Put(1, 12)
	LRUconstructor.Put(3, 32)
	LRUconstructor.Put(4, 42)
	fmt.Println("+++++++++++++")
	front := LRUconstructor.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}

	get := LRUconstructor.Get(3)
	fmt.Println(get)
	fmt.Println("+++++++++++++")
	front = LRUconstructor.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}
	LRUconstructor.Put(5, 52)
	fmt.Println("+++++++++++++")
	front = LRUconstructor.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}

	get = LRUconstructor.Get(3)
	fmt.Println(get)
	fmt.Println("+++++++++++++")
	front = LRUconstructor.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}
}

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}
type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if this.list.Len() >= this.capacity {
			delete(this.cache, this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(Pair{key, value})
		this.cache[key] = this.list.Front()
	}
}
