package conv

import (
	"fmt"
	"reflect"
)

func SliceToMap[T any](key string, data []T) map[string]T {
	resp := make(map[string]T)
	for _, v := range data {
		k := ""
		index := reflect.ValueOf(v)
		if index.Kind() == reflect.Pointer {
			k = fmt.Sprint(index.Elem().FieldByName(key))
		} else {
			k = fmt.Sprint(index.FieldByName(key))
		}
		resp[k] = v
	}
	return resp
}

// Push: 包含剔除能力的push
func Push[T any](rds []T, fun func(t T) bool, rd ...T) []T {
	for _, v := range rd {
		if fun(v) {
			continue
		}
		rds = append(rds, v)
	}
	return rds
}
