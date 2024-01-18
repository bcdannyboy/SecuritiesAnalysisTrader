package utils

import (
	"math/rand"
	"reflect"
	"time"
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

func InitStructWithRandomFloats(v interface{}) {
	rand.Seed(time.Now().UnixNano())

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		panic("InitStructWithRandomFloats: expected a non-nil pointer to a struct")
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		panic("InitStructWithRandomFloats: expected a pointer to a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.Float64:
			// Generate a random float64 between -1.0 and 1.0
			randomFloat := rand.Float64()*2 - 1
			field.SetFloat(randomFloat)
		case reflect.Struct:
			// Recursively initialize nested structs
			if field.CanAddr() {
				InitStructWithRandomFloats(field.Addr().Interface())
			}
		}
	}
}
