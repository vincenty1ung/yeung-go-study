package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/*
目标心率范围有一个相对复杂的计算公式：

［（最大心率-静止心率）×60% + 静止心率］~［（最大心率-静止心率）×80% + 静止心率］

假设小明的静止心率是70次，那他的目标心率范围就是：［（190-70）×60% +70］~［（190-70）×80% +70］，也就是142~166

静止心率: 59

max(190) = 220 - ”年龄“

(190-59)×0.6+59 =
(190-59)×0.8%+59 =
*/
var (
	maxHeartRateNum = 220
)

var (
	flagset         map[string]bool
	testConfig      bool
	help            bool
	isAlfred        bool
	staticHeartRate string
	age             string
)

func init() {
	flag.StringVar(&staticHeartRate, "static", "", "set static heart rate")
	flag.StringVar(&age, "age", "", "set age")
	flag.BoolVar(&help, "h", false, "获取使用帮助")
	flag.BoolVar(&isAlfred, "alfred", false, "支持alfred")
	flag.BoolVar(&testConfig, "t", false, "test configuration and exit")
	flag.Parse()

	flagset = map[string]bool{}
	flag.Visit(
		func(f *flag.Flag) {
			flagset[f.Name] = true
		},
	)
}

func main() {
	if isHelp() {
		return
	}
	if !alfredCheck() {
		return
	}
	staticHeartRateNum, ageNum, done := checkParameterLegality()
	if done {
		return
	}
	minRate, maxRate := recommendedHeartRate(maxHeartRate(ageNum), staticHeartRateNum)

	sprintf := fmt.Sprintf(
		"年龄:%d,静息心率:%d,推荐的最小心率:%d,最大心率:%d", ageNum, staticHeartRateNum, minRate, maxRate,
	)

	if flagset["alfred"] {
		alfredItems := AlfredItems{}
		alfredItems.Items = append(
			alfredItems.Items, AlfredItem{
				Uid:          "1",
				Arg:          sprintf,
				Valid:        "yes",
				Autocomplete: sprintf,
				Title:        sprintf,
				Subtitle:     "结果复制到剪切板",
				Icon:         "icon.png",
			},
		)
		marshal, _ := json.Marshal(alfredItems)
		fmt.Println(string(marshal))
	} else {
		fmt.Println(sprintf)
	}
}

func alfredCheck() bool {
	/*if flagset["alfred"] && flagset["age"] &&
	(flagset["static"] && len(staticHeartRate) == 0) {*/
	if flagset["alfred"] && len(staticHeartRate) == 0 {
		alfredItems := AlfredItems{}
		alfredItems.Items = append(
			alfredItems.Items, AlfredItem{
				Uid:          "1",
				Arg:          "请输入第二个参数(静息心率)",
				Valid:        "yes",
				Autocomplete: "请输入第二个参数(静息心率)",
				Title:        "请输入第二个参数(静息心率)",
				Subtitle:     "提示",
				Icon:         "icon.png",
			},
		)
		marshal, _ := json.Marshal(alfredItems)
		fmt.Println(string(marshal))
		return false
	}
	return true
}

// 检查参数合法性
func checkParameterLegality() (int, int, bool) {
	if !flagset["static"] && !flagset["age"] {
		fmt.Println("没有输入有效的参数")
		return 0, 0, true
	}
	staticHeartRateNum, err := strconv.Atoi(staticHeartRate)
	if err != nil {
		fmt.Println(err)
		return 0, 0, true
	}
	ageNum, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println(err)
		return 0, 0, true
	}
	return staticHeartRateNum, ageNum, false
}

func isHelp() bool {
	if flagset["h"] {
		var sb strings.Builder
		sb.WriteString("使用说明:\n")
		sb.WriteString("* 范例: ./running -age 30 -static 59 (执行)\n")
		sb.WriteString("* age :设置你当前的年龄\n")
		sb.WriteString("* static :设置你当月的平均静止心率\n")
		fmt.Println(sb.String())
		return true
	}
	return false
}

// 推荐心率
func recommendedHeartRate(maxHeartHate, staticHeartRate int) (min, max int) {
	return int(float32(maxHeartHate-staticHeartRate)*0.6 + float32(staticHeartRate)),
		int(float32(maxHeartHate-staticHeartRate)*0.8 + float32(staticHeartRate))
}

// 最大心率
func maxHeartRate(age int) int {
	if age > maxHeartRateNum {
		return 0
	}
	return maxHeartRateNum - age
}

type AlfredItemIcon struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

type AlfredItem struct {
	Uid          string `json:"uid"`
	Arg          string `json:"arg"`
	Valid        string `json:"valid"`
	Autocomplete string `json:"autocomplete"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Icon         string `json:"icon"`
}
type AlfredItems struct {
	Items []AlfredItem `json:"items"`
}
