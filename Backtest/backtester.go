package Backtest

import "github.com/spacecodewor/fmpcloud-go/objects"

func BackTest(Candles []map[string][]objects.StockCandle, RiskFreeRate float64, Strategy string, StartCash float64) map[string]PortfolioResults {
	ResultsMap := map[string]PortfolioResults{}

	switch Strategy {
	case "equalweightbuyandhold":
		ResultsMap["equalweightbuyandhold"] = EqualWeightBuyAndHold(Candles, RiskFreeRate, StartCash)
		break
	}

	return ResultsMap
}
