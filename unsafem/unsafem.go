package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
	"web-test/unsafem/empty1"
	"web-test/unsafem/empty2"
)

type Human struct {
	sex  bool   // 1byte
	age  uint8  // 1byte
	min  int    // 机器64位 8byte
	name string // 16byte
}
type service interface {
	handler()
}
type serviceImpl struct {
}

func (s serviceImpl) handler() {
	fmt.Println("handler")
}

func main() {
	Test1()

	// i :=6
test2:
	fmt.Println("321002902020-2")

	if true {
		goto test
	} else {
		goto test2
	}
	fmt.Println("321002902020-2")

test:
	fmt.Println("1232")
	fmt.Println("222")

	fmt.Println("333")
	fmt.Println("========")
	now := time.Now()
	bytes := unsafeFastStringToReadOnlyBytes("a")
	fmt.Println(bytes, len(bytes), cap(bytes))
	onlyString := unsafeFastBytesToReadOnlyString([]byte{98, 123})
	fmt.Println(onlyString, len(onlyString))
	fmt.Println(time.Since(now))

	now1 := time.Now()
	bytes1 := StringToBytes("a")
	fmt.Println(bytes1, len(bytes1), cap(bytes1))
	onlyString1 := BytesToString([]byte{98, 254})
	fmt.Println(onlyString1, len(onlyString1))
	fmt.Println(time.Since(now1))

	human := Human{}
	fmt.Println(fmt.Sprintf("%p", &human))
	fmt.Println(uintptr(unsafe.Pointer(&human)))                                                    // 内存地址
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&human.name)) - unsafe.Offsetof(human.name))) // 根据其中字段获取对象的起始内存地址
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&human.sex)) - unsafe.Offsetof(human.sex)))   // 根据其中字段获取对象的起始内存地址

	s := "string"
	lens := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(uint8(8))))
	fmt.Println(lens)
	lenptr := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(uint8(8))))
	// newlen := int(19)
	// i := 1 << 29
	// i2 := newlen | i
	fmt.Println(s)
	*lenptr = 29
	fmt.Println(len(s))
	fmt.Println(s)
	fmt.Println(len(s))

	fmt.Println(unsafe.Pointer(&s))                                              // string 的起始地址
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(uint8(8)))) // string->len 的起始地址
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s)))  // 溢出

	// fmt.Println(*(*string)(unsafe.Pointer(&s)))                                              // string 的起始地址

	h := Human{
		sex:  true,
		age:  uint8(1),
		name: "h",
	}

	h1 := Human{
		sex:  true,
		age:  uint8(1),
		name: "h1",
	}

	fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&h)) + unsafe.Offsetof(h.name))))
	fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&h1)) + unsafe.Offsetof(h1.name))))
	hname := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&h)) + unsafe.Offsetof(h.name)))
	*hname = *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&h1)) + unsafe.Offsetof(h1.name)))

	fmt.Println(*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&h)) + unsafe.Offsetof(h.name))))

	fmt.Println(unsafe.Pointer(&h))                                                    // 结构体的起始内存地址 0x14000134020
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&h)) + unsafe.Offsetof(h.sex)))  // 第一个字段的起始内存地址 0x14000134020
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&h)) + unsafe.Offsetof(h.age)))  // 第二个字段的起始内存地址 0x14000134021
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&h)) + unsafe.Offsetof(h.name))) // 第三个字段的起始内存地址 0x14000134030
	fmt.Println(unsafe.Pointer(uintptr(unsafe.Pointer(&h)) + unsafe.Sizeof(h)))        // 结构体的结束内存地址 0x14000134040

	// fmt.Println(reflect.DeepEqual(empty1.Tempty1Var,empty2.Tempty2Var))

	fmt.Println(
		unsafe.Pointer(&empty1.Tempty1Var),
		unsafe.Pointer(uintptr(unsafe.Pointer(&empty1.Tempty1Var))+uintptr(unsafe.Sizeof(empty1.Tempty1Var))),
	)
	fmt.Println(
		unsafe.Pointer(&empty2.Tempty2Var),
		unsafe.Pointer(uintptr(unsafe.Pointer(&empty2.Tempty2Var))+unsafe.Sizeof(empty2.Tempty2Var)),
	)

}

func Test1() {
	// head = {address, 10, 10}
	// body = [1,2,3,4,5,6,7,8,9,10]
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var address = (**[10]int)(unsafe.Pointer(&s))
	var len = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	var cap = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(address, *len, *cap)
	var body = **address
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", body[i])
	}
}
func Test2() {
	h := Human{
		true,
		30,
		1,
		"hello",
	}
	i := unsafe.Sizeof(h)
	j := unsafe.Alignof(h.age)
	k := unsafe.Offsetof(h.name)
	fmt.Println(unsafe.Sizeof("大"))
	fmt.Println(unsafe.Sizeof(int(1)))
	fmt.Println(i, j, k)
	fmt.Printf("%p\n", &h)
	fmt.Println(unsafe.Pointer(&h))

	// 空结构体/type学习
	tempty1Var := empty1.Tempty1Var
	// 占用空间宽度 8bit==1b为单位
	fmt.Println("tempty1Var占用空间宽度:")
	fmt.Println(unsafe.Sizeof(tempty1Var))
	fmt.Println("tempty1Var指针地址:")
	fmt.Println(fmt.Sprintf("%p", &tempty1Var))
	v1value := reflect.ValueOf(tempty1Var)
	// reflect.StringHeader{}
	fmt.Println(v1value.Pointer())

	tempty2Var := empty2.Tempty2Var
	fmt.Println("tempty2Var占用空间宽度:")
	fmt.Println(unsafe.Sizeof(tempty2Var))
	fmt.Println("tempty2Var指针地址:")
	pointer := *(*struct{})(unsafe.Pointer(&tempty2Var))
	fmt.Println(pointer)
	fmt.Println(fmt.Sprintf("%p", &tempty2Var))
	// v2value := reflect.ValueOf(tempty2Var)
	// fmt.Println(v2value.Pointer())

	empty1.Tempty1StringListVar = append(empty1.Tempty1StringListVar, "k")
	empty1.Tempty1StringListVar.Sort()
	fmt.Println(empty1.Tempty1StringListVar)

	strings := make([]string, 0, 10)
	strings = append(strings, "a")
	strings = append(strings, "2")
	strings = append(strings, "3")
	strings = append(strings, "6")
	strings = append(strings, "1")

	// 强转
	list := empty1.Tempty1StringList(strings)
	list.Sort()
	fmt.Println(list)

	// 断言
	var s interface{} = serviceImpl{}
	if s2, ok := s.(service); ok {
		s2.handler()
	}

}

func unsafeFastStringToReadOnlyBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
func unsafeFastBytesToReadOnlyString(b []byte) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	header := reflect.StringHeader{
		Data: sh.Data,
		Len:  sh.Len,
	}
	return *(*string)(unsafe.Pointer(&header))
}

// BytesToString converts byte slice to string.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes converts string to byte slice.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
