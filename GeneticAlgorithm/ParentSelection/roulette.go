package ParentSelection

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/GeneticAlgorithm/GAUtils"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"math/rand"
)

func performRouletteWheelSelection(populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64, ScaleTilt float64) []*Optimization.SecurityAnalysisWeights {
	numWorkers := GAUtils.DetermineOptimalBatchSize(len(populationWeights), len(companies))
	jobs := make(chan int, len(populationWeights))
	results := make(chan *Optimization.SecurityAnalysisWeights, len(populationWeights))

	for w := 1; w <= numWorkers; w++ {
		fmt.Printf("initializing roulette wheel worker %d\n", w)
		go rouletteWorker(w, jobs, results, populationWeights, companies, totalScore, ScaleTilt)
	}

	for j := 0; j < len(populationWeights); j++ {
		jobs <- j
	}
	close(jobs)

	rouletteWheel := []*Optimization.SecurityAnalysisWeights{}
	for i := 0; i < len(populationWeights); i++ {
		fmt.Printf("waiting on roulette wheel winner %d/%d\n", i, len(populationWeights))
		winner := <-results
		rouletteWheel = append(rouletteWheel, winner)
		fmt.Printf("got rot roulette wheel winner %d/%d\n", i, len(populationWeights))
	}
	fmt.Printf("returning %d roulette wheel winners\n", len(rouletteWheel))
	return rouletteWheel
}

func rouletteWorker(id int, jobs <-chan int, results chan<- *Optimization.SecurityAnalysisWeights, populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64, ScaleTilt float64) {
	for j := range jobs {
		seed := utils.GetRandomSeed()
		rnd := rand.New(rand.NewSource(seed + int64(j)))
		pick := rnd.Float64() * (totalScore + ScaleTilt)
		current := 0.0
		for _, weights := range populationWeights {
			current += GAUtils.CalculateTotalScore(weights, companies)
			if current >= pick {
				results <- weights
				break
			}
		}
	}
	fmt.Printf("exiting roulette wheel worker %d\n", id)
}
