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

	fundamentals, err := Analysis.PullCompanyFundamentals(APIClient, "MSFT", "quarter")
	if err != nil {
		panic(fmt.Sprintf("Failed to pull company fundamentals: %s", err.Error()))
	}

	CalculationResults := Analysis.PerformFundamentalsCalculations(fundamentals, "quarter")

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
