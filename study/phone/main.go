package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

const (
	SampleRate  = 44100
	BufferSize  = 256
	NumChannels = 1
)

func main() {
	/*atr := int32(11)
	mnum := test(atr, 0, 1, 2, 3, 4, 5, 6)
	_ = mnum
	go func() {
		a := int32(1 + 1)
		if a > mnum {
			a = 12
		}
		a = 10
	}()
	*/
	create, _ := os.Create("trace.out")
	defer create.Close()
	trace.Start(create)
	defer trace.Stop()

	go f()

	time.Sleep(30 * time.Second)
	runtime.GC()
	fmt.Println("ok")
}

// go tool trace trace.out
func test(a, b, c, d, e, f, g, h int32) int32 {
	k := int32(10)
	return a + k
}

func f() {
	for {
		_ = make([]byte, 1<<32)
	}
}
