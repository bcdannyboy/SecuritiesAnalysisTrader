package Optimization

import (
	"fmt"
	"math"
	"reflect"
)

func CalculateWeightedAverage(Weights interface{}, DataToWeight interface{}, Path string) (float64, error) {
	totalWeightedValue, totalWeight, err := calculateWeightedAverageRecursive(Weights, DataToWeight, Path, 0, map[uintptr]bool{})
	if err != nil {
		fmt.Printf("got error from CalculateWeightedAverage: %s\n", err)
		return math.NaN(), err
	}

	if totalWeight == 0 {
		fmt.Printf("got totalWeight == 0, returning 0\n")
		return 0, fmt.Errorf("got totalWeight == 0")
	}

	finalValue := totalWeightedValue / totalWeight
	return finalValue, nil
}

func calculateWeightedAverageRecursive(Weights interface{}, DataToWeight interface{}, Path string, depth int, visited map[uintptr]bool) (float64, float64, error) {
	const maxDepth = 1000000000
	if depth > maxDepth {
		return 0, 0, fmt.Errorf("maximum recursion depth exceeded at path: %s", Path)
	}

	var totalWeightedValue float64
	var totalWeight float64

	vWeights := reflect.Indirect(reflect.ValueOf(Weights))
	vDataToWeight := reflect.Indirect(reflect.ValueOf(DataToWeight))

	// Check for circular references only if the value is addressable
	if vWeights.CanAddr() {
		ptr := vWeights.UnsafeAddr()
		if visited[ptr] {
			return 0, 0, nil // Circular reference detected, skip processing
		}
		visited[ptr] = true
	}

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

				if !subWeightValue.IsValid() || !subDataValue.IsValid() {
					// Handle invalid values appropriately
					continue
				}

				subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(subWeightValue.Interface(), subDataValue.Interface(), newPath, depth+1, visited)
				if err != nil {
					return err
				}
				totalWeightedValue += subTotalWeightedValue
				totalWeight += subTotalWeight
			}

		case reflect.Slice, reflect.Array:
			for i := 0; i < weightValue.Len(); i++ {
				newPath := fmt.Sprintf("%s[%d]", path, i)
				subWeightValue := weightValue.Index(i)

				if dataValue.Kind() == reflect.Slice || dataValue.Kind() == reflect.Array {
					if i < dataValue.Len() {
						subDataValue := dataValue.Index(i)
						subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(subWeightValue.Interface(), subDataValue.Interface(), newPath, depth+1, visited)
						if err != nil {
							return err
						}
						totalWeightedValue += subTotalWeightedValue
						totalWeight += subTotalWeight
					} else {
						fmt.Printf("Warning: dataValue does not have index %d at path %s\n", i, newPath)
					}
				} else {
					// Recursively handle non-slice/array types
					if dataValue.Kind() == reflect.Struct {
						for j := 0; j < dataValue.NumField(); j++ {
							fieldPath := newPath + "." + dataValue.Type().Field(j).Name
							fieldValue := dataValue.Field(j)
							subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(subWeightValue.Interface(), fieldValue.Interface(), fieldPath, depth+1, visited)
							if err != nil {
								return err
							}
							totalWeightedValue += subTotalWeightedValue
							totalWeight += subTotalWeight
						}
					} else {
						// Handle other types as needed
						fmt.Printf("Warning: dataValue is not a slice, array, or struct at path %s\n", newPath)
					}
				}
			}

		case reflect.Map:
			if weightValue.Kind() == reflect.Map && dataValue.Kind() == reflect.Map {
				for _, key := range weightValue.MapKeys() {
					newPath := path + "[" + fmt.Sprint(key.Interface()) + "]"
					subWeightValue := weightValue.MapIndex(key)
					subDataValue := dataValue.MapIndex(key)
					if !subWeightValue.IsValid() || !subDataValue.IsValid() {
						fmt.Printf("Warning: Invalid map key %v at path %s\n", key, newPath)
						continue
					}
					subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(subWeightValue.Interface(), subDataValue.Interface(), newPath, depth+1, visited)
					if err != nil {
						return err
					}
					totalWeightedValue += subTotalWeightedValue
					totalWeight += subTotalWeight
				}
			} else {
				// Handle non-map types within the map
				if dataValue.Kind() == reflect.Struct {
					for i := 0; i < dataValue.NumField(); i++ {
						field := dataValue.Type().Field(i)
						newPath := path + "." + field.Name
						fieldValue := dataValue.Field(i)
						fieldWeightValue := weightValue.MapIndex(reflect.ValueOf(field.Name)) // Get the corresponding weight for the field
						if !fieldWeightValue.IsValid() {
							continue // Skip if no corresponding weight found
						}
						subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(fieldWeightValue.Interface(), fieldValue.Interface(), newPath, depth+1, visited)
						if err != nil {
							return err
						}
						totalWeightedValue += subTotalWeightedValue
						totalWeight += subTotalWeight
					}
				} else {
					// Handle other types as needed
					if dataValue.CanInterface() {
						// Check if the type of dataValue is one that should be recursively processed
						switch dataValue.Kind() {
						case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map:
							// Only recurse if the type is complex and requires further decomposition
							subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(weightValue.Interface(), reflect.ValueOf(dataValue.Interface()).Interface(), path, depth+1, visited)
							if err != nil {
								return err
							}
							totalWeightedValue += subTotalWeightedValue
							totalWeight += subTotalWeight
						default:
							// For simple types, handle them directly without recursion
							// You can add logic here to handle simple types like integers, floats, etc.
							fmt.Printf("Info: Encountered a simple type at path %s. WeightValue Kind: %s, DataValue Kind: %s\n", path, weightValue.Kind(), dataValue.Kind())
						}
					} else {
						fmt.Printf("Warning: Cannot handle type at path %s. WeightValue Kind: %s, DataValue Kind: %s\n", path, weightValue.Kind(), dataValue.Kind())
					}
				}
			}

		default:
			// ignore if string
			if weightValue.Kind() != reflect.String && weightValue.Kind() != reflect.Interface {
				fmt.Printf("Unhandled type: %s at path %s\n", weightValue.Kind(), path)
			}

			// handle if interface
			if weightValue.Kind() == reflect.Interface {
				if !weightValue.IsNil() {
					subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(weightValue.Elem().Interface(), dataValue.Interface(), path, depth+1, visited)
					if err != nil {
						return err
					}
					totalWeightedValue += subTotalWeightedValue
					totalWeight += subTotalWeight
				}
			}

			if dataValue.CanInterface() {
				subTotalWeightedValue, subTotalWeight, err := calculateWeightedAverageRecursive(weightValue.Interface(), reflect.ValueOf(dataValue.Interface()).Interface(), path, depth+1, visited)
				if err != nil {
					return err
				}
				totalWeightedValue += subTotalWeightedValue
				totalWeight += subTotalWeight
			} else {
				fmt.Printf("Warning: Cannot handle type at path %s. WeightValue Kind: %s, DataValue Kind: %s\n", path, weightValue.Kind(), dataValue.Kind())
			}
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
