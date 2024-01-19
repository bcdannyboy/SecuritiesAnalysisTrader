package GeneticAlgorithm

import (
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
)

func CreateGeneticAlgorithmInput(weights *Optimization.SecurityAnalysisWeights, parent GeneticAlgorithmInput) GeneticAlgorithmInput {
	// Calculate the new Security Analysis Score with the new weights
	newSecAnalysisScore, err := Optimization.CalculateWeightedAverage(*weights, *parent.SecurityAnalysisResults, "root")
	if err != nil {
		// Handle the error appropriately. For example, you might log the error and use the parent's score
		newSecAnalysisScore = parent.SecurityAnalysisScore
	}

	// Create a new GeneticAlgorithmInput with the new weights and other necessary data from the parent
	return GeneticAlgorithmInput{
		Ticker:                  parent.Ticker,
		SecurityAnalysisResults: parent.SecurityAnalysisResults,
		SecurityAnalysisWeights: weights,
		SecurityAnalysisScore:   newSecAnalysisScore,
		CandleSticks:            parent.CandleSticks,
	}
}
