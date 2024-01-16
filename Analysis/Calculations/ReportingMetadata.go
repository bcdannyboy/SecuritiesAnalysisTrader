package Calculations

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
)

func CalculateMeanSTDObjs(objects []interface{}) (map[string][]float64, error) {
	if len(objects) == 0 {
		return nil, fmt.Errorf("no objects provided")
	}

	fieldStats := make(map[string][]float64) // Map to store field stats

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
		count := stats[2]
		if count == 0 {
			continue
		}
		mean := stats[0] / count
		variance := (stats[1] - (stats[0]*stats[0])/count) / count
		stdDev := math.Sqrt(variance)

		fieldStats[key] = []float64{mean, stdDev}
	}

	return fieldStats, nil
}

func processElement(val reflect.Value, fieldStats map[string][]float64) error {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct or pointer to struct, got %s", val.Kind())
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		key := val.Type().Field(i).Name

		if field.Kind() == reflect.Float64 {
			fieldValue := field.Float()
			updateFieldStats(fieldStats, key, fieldValue)
		} else if field.Kind() == reflect.Interface && !field.IsNil() {
			if convertedValue, ok := field.Interface().(float64); ok {
				updateFieldStats(fieldStats, key, convertedValue)
			}
		}
	}

	return nil
}

func updateFieldStats(fieldStats map[string][]float64, key string, value float64) {
	stats, exists := fieldStats[key]
	if !exists {
		stats = []float64{0, 0, 0} // sum, sum of squares, count
	}
	stats[0] += value         // sum
	stats[1] += value * value // sum of squares
	stats[2]++                // count
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
