package Optimization

import (
	"math"
	"reflect"
)

func CalculateWeightedAverage(Weights interface{}, DataToWeight interface{}, Path string) (float64, error) {
	totalWeightedValue, totalWeight, err := calculateWeightedAverageRecursive(Weights, DataToWeight, Path)
	if err != nil {
		return math.NaN(), err
	}

	if totalWeight == 0 {
		return 0, nil
	}

	finalValue := totalWeightedValue / totalWeight
	return finalValue, nil
}

func calculateWeightedAverageRecursive(Weights interface{}, DataToWeight interface{}, Path string) (float64, float64, error) {
	var totalWeightedValue float64
	var totalWeight float64

	vWeights := reflect.ValueOf(Weights)
	vDataToWeight := reflect.ValueOf(DataToWeight)

	if vWeights.Kind() == reflect.Struct {
		for i := 0; i < vWeights.NumField(); i++ {
			weightField := vWeights.Type().Field(i)
			weightValue := vWeights.Field(i)

			if weightValue.Kind() == reflect.Float64 {
				dataValue := extractValueByName(vDataToWeight, weightField.Name)
				if dataValue.IsValid() {
					data := extractFloatValue(dataValue)
					normalizedData := normalizeData(data, weightValue.Float())
					if !math.IsNaN(normalizedData) {
						totalWeightedValue += normalizedData
					}
					totalWeight += math.Abs(weightValue.Float())
				}
			} else if weightValue.Kind() == reflect.Struct {
				dataValue := extractValueByName(vDataToWeight, weightField.Name)
				if dataValue.IsValid() {
					subTotal, subWeight, err := calculateWeightedAverageRecursive(weightValue.Interface(), dataValue.Interface(), Path+"/"+weightField.Name)
					if err != nil {
						return 0, 0, err
					}
					totalWeightedValue += subTotal
					totalWeight += subWeight
				}
			}
		}
	} else if vWeights.Kind() == reflect.Float64 {
		data := extractFloatValue(vDataToWeight)
		normalizedData := normalizeData(data, vWeights.Float())
		if !math.IsNaN(normalizedData) {
			totalWeightedValue += normalizedData
		}
		totalWeight += math.Abs(vWeights.Float())
	}

	if totalWeight == 0 {
		return 0, 0, nil
	}

	return totalWeightedValue, totalWeight, nil
}

func extractValueByName(v reflect.Value, name string) reflect.Value {
	switch v.Kind() {
	case reflect.Struct:
		return v.FieldByName(name)
	case reflect.Map:
		return v.MapIndex(reflect.ValueOf(name))
	case reflect.Slice:
		// Handle slices of maps
		for i := 0; i < v.Len(); i++ {
			sliceItem := v.Index(i)
			if sliceItem.Kind() == reflect.Map {
				if val := sliceItem.MapIndex(reflect.ValueOf(name)); val.IsValid() {
					return val
				}
			}
		}
	}
	return reflect.Value{}
}

func extractFloatValue(value reflect.Value) float64 {
	switch value.Kind() {
	case reflect.Float64:
		return value.Float()
	case reflect.Ptr:
		if value.IsNil() {
			return 0
		}
		return extractFloatValue(value.Elem())
	case reflect.Slice:
		var sum float64
		for i := 0; i < value.Len(); i++ {
			sum += extractFloatValue(value.Index(i))
		}
		return sum / float64(value.Len())
	case reflect.Map:
		var sum float64
		var count float64
		for _, key := range value.MapKeys() {
			mapValue := value.MapIndex(key)
			if mapValue.CanInterface() {
				sum += extractFloatValue(mapValue)
				count++
			}
		}
		if count == 0 {
			return 0
		}
		return sum / count
	default:
		return 0
	}
}

func normalizeData(data float64, weight float64) float64 {
	normalizedData := 1.0
	if data != 0 {
		normalizedData = data / math.Abs(data)
	}
	return normalizedData * weight
}
