package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"math/rand"
	"reflect"
)

func Mutate(s *Optimization.SecurityAnalysisWeights, mutationRate, maxWeightChange, minWeightChange float64) *Optimization.SecurityAnalysisWeights {
	if maxWeightChange < minWeightChange {
		panic("maxWeightChange must be greater than or equal to minWeightChange")
	}

	// Check if s is nil
	if s == nil {
		panic("Mutate: s must be non-nil")
	}

	// Create a new instance for the mutated struct
	val := reflect.ValueOf(s).Elem()
	mutatedVal := mutateStruct(val, mutationRate, maxWeightChange, minWeightChange)

	mutated := mutatedVal.Addr().Interface().(*Optimization.SecurityAnalysisWeights)

	fmt.Printf("Performed mutation\n")
	return mutated
}

func mutateStruct(v reflect.Value, mutationRate, maxWeightChange, minWeightChange float64) reflect.Value {
	vm := reflect.New(v.Type()).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		mutatedField := vm.Field(i)

		switch field.Kind() {
		case reflect.Float64:
			// Apply mutation for float64 fields
			newValue := field.Float()
			if rand.Float64() < mutationRate {
				mutationAmount := rand.Float64()*(maxWeightChange-minWeightChange) + minWeightChange
				newValue += mutationAmount
				if newValue > 1.0 {
					newValue = 1.0
				} else if newValue < -1.0 {
					newValue = -1.0
				}
			}
			mutatedField.SetFloat(newValue)
		case reflect.Struct:
			// Recursive mutation for nested structs
			mutatedField.Set(mutateStruct(field, mutationRate, maxWeightChange, minWeightChange))
		}
	}

	return vm
}
