package Optimization

import (
	"fmt"
	"math"
	"reflect"
)

func CalculateWeightedAverage(Weights interface{}, DataToWeight interface{}, Path string) (float64, error) {
	totalWeightedValue, totalWeight, err := calculateWeightedAverageRecursive(Weights, DataToWeight, Path)
	if err != nil {
		fmt.Printf("got error from CalculateWeightedAverage: %s\n", err)
		return math.NaN(), err
	}

	if totalWeight == 0 {
		fmt.Printf("got totalWeight == 0, returning 0\n")
		return 0, nil
	}

	finalValue := totalWeightedValue / totalWeight
	return finalValue, nil
}

func calculateWeightedAverageRecursive(Weights interface{}, DataToWeight interface{}, Path string) (float64, float64, error) {
	var totalWeightedValue float64
	var totalWeight float64

	vWeights := reflect.Indirect(reflect.ValueOf(Weights))
	vDataToWeight := reflect.Indirect(reflect.ValueOf(DataToWeight))

	// Calculate the maximum absolute value for normalization
	maxValue := findMaxAbsoluteValue(vDataToWeight)

	var handleValue func(reflect.Value, reflect.Value, float64, string) error
	handleValue = func(weightValue reflect.Value, dataValue reflect.Value, weight float64, path string) error {
		weightValue = reflect.Indirect(weightValue)
		dataValue = reflect.Indirect(dataValue)

		if !dataValue.IsValid() {
			return nil // Skip invalid data values
		}

		switch weightValue.Kind() {
		case reflect.Float64:
			data := extractFloatValue(dataValue)
			if !math.IsNaN(data) { // Check for NaN data
				normalizedData := normalizeData(data, maxValue, weightValue.Float())
				totalWeightedValue += normalizedData
				totalWeight += math.Abs(weightValue.Float())
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			data := extractFloatValue(dataValue)
			if !math.IsNaN(data) { // Check for NaN data
				normalizedData := normalizeData(data, maxValue, float64(weightValue.Int()))
				totalWeightedValue += normalizedData
				totalWeight += math.Abs(float64(weightValue.Int()))
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			data := extractFloatValue(dataValue)
			if !math.IsNaN(data) { // Check for NaN data
				normalizedData := normalizeData(data, maxValue, float64(weightValue.Uint()))
				totalWeightedValue += normalizedData
				totalWeight += math.Abs(float64(weightValue.Uint()))
			}

		case reflect.Struct:
			for i := 0; i < weightValue.NumField(); i++ {
				field := weightValue.Type().Field(i)
				newPath := path + "." + field.Name
				subWeightValue := weightValue.Field(i)
				subDataValue := extractValueByName(dataValue, field.Name)
				err := handleValue(subWeightValue, subDataValue, weight, newPath)
				if err != nil {
					return err
				}
			}

		case reflect.Slice, reflect.Array:
			for i := 0; i < weightValue.Len(); i++ {
				newPath := fmt.Sprintf("%s[%d]", path, i)
				subWeightValue := weightValue.Index(i)
				subDataValue := dataValue.Index(i)
				err := handleValue(subWeightValue, subDataValue, weight, newPath)
				if err != nil {
					return err
				}
			}

		case reflect.Map:
			for _, key := range weightValue.MapKeys() {
				newPath := path + "[" + fmt.Sprint(key.Interface()) + "]"
				subWeightValue := weightValue.MapIndex(key)
				subDataValue := dataValue.MapIndex(key)
				err := handleValue(subWeightValue, subDataValue, weight, newPath)
				if err != nil {
					return err
				}
			}

		default:
			fmt.Printf("Unhandled type: %s at path %s\n", weightValue.Kind(), path)
		}
		return nil
	}

	if err := handleValue(vWeights, vDataToWeight, 1.0, Path); err != nil {
		return 0, 0, err
	}

	if totalWeight == 0 {
		return 0, 0, nil
	}

	return totalWeightedValue, totalWeight, nil
}

func extractValueByName(v reflect.Value, name string) reflect.Value {
	v = reflect.Indirect(v) // Dereference if it's a pointer

	switch v.Kind() {
	case reflect.Float64, reflect.Struct:
		// Directly handle float64 and struct
		if v.Kind() == reflect.Float64 {
			return v
		}
		return v.FieldByName(name)

	case reflect.Slice, reflect.Array:
		// Iterate over slice or array elements
		for i := 0; i < v.Len(); i++ {
			sliceItem := reflect.Indirect(v.Index(i)) // Dereference if it's a pointer
			if sliceItem.Kind() == reflect.Map {
				// Handle slices of maps
				if val := sliceItem.MapIndex(reflect.ValueOf(name)); val.IsValid() {
					return val
				}
			} else {
				// Recursively handle other types in slices
				val := extractValueByName(sliceItem, name)
				if val.IsValid() {
					return val
				}
			}
		}

	case reflect.Map:
		// Handle map directly
		return v.MapIndex(reflect.ValueOf(name))

	case reflect.Interface:
		// Handle interface by recursing on its element
		if !v.IsNil() {
			return extractValueByName(v.Elem(), name)
		}

	default:
		// Print a message for unhandled types
		fmt.Printf("Unhandled type in extractValueByName: %s\n", v.Kind())
	}

	return reflect.Value{}
}

func extractFloatValue(value reflect.Value) float64 {
	value = reflect.Indirect(value) // Dereference if it's a pointer

	switch value.Kind() {
	case reflect.Float64:
		return value.Float()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(value.Uint())

	case reflect.Slice, reflect.Array:
		var sum float64
		for i := 0; i < value.Len(); i++ {
			sum += extractFloatValue(value.Index(i))
		}
		if value.Len() > 0 {
			return sum / float64(value.Len())
		}
		return 0

	case reflect.Map:
		var sum float64
		var count float64
		for _, key := range value.MapKeys() {
			mapValue := value.MapIndex(key)
			sum += extractFloatValue(mapValue)
			count++
		}
		if count > 0 {
			return sum / count
		}
		return 0

	case reflect.Struct:
		var sum float64
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			sum += extractFloatValue(field)
		}
		return sum

	case reflect.Interface:
		if !value.IsNil() {
			return extractFloatValue(value.Elem())
		}
		return 0

	default:
		return 0
	}
}

func normalizeData(data float64, maxValue float64, weight float64) float64 {
	if maxValue == 0 {
		return 0
	}

	// Normalize the data based on its proportion to the maximum absolute value.
	normalizedData := data / maxValue

	// Apply the weight. Since weights can be negative, they will affect the direction of the impact.
	weightedData := normalizedData * weight

	return weightedData
}

func findMaxAbsoluteValue(vData reflect.Value) float64 {
	var maxValue float64

	// Helper function to update maxValue if necessary
	updateMaxValue := func(value float64) {
		absValue := math.Abs(value)
		if absValue > maxValue {
			maxValue = absValue
		}
	}

	var findMax func(reflect.Value)
	findMax = func(value reflect.Value) {
		value = reflect.Indirect(value) // Dereference pointers

		switch value.Kind() {
		case reflect.Float64:
			updateMaxValue(value.Float())

		case reflect.Slice, reflect.Array:
			for i := 0; i < value.Len(); i++ {
				findMax(value.Index(i))
			}

		case reflect.Map:
			for _, key := range value.MapKeys() {
				mapValue := value.MapIndex(key)
				findMax(mapValue)
			}

		case reflect.Struct:
			for i := 0; i < value.NumField(); i++ {
				structFieldValue := value.Field(i)
				findMax(structFieldValue)
			}

		case reflect.Interface:
			if !value.IsNil() {
				findMax(value.Elem())
			}
		}
	}

	findMax(vData)
	return maxValue
}
