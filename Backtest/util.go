package Backtest

import "math"

func mean(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func stdDev(numbers []float64) float64 {
	m := mean(numbers)
	sum := 0.0
	for _, number := range numbers {
		sum += math.Pow(number-m, 2)
	}
	return math.Sqrt(sum / float64(len(numbers)-1))
}

func calculateSharpeRatio(avgReturn, riskFreeRate, stdDev float64) float64 {
	if stdDev == 0 {
		return 0.0
	}
	return (avgReturn - riskFreeRate) / stdDev * math.Sqrt(252)
}

func calculateMaxDrawdown(cumulativeReturns []float64) float64 {
	peak := cumulativeReturns[0]
	maxDrawdown := 0.0

	for _, value := range cumulativeReturns {
		if value > peak {
			peak = value
		}
		drawdown := (peak - value) / peak
		if drawdown > maxDrawdown {
			maxDrawdown = drawdown
		}
	}

	return maxDrawdown
}

func downsideDeviation(numbers []float64) float64 {
	var downsideSum float64
	var count float64

	for _, number := range numbers {
		if number < 0 {
			downsideSum += number * number
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return math.Sqrt(downsideSum / count)
}
