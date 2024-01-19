package GeneticAlgorithm

import (
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"math/rand"
	"reflect"
)

func Mutate(s *Optimization.SecurityAnalysisWeights, mutationRate, maxWeightChange, minWeightChange float64) {
	if maxWeightChange < minWeightChange {
		panic("maxWeightChange must be greater than or equal to minWeightChange")
	}
	mutateStruct(reflect.ValueOf(s).Elem(), mutationRate, maxWeightChange, minWeightChange)
}

func mutateStruct(v reflect.Value, mutationRate, maxWeightChange, minWeightChange float64) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Float64:
			if rand.Float64() < mutationRate {
				mutationAmount := rand.Float64()*(maxWeightChange-minWeightChange) + minWeightChange
				newValue := field.Float() + mutationAmount
				if newValue > 1.0 {
					newValue = 1.0
				} else if newValue < -1.0 {
					newValue = -1.0
				}
				field.SetFloat(newValue)
			}
		case reflect.Struct:
			mutateStruct(field, mutationRate, maxWeightChange, minWeightChange)
		}
	}
}
