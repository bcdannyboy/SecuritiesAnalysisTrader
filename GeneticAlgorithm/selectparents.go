package GeneticAlgorithm

import (
	"math/rand"
	"sort"
	"sync"
)

func SelectParents(ga *GeneticAlgorithm) (GeneticAlgorithmInput, GeneticAlgorithmInput) {
	// Create a wait group for managing concurrency
	var wg sync.WaitGroup

	// Channels for roulette wheel and tournament results
	rouletteWheelChan := make(chan []GeneticAlgorithmInput, 2)
	tournamentResultChan := make(chan GeneticAlgorithmInput, 2)

	// Total score calculation for roulette wheel
	totalScore := 0.0
	for _, input := range ga.AlgorithmInputs {
		totalScore += input.SecurityAnalysisScore
	}

	// Concurrent Roulette Wheel Selection
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			rouletteWheel := make([]GeneticAlgorithmInput, 0)
			for j := 0; j < len(ga.AlgorithmInputs); j++ {
				pick := rand.Float64() * totalScore
				current := 0.0
				for _, input := range ga.AlgorithmInputs {
					current += input.SecurityAnalysisScore
					if current >= pick {
						rouletteWheel = append(rouletteWheel, input)
						break
					}
				}
			}
			rouletteWheelChan <- rouletteWheel
		}()
	}

	// Concurrent Tournament Selection
	wg.Add(2)
	go func() {
		defer wg.Done()
		pool := <-rouletteWheelChan
		tournamentResultChan <- tournamentSelection(pool, ga.TournamentThreshold)
	}()
	go func() {
		defer wg.Done()
		pool := <-rouletteWheelChan
		tournamentResultChan <- tournamentSelection(pool, ga.TournamentThreshold)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(rouletteWheelChan)
	close(tournamentResultChan)

	// Retrieve tournament winners
	parent1 := <-tournamentResultChan
	parent2 := <-tournamentResultChan

	return parent1, parent2
}

func tournamentSelection(pool []GeneticAlgorithmInput, threshold float64) GeneticAlgorithmInput {
	sort.Slice(pool, func(i, j int) bool {
		return pool[i].SecurityAnalysisScore > pool[j].SecurityAnalysisScore
	})
	topCompetitors := int(float64(len(pool)) * threshold)
	if topCompetitors == 0 {
		topCompetitors = 1
	}
	winnerIndex := rand.Intn(topCompetitors)
	return pool[winnerIndex]
}
