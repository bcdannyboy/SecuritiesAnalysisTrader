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
	// Roulette Wheel Selection
	var rouletteWheelResults [2][]*Optimization.SecurityAnalysisWeights
	for i := 0; i < 2; i++ {
		fmt.Printf("initiating roulette wheel selection %d\n", i)
		rouletteWheelResults[i] = performRouletteWheelSelection(populationWeights, companies, totalScore)
		fmt.Printf("Completed roulette wheel selection %d\n", i)
	}

	// Tournament Selection
	var parent1, parent2 *Optimization.SecurityAnalysisWeights
	fmt.Printf("choosing parent 1 from roulette wheel results through a tournament with threshold %f\n", tournamentThreshold)
	parent1 = tournamentSelection(rouletteWheelResults[0], tournamentThreshold, companies)
	fmt.Printf("choosing parent 2 from roulette wheel results through a tournament with threshold %f\n", tournamentThreshold)
	parent2 = tournamentSelection(rouletteWheelResults[1], tournamentThreshold, companies)

	if parent1 == nil || parent2 == nil {
		panic("SelectParents: failed to select non-nil parents")
	}

	fmt.Printf("Selected parents for crossover\n")
	return parent1, parent2
}

func performRouletteWheelSelection(populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64) []*Optimization.SecurityAnalysisWeights {
	var wg sync.WaitGroup

	// Decide on a batch size
	batchSize := 100
	batches := (len(populationWeights) + batchSize - 1) / batchSize

	winnersChan := make(chan *Optimization.SecurityAnalysisWeights, len(populationWeights))

	for i := 0; i < batches; i++ {
		wg.Add(1)
		go func(batchStart int) {
			defer wg.Done()
			batchEnd := batchStart + batchSize
			if batchEnd > len(populationWeights) {
				batchEnd = len(populationWeights)
			}

			for j := batchStart; j < batchEnd; j++ {
				pick := rand.Float64() * totalScore
				current := 0.0
				for _, weights := range populationWeights {
					current += CalculateTotalScore(weights, companies)
					if current >= pick {
						winnersChan <- weights
						break
					}
				}
			}
		}(i * batchSize)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(winnersChan)

	// Collect winners from channel
	var rouletteWheel []*Optimization.SecurityAnalysisWeights
	for winner := range winnersChan {
		rouletteWheel = append(rouletteWheel, winner)
	}

	fmt.Printf("returning %d winners from roulette wheel\n", len(rouletteWheel))
	return rouletteWheel
}

func tournamentSelection(pool []*Optimization.SecurityAnalysisWeights, threshold float64, companies []Analysis.CompanyData) *Optimization.SecurityAnalysisWeights {
	var wg sync.WaitGroup

	// Decide on a batch size
	batchSize := 100 // Adjust this according to your needs
	batches := (len(pool) + batchSize - 1) / batchSize

	scoresChan := make(chan struct {
		index int
		score float64
	}, len(pool))

	// Concurrently calculate scores in batches
	for i := 0; i < batches; i++ {
		wg.Add(1)
		go func(batchStart int) {
			defer wg.Done()
			batchEnd := batchStart + batchSize
			if batchEnd > len(pool) {
				batchEnd = len(pool)
			}
			for j := batchStart; j < batchEnd; j++ {
				score := CalculateTotalScore(pool[j], companies)
				scoresChan <- struct {
					index int
					score float64
				}{j, score}
			}
		}(i * batchSize)
	}

	wg.Wait()
	close(scoresChan)

	// Collect scores
	scores := make([]struct {
		index int
		score float64
	}, len(pool))
	for s := range scoresChan {
		scores[s.index] = s
	}

	// Sort the scores
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	topCompetitors := int(float64(len(scores)) * threshold)
	if topCompetitors == 0 {
		topCompetitors = 1
	}

	// Select a random winner from the top competitors
	winnerIndex := rand.Intn(topCompetitors)
	fmt.Printf("returning winner from tournament selection\n")
	return pool[scores[winnerIndex].index]
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
