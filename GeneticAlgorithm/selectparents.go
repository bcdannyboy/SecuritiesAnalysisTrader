package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"math"
	"math/rand"
	"sort"
	"sync"
)

func SelectParents(populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64, tournamentThreshold float64) (*Optimization.SecurityAnalysisWeights, *Optimization.SecurityAnalysisWeights) {
	var wg sync.WaitGroup

	rouletteWheelChan := make(chan []*Optimization.SecurityAnalysisWeights, 2)
	tournamentResultChan := make(chan *Optimization.SecurityAnalysisWeights, 2)

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			rouletteWheelChan <- performRouletteWheelSelection(populationWeights, companies, totalScore)
		}()
	}

	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			pool := <-rouletteWheelChan
			tournamentResultChan <- tournamentSelection(pool, tournamentThreshold, companies)
		}()
	}

	wg.Wait()
	close(rouletteWheelChan)
	close(tournamentResultChan)

	parent1 := <-tournamentResultChan
	parent2 := <-tournamentResultChan

	if parent1 == nil || parent2 == nil {
		panic("SelectParents: failed to select non-nil parents")
	}

	fmt.Printf("Selected parents for crossover\n")
	return parent1, parent2
}

func performRouletteWheelSelection(populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64) []*Optimization.SecurityAnalysisWeights {
	rouletteWheel := make([]*Optimization.SecurityAnalysisWeights, 0)
	for j := 0; j < len(populationWeights); j++ {
		pick := rand.Float64() * totalScore
		current := 0.0
		for _, weights := range populationWeights {
			current += CalculateTotalScore(weights, companies)
			if current >= pick {
				rouletteWheel = append(rouletteWheel, weights)
				break
			}
		}
	}
	return rouletteWheel
}

func tournamentSelection(pool []*Optimization.SecurityAnalysisWeights, threshold float64, companies []Analysis.CompanyData) *Optimization.SecurityAnalysisWeights {
	sort.Slice(pool, func(i, j int) bool {
		return CalculateTotalScore(pool[i], companies) > CalculateTotalScore(pool[j], companies)
	})

	topCompetitors := int(float64(len(pool)) * threshold)
	if topCompetitors == 0 {
		topCompetitors = 1
	}
	winnerIndex := rand.Intn(topCompetitors)
	return pool[winnerIndex]
}

// CalculateTotalScore calculates the total score for a given set of weights.
func CalculateTotalScore(weights *Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData) float64 {
	if utils.IsStructEmpty(weights) {
		fmt.Printf("got empty weights in CalculateTotalScore\n")
		return math.Inf(-1)
	}
	var totalScore float64
	for _, company := range companies {
		score, err := Optimization.CalculateWeightedAverage(weights, company.Data, "root")
		if err != nil {
			fmt.Printf("Error calculating weighted average for company %s: %v\n", company.Ticker, err)
			// Decide how to handle the error. For example, you might continue with a default score.
			continue
		}
		if math.IsNaN(score) || math.IsInf(score, 0) {
			fmt.Printf("Invalid score for company %s\n", company.Ticker)
			continue
		}
		totalScore += score
	}
	return totalScore
}
