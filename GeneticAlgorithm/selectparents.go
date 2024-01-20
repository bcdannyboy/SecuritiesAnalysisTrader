package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"math"
	"math/rand"
	"sort"
	"sync"
)

func SelectParents(ga *GeneticAlgorithm) (*Optimization.SecurityAnalysisWeights, *Optimization.SecurityAnalysisWeights) {
	// Create a wait group for managing concurrency
	var wg sync.WaitGroup

	// Channels for roulette wheel and tournament results
	rouletteWheelChan := make(chan []*Optimization.SecurityAnalysisWeights, 2)
	tournamentResultChan := make(chan *Optimization.SecurityAnalysisWeights, 2)

	// Total score calculation for roulette wheel
	totalScore := 0.0
	for _, weights := range ga.PopulationWeights {
		totalScore += CalculateTotalScore(weights, ga.Companies)
	}

	// Concurrent Roulette Wheel Selection
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			rouletteWheel := make([]*Optimization.SecurityAnalysisWeights, 0)
			for j := 0; j < len(ga.PopulationWeights); j++ {
				pick := rand.Float64() * totalScore
				current := 0.0
				for _, weights := range ga.PopulationWeights {
					current += CalculateTotalScore(weights, ga.Companies)
					if current >= pick {
						rouletteWheel = append(rouletteWheel, weights)
						break
					}
				}
			}
			rouletteWheelChan <- rouletteWheel
		}()
	}

	// Concurrent Tournament Selection
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			pool := <-rouletteWheelChan
			tournamentResultChan <- tournamentSelection(pool, ga.TournamentThreshold, ga.Companies)
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(rouletteWheelChan)
	close(tournamentResultChan)

	// Retrieve tournament winners
	parent1 := <-tournamentResultChan
	parent2 := <-tournamentResultChan

	// Ensure that parents are not nil
	if parent1 == nil || parent2 == nil {
		panic("SelectParents: failed to select non-nil parents")
	}

	fmt.Printf("Selected parents for crossover\n")
	return parent1, parent2
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
