// func.go
package main

import "C"
import (
	"fmt"
)

//export Add
func Add(a C.int, b C.int) C.int {
	return a + b
}

//export Print
func Print(s *C.char) {
	/*
	   函数参数可以用 string, 但是用*C.char更通用一些。
	   由于string的数据结构，是可以被其它go程序调用的，
	   但其它语言（如 python）就不行了
	*/
	print("Hello ", C.GoString(s)) // 这里不能用fmt包，会报错，调了很久...
}

//export SayHello
func SayHello(name string) {
	fmt.Printf("func in Golang SayHello says: Hello, %s!\n", name)
}

//export SayHelloByte
func SayHelloByte(name []byte) {
	fmt.Printf("func in Golang SayHelloByte says: Hello, %s!\n", string(name))
}

//export SayBye
func SayBye() {
	fmt.Println("func in Golang SayBye says: Bye!")
}

func main() {
}
