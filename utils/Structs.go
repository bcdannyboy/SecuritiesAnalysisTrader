package utils

import (
	"math"
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

func removeNaNs(v reflect.Value) reflect.Value {
	// Check if the value is valid and can be nil before calling IsNil
	if !v.IsValid() || (v.Kind() != reflect.Chan && v.Kind() != reflect.Func && v.Kind() != reflect.Interface && v.Kind() != reflect.Map && v.Kind() != reflect.Ptr && v.Kind() != reflect.Slice) {
		return v
	}

	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return reflect.Zero(v.Type()) // Return a zero value for the type
		}
		v = v.Elem()
	}

	// Check for *string and string types and return as is
	if v.Kind() == reflect.String || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.String) {
		return v
	}

	// Create a new value of the same type as v
	newV := reflect.New(v.Type()).Elem()

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			newV.Field(i).Set(removeNaNs(field))
		}
	case reflect.Slice:
		newV = reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		for i := 0; i < v.Len(); i++ {
			newV.Index(i).Set(removeNaNs(v.Index(i)))
		}
	case reflect.Map:
		newV = reflect.MakeMapWithSize(v.Type(), v.Len())
		for _, key := range v.MapKeys() {
			mapValue := v.MapIndex(key)
			newMapValue := removeNaNs(mapValue)

			// Check if the map's element type is a pointer
			if v.Type().Elem().Kind() == reflect.Ptr {
				// Create a new pointer of the correct type
				newPtr := reflect.New(v.Type().Elem().Elem())
				if !newMapValue.IsNil() {
					newPtr.Elem().Set(newMapValue.Elem())
				}
				newV.SetMapIndex(key, newPtr)
			} else {
				newV.SetMapIndex(key, newMapValue)
			}
		}
	case reflect.Float64:
		if math.IsNaN(v.Float()) {
			newV.SetFloat(0) // Replace NaN with zero
		} else {
			newV.Set(v)
		}
	default:
		newV.Set(v)
	}

	return newV
}

func RemoveNaNsFromStruct(s interface{}) interface{} {
	v := reflect.ValueOf(s)
	newV := removeNaNs(v)
	return newV.Interface()
}
