package main

import (
	"bytes"
	"fmt"
	"sync"
	"unsafe"
)

// import "C"
var (
	w sync.WaitGroup
)
var done = make(chan bool)
var msg string

// [00001000,00000101,00000010...]
func aGoroutine() {
	var msg = "hello, world"
	fmt.Println(msg)

	var bytes = []byte(msg)
	fmt.Println(bytes)
	for _, v := range bytes {
		fmt.Println(fmt.Sprintf("%s ,计算机里是这样表示的:%b", string(v), v))

	}
	msg = "薛"
	fmt.Println(msg)
	bytes = []byte(msg)
	fmt.Println(bytes)

	for _, v := range bytes {
		fmt.Println(fmt.Sprintf("%s ,计算机里是这样表示的:%08b", string(v), uint8(v)))

	}

	t := newTestS("yangbo", addAge(12))
	fmt.Println(t)
	pointer := unsafe.Pointer(t)
	i := (*[1 << 2]testS)(pointer)
	fmt.Println(len(i))
	fmt.Println(i)
	for _, v := range i {
		fmt.Println(v)
	}

	<-done
}

type testS struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func addAge(age int) testOption {
	return func(a *testS) {
		a.Age = age
	}
}

type testOption func(a *testS)

func newTestS(name string, option ...testOption) *testS {
	t := new(testS)
	for _, apply := range option {
		apply(t)
	}
	t.Name = name
	return t
}

func main() {
	go aGoroutine()
	done <- true
	println(msg)
	buffer := bytes.Buffer{}
	buffer.Reset()
}
