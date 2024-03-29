package main

import (
	"encoding/json"
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/GeneticAlgorithm"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"math/rand"
	"os"
	"runtime"
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

	numCPU := runtime.NumCPU()
	fmt.Printf("found %d CPUs\n", numCPU)
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

	SymbolList := Backtest.SP500Tickers
	if Debug {
		RandSymbols := []string{}
		symbolMap := make(map[string]bool) // Map to track added symbols

		for len(RandSymbols) < 100 {
			Symbol := SymbolList[rand.Intn(len(SymbolList))]
			fmt.Printf("Debug picking random symbol: %s\n", Symbol)

			// Check if the symbol is already in the map
			if !symbolMap[Symbol] {
				RandSymbols = append(RandSymbols, Symbol)
				symbolMap[Symbol] = true // Mark this symbol as added
			}
		}

		SymbolList = RandSymbols
	}

	fmt.Printf("Retriving data for %d symbols\n", len(SymbolList))
	CompanyDataObjects := Analysis.PullCompanyData(APIClient, SymbolList, MaxRatePerMinute, WorkerCount, Debug, RiskFreeRate, MarketReturn, DefaultEffectiveTaxRate)
	if len(CompanyDataObjects) == 0 {
		panic("No company data objects returned")
	}
	finalCompanyDataObjects := []Analysis.CompanyData{}
	for _, CompanyDataObject := range CompanyDataObjects {
		newCDO := utils.RemoveNaNsFromStruct(CompanyDataObject)
		finalCompanyDataObjects = append(finalCompanyDataObjects, newCDO.(Analysis.CompanyData))
	}

	OptimizedSecurityAnalysisWeights := GeneticAlgorithm.InitGeneticAlgorithm(CompanyDataObjects, 100, 50, 0.1337, 0.787483, 0.73023, 0.1, 0.001, RiskFreeRate, 0.01)
	jOptimizedWeights, err := json.Marshal(OptimizedSecurityAnalysisWeights)

	outname := "optimized_weights.json"
	// output the optimized weights to a file
	f1, err := os.Create(outname)
	defer f1.Close()
	if err != nil {
		panic(fmt.Sprintf("Error creating file: %s", err.Error()))
	}
	_, err = f1.Write(jOptimizedWeights)
	if err != nil {
		fmt.Printf("failed to write to file: %s\n\nOptimized Weights:\n", err.Error())
		fmt.Println(string(jOptimizedWeights))
	}

}
