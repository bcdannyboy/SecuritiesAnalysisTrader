package Analysis

import (
	fundamentals "github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

type CompanyFundamentals struct {
	Symbol string

	BalanceSheetStatements                                             []objects.BalanceSheetStatement
	BalanceSheetStatementGrowth                                        []objects.BalanceSheetStatementGrowth
	BalanceSheetStatementAsReported                                    []objects.BalanceSheetStatementAsReported
	GrowthBalanceSheetStatementAsReported                              []*fundamentals.GrowthBalanceSheetStatementAsReported
	DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported []*fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported

	IncomeStatement                                        []objects.IncomeStatement
	IncomeStatementGrowth                                  []objects.IncomeStatementGrowth
	IncomeStatementAsReported                              []objects.IncomeStatementAsReported
	GrowthIncomeStatementAsReported                        []*fundamentals.GrowthIncomeStatementAsReported
	DiscrepancyIncomeStatementAndIncomeStatementAsReported []*fundamentals.DiscrepancyIncomeStatementAndIncomeStatementAsReported

	CashFlowStatement                                          []objects.CashFlowStatement
	CashFlowStatementGrowth                                    []objects.CashFlowStatementGrowth
	CashFlowStatementAsReported                                []objects.CashFlowStatementAsReported
	CashFlowStatementAsReportedGrowth                          []*fundamentals.CashFlowStatementAsReportedGrowth
	DiscrepancyCashFlowStatementAndCashFlowStatementAsReported []*fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported

	FinancialRatios          []objects.FinancialRatios
	FinancialRatiosTTM       []objects.FinancialRatiosTTM
	FinancialRatiosGrowth    []*fundamentals.FinancialRatiosGrowth
	FinancialRatiosTTMGrowth []*fundamentals.FinancialRatiosTTMGrowth

	FullFinancialStatement       []objects.FullFinancialStatementAsReported
	FullFinancialStatementGrowth []objects.FinancialStatementsGrowth
}

type FundamentalsCalculationsResults struct {
	Symbol           string
	Fundamentals     *CompanyFundamentals
	PeriodLength     objects.CompanyValuationPeriod
	CostOfEquity     float64
	Beta             float64
	EffectiveTaxRate float64

	BalanceSheet struct {
		DifferenceInLengthBetweenBalanceSheetStatementAndBalanceSheetStatementAsReported int
		StatementAndReportDiscrepancyGrowth                                              map[string][]float64

		TotalGapsInBalanceSheetStatementPeriods                  int
		TotalConsecutivePeriodsWithNoGapsInBalanceSheetStatement int
		TotalConsecutiveMissingPeriodsInBalanceSheetStatement    int

		TotalGapsInBalanceSheetStatementAsReportedPeriods                  int
		TotalConsecutivePeriodsWithNoGapsInBalanceSheetStatementAsReported int
		TotalConsecutiveMissingPeriodsInBalanceSheetStatementAsReported    int

		MeanSTDBalanceSheetStatement                          map[string][]float64
		MeanSTDBalanceSheetStatementAsReported                map[string][]float64
		MeanSTDBalanceSheetStatementGrowth                    map[string][]float64
		MeanSTDBalanceSheetStatementAsReportedGrowth          map[string][]float64
		MeanSTDBalanceSheetDiscrepancies                      map[string][]float64
		MeanZippedSTDBalanceSheetStatementAndAsReported       map[string][]float64
		MeanZippedSTDBalanceSheetStatementAndAsReportedGrowth map[string][]float64
	}

	IncomeStatement struct {
		DifferenceInLengthBetweenIncomeStatementAndIncomeStatementAsReported int
		StatementAndReportDiscrepancyGrowth                                  map[string][]float64

		TotalGapsInIncomeStatementPeriods                  int
		TotalConsecutivePeriodsWithNoGapsInIncomeStatement int
		TotalConsecutiveMissingPeriodsInIncomeStatement    int

		TotalGapsInIncomeStatementAsReportedPeriods                  int
		TotalConsecutivePeriodsWithNoGapsInIncomeStatementAsReported int
		TotalConsecutiveMissingPeriodsInIncomeStatementAsReported    int

		MeanSTDIncomeStatement                          map[string][]float64
		MeanSTDIncomeStatementAsReported                map[string][]float64
		MeanSTDIncomeStatementGrowth                    map[string][]float64
		MeanSTDIncomeStatementAsReportedGrowth          map[string][]float64
		MeanSTDIncomeStatementDiscrepancies             map[string][]float64
		MeanZippedSTDIncomeStatementAndAsReported       map[string][]float64
		MeanZippedSTDIncomeStatementAndAsReportedGrowth map[string][]float64
	}

	CashFlowStatement struct {
		DifferenceInLengthBetweenCashFlowStatementAndCashFlowStatementAsReported int
		StatementAndReportDiscrepancyGrowth                                      map[string][]float64

		TotalGapsInCashFlowStatementPeriods                  int
		TotalConsecutivePeriodsWithNoGapsInCashFlowStatement int
		TotalConsecutiveMissingPeriodsInCashFlowStatement    int

		TotalGapsInCashFlowStatementAsReportedPeriods                  int
		TotalConsecutivePeriodsWithNoGapsInCashFlowStatementAsReported int
		TotalConsecutiveMissingPeriodsInCashFlowStatementAsReported    int

		MeanSTDCashFlowStatement                          map[string][]float64
		MeanSTDCashFlowStatementAsReported                map[string][]float64
		MeanSTDCashFlowStatementGrowth                    map[string][]float64
		MeanSTDCashFlowStatementAsReportedGrowth          map[string][]float64
		MeanSTDCashFlowStatementDiscrepancies             map[string][]float64
		MeanZippedSTDCashFlowStatementAndAsReported       map[string][]float64
		MeanZippedSTDCashFlowStatementAndAsReportedGrowth map[string][]float64
	}

	FinancialRatios struct {
		FPMRatios          []objects.FinancialRatios
		FPMRatiosTTM       []objects.FinancialRatiosTTM
		FPMRatiosGrowth    []*fundamentals.FinancialRatiosGrowth
		FPMRatiosTTMGrowth []*fundamentals.FinancialRatiosTTMGrowth

		AverageSTDFPMRatios                     map[string][]float64
		AverageSTDFPMRatiosTTM                  map[string][]float64
		AverageSTDFPMRatiosGrowth               map[string][]float64
		AverageSTDFPMRatiosTTMGrowth            map[string][]float64
		AverageSTDFZippedFPMRationsAndTTMRatios map[string][]float64
	}

	CustomCalculations              []map[string]*float64
	CustomCalculationsGrowth        []map[string]float64
	MeanSTDCustomCalculations       map[string][]float64
	MeanSTDCustomCalculationsGrowth map[string][]float64

	CustomCalculationsAsReported              []map[string]*float64
	CustomCalculationsAsReportedGrowth        []map[string]float64
	MeanSTDCustomCalculationsAsReported       []map[string][]float64
	MeanSTDCustomCalculationsAsReportedGrowth []map[string][]float64

	MeanZippedSTDCustomCalculationsAndAsReported       []map[string][]float64
	MeanZippedSTDCustomCalculationsAndAsReportedGrowth map[string][]float64
}
