package GeneticAlgorithm

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Backtest"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math/rand"
	"sync"
	"time"
)

func getSubsections(data []map[string][]objects.StockCandle) [][]map[string][]objects.StockCandle {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	totalLength := len(data)
	var subsections [][]map[string][]objects.StockCandle

	if totalLength == 0 {
		return subsections // Return empty if no data is present
	}

	// Adjust the number of subsections based on the total length
	numSubsections := 4
	if totalLength < numSubsections {
		numSubsections = totalLength
	}

	// Define the lengths of each subsection
	lengths := make([]int, numSubsections)
	for i := range lengths {
		if totalLength > 0 {
			lengths[i] = rnd.Intn(totalLength/numSubsections) + 1
		} else {
			lengths[i] = 1 // Ensure at least 1 element
		}
	}

	start := 0
	for _, length := range lengths {
		end := start + length
		if end > totalLength {
			end = totalLength
		}
		subsections = append(subsections, data[start:end])
		start = end
	}

	return subsections
}

func CalculateFitness(ga *GeneticAlgorithm, candles []map[string][]objects.StockCandle, stockOrder []string) float64 {
	fmt.Println("Entering CalculateFitness function")

	// Check if Candles data is empty
	if len(candles) == 0 {
		fmt.Println("Candles data is empty")
		return 0.0 // Return a default low score or handle as needed
	}

	fmt.Printf("Calculating fitness for %d stocks\n", len(candles))
	subSections := getSubsections(candles)
	fmt.Printf("Testing %d stocks with %d subsections\n", len(candles), len(subSections))

	var wg sync.WaitGroup
	resultsChan := make(chan float64)

	for _, subsection := range subSections {
		wg.Add(1)
		go func(subsection []map[string][]objects.StockCandle) {
			defer wg.Done()

			// Prepare the backtest parameters
			backtestParams := Backtest.BackTestParameters{
				Strategies:   []string{"equalweightbuyandhold", "rankedweightbuyandhold"},
				StartingCash: 10000,
				RiskFreeRate: ga.RiskFreeRate,
				Candles:      subsection,
				StockOrder:   stockOrder,
			}

			// Perform the backtest
			backtestResults := Backtest.BackTest(backtestParams)

			// Evaluate the results and send to resultsChan
			for _, portfolioResults := range backtestResults {
				resultsChan <- Backtest.EvaluateResults(portfolioResults.Total)
			}

		}(subsection)
	}

	// Close the results channel once all goroutines are done
	go func() {
		wg.Wait()
		close(resultsChan)
		fmt.Println("All goroutines completed")
	}()

	resultScores := float64(0.0)
	totalResults := 0
	for result := range resultsChan {
		fmt.Printf("Got result: %f\n", result)
		resultScores += result
		totalResults++
	}

	// Check for division by zero
	if totalResults == 0 {
		fmt.Println("No valid results obtained, returning default score")
		return 0.0 // Return a default low score or handle as needed
	}

	finalScore := resultScores / float64(totalResults)
	fmt.Printf("Final fitness score: %f\n", finalScore)

	return finalScore
}
