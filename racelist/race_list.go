package main

import (
	"container/list"
	"fmt"
	"reflect"
	"sync"
)

var (
	lists = newlistener()
)

type (
	able interface {
		Test()
	}
	able1 struct {
	}
	able2 struct {
	}
	funcable func()
)

func (a able1) Test() {
	fmt.Println("able1")
}
func newAlbe1() able {
	return able1{}
}
func (a able2) Test() {
	fmt.Println("able2")
}
func newAlbe2() able {
	return able2{}
}
func (a funcable) Test() {
	fmt.Println("funcable")
	a()
}
func newAlbe3(f func()) able {
	return funcable(f)
}

type listener struct {
	mutexRw sync.RWMutex // 读写锁
	list    *list.List   // 当前list为环形的双向链表
}

func newlistener() listener {
	return listener{
		list: list.New(),
	}
}

func (lists *listener) add(ableargs able) {
	lists.mutexRw.Lock()
	defer lists.mutexRw.Unlock()
	lists.list.PushFront(
		ableargs,
	)
}
func (lists *listener) del(ableargs able) {
	lists.mutexRw.RLock()
	back := lists.list.Back()
	lists.mutexRw.RUnlock()
	for back != nil {
		value := back.Value
		listenerValueOf := reflect.ValueOf(value)
		listenerValueOfKind := listenerValueOf.Kind()
		switch listenerValueOfKind {
		case reflect.Struct:
			if value == ableargs {
				lists.mutexRw.Lock()
				lists.list.Remove(back)
				lists.mutexRw.Unlock()
				break
			}
		case reflect.Ptr:
		case reflect.Func:
			valueR := reflect.ValueOf(ableargs)
			if valueR.Kind() == reflect.Struct {
				if value == ableargs {
					lists.mutexRw.Lock()
					lists.list.Remove(back)
					lists.mutexRw.Unlock()
					break
				}
			} else if valueR.Kind() == reflect.Func {
				if valueR.Pointer() == listenerValueOf.Pointer() {
					lists.mutexRw.Lock()
					lists.list.Remove(back)
					lists.mutexRw.Unlock()
					break
				}
			}
		default:
			continue
		}
		back = back.Prev()
	}
}

// 出现问题 panic: runtime error: comparing uncomparable type main.funcable
func (lists *listener) del1(ableargs able) {
	lists.mutexRw.RLock()
	back := lists.list.Back()
	lists.mutexRw.RUnlock()
	for back != nil {
		value := back.Value
		v, ok := value.(able)
		if !ok {
			continue
		}
		// panic: runtime error: comparing uncomparable type main.funcable
		if v == ableargs {
			lists.mutexRw.Lock()
			lists.list.Remove(back)
			lists.mutexRw.Unlock()
			break
		}
		back = back.Prev()
	}
}

func (lists *listener) notify() {
	lists.mutexRw.RLock()
	defer lists.mutexRw.RUnlock()
	// 优先使用的最后一个
	back := lists.list.Back()
	for back != nil {
		value := back.Value
		if able, ok := value.(able); ok {
			able.Test()
		}
		back = back.Prev()
	}
}
