package Backtest

import "github.com/spacecodewor/fmpcloud-go/objects"

type BackTestParameters struct {
	Strategies           []string
	StartingCash         float64
	RiskFreeRate         float64
	Candles              []map[string][]objects.StockCandle
	StockOrder           []string
	StockSpecificWeights map[string]float64
}

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
