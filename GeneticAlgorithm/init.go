package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
)

func InitGeneticAlgorithm(
	Companies []Analysis.CompanyData,
	Population int,
	Generations int,
	MutationRate float64,
	CrossoverRate float64,
	MaxWeightChange float64,
	MinWeightChange float64,
	RiskFreeRate float64) {
	AlgorithmInputs := []GeneticAlgorithmInput{}

	// first we initialize a random weight for the genetic algorithm's initial inputs
	SecAnalysisWeights := Optimization.SecurityAnalysisWeights{}
	utils.InitStructWithRandomFloats(&SecAnalysisWeights)

	for _, Company := range Companies {
		SecAnalysisScore, err := Optimization.CalculateWeightedAverage(SecAnalysisWeights, Company.Data, "root")
		if err != nil {
			fmt.Printf("failed to calculate weighted average for company %s: %s\n", Company.Ticker, err.Error())
			continue // for now just ignore the company if we can't calculate the weighted average
		}

		AlgorithmInputs = append(AlgorithmInputs, GeneticAlgorithmInput{
			Ticker:                  Company.Ticker,
			SecurityAnalysisResults: &Company.Data,
			SecurityAnalysisWeights: &SecAnalysisWeights,
			SecurityAnalysisScore:   SecAnalysisScore,
			CandleSticks:            Company.CandleSticks,
		})
	}

	// sort the inputs by their security analysis score
	for i := 0; i < len(AlgorithmInputs); i++ {
		for j := i + 1; j < len(AlgorithmInputs); j++ {
			if AlgorithmInputs[i].SecurityAnalysisScore < AlgorithmInputs[j].SecurityAnalysisScore {
				AlgorithmInputs[i], AlgorithmInputs[j] = AlgorithmInputs[j], AlgorithmInputs[i]
			}
		}
	}

	// now we can initialize the genetic algorithm
	GeneticAlgorithm := GeneticAlgorithm{
		AlgorithmInputs: AlgorithmInputs,
		Population:      Population,
		Generations:     Generations,
		MutationRate:    MutationRate,
		CrossoverRate:   CrossoverRate,
		MaxWeightChange: MaxWeightChange,
		MinWeightChange: MinWeightChange,
		RiskFreeRate:    RiskFreeRate,
	}

	// run the genetic algorithm
}
