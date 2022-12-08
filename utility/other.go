package utility

import (
	"errors"
	"reflect"
)

// InArray 判断某个值是否存在与切片或者数组中，此函数的命名十分优秀ヾ(≧▽≦*)o
func InArray(v interface{}, i interface{}) (b bool, err error) {
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return false, errors.New("给定类型必须是切片或者数组")
	}
	// 将i转换为map
	var (
		value = reflect.ValueOf(i)
		m     = make(map[interface{}]bool, value.Len())
	)
	for j := 0; j < value.Len(); j++ {
		m[value.Index(j).Interface()] = true
	}

	return m[v], nil
}
