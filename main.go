package main

import (
	"encoding/json"
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/joho/godotenv"
	fmp "github.com/spacecodewor/fmpcloud-go"
	"os"
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
	Ticker := "MSFT"

	fundamentals, err := Analysis.PullCompanyFundamentals(APIClient, Ticker, "quarter")
	if err != nil {
		fmt.Printf("failed to pull fundamentals for %s: %s\n", Ticker, err.Error())
		// TODO
	}

	FMPDCF, FMPMeanSTDDCF, err := Analysis.PullCompanyDCFs(APIClient, Ticker)
	if err != nil {
		fmt.Printf("failed to pull DCFs for %s: %s\n", Ticker, err.Error())
		// TODO
	}

	Ratings, RatingsGrowth, RatingsMeanSTD, RatingsGrowthMeanSTD, err := Analysis.PullCompanyRatings(APIClient, Ticker)
	if err != nil {
		fmt.Printf("failed to pull ratings for %s: %s\n", Ticker, err.Error())
		// TODO
	}

	CompanyOutlookObj, err := Analysis.PullCompanyOutlook(APIClient, Ticker)
	if err != nil {
		fmt.Printf("failed to pull outlook for %s: %s\n", Ticker, err.Error())
		// TODO
	}

	EmployeeCount, err := Analysis.PullEmployeeCount(APIClient, Ticker)
	if err != nil {
		fmt.Printf("failed to pull employee count for %s: %s\n", Ticker, err.Error())
		// TODO
	}

	CalculationResults := Analysis.PerformFundamentalsCalculations(fundamentals, "quarter", RiskFreeRate, MarketReturn, CompanyOutlookObj, *EmployeeCount)

	jCalculationResults, err := json.MarshalIndent(CalculationResults, "", "	")
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal calculation results: %s", err.Error()))
	}

	filename := fmt.Sprintf("calculation_results_%s.json", fundamentals.Symbol)
	err = os.WriteFile(filename, jCalculationResults, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to write calculation results to file: %s", err.Error()))
	}

}
