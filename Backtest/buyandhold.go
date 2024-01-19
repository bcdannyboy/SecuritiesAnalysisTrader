package Backtest

import (
	"fmt"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
	"sort"
	"time"
)

func BuyAndHold(candles []objects.StockCandle, riskFreeRate float64) BackTestResults {
	firstDate := candles[0].Date
	lastDate := candles[len(candles)-1].Date
	fmt.Printf("Candlesticks from %s to %s\n", firstDate, lastDate)
	sort.Slice(candles, func(i, j int) bool {
		return candles[i].Date < candles[j].Date
	})

	// Total Profit/Loss
	totalProfitLoss := candles[len(candles)-1].Close - candles[0].Open

	// Calculating returns, volatility, and drawdowns
	var returns []float64
	var yoyProfitLoss = make(map[string]float64)
	maxDrawdown := 0.0
	peak := candles[0].Close
	prevYear, _ := time.Parse("2006-01-02 15:04:05", candles[0].Date)
	totalReturn := 0.0

	for i, candle := range candles {
		if i > 0 {
			dailyReturn := (candle.Close - candles[i-1].Close) / candles[i-1].Close
			returns = append(returns, dailyReturn)
			totalReturn += dailyReturn

			// Drawdown calculation
			if candle.Close > peak {
				peak = candle.Close
			}
			drawdown := (peak - candle.Close) / peak
			if drawdown > maxDrawdown {
				maxDrawdown = drawdown
			}
		}

		// Year-over-Year calculation
		currentYear, _ := time.Parse("2006-01-02 15:04:05", candle.Date)
		if currentYear.Year() != prevYear.Year() {
			yoyProfitLoss[fmt.Sprintf("%d", prevYear.Year())] = totalReturn
			totalReturn = 0
		}
		prevYear = currentYear
	}
	if totalReturn != 0 {
		yoyProfitLoss[fmt.Sprintf("%d", prevYear.Year())] = totalReturn
	}

	// Volatility (Annualized)
	avgReturn := mean(returns)
	stdDev := stdDev(returns)
	volatility := stdDev * math.Sqrt(252)

	// Sharpe Ratio
	sharpeRatio := 0.0
	if stdDev != 0 {
		sharpeRatio = (avgReturn - riskFreeRate) / stdDev * math.Sqrt(252)
	}

	avgYoYProfitLoss := 0.0
	for _, profitLoss := range yoyProfitLoss {
		avgYoYProfitLoss += profitLoss
	}
	avgYoYProfitLoss /= float64(len(yoyProfitLoss))

	return BackTestResults{
		TotalProfitLoss:  totalProfitLoss,
		AnnualizedReturn: avgReturn,
		Volatility:       volatility,
		SharpeRatio:      sharpeRatio,
		MaxDrawdown:      maxDrawdown,
		YoYProfitLoss:    yoyProfitLoss,
	}
}
