package GeneticAlgorithm

import (
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"math/rand"
	"reflect"
)

func Crossover(parent1, parent2 *Optimization.SecurityAnalysisWeights, crossoverRate float64) *Optimization.SecurityAnalysisWeights {
	offspring := &Optimization.SecurityAnalysisWeights{}
	crossoverStruct(reflect.ValueOf(parent1).Elem(), reflect.ValueOf(parent2).Elem(), reflect.ValueOf(offspring).Elem(), crossoverRate)
	return offspring
}

func crossoverStruct(v1, v2, vo reflect.Value, crossoverRate float64) {
	for i := 0; i < vo.NumField(); i++ {
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
			crossoverStruct(field1, field2, fieldOffspring, crossoverRate)
		}
	}
}
