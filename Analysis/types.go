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
	GrowthBalanceSheetStatementAsReported                              []fundamentals.GrowthBalanceSheetStatementAsReported
	DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported []fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported

	IncomeStatement                                        []objects.IncomeStatement
	IncomeStatementGrowth                                  []objects.IncomeStatementGrowth
	IncomeStatementAsReported                              []objects.IncomeStatementAsReported
	GrowthIncomeStatementAsReported                        []fundamentals.GrowthIncomeStatementAsReported
	DiscrepancyIncomeStatementAndIncomeStatementAsReported []fundamentals.DiscrepancyIncomeStatementAndIncomeStatementAsReported

	CashFlowStatement                                          []objects.CashFlowStatement
	CashFlowStatementGrowth                                    []objects.CashFlowStatementGrowth
	CashFlowStatementAsReported                                []objects.CashFlowStatementAsReported
	CashFlowStatementAsReportedGrowth                          []fundamentals.CashFlowStatementAsReportedGrowth
	DiscrepancyCashFlowStatementAndCashFlowStatementAsReported []fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported

	FinancialRatios          []objects.FinancialRatios
	FinancialRatiosTTM       []objects.FinancialRatiosTTM
	FinancialRatiosGrowth    []fundamentals.FinancialRatiosGrowth
	FinancialRatiosTTMGrowth []fundamentals.FinancialRatiosTTMGrowth

	FullFinancialStatement       []objects.FullFinancialStatementAsReported
	FullFinancialStatementGrowth []objects.FinancialStatementsGrowth
}

type FundamentalsCalculationsResults struct {
	Symbol           string
	Fundamentals     CompanyFundamentals
	PeriodLength     objects.CompanyValuationPeriod
	Outlook          CompanyOutlook
	NumEmployees     float64
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

		MeanSTDBalanceSheetStatement                          map[string][]interface{}
		MeanSTDBalanceSheetStatementAsReported                map[string][]interface{}
		MeanSTDBalanceSheetStatementGrowth                    map[string][]interface{}
		MeanSTDBalanceSheetStatementAsReportedGrowth          map[string][]interface{}
		MeanSTDBalanceSheetDiscrepancies                      map[string][]interface{}
		MeanZippedSTDBalanceSheetStatementAndAsReported       map[string][]interface{}
		MeanZippedSTDBalanceSheetStatementAndAsReportedGrowth map[string][]interface{}
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

		MeanSTDIncomeStatement                          map[string][]interface{}
		MeanSTDIncomeStatementAsReported                map[string][]interface{}
		MeanSTDIncomeStatementGrowth                    map[string][]interface{}
		MeanSTDIncomeStatementAsReportedGrowth          map[string][]interface{}
		MeanSTDIncomeStatementDiscrepancies             map[string][]interface{}
		MeanZippedSTDIncomeStatementAndAsReported       map[string][]interface{}
		MeanZippedSTDIncomeStatementAndAsReportedGrowth map[string][]interface{}
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

		MeanSTDCashFlowStatement                          map[string][]interface{}
		MeanSTDCashFlowStatementAsReported                map[string][]interface{}
		MeanSTDCashFlowStatementGrowth                    map[string][]interface{}
		MeanSTDCashFlowStatementAsReportedGrowth          map[string][]interface{}
		MeanSTDCashFlowStatementDiscrepancies             map[string][]interface{}
		MeanZippedSTDCashFlowStatementAndAsReported       map[string][]interface{}
		MeanZippedSTDCashFlowStatementAndAsReportedGrowth map[string][]interface{}
	}

	FinancialRatios struct {
		FPMRatios          []objects.FinancialRatios
		FPMRatiosTTM       []objects.FinancialRatiosTTM
		FPMRatiosGrowth    []fundamentals.FinancialRatiosGrowth
		FPMRatiosTTMGrowth []fundamentals.FinancialRatiosTTMGrowth

		AverageSTDFPMRatios                     map[string][]interface{}
		AverageSTDFPMRatiosTTM                  map[string][]interface{}
		AverageSTDFPMRatiosGrowth               map[string][]interface{}
		AverageSTDFPMRatiosTTMGrowth            map[string][]interface{}
		AverageSTDFZippedFPMRationsAndTTMRatios map[string][]interface{}
	}

	CustomCalculations              []map[string]*float64
	CustomCalculationsGrowth        map[string][]float64
	MeanSTDCustomCalculations       map[string][]interface{}
	MeanSTDCustomCalculationsGrowth map[string][]interface{}

	CustomCalculationsAsReported              []map[string]*float64
	CustomCalculationsAsReportedGrowth        map[string][]float64
	MeanSTDCustomCalculationsAsReported       map[string][]interface{}
	MeanSTDCustomCalculationsAsReportedGrowth map[string][]interface{}

	MeanZippedSTDCustomCalculationsAndAsReported       map[string][]interface{}
	MeanZippedSTDCustomCalculationsAndAsReportedGrowth map[string][]interface{}
}

type CompanyOutlook struct {
	Beta                       float64
	AvgVolume                  float64
	MarketCap                  float64
	LastDividend               float64
	Changes                    float64
	FMPDcfDiff                 float64
	FMPDcf                     float64
	StockPrice                 float64
	TotalExecutivePay          float64
	InsiderTrades              []map[string]*float64
	InsiderTradesGrowth        map[string][]float64
	InsiderTradesMeanSTD       map[string][]interface{}
	InsiderTradesGrowthMeanSTD map[string][]interface{}
	Splits                     []float64
	SplitMeanSTD               map[string][]interface{}
	DividendHist               []map[string]*float64
	DividendHistGrowth         map[string][]float64
	DividendHistMeanSTD        map[string][]interface{}
	DividendHistGrowthMeanSTD  map[string][]interface{}
	DividendYieldTTM           float64
	CurrentVolume              float64
	YearHigh                   float64
	YearLow                    float64
	Ratios                     []map[string]*float64
	RatiosGrowth               map[string][]float64
	RatiosMeanSTD              map[string][]interface{}
	RatiosGrowthMeanSTD        map[string][]interface{}
}

type FinalNumbers struct {
	CalculationsOutlookFundamentals FundamentalsCalculationsResults
	FMPDCF                          []objects.DiscountedCashFlow
	FMPDCFMeanSTD                   map[string][]interface{}
	FMPMeanSTDDCF                   map[string][]interface{}
	EmployeeCount                   float64
	FMPRatings                      []map[string]*float64
	FMPRatingsGrowth                map[string][]float64
	FMPRatingsMeanSTD               map[string][]interface{}
	FMPRatingsGrowthMeanSTD         map[string][]interface{}
}
