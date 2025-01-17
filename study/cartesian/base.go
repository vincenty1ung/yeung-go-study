package main

import (
	"encoding/json"
)

type MetricDataReq struct {
	// 公共等值条件
	MetricKey     string        `json:"metric_key" validate:"required" zh:"指标唯一键"`
	GroupDataType GroupDataType `json:"group_data_type,default=FEATURE_TEAM,options=[FEATURE_TEAM,DEPARTMENT]" zh:"团队数据类型"`
	GroupID       string        `json:"group_id" validate:"required" zh:"团队ID[FEATURE_TEAM（team_id）,DEPARTMENT（dept_code）]"`
	// 分组条件
	SelectorConditions []*SelectorCondition `json:"selector_conditions" zh:"选择条件列表"`
}

func (m MetricDataReq) String() string {
	marshal, _ := json.Marshal(m)
	return string(marshal)
}

type SelectorCondition struct {
	SelectorType   SelectorType `json:"selector_type" zh:"选择器类型"`
	SelectorValues []string     `json:"selector_values" zh:"选择器值集合"`
}

func buildReq() MetricDataReq {
	return MetricDataReq{
		MetricKey:     "OnTimeStoryDeliveryRate",
		GroupDataType: GroupData_FEATURE_TEAM_TYPE,
		GroupID:       "team_id:qFIUkLkPYxNCq5t_2NQS_",
		SelectorConditions: []*SelectorCondition{
			{
				SelectorType: Selector_ITERATION_SELECT,
				SelectorValues: []string{
					"1132571206001003938",
					"1132571206001003937",
					"1132571206001003936",
					"1132571206001003901",
					"1132571206001003900",
					"1132571206001003863",
				},
			},
			{
				SelectorType: Selector_STORY_TYPE_SELECT,
				SelectorValues: []string{
					"712060",
					"122360",
				},
			},
			{
				SelectorType: Selector_STORY_MAX_SIZE_SELECT,
				SelectorValues: []string{
					"XXS",
					"XS",
					"M",
				},
			},
		},
	}
}

type SelectorType int32
type GroupDataType int32

const (
	Selector_CT_UNKNOWN              SelectorType = 0  // NULL
	Selector_PERIOD_SELECT           SelectorType = 1  // 「周期」选择器
	Selector_ITERATION_SELECT        SelectorType = 2  // 「迭代」选择器
	Selector_FEATURE_TEAM_SELECT     SelectorType = 3  // 「特性团队」选择器
	Selector_DEPARTMENT_SELECT       SelectorType = 4  // 「行政组织」选择器
	Selector_GROUP_SELECT            SelectorType = 5  // 「组织」选择器
	Selector_STORY_TYPE_SELECT       SelectorType = 6  // 「需求类别」选择器
	Selector_BUG_ORIGIN_PHASE_SELECT SelectorType = 7  // 「缺陷发现阶段」选择器
	Selector_MONTH_WEEK_DATEPICKER   SelectorType = 8  // 「月、周」日期选择框
	Selector_STORY_MAX_SIZE_SELECT   SelectorType = 9  // 「最大需求尺码」选择器
	Selector_DELIVERY_ENV_SELECT     SelectorType = 12 // 交付环境选择器
	Selector_BUG_TYPE_ID_SELECT      SelectorType = 13 // 缺陷类型选择器
	Selector_PERIOD_MONTH_SELECT     SelectorType = 14 // 「周期-月」选择器
	Selector_PERIOD_WEEK_SELECT      SelectorType = 15 // 「周期-周」选择器
	Selector_PERIOD_QUARTER_SELECT   SelectorType = 16 // 「周期-季度」选择器
)
const (
	GroupData_UNKNOWN_TYPE      GroupDataType = 0 // NULL
	GroupData_FEATURE_TEAM_TYPE GroupDataType = 1 // 「周期」选择器
	GroupData_DEPARTMENT_TYPE   GroupDataType = 2 // 「迭代」选择器
)

// 列
func (receiver SelectorType) GetColumnName() string {
	switch receiver {
	case Selector_PERIOD_MONTH_SELECT:
		return "month"
	case Selector_PERIOD_WEEK_SELECT:
		return "week"
	case Selector_PERIOD_QUARTER_SELECT:
		return "quarter"
	case Selector_ITERATION_SELECT:
		return "iteration_id"
	case Selector_STORY_TYPE_SELECT:
		return "story_type_id"
	case Selector_STORY_MAX_SIZE_SELECT:
		return "story_max_size"
	default:
		return ""
	}
}

type MetricDataResp struct {
	MetricDataList     []*MetricData        `json:"metric_data_list" zh:"指标数据集合"`
	SelectorAttentions []*SelectorAttention `json:"selector_attentions" zh:"选择关注（在意）集合，其实不返回前端通过SelectorType也可以自行知道"`
}

func (m MetricDataResp) String() string {
	marshal, _ := json.Marshal(m)
	return string(marshal)
}

type MetricData struct {
	Data map[string]any `json:"data" zh:"数据,其中指标值的key是value"`
	Fx   string         `json:"fx" zh:"数据,其中指标值的key是value"`
}

type SelectorAttention struct {
	SelectorType         int8   `json:"selector_type" zh:"选择器类型"`
	SelectorAttentionKey string `json:"selector_attention_key" zh:"选择器关注(在意)key"`
}
