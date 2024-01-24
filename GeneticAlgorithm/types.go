package GeneticAlgorithm

import (
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

type GeneticAlgorithmInput struct {
	Ticker                  string
	SecurityAnalysisResults *Analysis.FinalNumbers
	SecurityAnalysisWeights *Optimization.SecurityAnalysisWeights
	SecurityAnalysisScore   float64
	CandleSticks            []objects.StockCandle
}

type GeneticAlgorithm struct {
	PopulationWeights   []*Optimization.SecurityAnalysisWeights
	Companies           []Analysis.CompanyData
	Population          int
	Generations         int
	MutationRate        float64
	CrossoverRate       float64
	MaxWeightChange     float64
	MinWeightChange     float64
	TournamentThreshold float64
	RiskFreeRate        float64
	RouletteScaleTilt   float64
}
