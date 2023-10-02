package pages

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewPageInfoByDataListGenerics(t *testing.T) {
	p1 := NewPageInfoByDataList(
		[]string{
			"string1", "string2",
		},
	)
	p2 := NewPageInfoByDataList(
		[]name{
			{
				"name1",
			},
			{
				"name2",
			},
		},
	)
	dataList1 := p1.source
	dataList2 := p2.source
	for _, s := range dataList1 {
		fmt.Println(s)
	}
	for _, s := range dataList2 {
		fmt.Println(s)
	}
}
func TestNewPageInfoByDataList(t *testing.T) {
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
	list1.ConvertMemoryPaginatedDataByPagNumAndSize(3, 2)
	marshal, err = json.Marshal(list1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(string("===========maps==========="))

}

type name struct {
	name string
}

type Human struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}
