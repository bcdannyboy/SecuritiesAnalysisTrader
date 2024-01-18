package main

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	MaxRatePerMinute = 20 // for each item we're doing ~10 API calls, so we need to limit the rate
	WorkerCount      = 10
)

type CompanyData struct {
	Ticker string
	Data   Analysis.FinalNumbers
}

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

	SymbolList, err := APIClient.Stock.AvalibleSymbols()
	if err != nil {
		panic(fmt.Sprintf("Error getting avalible symbols: %s", err.Error()))
	}
	if Debug {
		// Limit the number of symbols for debugging
		newList := []objects.StockSymbolList{}
		for i := 0; i < 5; i++ {
			Symbol := SymbolList[rand.Intn(len(SymbolList))]
			fmt.Printf("Debug: Choosing symbol %s\n", Symbol.Symbol)
			newList = append(newList, Symbol)
		}

		SymbolList = newList
	}

	ResultsMap := []CompanyData{}
	fmt.Printf("Resolving %d symbols\n", len(SymbolList))

	// Create channels for tasks and results
	tasks := make(chan objects.StockSymbolList, len(SymbolList))
	results := make(chan CompanyData, len(SymbolList))

	// Start worker goroutines
	for i := 0; i < WorkerCount; i++ {
		go worker(tasks, results, APIClient, Debug, RiskFreeRate, MarketReturn, DefaultEffectiveTaxRate, SecAnalysisWeights)
	}

	// Create a ticker for rate limiting
	ticker := time.NewTicker(time.Minute / MaxRatePerMinute)

	// Distribute tasks
	for _, SymbolObj := range SymbolList {
		<-ticker.C // Wait for the next tick
		tasks <- SymbolObj
	}
	close(tasks) // Close the tasks channel as no more tasks will be sent

	// Collect results
	for range SymbolList {
		result := <-results
		ResultsMap = append(ResultsMap, result)
	}

}

func worker(tasks <-chan objects.StockSymbolList, results chan<- CompanyData, APIClient *fmp.APIClient, Debug bool, RiskFreeRate float64, MarketReturn float64, DefaultEffectiveTaxRate float64, SecAnalysisWeights Optimization.SecurityAnalysisWeights) {
	for SymbolObj := range tasks {
		result, err := processSymbol(SymbolObj, APIClient, Debug, RiskFreeRate, MarketReturn, DefaultEffectiveTaxRate, SecAnalysisWeights)
		if err != nil {
			if Debug {
				fmt.Printf("Error processing symbol %s: %s\n", SymbolObj.Symbol, err.Error())
			}
			// Even in case of error, send a result back to avoid blocking
			results <- CompanyData{Ticker: SymbolObj.Symbol, Data: Analysis.FinalNumbers{}}
			continue
		}

		results <- result
	}
}

func processSymbol(SymbolObj objects.StockSymbolList, APIClient *fmp.APIClient, Debug bool, RiskFreeRate float64, MarketReturn float64, DefaultEffectiveTaxRate float64, SecAnalysisWeights Optimization.SecurityAnalysisWeights) (CompanyData, error) {
	Ticker := SymbolObj.Symbol

	fundamentals, err := Analysis.PullCompanyFundamentals(APIClient, Ticker, "quarter")
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull fundamentals for %s: %s\n", Ticker, err.Error())
		}
		return CompanyData{}, err
	}

	FMPDCF, FMPMeanSTDDCF, err := Analysis.PullCompanyDCFs(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull DCFs for %s: %s\n", Ticker, err.Error())
		}
		return CompanyData{}, err
	}

	Ratings, RatingsGrowth, RatingsMeanSTD, RatingsGrowthMeanSTD, err := Analysis.PullCompanyRatings(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull ratings for %s: %s\n", Ticker, err.Error())
		}
		return CompanyData{}, err
	}

	CompanyOutlookObj, err := Analysis.PullCompanyOutlook(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull outlook for %s: %s\n", Ticker, err.Error())
		}
		return CompanyData{}, err
	}

	EmployeeCount, err := Analysis.PullEmployeeCount(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull employee count for %s: %s\n", Ticker, err.Error())
		}
		return CompanyData{}, err
	}

	CalculationResults := Analysis.PerformFundamentalsCalculations(fundamentals, "quarter", RiskFreeRate, MarketReturn, CompanyOutlookObj, EmployeeCount, DefaultEffectiveTaxRate)

	FromDate := time.Now().AddDate(-20, 0, 0)
	Today := time.Now()

	CandleSticks, err := APIClient.Stock.Candles(objects.RequestStockCandleList{
		Period: "1min",
		Symbol: Ticker,
		From:   &FromDate,
		To:     &Today,
	})
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull candles for %s: %s\n", Ticker, err.Error())
		}
		return CompanyData{}, err
	}

	FinalResults := Analysis.FinalNumbers{
		CandleSticks:                    CandleSticks,
		CalculationsOutlookFundamentals: CalculationResults,
		FMPDCF:                          FMPDCF,
		FMPMeanSTDDCF:                   FMPMeanSTDDCF,
		FMPRatings:                      Ratings,
		FMPRatingsGrowth:                RatingsGrowth,
		FMPRatingsMeanSTD:               RatingsMeanSTD,
		FMPRatingsGrowthMeanSTD:         RatingsGrowthMeanSTD,
	}

	return CompanyData{Ticker: Ticker, Data: FinalResults}, nil
}
