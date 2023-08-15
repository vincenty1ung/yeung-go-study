package main

import (
	"fmt"
	"sync"
)

// go run -race  main.go
func main() {
	// mapSIOSafe()
	mapSIONoSafe()
	// mapRaceIO()
}

// map都在读的时候虽然存在线程安全问题 但是不会panic，但是出现读写时候出现panic
func mapRaceIO() {
	m := make(map[string]string)
	m["1"] = "1"
	m["2"] = "2"
	var w sync.WaitGroup
	w.Add(2)

	go func(m map[string]string, w *sync.WaitGroup) {
		fmt.Println(m["1"])
		w.Done()
	}(m, &w)
	go func(m map[string]string, w *sync.WaitGroup) {
		m["1"] = "2"
		fmt.Println(m["1"])
		w.Done()
	}(m, &w)

	w.Wait()
}

// 当前传递man副本无竞争关系 线程安全
func mapSIOSafe() {

	type man struct {
		name string
		age  int
	}
	m := &man{
		age: 10, name: "s",
	}
	var w sync.WaitGroup
	w.Add(2)

	go func(m man, w *sync.WaitGroup) {
		fmt.Println(m.name)
		w.Done()
	}(*m, &w)
	go func(m man, w *sync.WaitGroup) {
		m.name = "q"
		fmt.Println(m.name)
		w.Done()
	}(*m, &w)

	w.Wait()
}
func mapSIONoSafe() {

	type man struct {
		name string
		age  int
	}
	m := &man{
		age: 10, name: "s",
	}
	var w sync.WaitGroup
	w.Add(2)

	go func(m *man, w *sync.WaitGroup) {
		fmt.Println(m.name)
		w.Done()
	}(m, &w)
	go func(m *man, w *sync.WaitGroup) {
		m.name = "q"
		fmt.Println(m.name)
		w.Done()
	}(m, &w)

	w.Wait()
}
