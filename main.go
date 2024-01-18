package main

import (
	"encoding/json"
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"os"
	"sort"
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

	if os.Getenv("DEBUG") == "true" {
		APIClient.Debug = true
	}

	// Get rating by symbol
	_, err = APIClient.CompanyValuation.Rating("AAPL")
	if err != nil {
		panic(fmt.Sprintf("Failed to confirm initialization of API client: %s", err.Error()))
	}

	RiskFreeRate := float64(5.0)
	MarketReturn := float64(7.2)

	TickerList := []string{"MSFT", "AAPL", "GOOGL", "GME", "AMC"}
	SecAnalysisWeights := Optimization.SecurityAnalysisWeights{}
	utils.InitStructWithRandomFloats(&SecAnalysisWeights)

	type Result struct {
		Ticker string
		Value  float64
	}

	ResultsMap := []Result{}

	for _, Ticker := range TickerList {

		fundamentals, err := Analysis.PullCompanyFundamentals(APIClient, Ticker, "quarter")
		if err != nil {
			fmt.Printf("failed to pull fundamentals for %s: %s\n", Ticker, err.Error())
		}

		FMPDCF, FMPMeanSTDDCF, err := Analysis.PullCompanyDCFs(APIClient, Ticker)
		if err != nil {
			fmt.Printf("failed to pull DCFs for %s: %s\n", Ticker, err.Error())
		}

		Ratings, RatingsGrowth, RatingsMeanSTD, RatingsGrowthMeanSTD, err := Analysis.PullCompanyRatings(APIClient, Ticker)
		if err != nil {
			fmt.Printf("failed to pull ratings for %s: %s\n", Ticker, err.Error())
		}

		CompanyOutlookObj, err := Analysis.PullCompanyOutlook(APIClient, Ticker)
		if err != nil {
			fmt.Printf("failed to pull outlook for %s: %s\n", Ticker, err.Error())
		}

		EmployeeCount, err := Analysis.PullEmployeeCount(APIClient, Ticker)
		if err != nil {
			fmt.Printf("failed to pull employee count for %s: %s\n", Ticker, err.Error())
		}

		CalculationResults := Analysis.PerformFundamentalsCalculations(fundamentals, "quarter", RiskFreeRate, MarketReturn, CompanyOutlookObj, EmployeeCount)

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
			fmt.Printf("failed to calculate weighted average: %s\n", err.Error())
		} else {
			ResultsMap = append(ResultsMap, Result{Ticker: Ticker, Value: FinalValue})
		}
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
