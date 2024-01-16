package utils

import "reflect"

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
