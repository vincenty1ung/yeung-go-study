package yjson

import jsoniter "github.com/json-iterator/go"

var ji = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return ji.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	return ji.MarshalToString(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return ji.Unmarshal(data, v)
}

func UnmarshalFromString(str string, v interface{}) error {
	return ji.UnmarshalFromString(str, v)
}

/**
*	Get 获取深层嵌套JSON结构的值
*	data: 数据源
*	path: If string, it will lookup json map. If int, it will lookup json array. If '*', it will map to each element of array or each key of map.
*	It will be faster than parsing into map[string]interface{} and much easier to read data out.
 */
func Get(data []byte, path ...interface{}) jsoniter.Any {
	return ji.Get(data, path...)
}
