package main

import (
	"fmt"
	"time"
	"unsafe"
)

var (
	chan1 = make(chan struct{}, 1)
	// chan2 = make(chan struct{} ,1)
)

func main() {
	i := int64(12132313)
	go test(i)

	ticker := time.NewTicker(time.Second * 6)
	select {
	case <-chan1:
		fmt.Println("test结束")
	case <-ticker.C:
		fmt.Println("定时器")
	}

	fmt.Println("会阻塞吗?")

}

// /can only use //go:noescape with external func implementations
// /go:noescape
func test(arg int64) {
	abc(arg)
}

func abc(int642 int64) {
	time.Sleep(time.Second * 5)
	sprintf := fmt.Sprintf("int:%d", int642)
	fmt.Println(sprintf)
	fmt.Println(unsafe.Sizeof(int642))
	chan1 <- struct{}{}
}
