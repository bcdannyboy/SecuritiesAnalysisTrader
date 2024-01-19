package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func InitGeneticAlgorithm(
	Companies []Analysis.CompanyData,
	Population int,
	Generations int,
	MutationRate float64,
	CrossoverRate float64,
	TournamentThreshold float64,
	MaxWeightChange float64,
	MinWeightChange float64,
	RiskFreeRate float64) *Optimization.SecurityAnalysisWeights {
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
		AlgorithmInputs:     AlgorithmInputs,
		Population:          Population,
		Generations:         Generations,
		MutationRate:        MutationRate,
		CrossoverRate:       CrossoverRate,
		MaxWeightChange:     MaxWeightChange,
		MinWeightChange:     MinWeightChange,
		RiskFreeRate:        RiskFreeRate,
		TournamentThreshold: TournamentThreshold,
	}

	fmt.Printf("Starting evolution with %d inputs\n", len(GeneticAlgorithm.AlgorithmInputs))
	return startEvolution(&GeneticAlgorithm)
}

func startEvolution(ga *GeneticAlgorithm) *Optimization.SecurityAnalysisWeights {
	var bestWeights *Optimization.SecurityAnalysisWeights
	var bestScore float64
	firstScore := true

	for generation := 0; generation < ga.Generations; generation++ {
		newPopulation := make([]GeneticAlgorithmInput, 0, ga.Population)

		for i := 0; i < ga.Population; i++ {
			// Parent Selection
			parent1, parent2 := SelectParents(ga)

			// Crossover
			offspringWeights := Crossover(parent1.SecurityAnalysisWeights, parent2.SecurityAnalysisWeights, ga.CrossoverRate)

			// Mutation
			Mutate(offspringWeights, ga.MutationRate, ga.MaxWeightChange, ga.MinWeightChange)

			// Create new input with the new weights
			newInput := CreateGeneticAlgorithmInput(offspringWeights, parent1)

			// Convert CandleSticks to the required format for CalculateFitness
			candleMap := make(map[string][]objects.StockCandle)
			candleMap[newInput.Ticker] = newInput.CandleSticks
			candles := []map[string][]objects.StockCandle{candleMap}

			// Calculate fitness score for the new input
			newScore := CalculateFitness(ga.RiskFreeRate, candles, []string{newInput.Ticker})

			// Update the best score and weights if this is the first score or if it's better than the current best
			if firstScore || newScore > bestScore {
				bestScore = newScore
				bestWeights = offspringWeights
				firstScore = false
			}

			// Append new input to the new population
			newPopulation = append(newPopulation, newInput)
		}

		// Replace old population with the new one
		ga.AlgorithmInputs = newPopulation

		// Log the best score of this generation
		fmt.Printf("Generation %d completed. Best Score: %f\n", generation, bestScore)
	}

	// Log the overall best score after all generations
	fmt.Printf("Evolution completed. Best Score: %f\n", bestScore)

	return bestWeights
}
