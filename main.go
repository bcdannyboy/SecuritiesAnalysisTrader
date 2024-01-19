package main

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"os"
	"strconv"
)

const (
	MaxRatePerMinute = 17 // for each item we're doing ~15 API calls, so we need to limit the rate
	WorkerCount      = 10
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err.Error()))
	}

	api_key := os.Getenv("API_KEY")
	if api_key == "" {
		panic("API_KEY is empty")
	}

	APIClient, err := fmp.NewAPIClient(fmp.Config{APIKey: api_key})
	if err != nil {
		panic(fmt.Sprintf("Error creating API client: %s", err.Error()))
	}

	Debug := false
	if os.Getenv("DEBUG") == "true" {
		Debug = true
		APIClient.Debug = true
	}

	SecAnalysisWeights := Optimization.SecurityAnalysisWeights{}
	utils.InitStructWithRandomFloats(&SecAnalysisWeights)

	RiskFreeRateStr := os.Getenv("RiskFreeRate")
	if RiskFreeRateStr == "" {
		panic("RiskFreeRate is empty")
	}
	RiskFreeRate, err := strconv.ParseFloat(RiskFreeRateStr, 64)
	if err != nil {
		panic(fmt.Sprintf("Error parsing RiskFreeRate: %s", err.Error()))
	}

	MarketReturnStr := os.Getenv("MarketReturn")
	if MarketReturnStr == "" {
		panic("MarketReturn is empty")
	}
	MarketReturn, err := strconv.ParseFloat(MarketReturnStr, 64)
	if err != nil {
		panic(fmt.Sprintf("Error parsing MarketReturn: %s", err.Error()))
	}

	DefaultEffectiveTaxRateStr := os.Getenv("DefaultEffectiveTaxRate")
	if DefaultEffectiveTaxRateStr == "" {
		panic("DefaultEffectiveTaxRate is empty")
	}
	DefaultEffectiveTaxRate, err := strconv.ParseFloat(DefaultEffectiveTaxRateStr, 64)
	if err != nil {
		panic(fmt.Sprintf("Error parsing DefaultEffectiveTaxRate: %s", err.Error()))
	}

	TickerList := []string{}
	SymbolList, err := APIClient.Stock.AvalibleSymbols()
	if err != nil {
		panic(fmt.Sprintf("Error getting avalible symbols: %s", err.Error()))
	}
	if Debug {
		TickerList = []string{"MSFT", "AAPL", "GME"}
	} else {
		for _, Symbol := range SymbolList {
			TickerList = append(TickerList, Symbol.Symbol)
		}
	}

	CompanyDataObjects := Analysis.PullCompanyData(APIClient, TickerList, MaxRatePerMinute, WorkerCount, Debug, RiskFreeRate, MarketReturn, DefaultEffectiveTaxRate)
	if len(CompanyDataObjects) == 0 {
		panic("No company data objects returned")
	}

}
