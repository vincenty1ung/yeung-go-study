package main

import (
	"fmt"
	"testing"
)

var humans140 = []Human{
	{name: "xili", age: 12}, {name: "vincent", age: 12}, {name: "chenfeng", age: 12}, {name: "huangyongyi", age: 12},
	{name: "xili23413", age: 12}, {name: "vincent33123", age: 12}, {name: "chenfeng233333", age: 12},
	{name: "huangyongyi22222222222", age: 12},
	{name: "xili23134", age: 12}, {name: "yngbo331233", age: 12}, {name: "chefeng2333333", age: 12},
	{name: "huangyongi222222222222", age: 12},
	{name: "xili23434", age: 12}, {name: "vincent31233", age: 12}, {name: "chenfen2333333", age: 12},
	{name: "huangyongyi22222222222", age: 12},
	{name: "xili23134", age: 12}, {name: "vincent1233", age: 12}, {name: "chenng2333333", age: 12},
	{name: "huangyongyi2222222222", age: 12},
	{name: "xili2134", age: 12}, {name: "yan331233", age: 12}, {name: "chenfe2333333", age: 12},
	{name: "huangngyi22222222222", age: 12},
	{name: "xili4134", age: 12}, {name: "yang331233", age: 12}, {name: "nfeng2333333", age: 12},
	{name: "huangyongyi222222222222", age: 12},
	{name: "xili234", age: 12}, {name: "vincent333", age: 12}, {name: "chng2333333", age: 12},
	{name: "huani222222222222", age: 12},
	{name: "xil134", age: 12}, {name: "vincent333", age: 12}, {name: "chenllfeng2333333", age: 12},
	{name: "huangyongyi222222222222", age: 12},
	{name: "xili2134", age: 12}, {name: "yabo331233", age: 12}, {name: "cheg2333333", age: 12},
	{name: "huangyongyi222222222222", age: 12},
	{name: "xil4134", age: 12}, {name: "vincent233", age: 12}, {name: "cheg2333333", age: 12},
	{name: "huangyyi222222222222", age: 12},
	{name: "xili2334", age: 12}, {name: "yao331233", age: 12}, {name: "chenfe2333333", age: 12},
	{name: "huangyongyi222222222222", age: 12},
	{name: "xili23414", age: 12}, {name: "vincent31233", age: 12}, {name: "chenfeg2333333", age: 12},
	{name: "huangyonyi22222222222", age: 12},
	{name: "xii23134", age: 12}, {name: "vincent33133", age: 12}, {name: "chenfeg2333333", age: 12},
	{name: "huangongyi22222222222", age: 12},
	{name: "xili23134", age: 12}, {name: "vincent1233", age: 12}, {name: "chenfen33333", age: 12},
	{name: "huangyongyi222222222222", age: 12},
	{name: "ili23434", age: 12}, {name: "yagbo31233", age: 12}, {name: "chenng23333", age: 12},
	{name: "huangongyi222222222", age: 12},
	{name: "xili244", age: 12}, {name: "yan331233", age: 12}, {name: "ch33333", age: 12},
	{name: "huang2222222222", age: 12},
	{name: "xi34134", age: 12}, {name: "yabo331233", age: 12}, {name: "enfeng2333333", age: 12},
	{name: "huaongyi22222222", age: 12},
	{name: "xi34134", age: 12}, {name: "yab331233", age: 12}, {name: "enfen2333333", age: 12},
	{name: "huaongyi222222222", age: 12},
	{name: "xi3134", age: 12}, {name: "yabo331233", age: 12}, {name: "enfng2333333", age: 12},
	{name: "huangyi222222222", age: 12},
	{name: "xi3414", age: 12}, {name: "ybo33123", age: 12}, {name: "enfeng233333", age: 12},
	{name: "huaongyi22222222", age: 12},
	{name: "xi3134", age: 12}, {name: "yab331233", age: 12}, {name: "eeng2333333", age: 12},
	{name: "huaongyi222222222", age: 12},
	{name: "xi3414", age: 12}, {name: "yabo33123", age: 12}, {name: "enfe2333333", age: 12},
	{name: "huaongy222222", age: 12},
}
var humans4 = []Human{
	{name: "xili", age: 12}, {name: "vincent", age: 12}, {name: "chenfeng", age: 12}, {name: "huangyongyi", age: 12},
}

func BenchmarkBinarySearchCount4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binarySearch(humans4, "vincent")
	}
}
func BenchmarkMapSearchCount4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapSearch(humans4, "vincent")
	}
}
func BenchmarkBinarySearchCount140(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binarySearch(humans140, "vincent")
	}
}
func BenchmarkMapSearchCount140(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapSearch(humans140, "vincent")
	}
}
func TestMapSearch(t *testing.T) {
	ok, human := mapSearch(humans140, "vincent")
	fmt.Println(ok)
	fmt.Println(human)
}
func TestBinarySearch(t *testing.T) {
	ok, human := binarySearch(humans140, "vincent")
	fmt.Println(ok)
	fmt.Println(human)
}
