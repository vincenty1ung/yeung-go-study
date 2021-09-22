package main

import (
	"fmt"
	"sort"

	"github.com/tal-tech/go-zero/core/fx"
)

// Slice 线程不安全扩容出现null
func main() {
	TestNamec4()
	var ints [8]int
	is := ints[0:4]
	fmt.Println(is)
	fmt.Println(len(is))
	fmt.Println(cap(is))
	is = append(is, 2)
	fmt.Println(is)
	fmt.Println(len(is))
	fmt.Println(cap(is))
}

type man struct{ Age int }

func (m *man) String() string { return "age:" + fmt.Sprintf("%d", m.Age) }
func TestNamec4() {
	mens := make([]*man, 0, 1)
	funcs := make([]func(), 0)
	for i := 0; i < 100; i++ {
		i1 := i
		funcs = append(
			funcs, func() {
				bu(i1, &mens)
			},
		)
	}
	fmt.Println(len(funcs))
	fx.Parallel(funcs...)
	fmt.Println(len(mens))
	fmt.Println(fmt.Sprintf("%+v", mens))
	sort.Slice(
		mens, func(i, j int) bool {
			return mens[i].Age < mens[j].Age
		},
	)
	fmt.Println(fmt.Sprintf("%+v", mens))
}
func bu(age int, mans *[]*man) {
	m := new(man)
	if age == 3 {
		m = nil
	} else {
		m.Age = age
	}
	if m != nil {
		*mans = append(*mans, m)
	}
}

func addInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 判断 zlen 是否超过了x的cap
		// 没有超过 直接赋值
		z = x[:zlen]
	} else { // 如果zlen 超过了x的cap
		zcap := zlen // 将zlen赋值非zcap
		// 这儿相当于 将x的cap 扩容了一倍, 和原本的append函数一样
		if zcap <= 2*len(x) {
			zcap = 2 * len(x)
		}
		// 初始化新切片，将原来的值复制过去
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// 最后一个元素为y 返回z
	z[len(x)] = y

	return z
}
