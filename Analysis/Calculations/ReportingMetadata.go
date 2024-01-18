package Calculations

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
)

func CalculateMeanSTDObjs(objects []interface{}) (map[string][]interface{}, error) {
	if len(objects) == 0 {
		return nil, fmt.Errorf("no objects provided")
	}

	fieldStats := make(map[string][]interface{}) // Map to store field stats with interface{} to allow nil

	for _, obj := range objects {
		val := reflect.ValueOf(obj)

		if val.Kind() == reflect.Slice {
			for i := 0; i < val.Len(); i++ {
				elem := val.Index(i)
				err := processElement(elem, fieldStats)
				if err != nil {
					return nil, err
				}
			}
		} else {
			err := processElement(val, fieldStats)
			if err != nil {
				return nil, err
			}
		}
	}

	// Compute mean and standard deviation for each field
	for key, stats := range fieldStats {
		if len(stats) < 3 {
			continue // Skip if stats are incomplete
		}

		count, ok := stats[2].(float64)
		if !ok || count == 0 {
			// Skip or handle fields with no data points
			continue
		}

		if count == 1 {
			// If there's only one element, use it as the mean, and set stdDev to nil
			mean := stats[0].(float64)
			fieldStats[key] = []interface{}{mean, nil}
			continue
		}

		// Compute mean and standard deviation for fields with more than one data point
		mean := stats[0].(float64) / count
		variance := (stats[1].(float64) - mean*mean*count) / (count - 1)
		stdDev := math.Sqrt(variance)
		fieldStats[key] = []interface{}{mean, stdDev}
	}

	return fieldStats, nil
}

func processElement(val reflect.Value, fieldStats map[string][]interface{}) error {
	// Check if the value is a pointer and dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Handle the case for struct, map, and slice of maps
	switch val.Kind() {
	case reflect.Struct:
		processStruct(val, fieldStats)
	case reflect.Map:
		processMap(val, fieldStats)
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			sliceElem := val.Index(i)
			if sliceElem.Kind() == reflect.Map {
				processMap(sliceElem, fieldStats)
			}
		}
	default:
		return fmt.Errorf("expected struct, map, or slice of maps, got %s", val.Kind())
	}

	return nil
}

func processStruct(val reflect.Value, fieldStats map[string][]interface{}) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		key := val.Type().Field(i).Name

		processField(field, key, fieldStats)
	}
}

func processMap(val reflect.Value, fieldStats map[string][]interface{}) {
	for _, key := range val.MapKeys() {
		mapValue := val.MapIndex(key)
		if mapValue.Kind() == reflect.Ptr && !mapValue.IsNil() {
			mapValue = mapValue.Elem()
		}
		if mapValue.Kind() == reflect.Float64 {
			processField(mapValue, key.String(), fieldStats)
		} else if mapValue.Kind() == reflect.Slice {
			processSlice(mapValue, key.String(), fieldStats)
		}
	}
}

func processSlice(sliceVal reflect.Value, key string, fieldStats map[string][]interface{}) {
	var sum, sumOfSquares, count float64

	for i := 0; i < sliceVal.Len(); i++ {
		elem := sliceVal.Index(i)
		if elem.Kind() == reflect.Float64 {
			value := elem.Float()
			sum += value
			sumOfSquares += value * value
			count++
		}
	}

	if count > 0 {
		mean := sum / count
		variance := sumOfSquares - mean*mean*count
		var stdDev float64
		if count > 1 {
			stdDev = math.Sqrt(variance / (count - 1))
		}

		// Update field statistics for the key
		updateMeanStdDevStats(fieldStats, key, mean, stdDev)
	}
}

func updateMeanStdDevStats(fieldStats map[string][]interface{}, key string, mean, stdDev float64) {
	fieldStats[key+"Mean"] = []interface{}{mean, nil}
	if stdDev > 0 {
		fieldStats[key+"StdDev"] = []interface{}{stdDev, nil}
	} else {
		fieldStats[key+"StdDev"] = []interface{}{nil, nil}
	}
}

func processField(field reflect.Value, key string, fieldStats map[string][]interface{}) {
	if field.Kind() == reflect.Float64 {
		fieldValue := field.Float()
		updateFieldStats(fieldStats, key, fieldValue)
	} else if field.Kind() == reflect.Interface && !field.IsNil() {
		if convertedValue, ok := field.Interface().(float64); ok {
			updateFieldStats(fieldStats, key, convertedValue)
		}
	}
}

func updateFieldStats(fieldStats map[string][]interface{}, key string, value float64) {
	stats, exists := fieldStats[key]
	if !exists {
		// Initialize stats with zero values for sum, sum of squares, and count
		stats = []interface{}{0.0, 0.0, 0.0}
	}

	// Safely update the stats slice elements
	sum := stats[0].(float64) + value                // Update sum
	sumOfSquares := stats[1].(float64) + value*value // Update sum of squares
	count := stats[2].(float64) + 1                  // Update count

	// Assign updated values back to the stats slice
	stats[0] = sum
	stats[1] = sumOfSquares
	stats[2] = count

	// Update the fieldStats map
	fieldStats[key] = stats
}

func ProcessReportDates(dates []string, increment string) (string, string, int, int, int) {
	// Sort the dates
	sort.Strings(dates)

	// Parse the dates
	parsedDates := make([]time.Time, len(dates))
	for i, dateStr := range dates {
		date, _ := time.Parse("2006-01-02", dateStr)
		parsedDates[i] = date
	}

	// Get earliest and latest dates
	earliestDate := parsedDates[0].Format("2006-01-02")
	latestDate := parsedDates[len(parsedDates)-1].Format("2006-01-02")

	// Calculate periods
	missingPeriods, consecutivePeriods, gapPeriods := calculatePeriods(parsedDates, increment)

	return earliestDate, latestDate, missingPeriods, consecutivePeriods, gapPeriods
}

func calculatePeriods(dates []time.Time, increment string) (int, int, int) {
	// Initialize variables
	var missingPeriods, maxConsecutive, currentConsecutive, gapPeriods int
	var periodDuration time.Duration

	// Determine the period duration
	if increment == "quarterly" {
		periodDuration = 24 * time.Hour * 90 // Approximate quarter duration
	} else {
		periodDuration = 24 * time.Hour * 365 // Approximate year duration
	}

	// Iterate through dates to calculate periods
	for i := 0; i < len(dates)-1; i++ {
		diff := dates[i+1].Sub(dates[i])

		// Check for consecutive periods
		if diff <= periodDuration {
			currentConsecutive++
			if currentConsecutive > maxConsecutive {
				maxConsecutive = currentConsecutive
			}
		} else {
			currentConsecutive = 0
			gapPeriods++
		}

		// Calculate missing periods
		missingPeriods += int(diff/periodDuration) - 1
	}

	return missingPeriods, maxConsecutive + 1, gapPeriods
}

func CalculateGrowthF64P(data []map[string]*float64) map[string][]float64 {
	growthMap := make(map[string][]float64)

	// Previous values for each key to calculate growth
	prevValues := make(map[string]float64)

	for _, entry := range data {
		for key, valuePtr := range entry {
			if valuePtr != nil {
				value := *valuePtr

				// Check if we have a previous value
				if prevValue, exists := prevValues[key]; exists && prevValue != 0 {
					// Calculate growth
					growth := (value - prevValue) / prevValue
					growthMap[key] = append(growthMap[key], growth)
				} else {
					// For the first non-nil value, growth is 0
					growthMap[key] = append(growthMap[key], 0)
				}

				// Update previous value
				prevValues[key] = value
			}
			// If value is nil, do nothing for this iteration
		}
	}

	// Remove keys that only had nil values
	for key, growths := range growthMap {
		if len(growths) == 0 {
			delete(growthMap, key)
		}
	}

	return growthMap
}
