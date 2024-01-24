package GAUtils

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"math"
)

func CalculateTotalScore(weights *Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData) float64 {
	// Generate a unique key for the cache based on weights and companies
	cacheKey := GenerateCacheKey(weights, companies)

	// First check without acquiring the lock
	score, found := GetFromCache(cacheKey)
	if found {
		return score
	}

	// Double check if the score was calculated and added to the cache while acquiring the lock
	if score, found := GetFromCache(cacheKey); found {
		return score
	}

	// If not in cache, calculate the score
	if utils.IsStructEmpty(weights) {
		fmt.Printf("got empty weights in CalculateTotalScore\n")
		return math.Inf(-1)
	}

	var totalScore float64
	for _, company := range companies {
		score, err := Optimization.CalculateWeightedAverage(weights, company.Data, "root")
		if err != nil {
			fmt.Printf("Error in select parents calculating weighted average for company %s: %s\n", company.Ticker, err.Error())
			continue
		}
		if math.IsNaN(score) || math.IsInf(score, 0) {
			fmt.Printf("Invalid score for csompany %s\n", company.Ticker)
			continue
		}
		totalScore += score
	}

	// Store the calculated score in the cache
	AddToCache(cacheKey, totalScore)

	return totalScore
}
