package utils

import (
	"reflect"
)

func InterfaceToFloat64Ptr(value interface{}) *float64 {
	switch v := value.(type) {
	case float64:
		return &v
	case int: // Handle int as it can be safely converted to float64
		f := float64(v)
		return &f
	case int64: // Handle int64, if needed
		f := float64(v)
		return &f
	// Add more cases here to handle more types
	default:
		// Check if it's a convertible type using reflection
		val := reflect.ValueOf(value)
		if val.Kind() == reflect.Float32 || val.Kind() == reflect.Int || val.Kind() == reflect.Int64 {
			floatVal := val.Convert(reflect.TypeOf(float64(0))).Float()
			return &floatVal
		}
		return nil // Not a type that can be converted to float64
	}
}

func GetFloat64PtrIfNotEmpty(inp interface{}, fieldName string) *float64 {
	if inp == nil {
		return nil
	}

	rv := reflect.ValueOf(inp)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	// Check if the struct is valid and the field exists
	if rv.Kind() != reflect.Struct || !rv.FieldByName(fieldName).IsValid() {
		return nil
	}

	fieldValue := rv.FieldByName(fieldName)

	// Check for *float64 and float64 types
	if fieldValue.Kind() == reflect.Ptr && fieldValue.Elem().Kind() == reflect.Float64 {
		if !fieldValue.IsNil() {
			return fieldValue.Interface().(*float64)
		}
	} else if fieldValue.Kind() == reflect.Float64 {
		if fieldValue.CanAddr() {
			value := fieldValue.Float()
			return &value
		}
	}

	return nil
}
