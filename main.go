package main

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
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

	SymbolList := Backtest.NASDAQStockTickers
	if Debug {
		SymbolList = Backtest.NASDAQStockTickers[:5]
	}

	CompanyDataObjects := Analysis.PullCompanyData(APIClient, SymbolList, MaxRatePerMinute, WorkerCount, Debug, RiskFreeRate, MarketReturn, DefaultEffectiveTaxRate)
	if len(CompanyDataObjects) == 0 {
		panic("No company data objects returned")
	}

	Sticks := []map[string][]objects.StockCandle{}

	for _, CompanyDataObject := range CompanyDataObjects {
		fmt.Printf("[+] Adding %s to backtest portfolio\n", CompanyDataObject.Ticker)
		Sticks = append(Sticks, map[string][]objects.StockCandle{CompanyDataObject.Ticker: CompanyDataObject.CandleSticks})
	}

	StickKeys := []string{}
	for _, Stick := range Sticks {
		for key, _ := range Stick {
			StickKeys = append(StickKeys, key)
		}
	}

	BacktestResults := Backtest.BackTest(Backtest.BackTestParameters{
		Strategies:   []string{"equalweightbuyandhold", "rankedweightbuyandhold"},
		Candles:      Sticks,
		StartingCash: 10000,
		RiskFreeRate: RiskFreeRate,
		StockOrder:   StickKeys,
	})
	for key, value := range BacktestResults {
		fmt.Printf("Results for Strategy: %s\n", key)
		fmt.Printf("\tTotal Portfolio Profit/Loss: %f\n", value.Total.TotalProfitLoss)
		fmt.Printf("\tPortfolio Annualized Return: %f\n", value.Total.AnnualizedReturn)
		fmt.Printf("\tPortfolio Volatility: %f\n", value.Total.Volatility)
		fmt.Printf("\tPortfolio Sharpe Ratio: %f\n", value.Total.SharpeRatio)
		fmt.Printf("\tPortfolio Sortino Ratio: %f\n", value.Total.SortinoRatio)
		fmt.Printf("\tPortfolio Max Drawdown: %f\n", value.Total.MaxDrawdown)
	}
}
