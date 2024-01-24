package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/GeneticAlgorithm/Crossover"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/GeneticAlgorithm/GAUtils"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/GeneticAlgorithm/Mutation"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/GeneticAlgorithm/ParentSelection"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"math"
	"sync"
	"time"
)

func InitGeneticAlgorithm(Companies []Analysis.CompanyData, Population int, Generations int, MutationRate float64, CrossoverRate float64, TournamentThreshold float64, MaxWeightChange float64, MinWeightChange float64, RiskFreeRate float64, RouletteScaleTilt float64) *Optimization.SecurityAnalysisWeights {
	if Population <= 0 {
		fmt.Println("Error: Population size must be greater than 0.")
		return nil
	}

	PopulationWeights := make([]*Optimization.SecurityAnalysisWeights, Population)
	for i := 0; i < Population; i++ {
		// The InitStructWithRandomFloats function should return a pointer
		SecAnalysisWeights := utils.InitStructWithRandomFloats(new(Optimization.SecurityAnalysisWeights)).(*Optimization.SecurityAnalysisWeights)
		PopulationWeights[i] = SecAnalysisWeights
	}

	GeneticAlgorithm := GeneticAlgorithm{
		PopulationWeights:   PopulationWeights,
		Companies:           Companies,
		Population:          Population,
		Generations:         Generations,
		MutationRate:        MutationRate,
		CrossoverRate:       CrossoverRate,
		MaxWeightChange:     MaxWeightChange,
		MinWeightChange:     MinWeightChange,
		RiskFreeRate:        RiskFreeRate,
		TournamentThreshold: TournamentThreshold,
		RouletteScaleTilt:   RouletteScaleTilt,
	}

	fmt.Printf("Genetic Algorithm initialized with %d weight sets\n", len(PopulationWeights))
	return startEvolution(&GeneticAlgorithm)
}

func startEvolution(ga *GeneticAlgorithm) *Optimization.SecurityAnalysisWeights {
	var bestWeights *Optimization.SecurityAnalysisWeights
	var bestScore float64 = math.Inf(-1)
	var mutex sync.Mutex

	fmt.Printf("Starting evolution for %d generations with a population of %d\n", ga.Generations, ga.Population)

	type offspringResult struct {
		weights *Optimization.SecurityAnalysisWeights
		score   float64
		index   int
	}

	batchSize := GAUtils.DetermineOptimalBatchSize(len(ga.PopulationWeights), len(ga.Companies))
	fmt.Printf("got batch size %d\n", batchSize)

	for generation := 0; generation < ga.Generations; generation++ {
		genStartTime := time.Now()
		fmt.Printf("Starting Generation %d with %d weight sets\n", generation, len(ga.PopulationWeights))

		newGenerationWeights := make([]*Optimization.SecurityAnalysisWeights, len(ga.PopulationWeights))

		// Calculate totalScore for the current generation
		var totalScore float64
		for _, weights := range ga.PopulationWeights {
			score := GAUtils.CalculateTotalScore(weights, ga.Companies)
			totalScore += score
		}
		fmt.Printf("Total score for generation %d: %f\n", generation, totalScore)

		for batchStart := 0; batchStart < len(ga.PopulationWeights); batchStart += batchSize {
			batchStartTime := time.Now()
			batchEnd := utils.Min(batchStart+batchSize, len(ga.PopulationWeights))
			resultsChan := make(chan offspringResult, batchSize)
			var batchWg sync.WaitGroup

			for i := batchStart; i < batchEnd; i++ {
				batchWg.Add(1)
				go func(index int, batchtot int) {
					indexStartTime := time.Now()
					defer batchWg.Done()
					fmt.Printf("getting parent weights for index %d/%d in generation %d\n", index, batchtot, generation)
					parent1Weights, parent2Weights := ParentSelection.SelectParents(ga.PopulationWeights, ga.Companies, totalScore, ga.TournamentThreshold, ga.RouletteScaleTilt)

					fmt.Printf("getting crossover weights for index %d/%d in generation %d\n", index, batchtot, generation)
					offspringWeights := Crossover.Crossover(parent1Weights, parent2Weights, ga.CrossoverRate)

					fmt.Printf("getting mutation weights for index %d/%d in generation %d\n", index, batchtot, generation)
					offspringWeights = Mutation.Mutate(offspringWeights, ga.MutationRate, ga.MaxWeightChange, ga.MinWeightChange)

					fmt.Printf("generating weights key for index %d/%d in generation %d\n", index, batchtot, generation)
					weightsKey := GAUtils.GenerateCacheKey(offspringWeights, ga.Companies)

					var score float64
					var found bool
					if score, found = GAUtils.GetFromCache(weightsKey); !found {
						fitnessStart := time.Now()
						fmt.Printf("getting top 10 companies for index %d/%d in generation %d\n", index, batchtot, generation)
						top10Candles := GAUtils.GetTop10Companies(ga.Companies, offspringWeights)
						fmt.Printf("extracting tickers for index %d/%d in generation %d\n", index, batchtot, generation)
						top10Tickers := GAUtils.ExtractTickers(top10Candles)
						fmt.Printf("calculating fitness for index %d/%d in generation %d\n", index, batchtot, generation)
						score = CalculateFitness(ga, top10Candles, top10Tickers)
						fmt.Printf("fitness for index %d/%d in generation %d is %f\n", index, batchtot, generation, score)
						GAUtils.AddToCache(weightsKey, score)
						fmt.Printf("calculated fitness for index %d/%d in generation %d in %s\n", index, batchtot, generation, time.Since(fitnessStart))
					}
					fmt.Printf("submitting to offspring results channel for index %d/%d in generation %d\n", index, batchtot, generation)
					resultsChan <- offspringResult{offspringWeights, score, index}
					fmt.Printf("completed index %d/%d in generation %d in %s\n", index, batchtot, generation, time.Since(indexStartTime))
				}(i, batchEnd-batchStart)
			}

			batchWg.Wait()
			close(resultsChan) // Close the new channel

			localBestScore := bestScore
			localBestWeights := bestWeights

			for result := range resultsChan {
				if result.weights != nil && result.score > localBestScore {
					localBestScore = result.score
					localBestWeights = result.weights
				}
				newGenerationWeights[result.index] = result.weights
			}

			mutex.Lock()
			if localBestScore > bestScore {
				bestScore = localBestScore
				bestWeights = localBestWeights
				fmt.Printf("got local best score %f in generation %d\n", bestScore, generation)
			}
			mutex.Unlock()
			fmt.Printf("completed batch %d/%d in generation %d in %s\n", batchStart/batchSize, len(ga.PopulationWeights)/batchSize, generation, time.Since(batchStartTime))
		}

		ga.PopulationWeights = newGenerationWeights
		fmt.Printf("Generation %d completed with best score %f in %s\n", generation, bestScore, time.Since(genStartTime))
	}

	// Perform final backtest with the best weights
	if bestWeights != nil {
		top10Candles := GAUtils.GetTop10Companies(ga.Companies, bestWeights)
		top10Tickers := GAUtils.ExtractTickers(top10Candles)

		// Prepare the backtest parameters
		backtestParams := Backtest.BackTestParameters{
			Strategies:   []string{"equalweightbuyandhold", "rankedweightbuyandhold"},
			StartingCash: 10000,
			RiskFreeRate: ga.RiskFreeRate,
			Candles:      top10Candles,
			StockOrder:   top10Tickers,
		}

		// Perform the backtest
		backtestResults := Backtest.BackTest(backtestParams)

		// Print out the backtest statistics
		GAUtils.PrintBacktestResults(backtestResults)
	}

	return bestWeights
}
