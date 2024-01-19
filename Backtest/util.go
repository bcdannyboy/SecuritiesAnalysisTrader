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
