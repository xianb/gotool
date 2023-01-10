package utils

import (
	"reflect"
)

// Contain 判断obj是否在target中，target支持的类型arrary,slice,map
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func DeleteSlice(slice []interface{}, elem int) []interface{} {
	i := 0
	for _, v := range slice {
		if v != elem {
			slice[i] = v
			i++
		}
	}
	return slice[:i]
}
