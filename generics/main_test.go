package main

import (
	"fmt"
	"testing"
)

var humans140 = []Human{
	{Name: "xili", Age: 12}, {Name: "vincent", Age: 12}, {Name: "chenfeng", Age: 12}, {Name: "huangyongyi", Age: 12},
	{Name: "xili23413", Age: 12}, {Name: "vincent33123", Age: 12}, {Name: "chenfeng233333", Age: 12},
	{Name: "huangyongyi22222222222", Age: 12},
	{Name: "xili23134", Age: 12}, {Name: "yngbo331233", Age: 12}, {Name: "chefeng2333333", Age: 12},
	{Name: "huangyongi222222222222", Age: 12},
	{Name: "xili23434", Age: 12}, {Name: "vincent31233", Age: 12}, {Name: "chenfen2333333", Age: 12},
	{Name: "huangyongyi22222222222", Age: 12},
	{Name: "xili23134", Age: 12}, {Name: "vincent1233", Age: 12}, {Name: "chenng2333333", Age: 12},
	{Name: "huangyongyi2222222222", Age: 12},
	{Name: "xili2134", Age: 12}, {Name: "yan331233", Age: 12}, {Name: "chenfe2333333", Age: 12},
	{Name: "huangngyi22222222222", Age: 12},
	{Name: "xili4134", Age: 12}, {Name: "yang331233", Age: 12}, {Name: "nfeng2333333", Age: 12},
	{Name: "huangyongyi222222222222", Age: 12},
	{Name: "xili234", Age: 12}, {Name: "vincent333", Age: 12}, {Name: "chng2333333", Age: 12},
	{Name: "huani222222222222", Age: 12},
	{Name: "xil134", Age: 12}, {Name: "vincent333", Age: 12}, {Name: "chenllfeng2333333", Age: 12},
	{Name: "huangyongyi222222222222", Age: 12},
	{Name: "xili2134", Age: 12}, {Name: "yabo331233", Age: 12}, {Name: "cheg2333333", Age: 12},
	{Name: "huangyongyi222222222222", Age: 12},
	{Name: "xil4134", Age: 12}, {Name: "vincent233", Age: 12}, {Name: "cheg2333333", Age: 12},
	{Name: "huangyyi222222222222", Age: 12},
	{Name: "xili2334", Age: 12}, {Name: "yao331233", Age: 12}, {Name: "chenfe2333333", Age: 12},
	{Name: "huangyongyi222222222222", Age: 12},
	{Name: "xili23414", Age: 12}, {Name: "vincent31233", Age: 12}, {Name: "chenfeg2333333", Age: 12},
	{Name: "huangyonyi22222222222", Age: 12},
	{Name: "xii23134", Age: 12}, {Name: "vincent33133", Age: 12}, {Name: "chenfeg2333333", Age: 12},
	{Name: "huangongyi22222222222", Age: 12},
	{Name: "xili23134", Age: 12}, {Name: "vincent1233", Age: 12}, {Name: "chenfen33333", Age: 12},
	{Name: "huangyongyi222222222222", Age: 12},
	{Name: "ili23434", Age: 12}, {Name: "yagbo31233", Age: 12}, {Name: "chenng23333", Age: 12},
	{Name: "huangongyi222222222", Age: 12},
	{Name: "xili244", Age: 12}, {Name: "yan331233", Age: 12}, {Name: "ch33333", Age: 12},
	{Name: "huang2222222222", Age: 12},
	{Name: "xi34134", Age: 12}, {Name: "yabo331233", Age: 12}, {Name: "enfeng2333333", Age: 12},
	{Name: "huaongyi22222222", Age: 12},
	{Name: "xi34134", Age: 12}, {Name: "yab331233", Age: 12}, {Name: "enfen2333333", Age: 12},
	{Name: "huaongyi222222222", Age: 12},
	{Name: "xi3134", Age: 12}, {Name: "yabo331233", Age: 12}, {Name: "enfng2333333", Age: 12},
	{Name: "huangyi222222222", Age: 12},
	{Name: "xi3414", Age: 12}, {Name: "ybo33123", Age: 12}, {Name: "enfeng233333", Age: 12},
	{Name: "huaongyi22222222", Age: 12},
	{Name: "xi3134", Age: 12}, {Name: "yab331233", Age: 12}, {Name: "eeng2333333", Age: 12},
	{Name: "huaongyi222222222", Age: 12},
	{Name: "xi3414", Age: 12}, {Name: "yabo33123", Age: 12}, {Name: "enfe2333333", Age: 12},
	{Name: "huaongy222222", Age: 12},
}
var humans4 = []Human{
	{Name: "xili", Age: 12}, {Name: "vincent", Age: 12}, {Name: "chenfeng", Age: 12}, {Name: "huangyongyi", Age: 12},
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
