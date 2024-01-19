package Backtest

import (
	"fmt"
	"math"
	"sync"
)

// Weights for each component. Adjust these based on your preference.
// TODO: Make these adjustable for optimization
const (
	WeightTotalProfitLoss  = 1.0
	WeightAnnualizedReturn = 1.0
	WeightVolatility       = -1.0 // Negative because we want to minimize volatility
	WeightSharpeRatio      = 1.0
	WeightSortinoRatio     = 1.0
	WeightMaxDrawdown      = -1.0 // Negative because we want to minimize max drawdown
	WeightYoYProfitLoss    = 1.0
)

func BackTest(Params BackTestParameters) map[string]PortfolioResults {
	ResultsMap := make(map[string]PortfolioResults)
	var wg sync.WaitGroup
	resultsChan := make(chan map[string]PortfolioResults)

	for _, strategy := range Params.Strategies {
		wg.Add(1)
		go func(strategy string) {
			defer wg.Done()
			switch strategy {
			case "equalweightbuyandhold":
				fmt.Printf("Initiating Equal Weight Buy and Hold Strategy\n")
				resultsChan <- map[string]PortfolioResults{"equalweightbuyandhold": EqualWeightBuyAndHold(Params.Candles, Params.RiskFreeRate, Params.StartingCash)}
			case "rankedweightbuyandhold":
				fmt.Printf("Initiating Ranked Weight Buy and Hold Strategy\n")
				resultsChan <- map[string]PortfolioResults{"rankedweightbuyandhold": RankedWeightBuyAndHold(Params.Candles, Params.StockOrder, Params.RiskFreeRate, Params.StartingCash)}
			}
		}(strategy)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		for key, value := range result {
			ResultsMap[key] = value
		}
	}

	return ResultsMap
}

func EvaluateResults(results StockResults) float64 {
	// Normalize or constrain values to avoid extreme effects
	normalizedSharpeRatio := math.Max(results.SharpeRatio, -10)   // anything beyond -10 is just as bad
	normalizedSortinoRatio := math.Max(results.SortinoRatio, -10) // anything beyond -10 is just as bad

	// Calculate average YoY Profit Loss
	var totalYoYProfitLoss float64
	for _, profitLoss := range results.YoYProfitLoss {
		totalYoYProfitLoss += profitLoss
	}
	avgYoYProfitLoss := 0.0
	if len(results.YoYProfitLoss) > 0 {
		avgYoYProfitLoss = totalYoYProfitLoss / float64(len(results.YoYProfitLoss))
	}

	// Calculate the weighted score
	score := WeightTotalProfitLoss*results.TotalProfitLoss +
		WeightAnnualizedReturn*results.AnnualizedReturn +
		WeightVolatility*(1/math.Max(results.Volatility, 0.0001)) + // Avoid division by zero
		WeightSharpeRatio*normalizedSharpeRatio +
		WeightSortinoRatio*normalizedSortinoRatio +
		WeightMaxDrawdown*(1/math.Max(results.MaxDrawdown, 0.0001)) + // Avoid division by zero
		WeightYoYProfitLoss*avgYoYProfitLoss

	return score
}
