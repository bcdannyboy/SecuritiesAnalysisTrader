package utils

import (
	"reflect"
)

func IsStructEmpty(v interface{}) bool {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			if !reflect.DeepEqual(val.Field(i).Interface(), reflect.Zero(val.Field(i).Type()).Interface()) {
				return false
			}
		}
		return true

	}
	// handle non-struct types differently
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
