package main

import (
	"container/list"
	"fmt"
	"strings"
)

type (
	nweman struct {
		name string // name `json:"name"`
		age  int64  // age `json:"age"`
	}
	man struct {
	}
)

// 我是一个接口
type bbk interface {
	// 这是我的处理方法
	handle() error
}

func init() {
	fmt.Println("asdad")
}

func (n nweman) handle() error {
	return nil
}
func (n man) handle() error {
	return nil
}
func main() {
	fmt.Println("这是第一个用vim写的程序")
	// asda
	n := new(nweman)
	n.age = 10
	n.name = "asda"
	strlsif := strings.Split(string("as.dad"), ".")
	for i, v := range strlsif {
		fmt.Println(i)
		fmt.Println(v)
	}
	_ = n.handle()

	// `开始一个lru`
	fmt.Println("`开始lru检查``")
	lt := newLruObj(5)
	lt.Put("key1", "key1")
	lt.Put("key2", "key2")
	lt.Put("key3", "key3")
	flo := lt.listLruTmp.Front()
	for flo != nil {
		fmt.Printf("flo.Value: %v\n", flo.Value)
		flo = flo.Next()
	}
	lt.Get("key1")
	flo = lt.listLruTmp.Front()
	for flo != nil {
		fmt.Printf("flo.Value: %v\n", flo.Value)
		flo = flo.Next()
	}
	lt.Put("key4", "key4")
	flo = lt.listLruTmp.Front()
	for flo != nil {
		fmt.Printf("flo.Value: %v\n", flo.Value)
		flo = flo.Next()
	}
	flo = lt.listLruTmp.Front()
	for flo != nil {
		fmt.Printf("flo.Value: %v\n", flo.Value)
		flo = flo.Next()
	}
	lt.Get("key3")

	lt.Put("key6", "key6")
	lt.Put("key7", "key7")

	lt.Put("key8", "key8")

	flo = lt.listLruTmp.Front()
	for flo != nil {
		fmt.Printf("flo.Value: %v\n", flo.Value)
		flo = flo.Next()
	}

}

type LruHandler interface {
	Get(key string) string
	Put(key, value string)
}

// LruObj map[key]ele;list
type LruObj struct {
	cap        uint                     // 阈值
	listLruTmp *list.List               // 缓存队列
	mapTmp     map[string]*list.Element // map唯一性
}

func newLruObj(cap uint) LruObj {
	lo := LruObj{}
	lo.listLruTmp = list.New()
	lo.mapTmp = make(map[string]*list.Element)
	lo.cap = cap
	return lo
}

func (l LruObj) Get(key string) string {
	// `如果存在,将结果放置对头`
	// 将当前节点放置为头节点
	// 当前节点next指向当前的head节点,当前节点的prev节点指向tail
	// 当前的节点的prev节点指向当前节点的next
	if v, ok := l.mapTmp[key]; ok {
		e := v
		// prevEle := v.Prev()
		// nextEle := v.Next()
		// prevEle.Next() = v.n
		// v.Prev().Next() = nextEle
		l.listLruTmp.MoveToFront(v)
		s, _ := e.Value.(string)
		return s
	}
	return ""
}

func (l LruObj) Put(key, value string) {
	// 获取当前元素是否存在,
	// 存在,讲当前元素添加到head
	// 不存在,直接讲元素添加到head
	// 检查当前队列大小是否超出阈值
	if v, ok := l.mapTmp[key]; ok {
		l.listLruTmp.PushFront(v.Value)
		delete(l.mapTmp, key)
		l.mapTmp[key] = l.listLruTmp.Front()
	} else {
		// 检查当前缓存是否大于阈值:大于将队尾nil,当前值插入对头,小余直接讲数据插入对头
		if l.listLruTmp.Len() > int(l.cap) {
			fmt.Println("chuxian len >")
			l.listLruTmp.Remove(l.listLruTmp.Back())
		}
		l.listLruTmp.PushFront(value)
		l.mapTmp[key] = l.listLruTmp.Front()
	}
}

func (receive LruObj) handle() error {
	return nil
}

func handle(args nweman) {
	fmt.Println(args)
}
