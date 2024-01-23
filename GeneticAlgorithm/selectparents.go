package GeneticAlgorithm

import (
	"encoding/json"
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"hash/fnv"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"sync"
)

var scoreCache = make(map[string]float64)
var cacheMutex sync.RWMutex

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
	parent1 = tournamentSelection(rouletteWheelResults[0], tournamentThreshold, companies)
	fmt.Printf("choosing parent 2 from roulette wheel results through a tournament with threshold %f\n", tournamentThreshold)
	parent2 = tournamentSelection(rouletteWheelResults[1], tournamentThreshold, companies)

	if parent1 == nil || parent2 == nil {
		panic("SelectParents: failed to select non-nil parents")
	}

	fmt.Printf("Selected parents for crossover\n")
	return parent1, parent2
}

func performRouletteWheelSelection(populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64, ScaleTilt float64) []*Optimization.SecurityAnalysisWeights {
	numWorkers := 10 // Adjust as needed
	jobs := make(chan int, len(populationWeights))
	results := make(chan *Optimization.SecurityAnalysisWeights, len(populationWeights))

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results, populationWeights, companies, totalScore, ScaleTilt)
	}

	for j := 0; j < len(populationWeights); j++ {
		jobs <- j
	}
	close(jobs)

	var rouletteWheel []*Optimization.SecurityAnalysisWeights
	for a := 1; a <= len(populationWeights); a++ {
		winner := <-results
		rouletteWheel = append(rouletteWheel, winner)
	}
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
	fmt.Printf("calculating scores for %d weights in %d batches for tournaments\n", len(pool), batches)
	for i := 0; i < batches; i++ {
		fmt.Printf("submitting batch %d/%d\n", i, batches)
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
		fmt.Printf("getting score for weight %d in tournament\n", s.index)
		scores[s.index] = s
	}

	fmt.Printf("sorting scores for tournament selection\n")
	// Sort the scores
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	topCompetitors := int(float64(len(scores)) * threshold)
	if topCompetitors == 0 {
		topCompetitors = 1
	}
	fmt.Printf("got %d top competitors for tournament selection\n", topCompetitors)

	seed := utils.GetRandomSeed()
	rnd := rand.New(rand.NewSource(seed))
	// Select a random winner from the top competitors
	winnerIndex := rnd.Intn(topCompetitors)
	fmt.Printf("returning winner from tournament selection with seed: %d\n", seed)
	return pool[scores[winnerIndex].index]
}

func CalculateTotalScore(weights *Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData) float64 {
	// Generate a unique key for the cache based on weights and companies
	cacheKey := generateCacheKey(weights, companies)

	// Check if the score is already in the cache
	cacheMutex.RLock()
	if score, found := scoreCache[cacheKey]; found {
		cacheMutex.RUnlock()
		return score
	}
	cacheMutex.RUnlock()

	// If not in cache, calculate the score
	if utils.IsStructEmpty(weights) {
		fmt.Printf("got empty weights in CalculateTotalScore\n")
		return math.Inf(-1)
	}

	var totalScore float64
	for _, company := range companies {
		score, err := Optimization.CalculateWeightedAverage(weights, company.Data, "root")
		if err != nil {
			fmt.Printf("Error calculating weighted average for company %s: %v\n", company.Ticker, err)
			continue
		}
		if math.IsNaN(score) || math.IsInf(score, 0) {
			fmt.Printf("Invalid score for company %s\n", company.Ticker)
			continue
		}
		totalScore += score
	}

	// Store the calculated score in the cache
	cacheMutex.Lock()
	scoreCache[cacheKey] = totalScore
	cacheMutex.Unlock()

	return totalScore
}

func worker(id int, jobs <-chan int, results chan<- *Optimization.SecurityAnalysisWeights, populationWeights []*Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData, totalScore float64, ScaleTilt float64) {
	for j := range jobs {
		fmt.Printf("in roulette worker %d with job %d\n", id, j)
		seed := utils.GetRandomSeed()
		rnd := rand.New(rand.NewSource(seed))
		pick := rnd.Float64() * (totalScore + ScaleTilt)
		current := 0.0
		for _, weights := range populationWeights {
			current += CalculateTotalScore(weights, companies)
			if current >= pick {
				results <- weights
				break
			}
		}
	}
}

func generateCacheKey(weights *Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData) string {
	hash := fnv.New64a()

	// Try JSON Marshaling
	weightsBytes, err := json.Marshal(weights)
	if err != nil {
		// Log the error
		fmt.Printf("Failed to marshal weights to JSON: %s\n", err.Error())

		// Fallback to using fmt.Sprintf
		fallbackKey := fmt.Sprintf("%v%v", weights, companies)
		hash.Write([]byte(fallbackKey))
	} else {
		// If marshaling is successful, use the JSON bytes
		hash.Write(weightsBytes)
	}

	cacheKey := strconv.FormatUint(hash.Sum64(), 10)

	return cacheKey
}
