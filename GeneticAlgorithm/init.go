package GeneticAlgorithm

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
	"sort"
	"sync"
)

func InitGeneticAlgorithm(Companies []Analysis.CompanyData, Population int, Generations int, MutationRate float64, CrossoverRate float64, TournamentThreshold float64, MaxWeightChange float64, MinWeightChange float64, RiskFreeRate float64) *Optimization.SecurityAnalysisWeights {
	if Population <= 0 {
		fmt.Println("Error: Population size must be greater than 0.")
		return nil
	}

	PopulationWeights := make([]*Optimization.SecurityAnalysisWeights, Population)
	for i := 0; i < Population; i++ {
		// The InitStructWithRandomFloats function should return a pointer
		SecAnalysisWeights := utils.InitStructWithRandomFloats(new(Optimization.SecurityAnalysisWeights)).(*Optimization.SecurityAnalysisWeights)
		PopulationWeights[i] = SecAnalysisWeights
		fmt.Printf("Initialized weight set %d\n", i)
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

	// Cache for storing fitness scores
	fitnessCache := make(map[string]float64)

	batchSize := 5000
	for generation := 0; generation < ga.Generations; generation++ {
		fmt.Printf("Starting Generation %d with %d weight sets\n", generation, len(ga.PopulationWeights))

		newGenerationWeights := make([]*Optimization.SecurityAnalysisWeights, len(ga.PopulationWeights))

		for batchStart := 0; batchStart < len(ga.PopulationWeights); batchStart += batchSize {
			batchEnd := batchStart + batchSize
			if batchEnd > len(ga.PopulationWeights) {
				batchEnd = len(ga.PopulationWeights)
			}

			resultsChan := make(chan offspringResult, batchEnd-batchStart)
			var wg sync.WaitGroup

			var totalScore float64
			var scoreMutex sync.Mutex
			// Calculate the total score for the current generation in batches
			chunkSize := (len(ga.PopulationWeights) + batchSize - 1) / batchSize
			for i := 0; i < len(ga.PopulationWeights); i += chunkSize {
				end := i + chunkSize
				if end > len(ga.PopulationWeights) {
					end = len(ga.PopulationWeights)
				}

				wg.Add(1)
				go func(start, end int) {
					defer wg.Done()
					localTotalScore := 0.0
					for j := start; j < end; j++ {
						lScore := CalculateTotalScore(ga.PopulationWeights[j], ga.Companies)
						fmt.Printf("Calculated score %f for weight set %d in generation %d\n", lScore, j, generation)
						localTotalScore += lScore
					}
					fmt.Printf("got local total score %f for start %d to end %d in generation %d\n", localTotalScore, start, end, generation)
					scoreMutex.Lock()
					fmt.Printf("Adding %f to total score from start %d to end %d in generation %d\n", localTotalScore, start, end, generation)
					totalScore += localTotalScore
					scoreMutex.Unlock()
				}(i, end)
			}

			wg.Wait()

			for i := batchStart; i < batchEnd; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()
					fmt.Printf("Creating offspring %d in generation %d\n", index, generation)

					// Parent Selection
					parent1Weights, parent2Weights := SelectParents(ga.PopulationWeights, ga.Companies, totalScore, ga.TournamentThreshold)
					fmt.Printf("Selected parents for offspring %d\n", index)

					// Crossover
					offspringWeights := Crossover(parent1Weights, parent2Weights, ga.CrossoverRate)
					fmt.Printf("Crossover completed for offspring %d\n", index)

					// Mutation
					offspringWeights = Mutate(offspringWeights, ga.MutationRate, ga.MaxWeightChange, ga.MinWeightChange)
					fmt.Printf("Mutation completed for offspring %d\n", index)

					// Generate a unique key for the current weights
					weightsKey := generateWeightsKey(offspringWeights)

					// Check if the fitness score is already in the cache
					if score, found := fitnessCache[weightsKey]; found {
						fmt.Printf("Using cached fitness score for offspring %d\n", index)
						resultsChan <- offspringResult{offspringWeights, score, index}
					} else {
						// Select top 10 companies based on the security analysis score
						top10Candles := getTop10Companies(ga.Companies, offspringWeights)
						fmt.Printf("Selected top 10 companies for offspring %d\n", index)

						// Extract the tickers from the top 10 companies
						top10Tickers := extractTickers(top10Candles)
						fmt.Printf("Top 10 tickers for offspring %d\n", index)

						// Calculate fitness for the top 10 companies
						fitnessScore := CalculateFitness(ga, top10Candles, top10Tickers)
						if math.IsNaN(fitnessScore) || math.IsInf(fitnessScore, 0) {
							fmt.Printf("Invalid fitness score for offspring %d in generation %d\n", index, generation)
							resultsChan <- offspringResult{nil, 0, index}
							return
						}
						fmt.Printf("Calculated fitness score %f for offspring %d\n", fitnessScore, index)

						// Store the calculated fitness score in the cache
						fitnessCache[weightsKey] = fitnessScore
						resultsChan <- offspringResult{offspringWeights, fitnessScore, index}
					}
				}(i)
			}

			wg.Wait()
			close(resultsChan)

			localBestScore := bestScore
			localBestWeights := bestWeights

			for result := range resultsChan {
				if result.weights != nil && result.score > localBestScore {
					localBestScore = result.score
					localBestWeights = result.weights
					fmt.Printf("New best score %f found in generation %d for offspring %d\n", localBestScore, generation, result.index)
				}
				newGenerationWeights[result.index] = result.weights
			}

			mutex.Lock()
			fmt.Printf("Locking mutex in generation %d to calculate best score \n", generation)
			if localBestScore > bestScore {
				bestScore = localBestScore
				bestWeights = localBestWeights
			}
			mutex.Unlock()
		}

		ga.PopulationWeights = newGenerationWeights
		fmt.Printf("Generation %d completed with best score %f\n", generation, bestScore)
	}

	// Perform final backtest with the best weights
	if bestWeights != nil {
		top10Candles := getTop10Companies(ga.Companies, bestWeights)
		top10Tickers := extractTickers(top10Candles)

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
		printBacktestResults(backtestResults)
	}

	return bestWeights
}

func getTop10Companies(companies []Analysis.CompanyData, weights *Optimization.SecurityAnalysisWeights) []map[string][]objects.StockCandle {
	weightedCompanies := make([]weightedCompany, len(companies))
	for i, company := range companies {
		score, err := Optimization.CalculateWeightedAverage(company, weights, "path")
		if err != nil {
			fmt.Printf("failed to calculate weighted average for company %s\n", company.Ticker)
			continue
		}
		weightedCompanies[i] = weightedCompany{
			Company: company,
			Weight:  score,
		}
	}

	// Sort companies based on weight
	sort.Slice(weightedCompanies, func(i, j int) bool {
		return weightedCompanies[i].Weight > weightedCompanies[j].Weight
	})

	// Select top 10 companies and create a slice of map[string][]objects.StockCandle
	top10Candles := make([]map[string][]objects.StockCandle, 0, 10)
	for i := 0; i < 10 && i < len(weightedCompanies); i++ {
		company := weightedCompanies[i].Company
		top10Candles = append(top10Candles, map[string][]objects.StockCandle{company.Ticker: company.CandleSticks})
	}

	return top10Candles
}

func extractTickers(candles []map[string][]objects.StockCandle) []string {
	tickers := make([]string, 0, len(candles))
	for _, candleMap := range candles {
		for ticker := range candleMap {
			tickers = append(tickers, ticker)
		}
	}
	return tickers
}

func printBacktestResults(results map[string]Backtest.PortfolioResults) {
	for strategy, backtestResult := range results {
		fmt.Printf("Result for strategy: %s\n", strategy)
		fmt.Printf("Total Return: %f\n", backtestResult.Total.TotalProfitLoss)
		fmt.Printf("Annualized Return: %f\n", backtestResult.Total.AnnualizedReturn)
		fmt.Printf("Volatility: %f\n", backtestResult.Total.Volatility)
		fmt.Printf("Sharpe Ratio: %f\n", backtestResult.Total.SharpeRatio)
		fmt.Printf("Sortino Ratio: %f\n", backtestResult.Total.SortinoRatio)
		fmt.Printf("Max Drawdown: %f\n", backtestResult.Total.MaxDrawdown)
		fmt.Printf("YoY Profit/Loss:\n")
		for year, profitLoss := range backtestResult.Total.YoYProfitLoss {
			fmt.Printf("\t%s: %f\n", year, profitLoss)
		}

		fmt.Printf("Tickers in portfolio:\n")
		for ticker, _ := range backtestResult.IndividualStocks {
			fmt.Printf("\t%s\n", ticker)
		}

	}
}

func generateWeightsKey(weights *Optimization.SecurityAnalysisWeights) string {
	weightsStr := fmt.Sprintf("%v", weights)
	hash := sha256.New()
	hash.Write([]byte(weightsStr))
	return hex.EncodeToString(hash.Sum(nil))
}
