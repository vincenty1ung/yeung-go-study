package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Products []*Product

type Product struct {
	Name   string
	Weight int
	Values []string
	Value  string
}

func (p Products) Len() int {
	return len(p)
}

func (p Products) Less(i, j int) bool {
	return p[i].Weight > p[j].Weight
}

func (p Products) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Products) Sort() {
	sort.Sort(p)
}

func NewProduct(Name string, Values []string, Value string) *Product {
	return &Product{
		Name:   Name,
		Values: Values,
		Value:  Value,
	}
}
func (m *Product) SetWeight() *Product {
	m.Weight = getProductWeight(m.Name)
	return m
}

func getProductWeight(name string) int {
	switch name {
	case "team_id":
		return 12
	case "month":
		return 11
	case "week":
		return 11
	case "quarter":
		return 11
	case "iteration_id":
		return 11
	case "story_type_id":
		return 8
	case "story_max_size":
		return 6
		// ...
	default:
		return 0
	}
}
func main() {
	now := time.Now()
	storyTypeIdProduct := NewProduct(
		"story_type_id", []string{"技术需求", "测试需求", "功能需求", "上级需求", "甲方需求", "紧急需求", "故障需求"},
		"",
	).SetWeight() // len：7
	storyMaxSizeProduct := NewProduct("story_max_size", []string{"M", "L", "XL", "XXL"}, "").SetWeight() // len：4
	monthProduct := NewProduct(
		"month", []string{
			"202401", "202402", "202403", "202404", "202405", "202406", "202407", "202408", "202409", "202410",
			"202411", "202412",
		}, "",
	).SetWeight() // len：12
	teamIdProduct := NewProduct(
		"team_id",
		[]string{"teamid1", "teamid2", "teamid3", "teamid4", "teamid5", "teamid6", "teamid7", "teamid8", "teamid9"},
		"",
	).SetWeight() // len：9
	products := Products{teamIdProduct, storyMaxSizeProduct, monthProduct, storyTypeIdProduct}
	products.Sort()
	cartesianProductResult := cartesianProduct(products)
	fmt.Println(len(cartesianProductResult)) // 3024
	fmt.Println(time.Since(now))
	for _, productList := range cartesianProductResult {
		ps := Products(productList)
		ps.Sort()
		result := ""
		for _, p := range ps {
			result += fmt.Sprintf("name:%s,value:%s | ", p.Name, p.Value)
		}
		fmt.Println(result)
	}

	storyTypeIdProduct1 := NewProduct(
		"story_type_id", []string{"趣丸需求"},
		"",
	).SetWeight() // len：1
	storyMaxSizeProduct1 := NewProduct("story_max_size", []string{"XXS", "XS"}, "").SetWeight() // len：0
	monthProduct1 := NewProduct(
		"month", []string{"202404", "202405", "202406", "202407", "202408", "202409"}, "",
	).SetWeight() // len：3
	teamIdProduct1 := NewProduct("team_id", []string{"di平台"}, "").SetWeight() // len：1
	products1 := Products{teamIdProduct1, storyMaxSizeProduct1, monthProduct1, storyTypeIdProduct1}
	products1.Sort()

	cartesianProductResult1 := cartesianProduct(products1)
	fmt.Println(len(cartesianProductResult1))
	for _, productList := range cartesianProductResult1 {
		ps := Products(productList)
		ps.Sort()
		result := ""
		for _, p := range ps {
			result += fmt.Sprintf("name:%s,value:%s | ", p.Name, p.Value)
		}
		fmt.Println(result)
	}

	// =========
	req := buildReq()
	fmt.Print("请求：")
	fmt.Println(req)
	var ps Products
	productList := make([]*Product, 0, len(req.SelectorConditions))
	productList = append(
		productList, NewProduct(
			"team_id",
			[]string{req.GroupID}, "",
		).SetWeight(),
	)
	for _, condition := range req.SelectorConditions {
		productList = append(
			productList, NewProduct(
				condition.SelectorType.GetColumnName(),
				condition.SelectorValues, "",
			).SetWeight(),
		)
	}
	ps = productList
	ps.Sort()
	cartesianProductResult2 := cartesianProduct(ps)
	fmt.Println(len(cartesianProductResult2))
	resp := MetricDataResp{
		MetricDataList:     make([]*MetricData, 0),
		SelectorAttentions: nil,
	}
	for _, productList := range cartesianProductResult2 {
		resp.MetricDataList = append(resp.MetricDataList, buildMetricData(productList))
	}

	fmt.Print("响应：")
	s := resp.String()
	fmt.Println(s)
}

func buildMetricData(ps Products) *MetricData {
	ps.Sort()
	dataMap := make(map[string]any, ps.Len())
	for _, p := range ps {
		dataMap[p.Name] = p.Value
	}
	dataMap["value"] = rand.Float32()
	return &MetricData{
		Data: dataMap,
	}
}

func cartesianProduct(products Products) [][]*Product {
	// fix
	for i, product := range products {
		if len(product.Values) == 0 {
			products = append(products[:i], products[i+1:]...)
		}
	}
	var result [][]*Product
	if len(products) == 0 {
		return result
	}

	CartesianHelper(products, []*Product{}, &result)
	return result
}

func CartesianHelper(products Products, current []*Product, result *[][]*Product) {
	if len(products) == 1 {
		p0 := products[0]
		for _, value := range p0.Values {
			product := NewProduct(p0.Name, p0.Values, value).SetWeight()
			newCurrent := make([]*Product, len(current)+1)
			copy(newCurrent, current)
			newCurrent[len(current)] = product
			*result = append(*result, newCurrent)
		}
		return
	}
	p0 := products[0]
	for _, value := range p0.Values {
		product := NewProduct(p0.Name, p0.Values, value).SetWeight()
		CartesianHelper(products[1:], append(current, product), result)
	}
}
