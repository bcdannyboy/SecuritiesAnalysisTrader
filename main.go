package main

import (
	"encoding/json"
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"os"
	"sort"
	"time"
)

const (
	MaxRatePerMinute = 15 // for each item we're doing ~10 API calls, so we need to limit the rate
	WorkerCount      = 10 // Adjust the number of workers as needed
)

type Result struct {
	Ticker string
	Value  float64
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

	RiskFreeRate := float64(5.0)
	MarketReturn := float64(7.2)
	DefaultEffectiveTaxRate := float64(0.21)

	SymbolList, err := APIClient.Stock.AvalibleSymbols()
	if err != nil {
		panic(fmt.Sprintf("Error getting avalible symbols: %s", err.Error()))
	}

	ResultsMap := []Result{}
	fmt.Printf("Resolving %d symbols\n", len(SymbolList))

	// Create channels for tasks and results
	tasks := make(chan objects.StockSymbolList, len(SymbolList))
	results := make(chan Result, len(SymbolList))

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
		fmt.Printf("Submitted task for %s\n", SymbolObj.Symbol)
	}

	close(tasks) // Close the tasks channel as no more tasks will be sent

	// Collect results
	for range SymbolList {
		result := <-results
		fmt.Printf("Got result for %s: %f\n", result.Ticker, result.Value)
		ResultsMap = append(ResultsMap, result)
	}

	sort.Slice(ResultsMap, func(i, j int) bool {
		return ResultsMap[i].Value > ResultsMap[j].Value
	})

	jResultsMap, err := json.MarshalIndent(ResultsMap, "", "    ")
	if err != nil {
		fmt.Printf("failed to marshal results map: %s\n", err.Error())
	} else {
		fmt.Printf("%s\n", string(jResultsMap))
	}
}

func worker(tasks <-chan objects.StockSymbolList, results chan<- Result, APIClient *fmp.APIClient, Debug bool, RiskFreeRate float64, MarketReturn float64, DefaultEffectiveTaxRate float64, SecAnalysisWeights Optimization.SecurityAnalysisWeights) {
	for SymbolObj := range tasks {
		result, err := processSymbol(SymbolObj, APIClient, Debug, RiskFreeRate, MarketReturn, DefaultEffectiveTaxRate, SecAnalysisWeights)
		if err != nil {
			if Debug {
				fmt.Printf("Error processing symbol %s: %s\n", SymbolObj.Symbol, err.Error())
			}
			continue // Skip this symbol on error
		}
		results <- result
	}
}

func processSymbol(SymbolObj objects.StockSymbolList, APIClient *fmp.APIClient, Debug bool, RiskFreeRate float64, MarketReturn float64, DefaultEffectiveTaxRate float64, SecAnalysisWeights Optimization.SecurityAnalysisWeights) (Result, error) {
	Ticker := SymbolObj.Symbol

	fundamentals, err := Analysis.PullCompanyFundamentals(APIClient, Ticker, "quarter")
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull fundamentals for %s: %s\n", Ticker, err.Error())
		}
		return Result{}, err
	}

	FMPDCF, FMPMeanSTDDCF, err := Analysis.PullCompanyDCFs(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull DCFs for %s: %s\n", Ticker, err.Error())
		}
		return Result{}, err
	}

	Ratings, RatingsGrowth, RatingsMeanSTD, RatingsGrowthMeanSTD, err := Analysis.PullCompanyRatings(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull ratings for %s: %s\n", Ticker, err.Error())
		}
		return Result{}, err
	}

	CompanyOutlookObj, err := Analysis.PullCompanyOutlook(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull outlook for %s: %s\n", Ticker, err.Error())
		}
		return Result{}, err
	}

	EmployeeCount, err := Analysis.PullEmployeeCount(APIClient, Ticker)
	if err != nil {
		if Debug {
			fmt.Printf("failed to pull employee count for %s: %s\n", Ticker, err.Error())
		}
		return Result{}, err
	}

	CalculationResults := Analysis.PerformFundamentalsCalculations(fundamentals, "quarter", RiskFreeRate, MarketReturn, CompanyOutlookObj, EmployeeCount, DefaultEffectiveTaxRate)

	FinalResults := Analysis.FinalNumbers{
		CalculationsOutlookFundamentals: CalculationResults,
		FMPDCF:                          FMPDCF,
		FMPMeanSTDDCF:                   FMPMeanSTDDCF,
		FMPRatings:                      Ratings,
		FMPRatingsGrowth:                RatingsGrowth,
		FMPRatingsMeanSTD:               RatingsMeanSTD,
		FMPRatingsGrowthMeanSTD:         RatingsGrowthMeanSTD,
	}

	FinalValue, err := Optimization.CalculateWeightedAverage(SecAnalysisWeights, FinalResults, "root")
	if err != nil {
		if Debug {
			fmt.Printf("failed to calculate weighted average for %s: %s\n", Ticker, err.Error())
		}
		return Result{}, err
	}

	return Result{Ticker: Ticker, Value: FinalValue}, nil
}
