package Backtest

import "github.com/spacecodewor/fmpcloud-go/objects"

func BackTest(Candles []objects.StockCandle, RiskFreeRate float64, Strategy string) map[string]BackTestResults {
	ResultsMap := map[string]BackTestResults{}

	switch Strategy {
	case "buyandhold":
		ResultsMap["buyandhold"] = BuyAndHold(Candles, RiskFreeRate)
		break
	}

	return ResultsMap
}
