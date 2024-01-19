package Backtest

type StockResults struct {
	TotalProfitLoss  float64
	AnnualizedReturn float64
	Volatility       float64
	SharpeRatio      float64
	SortinoRatio     float64
	MaxDrawdown      float64
	YoYProfitLoss    map[string]float64
}

type PortfolioResults struct {
	IndividualStocks map[string]StockResults
	Total            StockResults
}
