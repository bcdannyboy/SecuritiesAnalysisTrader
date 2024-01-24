package GAUtils

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"sort"
)

func ExtractTickers(candles []map[string][]objects.StockCandle) []string {
	tickers := make([]string, 0, len(candles))
	for _, candleMap := range candles {
		for ticker := range candleMap {
			tickers = append(tickers, ticker)
		}
	}
	return tickers
}

func GetTop10Companies(companies []Analysis.CompanyData, weights *Optimization.SecurityAnalysisWeights) []map[string][]objects.StockCandle {
	type weightedCompany struct {
		Company Analysis.CompanyData
		Weight  float64
	}
	weightedCompanies := []weightedCompany{}
	for _, company := range companies {
		score, err := Optimization.CalculateWeightedAverage(company.Data, weights, "path")
		if err != nil {
			fmt.Printf("failed to calculate weighted average for company %s: %s\n", company.Ticker, err.Error())
			continue
		}
		weightedCompanies = append(weightedCompanies, weightedCompany{
			Company: company,
			Weight:  score,
		})
	}

	// Sort companies based on weight
	sort.Slice(weightedCompanies, func(i, j int) bool {
		return weightedCompanies[i].Weight > weightedCompanies[j].Weight
	})

	// Select top 10 companies and create a slice of map[string][]objects.StockCandle
	top10Candles := make([]map[string][]objects.StockCandle, 0, 10)
	for i := 0; i < 10 && i < len(weightedCompanies); i++ {
		company := weightedCompanies[i].Company
		fmt.Printf("got top company %s with weight %f and %d candles\n", company.Ticker, weightedCompanies[i].Weight, len(company.CandleSticks))
		top10Candles = append(top10Candles, map[string][]objects.StockCandle{company.Ticker: company.CandleSticks})
	}

	return top10Candles
}
