package GeneticAlgorithm

import (
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math/rand"
	"sync"
	"time"
)

func getSubsections(data []map[string][]objects.StockCandle) [][]map[string][]objects.StockCandle {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	totalLength := len(data)
	var subsections [][]map[string][]objects.StockCandle

	if totalLength == 0 {
		return subsections // Return empty if no data is present
	}

	// Adjust the number of subsections based on the total length
	numSubsections := 4
	if totalLength < numSubsections {
		numSubsections = totalLength
	}

	// Define the lengths of each subsection
	lengths := make([]int, numSubsections)
	for i := range lengths {
		if totalLength > 0 {
			lengths[i] = rnd.Intn(totalLength/numSubsections) + 1
		} else {
			lengths[i] = 1 // Ensure at least 1 element
		}
	}

	start := 0
	for _, length := range lengths {
		end := start + length
		if end > totalLength {
			end = totalLength
		}
		subsections = append(subsections, data[start:end])
		start = end
	}

	return subsections
}

func CalculateFitness(RiskFreeRate float64, Candles []map[string][]objects.StockCandle, StockOrder []string) float64 {
	SubSections := getSubsections(Candles)

	BacktestInputs := []Backtest.BackTestParameters{
		{
			Strategies:   []string{"equalweightbuyandhold", "rankedweightbuyandhold"},
			StartingCash: 10000,
			RiskFreeRate: RiskFreeRate,
			Candles:      Candles,
			StockOrder:   StockOrder,
		},
	}

	for _, SubSection := range SubSections {
		BacktestInputs = append(BacktestInputs, Backtest.BackTestParameters{
			Strategies:   []string{"equalweightbuyandhold", "rankedweightbuyandhold"},
			StartingCash: 10000,
			RiskFreeRate: RiskFreeRate,
			Candles:      SubSection,
			StockOrder:   StockOrder,
		})
	}

	var wg sync.WaitGroup
	resultsChan := make(chan float64)

	for _, BacktestInput := range BacktestInputs {
		wg.Add(1)
		go func(input Backtest.BackTestParameters) {
			defer wg.Done()
			for _, StrategyResult := range Backtest.BackTest(input) {
				resultsChan <- Backtest.EvaluateResults(StrategyResult.Total)
				for _, IndividualResult := range StrategyResult.IndividualStocks {
					resultsChan <- Backtest.EvaluateResults(IndividualResult)
				}
			}
		}(BacktestInput)
	}

	// Close the results channel once all goroutines are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	resultScores := float64(0.0)
	totalResults := 0
	for result := range resultsChan {
		resultScores += result
		totalResults++
	}

	return resultScores / float64(totalResults)
}
