package main

// #cgo CFLAGS: -I./func
// #cgo LDFLAGS: -L${SRCDIR}/func -lfunc
// #include "libfunc.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type fakeString struct {
	Data *C.char
	Len  int
}

func main() {
	var s = "helloworld"
	cString := (*fakeString)(unsafe.Pointer(&s))
	cData := cString.Data
	cLen := C.uint(len(s))
	fmt.Println(cLen)

	// load c文件
	C.Add(1, 1)
	C.Print(cData)
}
