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

}
