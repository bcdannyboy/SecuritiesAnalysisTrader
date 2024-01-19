package Backtest

import (
	"fmt"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
	"sort"
	"time"
)

func EqualWeightBuyAndHold(portfolio []map[string][]objects.StockCandle, riskFreeRate float64, startAmount float64) PortfolioResults {
	portfolioResults := PortfolioResults{
		IndividualStocks: make(map[string]StockResults),
		Total:            StockResults{YoYProfitLoss: make(map[string]float64)},
	}

	var totalReturns, cumulativeReturns []float64

	for _, stockMap := range portfolio {
		for ticker, candles := range stockMap {
			if len(candles) == 0 {
				continue
			}

			sort.Slice(candles, func(i, j int) bool {
				return candles[i].Date < candles[j].Date
			})

			// Calculations for each stock
			stockWeight := startAmount / float64(len(stockMap))
			initialPrice := candles[0].Open
			sharesBought := stockWeight / initialPrice // Calculate the number of shares that can be bought

			finalPrice := candles[len(candles)-1].Close
			stockProfitLoss := sharesBought * (finalPrice - initialPrice) // Profit/loss based on the number of shares bought

			var stockReturns []float64
			cumulativeReturn := 1.0
			prevYear, _ := time.Parse("2006-01-02 15:04:05", candles[0].Date)
			stockYoYProfitLoss := make(map[string]float64)

			for i, candle := range candles {
				if i > 0 {
					dailyReturn := (candle.Close - candles[i-1].Close) / candles[i-1].Close
					stockReturns = append(stockReturns, dailyReturn)
					cumulativeReturn *= (1 + dailyReturn)
					cumulativeReturns = append(cumulativeReturns, cumulativeReturn)
				}

				// Year-over-Year calculation
				currentYear, _ := time.Parse("2006-01-02 15:04:05", candle.Date)
				if currentYear.Year() != prevYear.Year() {
					stockYoYProfitLoss[fmt.Sprintf("%d", prevYear.Year())] = stockWeight * cumulativeReturn
					cumulativeReturn = 1.0
				}
				prevYear = currentYear
			}
			if len(stockReturns) > 0 {
				stockYoYProfitLoss[fmt.Sprintf("%d", prevYear.Year())] = stockWeight * cumulativeReturn
			}

			avgReturn := mean(stockReturns)
			stdDev := stdDev(stockReturns)
			volatility := stdDev * math.Sqrt(252)
			sharpeRatio := calculateSharpeRatio(avgReturn, riskFreeRate, stdDev)
			downsideDev := downsideDeviation(stockReturns)
			sortinoRatio := calculateSharpeRatio(avgReturn, riskFreeRate, downsideDev) // Reusing Sharpe Ratio calculation as the formula is similar
			maxDrawdown := calculateMaxDrawdown(cumulativeReturns)

			// Store individual stock results
			portfolioResults.IndividualStocks[ticker] = StockResults{
				TotalProfitLoss:  stockProfitLoss,
				AnnualizedReturn: avgReturn,
				Volatility:       volatility,
				SharpeRatio:      sharpeRatio,
				SortinoRatio:     sortinoRatio,
				MaxDrawdown:      maxDrawdown,
				YoYProfitLoss:    stockYoYProfitLoss,
			}

			// Aggregate returns for total portfolio calculation
			totalReturns = append(totalReturns, stockReturns...)
		}
	}

	// Total portfolio calculations
	portfolioProfitLoss := 0.0
	for _, stockResult := range portfolioResults.IndividualStocks {
		portfolioProfitLoss += stockResult.TotalProfitLoss
	}

	avgReturn := mean(totalReturns)
	stdDev := stdDev(totalReturns)
	volatility := stdDev * math.Sqrt(252)
	sharpeRatio := calculateSharpeRatio(avgReturn, riskFreeRate, stdDev)
	downsideDev := downsideDeviation(totalReturns)
	sortinoRatio := calculateSharpeRatio(avgReturn, riskFreeRate, downsideDev)
	maxDrawdown := calculateMaxDrawdown(cumulativeReturns)

	// Aggregate YoY Profit/Loss for the portfolio
	for _, stockResult := range portfolioResults.IndividualStocks {
		for year, profitLoss := range stockResult.YoYProfitLoss {
			portfolioResults.Total.YoYProfitLoss[year] += profitLoss
		}
	}

	portfolioResults.Total = StockResults{
		TotalProfitLoss:  portfolioProfitLoss,
		AnnualizedReturn: avgReturn,
		Volatility:       volatility,
		SharpeRatio:      sharpeRatio,
		SortinoRatio:     sortinoRatio,
		MaxDrawdown:      maxDrawdown,
	}

	return portfolioResults
}
