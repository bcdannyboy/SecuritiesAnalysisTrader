package ParentSelection

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/GeneticAlgorithm/GAUtils"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"math/rand"
	"sort"
)

func tournamentSelection(tourny int, pool []*Optimization.SecurityAnalysisWeights, threshold float64, companies []Analysis.CompanyData) *Optimization.SecurityAnalysisWeights {
	// Use a worker pool to limit the number of concurrent goroutines
	numWorkers := GAUtils.DetermineOptimalBatchSize(len(pool), len(companies))
	jobs := make(chan int, len(pool))
	results := make(chan struct {
		index int
		score float64
	}, len(pool))

	// Initialize workers
	for w := 1; w <= numWorkers; w++ {
		go tournamentWorker(tourny, w, jobs, results, pool, companies)
	}

	// Distribute jobs
	for j := 0; j < len(pool); j++ {
		jobs <- j
	}
	close(jobs)

	fmt.Printf("completed all worker jobs in tournament %d\n", tourny)

	// Collect results
	scores := make([]struct {
		index int
		score float64
	}, 0, len(pool))
	for a := 0; a < len(pool); a++ {
		scores = append(scores, <-results)
	}

	fmt.Printf("received %d scores, about to sort in tournament %d\n", len(scores), tourny)
	// Sort the scores
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	fmt.Printf("sorted %d scores in tournament %d\n", len(scores), tourny)

	topCompetitors := int(float64(len(scores)) * threshold)
	if topCompetitors == 0 {
		topCompetitors = 1
	}
	fmt.Printf("grabbing top %d competitors from tournament %d\n", topCompetitors, tourny)

	// Select a random winner from the top competitors
	seed := utils.GetRandomSeed()
	rnd := rand.New(rand.NewSource(seed))
	winnerIndex := rnd.Intn(topCompetitors)

	fmt.Printf("returning tournament %d winner at index %d\n", tourny, winnerIndex)
	return pool[scores[winnerIndex].index]
}

func tournamentWorker(tourny int, id int, jobs <-chan int, results chan<- struct {
	index int
	score float64
}, pool []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData) {
	for j := range jobs {
		fmt.Printf("received job %d for tournament %d worker %d\n", j, tourny, id)
		score := GAUtils.CalculateTotalScore(pool[j], companies)
		fmt.Printf("got score for job %d in tournament %d worker %d\n", j, tourny, id)
		results <- struct {
			index int
			score float64
		}{j, score}
	}
}
