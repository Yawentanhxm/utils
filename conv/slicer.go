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
