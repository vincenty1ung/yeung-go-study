package main

import (
	"fmt"
	"syscall"
)

func main() {
	returnAny := addReturnAny(12, 13)
	fmt.Println(returnAny)
	num := returnAny.(int64)
	fmt.Println(num)
	// 泛型数组
	var a GSlice[int64] = []int64{1, 12}
	for _, t := range a {
		fmt.Println(t)
	}
	var b GSlice[float32] = []float32{1.1, 12.23}
	for _, t := range b {
		fmt.Println(t)
	}
	var c GSlice[string] = []string{"nihao", "haha"}
	for _, t := range c {
		fmt.Println(t)
	}
	fmt.Println("foech")
	forFmtP(c)
	forFmtP[string](c)
	// forFmtP[int64](c)
	fmt.Println("m foech")
	c.forFmtP()

	// type GSliceAany[T any] []T // 代码等价于 type Slice[T interface{}] []T
	var gsa1 GSliceAany[int64] = []int64{1, 1}
	gsa1.forFmtP()
	var gsa2 GSliceAany[string] = []string{"1", "2"}
	gsa2.forFmtP()
	m := map[string]string{}
	m["k"] = "v"
	m1 := map[string]string{}
	m1["k1"] = "v1"
	var gsa3 GSliceAany[map[string]string] = []map[string]string{m, m1}
	gsa3.forFmtP()

	// 泛型map
	gmap := make(GMap[string, string])
	gmap["s"] = "key"
	fmt.Println(gmap)
	nm := make(map[string]string)
	nm["s"] = "key"
	fmt.Println(nm)
	for k, v := range gmap {
		fmt.Println(k)
		fmt.Println(v)
	}
	gm := GMan[int64, string, map[int64]string]{}
	gm.Age = 123
	gm.Name = "name"
	gm.Able = make(map[int64]string)
	gm.Able[123] = "123"
	gm.Able[125] = "125"
	fmt.Println(gm)
	// newMultimediaMusicWorksDao := newMultimediaMusicWorksDao()
	// var dao MultimediaMusicWorksDao[*MultimediaMusicWorks]
	// dao = newMultimediaMusicWorksDao()
	// works := MultimediaMusicWorks{}
	_, _ = newMultimediaMusicWorksDao().Insert(new(MultimediaMusicWorks))
	get := newMultimediaMusicWorksDao().Get(1)
	get.name = "aaa"
	// /works := get.(*MultimediaMusicWorks)
	fmt.Println(get)

	multimediaMusicWorksDao2 := newMultimediaMusicWorksDao2()
	_, _ = multimediaMusicWorksDao2.Insert(new(MultimediaMusicWorks))
	t := multimediaMusicWorksDao2.Get(1)
	t.name = "asd"
	fmt.Println(t)
	_, err := multimediaMusicWorksDao2.Update(new(MultimediaMusicWorks))
	if err != nil {
		return
	}
	multimediaMusicWorksDao2.Delete(1)

	// CSVProcessor实现了接口 DataProcessor[string] ，所以可赋值
	processor := newDataProcessor()
	processor.Process("name,age\nbob,12\njack,30")
	_ = processor.Save("name,age\nbob,13\njack,31")

	getgroups, err := syscall.Getgroups()
	fmt.Println(getgroups)
	fmt.Println(err)
	getgid := syscall.Getgid()
	fmt.Println(getgid)
}

// any func
func addReturnAny(a, b int64) any {
	return a + b
}

// GSlice 泛型数组
type GSlice[T int64 | uint32 | float32 | string] []T

// 泛型方法
func (receiver GSlice[T]) forFmtP() {
	for _, t := range receiver {
		fmt.Println(t)
	}
}

type GSliceAany[T any] []T // 代码等价于 type Slice[T interface{}] []T

func (receiver GSliceAany[T]) forFmtP() {
	for _, t := range receiver {
		fmt.Println(t)
	}
}

// 泛型函数
func forFmtP[T int64 | uint32 | float32 | string](slice GSlice[T]) {
	for _, t := range slice {
		fmt.Println(t)
	}
}

// GMap 泛型map
type GMap[K string | rune | float32, V string | int64] map[K]V

func (receiver GMap[K, V]) forFmtP() {
	for k, v := range receiver {
		fmt.Println(k)
		fmt.Println(v)
	}
}

// 泛型对象
type GMan[T int64 | uint32, V string | rune, S map[T]V] struct {
	Age  T
	Name V
	Able S
}
type (
	MultimediaMusicTag struct {
	}
	MultimediaMusicWorks struct {
		name string
	}
	MultimediaMusicHotMedia struct {
	}
	DB struct {
	}
	// 基础化泛型type
	baseType interface {
		*MultimediaMusicTag | *MultimediaMusicWorks | *MultimediaMusicHotMedia
	}
	// BaseDao 泛型接口
	BaseDao[T baseType] interface {
		Get(id int64) T
		Delete(id int64) T
		Insert(data T) (int64, error)
		Update(data T) (int64, error)
	}
	MultimediaMusicWorksDao[T interface{ *MultimediaMusicWorks }] interface {
		BaseDao[T]
		ListByTagIds(ids []int64) ([]*MultimediaMusicWorks, error)
		ListByTagIds2Map(ids []int64) (map[int64]*MultimediaMusicWorks, error)
	}

	MultimediaMusicWorksDao2 interface {
		BaseDao[*MultimediaMusicWorks]
		ListByTagIds(ids []int64) ([]*MultimediaMusicWorks, error)
		ListByTagIds2Map(ids []int64) (map[int64]*MultimediaMusicWorks, error)
	}
	defaultMultimediaMusicTagDao struct {
		mizaDB *DB
		// redisClient *uredis.Client
		table string
	}
	defaultMultimediaMusicWorksDao struct {
		mizaDB *DB
		// redisClient *uredis.Client
		table string
	}

	MultimediaMusicTagDao interface {
		BaseDao[*MultimediaMusicTag]
		GetByTagName(name string) (*MultimediaMusicTag, error)
		ListAllPage(name string, types int32, page int32, size int32) ([]*MultimediaMusicTag, int32, error)
		ListByAllStatus(types, status, recommendStatus int32) ([]*MultimediaMusicTag, error)
		ListByTagIds(ids []int64) ([]*MultimediaMusicTag, error)
		ListByTagIds2Map(ids []int64) (map[int64]*MultimediaMusicTag, error)
		ListByAllRecommend() ([]*MultimediaMusicTag, error)
		UpdateTagWorksNum(idList string, isAdd bool) error
	}
)

func newMultimediaMusicWorksDao() MultimediaMusicWorksDao[*MultimediaMusicWorks] {
	var t MultimediaMusicWorksDao[*MultimediaMusicWorks] = &defaultMultimediaMusicWorksDao{}
	return t
}
func newMultimediaMusicWorksDao2() MultimediaMusicWorksDao2 {
	// var t MultimediaMusicWorksDao[*MultimediaMusicWorks] =
	return &defaultMultimediaMusicWorksDao{}
}

// Get defaultMultimediaMusicWorksDao
func (d *defaultMultimediaMusicWorksDao) Get(id int64) *MultimediaMusicWorks {
	fmt.Println("Get")
	return &MultimediaMusicWorks{
		name: "ni",
	}
}

func (d *defaultMultimediaMusicWorksDao) Delete(id int64) *MultimediaMusicWorks {
	fmt.Println("Delete")
	return nil
}

func (d *defaultMultimediaMusicWorksDao) Insert(data *MultimediaMusicWorks) (int64, error) {
	fmt.Println("Insert")
	return 0, nil
}

func (d *defaultMultimediaMusicWorksDao) Update(data *MultimediaMusicWorks) (int64, error) {
	fmt.Println("Update")
	return 0, nil
}

func (d *defaultMultimediaMusicWorksDao) ListByTagIds(ids []int64) ([]*MultimediaMusicWorks, error) {
	fmt.Println("ListByTagIds")
	return nil, nil
}

func (d *defaultMultimediaMusicWorksDao) ListByTagIds2Map(ids []int64) (map[int64]*MultimediaMusicWorks, error) {
	fmt.Println("ListByTagIds2Map")
	return nil, nil
}

// ++++++++++++++++++
func (d defaultMultimediaMusicTagDao) Get(id int64) *MultimediaMusicTag {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) Delete(id int64) bool {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) Insert(data *MultimediaMusicTag) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) Update(data *MultimediaMusicTag) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) GetByTagName(name string) (*MultimediaMusicTag, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) ListAllPage(
	name string, types int32, page int32, size int32,
) ([]*MultimediaMusicTag, int32, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) ListByAllStatus(types, status, recommendStatus int32) (
	[]*MultimediaMusicTag, error,
) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) ListByTagIds(ids []int64) ([]*MultimediaMusicTag, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) ListByTagIds2Map(ids []int64) (map[int64]*MultimediaMusicTag, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) ListByAllRecommend() ([]*MultimediaMusicTag, error) {
	// TODO implement me
	panic("implement me")
}

func (d defaultMultimediaMusicTagDao) UpdateTagWorksNum(idList string, isAdd bool) error {
	// TODO implement me
	panic("implement me")
}

// 基本接口/一般接口
type ReadWriter interface {
	~string | ~[]rune

	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 类型 StringReadWriter 实现了接口 Readwriter
type StringReadWriter string

func (s StringReadWriter) Read(p []byte) (n int, err error) {
	return 0, err
}

func (s StringReadWriter) Write(p []byte) (n int, err error) {
	return 0, err
}

//  类型BytesReadWriter 没有实现接口 Readwriter
type BytesReadWriter []byte

func (s BytesReadWriter) Read(p []byte) (n int, err error) {
	return 0, err
}

func (s BytesReadWriter) Write(p []byte) (n int, err error) {
	return 0, err
}

// 泛型接口
type DataProcessor[T any] interface {
	Process(oriData T) (newData T)
	Save(data T) error
}

type CSVProcessor struct {
}

func newDataProcessor() DataProcessor[string] {
	var processor DataProcessor[string] = CSVProcessor{}
	return processor
}

// 注意，方法中 oriData 等的类型是 string
func (c CSVProcessor) Process(oriData string) (newData string) {
	fmt.Println(oriData)
	return ""
}

func (c CSVProcessor) Save(oriData string) error {
	fmt.Println(oriData)
	return nil
}

type DataProcessor2[T any] interface {
	int | ~struct{ Data any }

	Process(data T) (newData T)
	Save(data T) error
}

/*func newXMLProcessor() DataProcessor2[string] {
	var processor DataProcessor2[string] = XMLProcessor{}
	return processor
}*/

// 错误。DataProcessor2[string]是一般接口不能用于创建变量
// var processor DataProcessor2[string]
// var processor1 DataProcessor2[string] = JsonProcessor{Data: "s"}

// var processorxml DataProcessor2[string] = XMLProcessor{}
var processorcsv DataProcessor[string] = CSVProcessor{}

// XMLProcessor 虽然实现了接口 DataProcessor2[string] 的两个方法，但是因为它的底层类型是 []byte，所以依旧是未实现 DataProcessor2[string]
type XMLProcessor []byte

func (c XMLProcessor) Process(oriData string) (newData string) {
	return ""
}

func (c XMLProcessor) Save(oriData string) error {
	return nil
}

// JsonProcessor 实现了接口 DataProcessor2[string] 的两个方法，同时底层类型是 struct{ Data interface{} }。所以实现了接口 DataProcessor2[string]
type JsonProcessor struct {
	Data any
}

/*func newJsonProcessor() DataProcessor2[string] {
	var processor DataProcessor2[string] = JsonProcessor{}
	return processor
}*/
func (c JsonProcessor) Process(oriData string) (newData string) {
	return ""
}

func (c JsonProcessor) Save(oriData string) error {
	return nil
}

// 正确，实例化之后的 DataProcessor2[string] 可用于泛型的类型约束
type ProcessorList[T DataProcessor2[string]] []T

type MyInf[T int | Strings] interface {
	~int | ~struct{}

	bac() T
}
type bacc struct {
}

func (b bacc) bac() int {
	return 1
}

type myint int

// var baccc MyInf[int] = bacc{}

type abc struct {
}

func (a abc) bac() string {
	return "string()"
}

// var abcv MyInf[string] = abc{}

// var baccc1 MyInf[myint] = bacc{}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type Float interface {
	~float32 | ~float64
}
type Strings interface {
	~string | ~rune
}
