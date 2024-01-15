package main

import (
	"fmt"
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

	//balance_sheet, balance_sheet_growth, balance_sheet_as_reported, balance_sheet_as_reported_growth, discrepancies_between_both_sheet_types, err := Fundamentals.AnalyzeBalanceSheet(APIClient, "AAPL", "quarter")
	//if err != nil {
	//	panic(fmt.Sprintf("Failed to analyze balance sheet: %s", err.Error()))
	//}

	//cashflow_statement, cashflow_statement_growth, cashflow_statement_as_reported, cashflow_statement_as_reported_growth, discrepancies_between_both_sheet_types, err := Fundamentals.AnalyzeCashFlow(APIClient, "AAPL")
	//if err != nil {
	//	panic(fmt.Sprintf("Failed to analyze cash flow: %s", err.Error()))
	//}

	//income_statement, income_statement_growth, income_statement_as_reported, income_statement_as_reported_growth, discrepancies_between_both_sheet_types, err := Fundamentals.AnalyzeIncomeStatement(APIClient, "AAPL", "quarter")
	//if err != nil {
	//	panic(fmt.Sprintf("Failed to analyze income statement: %s", err.Error()))
	//}

}
