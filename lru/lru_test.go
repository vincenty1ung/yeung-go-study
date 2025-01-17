package lru

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	LRUconstructorInt := Constructor[int](3)
	LRUconstructorInt.Put(1, 12)
	LRUconstructorInt.Put(3, 32)
	LRUconstructorInt.Put(4, 42)
	fmt.Println("+++++++++++++")
	front := LRUconstructorInt.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}

	get := LRUconstructorInt.Get(3)
	fmt.Println(get)
	fmt.Println("+++++++++++++")
	front = LRUconstructorInt.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}
	LRUconstructorInt.Put(5, 52)
	fmt.Println("+++++++++++++")
	front = LRUconstructorInt.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}

	get = LRUconstructorInt.Get(3)
	fmt.Println(get)
	fmt.Println("+++++++++++++")
	front = LRUconstructorInt.list.Front()
	for front != nil {
		fmt.Println(front.Value)
		front = front.Next()
	}

	constructor := Constructor[string](3)
	constructor.Put("nihao", "nihao")
	fmt.Println(constructor.Get("nihao"))
	fmt.Println(constructor.Get("no"))
}
