package main

import (
	"encoding/json"
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

func TestName(t *testing.T) {
	hums := []Human{
		{Name: "1", Age: 0}, {Name: "2", Age: 0}, {Name: "3", Age: 0}, {Name: "4", Age: 0},
		{Name: "5", Age: 0}, {Name: "6", Age: 0}, {Name: "7", Age: 0}, {Name: "8", Age: 0},
		{Name: "9", Age: 0}, {Name: "10", Age: 0}, {Name: "11", Age: 0}, {Name: "12", Age: 0},
	}
	// 超过size
	list := NewPageInfoByDataList(hums)
	list.ConvertMemoryPaginatedDataByPagNumAndSize(1, 50)
	marshal, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("===========1==========="))

	list = NewPageInfoByDataList(hums)
	list.ConvertMemoryPaginatedDataByPagNumAndSize(2, 50)
	marshal, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("===========2==========="))

	// 超过pageNum
	list = NewPageInfoByDataList(hums)
	list.ConvertMemoryPaginatedDataByPagNumAndSize(1, 3)
	marshal, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("============3=========="))

	list = NewPageInfoByDataList(hums)
	list.ConvertMemoryPaginatedDataByPagNumAndSize(2, 3)
	marshal, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("=============4========="))

	list = NewPageInfoByDataList(hums)
	list.ConvertMemoryPaginatedDataByPagNumAndSize(3, 3)
	marshal, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("==========5============"))

	list = NewPageInfoByDataList(hums)
	list.ConvertMemoryPaginatedDataByPagNumAndSize(4, 3)
	marshal, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("==========6============"))

	list = NewPageInfoByDataList(hums)
	list.ConvertMemoryPaginatedDataByPagNumAndSize(5, 3)
	marshal, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("===========7==========="))

	// 0对象
	list = NewPageInfoByDataList[Human](make([]Human, 0))
	list.ConvertMemoryPaginatedDataByPagNumAndSize(1, 50)
	marshal, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("============8=========="))

	// maps
	maps := []map[string]any{
		{
			"name":  "1",
			"age":   1,
			"phone": "176",
		},
		{
			"name":  "2",
			"age":   2,
			"phone": "176",
		},
		{
			"name":  "3",
			"age":   3,
			"phone": "176",
		},
		{
			"name":  "4",
			"age":   4,
			"phone": "176",
		},
		{
			"name":  "5",
			"age":   5,
			"phone": "176",
		},
		{
			"name":  "6",
			"age":   6,
			"phone": "176",
		},
		{
			"name":  "7",
			"age":   7,
			"phone": "176",
		},
	}
	list1 := NewPageInfoByDataList[map[string]any](maps)
	list1.ConvertMemoryPaginatedDataByPagNumAndSize(4, 2)
	marshal, err = json.Marshal(list1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("===========maps==========="))

}
