package main

import (
	"fmt"
)

func addSub(a, b int) (int, int) {
	return a + b, a - b
}

func add(a, b int) int // 汇编函数声明

func sub(a, b int) int // 汇编函数声明

func mul(a, b int) int // 汇编函数声明

func main() {
	sub, i := addSub(333, 222)
	fmt.Println(sub, i)
	fmt.Println(mul(1, 2))

}

// go tool compile -S -N -l main.go
