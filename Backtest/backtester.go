package Backtest

import (
	"fmt"
	"sync"
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
