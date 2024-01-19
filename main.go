package main

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"math/rand"
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

	SymbolList := Backtest.NASDAQStockTickers
	if Debug {
		RandSymbols := []string{}
		// get 10 random symbols from SymbolList
		for i := 0; i < 10; i++ {
			RandSymbols = append(RandSymbols, SymbolList[rand.Intn(len(SymbolList))])
		}

		SymbolList = RandSymbols
	}

	CompanyDataObjects := Analysis.PullCompanyData(APIClient, SymbolList, MaxRatePerMinute, WorkerCount, Debug, RiskFreeRate, MarketReturn, DefaultEffectiveTaxRate)
	if len(CompanyDataObjects) == 0 {
		panic("No company data objects returned")
	}

}
