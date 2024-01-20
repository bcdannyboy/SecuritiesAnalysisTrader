package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"math/rand"
	"reflect"
)

func Crossover(parent1, parent2 *Optimization.SecurityAnalysisWeights, crossoverRate float64) *Optimization.SecurityAnalysisWeights {
	// Check if parents are nil
	if parent1 == nil || parent2 == nil {
		panic("Crossover: parent1 and parent2 must be non-nil")
	}

	// Create a new instance for the offspring
	offspring := &Optimization.SecurityAnalysisWeights{}

	// Perform crossover on the struct fields
	val1 := reflect.ValueOf(parent1).Elem()
	val2 := reflect.ValueOf(parent2).Elem()
	offspringVal := reflect.ValueOf(offspring).Elem()

	for i := 0; i < val1.NumField(); i++ {
		field1 := val1.Field(i)
		field2 := val2.Field(i)
		fieldOffspring := offspringVal.Field(i)

		switch field1.Kind() {
		case reflect.Float64:
			// Crossover for float64 fields
			if rand.Float64() < crossoverRate {
				fieldOffspring.Set(field2)
			} else {
				fieldOffspring.Set(field1)
			}
		case reflect.Struct:
			// Recursive crossover for nested structs
			fieldOffspring.Set(crossoverStruct(field1, field2, crossoverRate))
		}
	}

	fmt.Printf("Performed crossover\n")
	return offspring
}

func crossoverStruct(v1, v2 reflect.Value, crossoverRate float64) reflect.Value {
	// Create a new struct of the same type as the parents
	vo := reflect.New(v1.Type()).Elem()

	for i := 0; i < v1.NumField(); i++ {
		field1 := v1.Field(i)
		field2 := v2.Field(i)
		fieldOffspring := vo.Field(i)

		switch field1.Kind() {
		case reflect.Float64:
			// Crossover for float64 fields
			if rand.Float64() < crossoverRate {
				fieldOffspring.Set(field2)
			} else {
				fieldOffspring.Set(field1)
			}
		case reflect.Struct:
			// Recursive crossover for nested structs
			fieldOffspring.Set(crossoverStruct(field1, field2, crossoverRate))
		}
	}

	return vo
}
