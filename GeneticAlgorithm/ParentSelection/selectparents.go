package ParentSelection

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
)

func SelectParents(populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64, tournamentThreshold float64, roletteScaleTile float64) (*Optimization.SecurityAnalysisWeights, *Optimization.SecurityAnalysisWeights) {
	// Roulette Wheel Selection
	var rouletteWheelResults [2][]*Optimization.SecurityAnalysisWeights
	for i := 0; i < 2; i++ {
		fmt.Printf("initiating roulette wheel selection %d\n", i)
		rouletteWheelResults[i] = performRouletteWheelSelection(populationWeights, companies, totalScore, roletteScaleTile)
		fmt.Printf("Completed roulette wheel selection %d\n", i)
	}

	// Tournament Selection
	var parent1, parent2 *Optimization.SecurityAnalysisWeights
	fmt.Printf("choosing parent 1 from roulette wheel results through a tournament with threshold %f\n", tournamentThreshold)
	parent1 = tournamentSelection(1, rouletteWheelResults[0], tournamentThreshold, companies)
	fmt.Printf("choosing parent 2 from roulette wheel results through a tournament with threshold %f\n", tournamentThreshold)
	parent2 = tournamentSelection(2, rouletteWheelResults[1], tournamentThreshold, companies)

	if parent1 == nil || parent2 == nil {
		panic("SelectParents: failed to select non-nil parents")
	}

	fmt.Printf("Selected parents for crossover\n")
	return parent1, parent2
}
