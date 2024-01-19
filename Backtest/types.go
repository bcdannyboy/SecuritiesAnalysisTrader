package Backtest

type BackTestResults struct {
	TotalProfitLoss  float64
	SharpeRatio      float64
	MaxDrawdown      float64
	YoYProfitLoss    map[string]float64
	AnnualizedReturn float64
	Volatility       float64
}
