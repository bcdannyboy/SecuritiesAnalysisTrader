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

func InitStructWithRandomFloats(v interface{}) interface{} {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		// Handle pointer to struct
		val = val.Elem()
		if val.Kind() != reflect.Struct {
			panic("InitStructWithRandomFloats: expected a pointer to a struct")
		}
		return initStructWithRand(val, r).Addr().Interface()
	} else if val.Kind() == reflect.Struct {
		// Handle struct
		return initStructWithRand(val, r).Interface()
	} else {
		panic("InitStructWithRandomFloats: expected a struct or a non-nil pointer to a struct")
	}
}

func initStructWithRand(v reflect.Value, r *rand.Rand) reflect.Value {
	newStruct := reflect.New(v.Type()).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := newStruct.Field(i)
		switch field.Kind() {
		case reflect.Float64:
			// Generate a random float64 between -1.0 and 1.0
			randomFloat := r.Float64()*2 - 1
			field.SetFloat(randomFloat)
		case reflect.Struct:
			// Recursively initialize nested structs
			field.Set(initStructWithRand(v.Field(i), r))
		}
	}

	return newStruct
}
