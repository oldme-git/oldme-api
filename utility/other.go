package utility

import (
	"errors"
	"net"
	"net/http"
	"reflect"
	"strings"
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

// GetIP 获取http客户端真实ip
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
