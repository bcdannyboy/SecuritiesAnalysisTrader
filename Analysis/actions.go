package Analysis

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Calculations"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
)

func PerformFundamentalsCalculations(Fundamentals CompanyFundamentals, Period string, RiskFreeRate float64, MarketReturn float64, Outlook CompanyOutlook, NumEmployees *float64, DefaultEffectiveTaxRate float64) FundamentalsCalculationsResults {
	TotEmployees := float64(0)
	if NumEmployees != nil {
		TotEmployees = *NumEmployees
	}
	CalculationResults := FundamentalsCalculationsResults{
		Symbol:       Fundamentals.Symbol,
		Fundamentals: Fundamentals,
		Outlook:      Outlook,
		NumEmployees: TotEmployees,
	}

	Beta := Outlook.Beta

	EffectiveTaxRate := DefaultEffectiveTaxRate
	if len(Fundamentals.FinancialRatiosTTM) > 0 {
		EffectiveTaxRate = Fundamentals.FinancialRatiosTTM[len(Fundamentals.FinancialRatiosTTM)-1].EffectiveTaxRateTTM
	}

	PricePerShare := Outlook.StockPrice

	CalculationResults.BalanceSheet.DifferenceInLengthBetweenBalanceSheetStatementAndBalanceSheetStatementAsReported = len(Fundamentals.BalanceSheetStatements) - len(Fundamentals.BalanceSheetStatementAsReported)
	CalculationResults.IncomeStatement.DifferenceInLengthBetweenIncomeStatementAndIncomeStatementAsReported = len(Fundamentals.IncomeStatement) - len(Fundamentals.IncomeStatementAsReported)
	CalculationResults.CashFlowStatement.DifferenceInLengthBetweenCashFlowStatementAndCashFlowStatementAsReported = len(Fundamentals.CashFlowStatement) - len(Fundamentals.CashFlowStatementAsReported)

	BalanceSheetStatementReportDates := []string{}
	BalanceSheetAsReportedReportDates := []string{}

	IncomeStatementReportDates := []string{}
	IncomeStatementAsReportedReportDates := []string{}

	CashFlowReportDates := []string{}
	CashFlowAsReportedReportDates := []string{}

	for _, balance_sheet_stmt := range Fundamentals.BalanceSheetStatements {
		BalanceSheetStatementReportDates = append(BalanceSheetStatementReportDates, balance_sheet_stmt.Date)
	}

	for _, balance_sheet_stmt_as_reported := range Fundamentals.BalanceSheetStatementAsReported {
		BalanceSheetAsReportedReportDates = append(BalanceSheetAsReportedReportDates, balance_sheet_stmt_as_reported.Date)
	}

	for _, income_stmt := range Fundamentals.IncomeStatement {
		IncomeStatementReportDates = append(IncomeStatementReportDates, income_stmt.Date)
	}

	for _, income_stmt_as_reported := range Fundamentals.IncomeStatementAsReported {
		IncomeStatementAsReportedReportDates = append(IncomeStatementAsReportedReportDates, income_stmt_as_reported.Date)
	}

	for _, cash_flow_stmt := range Fundamentals.CashFlowStatement {
		CashFlowReportDates = append(CashFlowReportDates, cash_flow_stmt.Date)
	}

	for _, cash_flow_stmt_as_reported := range Fundamentals.CashFlowStatementAsReported {
		CashFlowAsReportedReportDates = append(CashFlowAsReportedReportDates, cash_flow_stmt_as_reported.Date)
	}

	for _, income_stmt := range Fundamentals.IncomeStatement {
		IncomeStatementReportDates = append(IncomeStatementReportDates, income_stmt.Date)
	}

	BalanceSheetStatementMissingPeriods, BalanceSheetStatementConsecutivePeriods, BalanceSheetStatementGapPeriods := 0, 0, 0
	BalanceSheetStatementAsReportedMissingPeriods, BalanceSheetStatementAsReportedConsecutivePeriods, BalanceSheetStatementAsReportedGapPeriods := 0, 0, 0
	IncomeStatementMissingPeriods, IncomeStatementConsecutivePeriods, IncomeStatementGapPeriods := 0, 0, 0
	IncomeStatementAsReportedMissingPeriods, IncomeStatementAsReportedConsecutivePeriods, IncomeStatementAsReportedGapPeriods := 0, 0, 0
	CashFlowStatementMissingPeriods, CashFlowStatementConsecutivePeriods, CashFlowStatementGapPeriods := 0, 0, 0
	CashFlowStatementAsReportedMissingPeriods, CashFlowStatementAsReportedConsecutivePeriods, CashFlowStatementAsReportedGapPeriods := 0, 0, 0

	if len(BalanceSheetStatementReportDates) > 0 {
		_, _, BalanceSheetStatementMissingPeriods, BalanceSheetStatementConsecutivePeriods, BalanceSheetStatementGapPeriods = Calculations.ProcessReportDates(BalanceSheetStatementReportDates, Period)
	}
	if len(BalanceSheetAsReportedReportDates) > 0 {
		_, _, BalanceSheetStatementAsReportedMissingPeriods, BalanceSheetStatementAsReportedConsecutivePeriods, BalanceSheetStatementAsReportedGapPeriods = Calculations.ProcessReportDates(BalanceSheetAsReportedReportDates, Period)
	}

	if len(IncomeStatementReportDates) > 0 {
		_, _, IncomeStatementMissingPeriods, IncomeStatementConsecutivePeriods, IncomeStatementGapPeriods = Calculations.ProcessReportDates(IncomeStatementReportDates, Period)
	}
	if len(IncomeStatementAsReportedReportDates) > 0 {
		_, _, IncomeStatementAsReportedMissingPeriods, IncomeStatementAsReportedConsecutivePeriods, IncomeStatementAsReportedGapPeriods = Calculations.ProcessReportDates(IncomeStatementAsReportedReportDates, Period)
	}

	if len(CashFlowReportDates) > 0 {
		_, _, CashFlowStatementMissingPeriods, CashFlowStatementConsecutivePeriods, CashFlowStatementGapPeriods = Calculations.ProcessReportDates(CashFlowReportDates, Period)
	}
	if len(CashFlowAsReportedReportDates) > 0 {
		_, _, CashFlowStatementAsReportedMissingPeriods, CashFlowStatementAsReportedConsecutivePeriods, CashFlowStatementAsReportedGapPeriods = Calculations.ProcessReportDates(CashFlowAsReportedReportDates, Period)
	}

	CalculationResults.BalanceSheet.TotalGapsInBalanceSheetStatementPeriods = BalanceSheetStatementGapPeriods
	CalculationResults.BalanceSheet.TotalConsecutivePeriodsWithNoGapsInBalanceSheetStatement = BalanceSheetStatementConsecutivePeriods
	CalculationResults.BalanceSheet.TotalConsecutiveMissingPeriodsInBalanceSheetStatement = BalanceSheetStatementMissingPeriods
	CalculationResults.BalanceSheet.TotalGapsInBalanceSheetStatementAsReportedPeriods = BalanceSheetStatementAsReportedGapPeriods
	CalculationResults.BalanceSheet.TotalConsecutivePeriodsWithNoGapsInBalanceSheetStatementAsReported = BalanceSheetStatementAsReportedConsecutivePeriods
	CalculationResults.BalanceSheet.TotalConsecutiveMissingPeriodsInBalanceSheetStatementAsReported = BalanceSheetStatementAsReportedMissingPeriods

	CalculationResults.IncomeStatement.TotalGapsInIncomeStatementPeriods = IncomeStatementGapPeriods
	CalculationResults.IncomeStatement.TotalConsecutivePeriodsWithNoGapsInIncomeStatement = IncomeStatementConsecutivePeriods
	CalculationResults.IncomeStatement.TotalConsecutiveMissingPeriodsInIncomeStatement = IncomeStatementMissingPeriods
	CalculationResults.IncomeStatement.TotalGapsInIncomeStatementAsReportedPeriods = IncomeStatementAsReportedGapPeriods
	CalculationResults.IncomeStatement.TotalConsecutivePeriodsWithNoGapsInIncomeStatementAsReported = IncomeStatementAsReportedConsecutivePeriods
	CalculationResults.IncomeStatement.TotalConsecutiveMissingPeriodsInIncomeStatementAsReported = IncomeStatementAsReportedMissingPeriods

	CalculationResults.CashFlowStatement.TotalGapsInCashFlowStatementPeriods = CashFlowStatementGapPeriods
	CalculationResults.CashFlowStatement.TotalConsecutivePeriodsWithNoGapsInCashFlowStatement = CashFlowStatementConsecutivePeriods
	CalculationResults.CashFlowStatement.TotalConsecutiveMissingPeriodsInCashFlowStatement = CashFlowStatementMissingPeriods
	CalculationResults.CashFlowStatement.TotalGapsInCashFlowStatementAsReportedPeriods = CashFlowStatementAsReportedGapPeriods
	CalculationResults.CashFlowStatement.TotalConsecutivePeriodsWithNoGapsInCashFlowStatementAsReported = CashFlowStatementAsReportedConsecutivePeriods
	CalculationResults.CashFlowStatement.TotalConsecutiveMissingPeriodsInCashFlowStatementAsReported = CashFlowStatementAsReportedMissingPeriods

	CalculationResults.FinancialRatios.FPMRatios = Fundamentals.FinancialRatios
	CalculationResults.FinancialRatios.FPMRatiosTTM = Fundamentals.FinancialRatiosTTM
	CalculationResults.FinancialRatios.FPMRatiosGrowth = Fundamentals.FinancialRatiosGrowth
	CalculationResults.FinancialRatios.FPMRatiosTTMGrowth = Fundamentals.FinancialRatiosTTMGrowth

	AverageSTDFPMRatios, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.FinancialRatios})
	if err != nil {
		print("Failed to calculate mean and standard deviation for financial ratios: %s\n", err.Error())
	} else {
		CalculationResults.FinancialRatios.AverageSTDFPMRatios = AverageSTDFPMRatios
	}

	AverageSTDFPMRatiosTTM, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.FinancialRatiosTTM})
	if err != nil {
		fmt.Printf("Failed to calculate mean and standard deviation for financial ratios TTM: %s, Data: %+v\n", err.Error(), Fundamentals.FinancialRatiosTTM)
	} else {
		CalculationResults.FinancialRatios.AverageSTDFPMRatiosTTM = AverageSTDFPMRatiosTTM
	}

	AverageSTDFPMRatiosGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.FinancialRatiosGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for financial ratios growth: %s\n", err.Error())
	} else {
		CalculationResults.FinancialRatios.AverageSTDFPMRatiosGrowth = AverageSTDFPMRatiosGrowth
	}

	AverageSTDFPMRatiosTTMGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.FinancialRatiosTTMGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for financial ratios TTM growth: %s\n", err.Error())
	} else {
		CalculationResults.FinancialRatios.AverageSTDFPMRatiosTTMGrowth = AverageSTDFPMRatiosTTMGrowth
	}

	AverageSTDFZippedFPMRationsAndTTMRatios, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.FinancialRatios, Fundamentals.FinancialRatiosTTM})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped financial ratios and TTM ratios: %s\n", err.Error())
	} else {
		CalculationResults.FinancialRatios.AverageSTDFZippedFPMRationsAndTTMRatios = AverageSTDFZippedFPMRationsAndTTMRatios
	}

	MeanSTDBalanceSheetStatement, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.BalanceSheetStatements})
	if err != nil {
		print("Failed to calculate mean and standard deviation for balance sheet statement: %s\n", err.Error())
	} else {
		CalculationResults.BalanceSheet.MeanSTDBalanceSheetStatement = MeanSTDBalanceSheetStatement
	}

	MeanSTDBalanceSheetStatementAsReported, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.BalanceSheetStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for balance sheet statement as reported: %s\n", err.Error())
	} else {
		CalculationResults.BalanceSheet.MeanSTDBalanceSheetStatementAsReported = MeanSTDBalanceSheetStatementAsReported
	}

	MeanSTDBalanceSheetStatementGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.BalanceSheetStatementGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for balance sheet statement growth: %s\n", err.Error())
	} else {
		CalculationResults.BalanceSheet.MeanSTDBalanceSheetStatementGrowth = MeanSTDBalanceSheetStatementGrowth
	}

	MeanSTDBalanceSheetStatementAsReportedGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.GrowthBalanceSheetStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for balance sheet statement as reported growth: %s\n", err.Error())
	} else {
		CalculationResults.BalanceSheet.MeanSTDBalanceSheetStatementAsReportedGrowth = MeanSTDBalanceSheetStatementAsReportedGrowth
	}

	MeanSTDBalanceSheetDiscrepancies, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for balance sheet statement discrepancies: %s\n", err.Error())
	} else {
		CalculationResults.BalanceSheet.MeanSTDBalanceSheetDiscrepancies = MeanSTDBalanceSheetDiscrepancies
	}

	MeanZippedSTDBalanceSheetStatementAndAsReported, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.BalanceSheetStatements, Fundamentals.BalanceSheetStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped balance sheet statement and as reported: %s\n", err.Error())
	} else {
		CalculationResults.BalanceSheet.MeanZippedSTDBalanceSheetStatementAndAsReported = MeanZippedSTDBalanceSheetStatementAndAsReported
	}

	MeanZippedSTDBalanceSheetStatementAndAsReportedGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.BalanceSheetStatementGrowth, Fundamentals.GrowthBalanceSheetStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped balance sheet statement growth and as reported growth: %s\n", err.Error())
	} else {
		CalculationResults.BalanceSheet.MeanZippedSTDBalanceSheetStatementAndAsReportedGrowth = MeanZippedSTDBalanceSheetStatementAndAsReportedGrowth
	}

	MeanSTDIncomeStatement, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.IncomeStatement})
	if err != nil {
		print("Failed to calculate mean and standard deviation for income statement: %s\n", err.Error())
	} else {
		CalculationResults.IncomeStatement.MeanSTDIncomeStatement = MeanSTDIncomeStatement
	}

	MeanSTDIncomeStatementAsReported, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.IncomeStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for income statement as reported: %s\n", err.Error())
	} else {
		CalculationResults.IncomeStatement.MeanSTDIncomeStatementAsReported = MeanSTDIncomeStatementAsReported
	}

	MeanSTDIncomeStatementGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.IncomeStatementGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for income statement growth: %s\n", err.Error())
	} else {
		CalculationResults.IncomeStatement.MeanSTDIncomeStatementGrowth = MeanSTDIncomeStatementGrowth
	}

	MeanSTDIncomeStatementAsReportedGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.GrowthIncomeStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for income statement as reported growth: %s\n", err.Error())
	} else {
		CalculationResults.IncomeStatement.MeanSTDIncomeStatementAsReportedGrowth = MeanSTDIncomeStatementAsReportedGrowth
	}

	MeanSTDIncomeStatementDiscrepancies, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.DiscrepancyIncomeStatementAndIncomeStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for income statement discrepancies: %s\n", err.Error())
	} else {
		CalculationResults.IncomeStatement.MeanSTDIncomeStatementDiscrepancies = MeanSTDIncomeStatementDiscrepancies
	}

	MeanZippedSTDIncomeStatementAndAsReported, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.IncomeStatement, Fundamentals.IncomeStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped income statement and as reported: %s\n", err.Error())
	} else {
		CalculationResults.IncomeStatement.MeanZippedSTDIncomeStatementAndAsReported = MeanZippedSTDIncomeStatementAndAsReported
	}

	MeanZippedSTDIncomeStatementAndAsReportedGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.IncomeStatementGrowth, Fundamentals.GrowthIncomeStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped income statement growth and as reported growth: %s\n", err.Error())
	} else {
		CalculationResults.IncomeStatement.MeanZippedSTDIncomeStatementAndAsReportedGrowth = MeanZippedSTDIncomeStatementAndAsReportedGrowth
	}

	MeanSTDCashFlowStatement, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.CashFlowStatement})
	if err != nil {
		print("Failed to calculate mean and standard deviation for cash flow statement: %s\n", err.Error())
	} else {
		CalculationResults.CashFlowStatement.MeanSTDCashFlowStatement = MeanSTDCashFlowStatement
	}

	MeanSTDCashFlowStatementAsReported, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.CashFlowStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for cash flow statement as reported: %s\n", err.Error())
	} else {
		CalculationResults.CashFlowStatement.MeanSTDCashFlowStatementAsReported = MeanSTDCashFlowStatementAsReported
	}

	MeanSTDCashFlowStatementGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.CashFlowStatementGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for cash flow statement growth: %s\n", err.Error())
	} else {
		CalculationResults.CashFlowStatement.MeanSTDCashFlowStatementGrowth = MeanSTDCashFlowStatementGrowth
	}

	MeanSTDCashFlowStatementAsReportedGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.CashFlowStatementAsReportedGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for cash flow statement as reported growth: %s\n", err.Error())
	} else {
		CalculationResults.CashFlowStatement.MeanSTDCashFlowStatementAsReportedGrowth = MeanSTDCashFlowStatementAsReportedGrowth
	}

	MeanSTDCashFlowStatementDiscrepancies, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for cash flow statement discrepancies: %s\n", err.Error())
	} else {
		CalculationResults.CashFlowStatement.MeanSTDCashFlowStatementDiscrepancies = MeanSTDCashFlowStatementDiscrepancies
	}

	MeanZippedSTDCashFlowStatementAndAsReported, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.CashFlowStatement, Fundamentals.CashFlowStatementAsReported})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped cash flow statement and as reported: %s\n", err.Error())
	} else {
		CalculationResults.CashFlowStatement.MeanZippedSTDCashFlowStatementAndAsReported = MeanZippedSTDCashFlowStatementAndAsReported
	}

	MeanZippedSTDCashFlowStatementAndAsReportedGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{Fundamentals.CashFlowStatementGrowth, Fundamentals.CashFlowStatementAsReportedGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped cash flow statement growth and as reported growth: %s\n", err.Error())
	} else {
		CalculationResults.CashFlowStatement.MeanZippedSTDCashFlowStatementAndAsReportedGrowth = MeanZippedSTDCashFlowStatementAndAsReportedGrowth
	}

	CostOfEquity := RiskFreeRate + (Beta * (MarketReturn - RiskFreeRate))
	cvp := objects.CompanyValuationPeriod(Period)
	CustomCalculationResults, CustomCalculationAsReportedResults := PerformCustomCalculations(Fundamentals, cvp, PricePerShare, EffectiveTaxRate, TotEmployees, CostOfEquity)

	CalculationResults.PeriodLength = cvp
	CalculationResults.CostOfEquity = CostOfEquity
	CalculationResults.Beta = Beta
	CalculationResults.EffectiveTaxRate = EffectiveTaxRate
	CalculationResults.CustomCalculations = CustomCalculationResults
	CalculationResults.CustomCalculationsAsReported = CustomCalculationAsReportedResults

	CalculationResults.CustomCalculationsGrowth = Calculations.CalculateGrowthF64P(CustomCalculationResults)
	CalculationResults.CustomCalculationsAsReportedGrowth = Calculations.CalculateGrowthF64P(CustomCalculationAsReportedResults)

	MeanSTDCustomCalculations, err := Calculations.CalculateMeanSTDObjs([]interface{}{CustomCalculationResults})
	if err != nil {
		print("Failed to calculate mean and standard deviation for custom calculations: %s\n", err.Error())
	} else {
		CalculationResults.MeanSTDCustomCalculations = MeanSTDCustomCalculations
	}

	MeanSTDCustomCalculationsGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{CalculationResults.CustomCalculationsGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for custom calculations growth: %s\n", err.Error())
	} else {
		CalculationResults.MeanSTDCustomCalculationsGrowth = MeanSTDCustomCalculationsGrowth
	}

	MeanZippedSTDCustomCalculationsAndAsReported, err := Calculations.CalculateMeanSTDObjs([]interface{}{CustomCalculationResults, CustomCalculationAsReportedResults})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped custom calculations and as reported: %s\n", err.Error())
	} else {
		CalculationResults.MeanZippedSTDCustomCalculationsAndAsReported = MeanZippedSTDCustomCalculationsAndAsReported
	}

	MeanZippedSTDCustomCalculationsAndAsReportedGrowth, err := Calculations.CalculateMeanSTDObjs([]interface{}{CalculationResults.CustomCalculationsGrowth, CalculationResults.CustomCalculationsAsReportedGrowth})
	if err != nil {
		print("Failed to calculate mean and standard deviation for zipped custom calculations growth and as reported growth: %s\n", err.Error())
	} else {
		CalculationResults.MeanZippedSTDCustomCalculationsAndAsReportedGrowth = MeanZippedSTDCustomCalculationsAndAsReportedGrowth
	}

	return CalculationResults
}

func PerformCustomCalculations(Fundamentals CompanyFundamentals, Period objects.CompanyValuationPeriod, PricePerShare float64, ETR float64, NumEmployees float64, CostOfEquity float64) ([]map[string]*float64, []map[string]*float64) {
	FinalCalcResults := []map[string]*float64{}
	FinalCalcResultsAsReported := []map[string]*float64{}

	LenBalanceSheetStatements := len(Fundamentals.BalanceSheetStatements)
	LenBalanceSheetStatementAsReported := len(Fundamentals.BalanceSheetStatementAsReported)
	LenIncomeStatement := len(Fundamentals.IncomeStatement)
	LenIncomeStatementAsReported := len(Fundamentals.IncomeStatementAsReported)
	LenCashFlowStatement := len(Fundamentals.CashFlowStatement)
	LenCashFlowStatementAsReported := len(Fundamentals.CashFlowStatementAsReported)

	var DaysInPeriod *float64 = nil
	if Period == "quarter" {
		DaysInPeriod = utils.InterfaceToFloat64Ptr(91)
	} else {
		DaysInPeriod = utils.InterfaceToFloat64Ptr(365)
	}

	// we need to handle if some companies report during different periods or missed periods for different reports
	LongestLength := 0
	if LenBalanceSheetStatements > LongestLength {
		LongestLength = LenBalanceSheetStatements
	}
	if LenBalanceSheetStatementAsReported > LongestLength {
		LongestLength = LenBalanceSheetStatementAsReported
	}
	if LenIncomeStatement > LongestLength {
		LongestLength = LenIncomeStatement
	}
	if LenIncomeStatementAsReported > LongestLength {
		LongestLength = LenIncomeStatementAsReported
	}
	if LenCashFlowStatement > LongestLength {
		LongestLength = LenCashFlowStatement
	}
	if LenCashFlowStatementAsReported > LongestLength {
		LongestLength = LenCashFlowStatementAsReported
	}

	for current_iteration := 0; current_iteration < LongestLength; current_iteration++ {
		FullCalcResults := map[string]*float64{}
		FullCalcResultsAsReported := map[string]*float64{}

		curBalanceSheet := objects.BalanceSheetStatement{}
		if current_iteration < LenBalanceSheetStatements-1 {
			curBalanceSheet = Fundamentals.BalanceSheetStatements[current_iteration]
		}

		curBalanceSheetAsReported := objects.BalanceSheetStatementAsReported{}
		if current_iteration < LenBalanceSheetStatementAsReported-1 {
			curBalanceSheetAsReported = Fundamentals.BalanceSheetStatementAsReported[current_iteration]
		}

		curIncomeStatement := objects.IncomeStatement{}
		if current_iteration < LenIncomeStatement-1 {
			curIncomeStatement = Fundamentals.IncomeStatement[current_iteration]
		}

		curIncomeStatementAsReported := objects.IncomeStatementAsReported{}
		if current_iteration < LenIncomeStatementAsReported-1 {
			curIncomeStatementAsReported = Fundamentals.IncomeStatementAsReported[current_iteration]
		}

		curCashFlowStatement := objects.CashFlowStatement{}
		if current_iteration < LenCashFlowStatement-1 {
			curCashFlowStatement = Fundamentals.CashFlowStatement[current_iteration]
		}

		curCashFlowStatementAsReported := objects.CashFlowStatementAsReported{}
		if current_iteration < LenCashFlowStatementAsReported-1 {
			curCashFlowStatementAsReported = Fundamentals.CashFlowStatementAsReported[current_iteration]
		}

		gotCurBalanceSheet := false
		gotCurBalanceSheetAsReported := false
		gotCurIncomeStatement := false
		gotCurIncomeStatementAsReported := false
		gotCurCashFlowStatement := false
		gotCurCashFlowStatementAsReported := false

		if curBalanceSheet.Date != "" {
			gotCurBalanceSheet = true
		}
		if curBalanceSheetAsReported.Date != "" {
			gotCurBalanceSheetAsReported = true
		}
		if curIncomeStatement.Date != "" {
			gotCurIncomeStatement = true
		}
		if curIncomeStatementAsReported.Date != "" {
			gotCurIncomeStatementAsReported = true
		}
		if curCashFlowStatement.Date != "" {
			gotCurCashFlowStatement = true
		}
		if curCashFlowStatementAsReported.Date != "" {
			gotCurCashFlowStatementAsReported = true
		}

		if !gotCurBalanceSheet && !gotCurBalanceSheetAsReported && !gotCurIncomeStatement && !gotCurIncomeStatementAsReported && !gotCurCashFlowStatement && !gotCurCashFlowStatementAsReported {
			continue
		}

		TotalAssets := utils.InterfaceToFloat64Ptr(curBalanceSheet.TotalAssets)
		TotalLiabilities := utils.InterfaceToFloat64Ptr(curBalanceSheet.TotalLiabilities)
		Inventory := utils.InterfaceToFloat64Ptr(curBalanceSheet.Inventory)
		IntangibleAssets := utils.InterfaceToFloat64Ptr(curBalanceSheet.IntangibleAssets)
		NetDebt := utils.InterfaceToFloat64Ptr(curBalanceSheet.NetDebt)
		CashAndCashEquivalents := utils.InterfaceToFloat64Ptr(curBalanceSheet.CashAndCashEquivalents)
		NetReceivables := utils.InterfaceToFloat64Ptr(curBalanceSheet.NetReceivables)
		NetFixedAssets := utils.InterfaceToFloat64Ptr(curBalanceSheet.PropertyPlantEquipmentNet)
		DeferredTaxLiabilities := utils.InterfaceToFloat64Ptr(curBalanceSheet.DeferredTaxLiabilitiesNonCurrent)
		ShareholderEquity := utils.InterfaceToFloat64Ptr(curBalanceSheet.TotalStockholdersEquity)
		AccountsPayable := utils.InterfaceToFloat64Ptr(curBalanceSheet.AccountPayables)
		CommonStock := utils.InterfaceToFloat64Ptr(curBalanceSheet.CommonStock)
		ShortTermInvestments := utils.InterfaceToFloat64Ptr(curBalanceSheet.ShortTermInvestments)
		LongTermInvestments := utils.InterfaceToFloat64Ptr(curBalanceSheet.LongTermInvestments)
		TotalInvestments := utils.InterfaceToFloat64Ptr(curBalanceSheet.TotalInvestments)
		ShortTermDebt := utils.InterfaceToFloat64Ptr(curBalanceSheet.ShortTermDebt)
		LongTermDebt := utils.InterfaceToFloat64Ptr(curBalanceSheet.LongTermDebt)

		TotalAssetsAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Assets)
		AssetsNonCurrentAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Assetsnoncurrent)
		OtherAssetsNonCurrentAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Otherassetsnoncurrent)
		TotalLiabilitiesAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Liabilities)
		InventoryAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Inventorynet)
		CurrentLongTermDebtAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Longtermdebtcurrent)
		NonCurrentLongTermDebtAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Longtermdebtnoncurrent)
		CashAndCashEquivalentsAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Cashandcashequivalentsatcarryingvalue)
		AccountsReceivableAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Accountsreceivablenetcurrent)
		NonTradeReceivablesAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Nontradereceivablescurrent)
		NetFixedAssetsAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Propertyplantandequipmentnet)
		ShareholderEquityAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Stockholdersequity)
		AccountsPayableAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Accountspayablecurrent)
		SharesOutstandingAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Commonstocksharesoutstanding)
		CurrentMarketableSecuritiesAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Marketablesecuritiescurrent)
		NonCurrentMarketableSecuritiesAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Marketablesecuritiesnoncurrent)
		ShortTermDebtAsReported := utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Othershorttermborrowings)

		NetIncome := utils.InterfaceToFloat64Ptr(curIncomeStatement.NetIncome)
		GrossProfit := utils.InterfaceToFloat64Ptr(curIncomeStatement.GrossProfit)
		NetRevenue := utils.InterfaceToFloat64Ptr(curIncomeStatement.Revenue)
		GrossProfitMargin := utils.InterfaceToFloat64Ptr(curIncomeStatement.GrossProfitRatio)
		OperatingExpenses := utils.InterfaceToFloat64Ptr(curIncomeStatement.OperatingExpenses)
		OperatingIncome := utils.InterfaceToFloat64Ptr(curIncomeStatement.OperatingIncome)
		EBITDA := utils.InterfaceToFloat64Ptr(curIncomeStatement.Ebitda)
		CostAndExpenses := utils.InterfaceToFloat64Ptr(curIncomeStatement.CostAndExpenses)
		CostOfRevenue := utils.InterfaceToFloat64Ptr(curIncomeStatement.CostOfRevenue)

		SellingGeneralAndAdministrativeExpenses := utils.InterfaceToFloat64Ptr(curIncomeStatement.SellingGeneralAndAdministrativeExpenses)
		NetIncomeAsReported := utils.InterfaceToFloat64Ptr(curIncomeStatementAsReported.Comprehensiveincomenetoftax)
		GrossProfitAsReported := utils.InterfaceToFloat64Ptr(curIncomeStatementAsReported.Grossprofit)
		NetRevenueAsReported := utils.InterfaceToFloat64Ptr(curIncomeStatementAsReported.Revenuefromcontractwithcustomerexcludingassessedtax)
		OperatingExpensesAsReported := utils.InterfaceToFloat64Ptr(curIncomeStatementAsReported.Operatingexpenses)
		CostOfGoodsSoldAsReported := utils.InterfaceToFloat64Ptr(curIncomeStatementAsReported.Costofgoodsandservicessold)
		SellingGeneralAndAdministrativeExpensesAsReported := utils.InterfaceToFloat64Ptr(curIncomeStatementAsReported.Sellinggeneralandadministrativeexpense)

		DepreciationAndAmortization := utils.InterfaceToFloat64Ptr(curCashFlowStatement.DepreciationAndAmortization)
		TotalTaxesPaid := utils.InterfaceToFloat64Ptr(curCashFlowStatement.DeferredIncomeTax)
		ChangeInWorkingCapital := utils.InterfaceToFloat64Ptr(curCashFlowStatement.ChangeInWorkingCapital)
		CapitalExpenditures := utils.InterfaceToFloat64Ptr(curCashFlowStatement.CapitalExpenditure)
		NetCashInvestingActivitiesAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatement.NetCashUsedForInvestingActivites)
		NetCashFinancingActivitiesAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatement.NetCashUsedProvidedByFinancingActivities)
		OperatingCashflow := utils.InterfaceToFloat64Ptr(curCashFlowStatement.OperatingCashFlow)
		FreeCashFlow := utils.InterfaceToFloat64Ptr(curCashFlowStatement.FreeCashFlow)
		NonCashCharges := utils.InterfaceToFloat64Ptr(curCashFlowStatement.OtherNonCashItems)
		DebtService := utils.InterfaceToFloat64Ptr(curCashFlowStatement.DebtRepayment)
		DividendsPaid := utils.InterfaceToFloat64Ptr(curCashFlowStatement.DividendsPaid)

		DepreciationAndAmortizationAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Depreciationdepletionandamortization)
		TotalInterestPaymentsAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Interestpaidnet)
		TotalTaxesPaidAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Incometaxespaidnet)
		CapitalExpendituresAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Paymentstoacquirepropertyplantandequipment)
		NetCashOperatingActivitiesAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Netcashprovidedbyusedinoperatingactivities)
		NonCashChargesAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Othernoncashincomeexpense)
		DebtServiceAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Repaymentsoflongtermdebt)
		DividendsPaidAsReported := utils.InterfaceToFloat64Ptr(curCashFlowStatementAsReported.Paymentsofdividends)

		var EffectiveTaxRate = &ETR
		var IntangibleAssetsAsReported *float64 = nil
		var NetDebtAsReported *float64 = nil
		var NetReceivablesAsReported *float64 = nil
		var SharesOutstanding *float64 = nil
		var HighQualityLiquidAssets *float64 = nil
		var HighQualityLiquidAssetsAsReported *float64 = nil
		var WorkingCapital *float64 = nil
		var WorkingCapitalAsReported *float64 = nil
		var TangibleNetWorth *float64 = nil
		var TangibleNetWorthAsReported *float64 = nil
		var BookValueOfEquity *float64 = nil
		var BookValueOfEquityAsReported *float64 = nil
		var BookValueOfDebt *float64 = nil
		var BookValueOfDebtAsReported *float64 = nil
		var EquityBookValue *float64 = nil
		var EquityBookValueAsReported *float64 = nil
		var LiabilitiesBookValue *float64 = nil
		var LiabilitiesBookValueAsReported *float64 = nil
		var TotalAccrualsToTotalAssets *float64 = nil
		var TotalAccrualsToTotalAssetsAsReported *float64 = nil
		var TotalMarketableSecuritiesAsReported *float64 = nil
		var LongTermInvestmentsAsReported *float64 = nil
		var TotalMarketableSecurities *float64 = nil
		var TotalInvestmentsAsReported *float64 = nil
		var NetProfitMarginAsReported *float64 = nil
		var NetProfitMargin *float64 = nil
		var OperatingIncomeAsReported *float64 = nil
		var OperatingCashFlowAsReported *float64 = nil
		var FreeCashFlowAsReported *float64 = nil
		var EBITDAAsReported *float64 = nil
		var TotalInterestPayments *float64 = nil
		var EBIT *float64 = nil
		var EBITAsReported *float64 = nil
		var MarketValueOfEquity *float64 = nil
		var MarketValueOfEquityAsReported *float64 = nil
		var TotalDebt *float64 = nil
		var TotalDebtAsReported *float64 = nil
		var CostOfDebt *float64 = nil
		var CostOfDebtAsReported *float64 = nil
		var UnleveredFirmValue *float64 = nil
		var UnleveredFirmValueAsReported *float64 = nil
		var TaxShieldBenefits *float64 = nil
		var TaxShieldBenefitsAsReported *float64 = nil
		var NetEffectOfDebt *float64 = nil
		var NetEffectOfDebtAsReported *float64 = nil
		var NonInterestExpenses *float64 = nil
		var NonInterestExpensesAsReported *float64 = nil
		var MarketCapitalization *float64 = nil
		var MarketCapitalizationAsReported *float64 = nil
		var EnterpriseValue *float64 = nil
		var EnterpriseValueAsReported *float64 = nil
		var DebtOutstanding *float64 = nil
		var DebtOutstandingAsReported *float64 = nil
		var AssetTurnoverRatio *float64 = nil
		var AssetTurnoverRatioAsReported *float64 = nil
		var RetentionRatio *float64 = nil
		var RetentionRatioAsReported *float64 = nil
		var ReturnOnEquity *float64 = nil
		var ReturnOnEquityAsReported *float64 = nil
		var ExplicitCosts *float64 = nil
		var ExplicitCostsAsReported *float64 = nil
		var DaysInventoryOutstanding *float64 = nil
		var DaysInventoryOutstandingAsReported *float64 = nil
		var TotalCapital *float64 = nil
		var TotalCapitalAsReported *float64 = nil
		var NetMargin *float64 = nil
		var NetMarginAsReported *float64 = nil
		var FreeCashFlowToEquity *float64 = nil
		var AdjustedPresentValue *float64 = nil
		var AdjustedPresentValueAsReported *float64 = nil
		var InterestCoverageRatio *float64 = nil
		var InterestCoverageRatioAsReported *float64 = nil
		var FixedChargeCoverageRatio *float64 = nil
		var FixedChargeCoverageRatioAsReported *float64 = nil
		var DebtServiceCoverageRatio *float64 = nil
		var DebtServiceCoverageRatioAsReported *float64 = nil
		var AssetCoverageRatio *float64 = nil
		var AssetCoverageRatioAsReported *float64 = nil
		var EBITDAToInterestCoverageRatio *float64 = nil
		var EBITDAToInterestCoverageRatioAsReported *float64 = nil
		var PreferredDividendCoverageRatio *float64 = nil
		var PreferredDividendCoverageRatioAsReported *float64 = nil
		var LiquidityCoverageRatio *float64 = nil
		var LiquidityCoverageRatioAsReported *float64 = nil
		var InventoryTurnoverRatio *float64 = nil
		var InventoryTurnoverRatioAsReported *float64 = nil
		var ReturnOnCapitalEmployed *float64 = nil
		var ReturnOnCapitalEmployedAsReported *float64 = nil
		var EfficiencyRatio *float64 = nil
		var EfficiencyRatioAsReported *float64 = nil
		var RevenuePerEmployee *float64 = nil
		var RevenuePerEmployeeAsReported *float64 = nil
		var CapitalExpenditureRatio *float64 = nil
		var CapitalExpenditureRatioAsReported *float64 = nil
		var OperatingCashFlowRatio *float64 = nil
		var OperatingCashFlowRatioAsReported *float64 = nil
		var EBITDAToEVRatio *float64 = nil
		var EBITDAToEVRatioAsReported *float64 = nil
		var TangibleNetWorthRatio *float64 = nil
		var TangibleNetWorthRatioAsReported *float64 = nil
		var DeferredTaxLiabilityToEquityRatio *float64 = nil
		var DeferredTaxLiabilityToEquityRatioAsReported *float64 = nil
		var TangibleEquityRatio *float64 = nil
		var TangibleEquityRatioAsReported *float64 = nil
		var WACC *float64 = nil
		var WACCAsReported *float64 = nil
		var FixedAssetTurnoverRatio *float64 = nil
		var FixedAssetTurnoverRatioAsReported *float64 = nil
		var PPETurnoverRatio *float64 = nil
		var PPETurnoverRatioAsReported *float64 = nil
		var InvestmentTurnoverRatio *float64 = nil
		var InvestmentTurnoverRatioAsReported *float64 = nil
		var WorkingCapitalTurnoverRatio *float64 = nil
		var WorkingCapitalTurnoverRatioAsReported *float64 = nil
		var ReturnOnAssetRatio *float64 = nil
		var ReturnOnAssetRatioAsReported *float64 = nil
		var GrossProfitMarginAsReported *float64 = nil
		var OperatingProfitMargin *float64 = nil
		var OperatingProfitMarginAsReported *float64 = nil
		var EBITDAMarginRatio *float64 = nil
		var EBITDAMarginRatioAsReported *float64 = nil
		var DividendPayoutRatio *float64 = nil
		var DividendPayoutRatioAsReported *float64 = nil
		var RetentionRate *float64 = nil
		var RetentionRateAsReported *float64 = nil
		var SustainableGrowthRate *float64 = nil
		var SustainableGrowthRateAsReported *float64 = nil
		var GrossMarginOnInventory *float64 = nil
		var GrossMarginOnInventoryAsReported *float64 = nil
		var CashFlowReturnOnEquity *float64 = nil
		var CashFlowReturnOnEquityAsReported *float64 = nil
		var OperatingMargin *float64 = nil
		var OperatingMarginAsReported *float64 = nil
		var OperatingExpenseRatio *float64 = nil
		var OperatingExpenseRatioAsReported *float64 = nil
		var CurrentRatio *float64 = nil
		var CurrentRatioAsReported *float64 = nil
		var AcidTestRatio *float64 = nil
		var AcidTestRatioAsReported *float64 = nil
		var CashRatio *float64 = nil
		var CashRatioAsReported *float64 = nil
		var DefensiveIntervalRatio *float64 = nil
		var DefensiveIntervalRatioAsReported *float64 = nil
		var DrySalesRatio *float64 = nil
		var DrySalesRatioAsReported *float64 = nil
		var PriceToBookValueRatio *float64 = nil
		var PriceToBookValueRatioAsReported *float64 = nil
		var EarningsPerShare *float64 = nil
		var EarningsPerShareAsReported *float64 = nil
		var EBITDAPerShare *float64 = nil
		var EBITDAPerShareAsReported *float64 = nil
		var BookValuePerShare *float64 = nil
		var BookValuePerShareAsReported *float64 = nil
		var NetTangibleAssetsPerShare *float64 = nil
		var NetTangibleAssetsPerShareAsReported *float64 = nil
		var MarketValueOfDebt *float64 = nil
		var MarketValueOfDebtAsReported *float64 = nil
		var MarketToBookRatio *float64 = nil
		var MarketToBookRatioAsReported *float64 = nil
		var IntangiblesRatio *float64 = nil
		var IntangiblesRatioAsReported *float64 = nil
		var PriceToSalesRatio *float64 = nil
		var PriceToSalesRatioAsReported *float64 = nil
		var PriceToBookRatio *float64 = nil
		var PriceToBookRatioAsReported *float64 = nil
		var PricetoSalesValue *float64 = nil
		var PricetoSalesValueAsReported *float64 = nil
		var OperatingCashFlowPerShare *float64 = nil
		var OperatingCashFlowPerShareAsReported *float64 = nil
		var PriceToCashFlowRatio *float64 = nil
		var PriceToCashFlowRatioAsReported *float64 = nil
		var FreeCashFlowPerShare *float64 = nil
		var FreeCashFlowPerShareAsReported *float64 = nil
		var PriceToFreeCashFlowRatio *float64 = nil
		var PriceToFreeCashFlowRatioAsReported *float64 = nil
		var PriceToCashFlowValuation *float64 = nil
		var PriceToCashFlowValuationAsReported *float64 = nil
		var PriceToFreeCashFlowValuation *float64 = nil
		var PriceToFreeCashFlowValuationAsReported *float64 = nil
		var PriceToEarningsValuation *float64 = nil
		var PriceToEarningsValuationAsReported *float64 = nil
		var LiabilitiesMarketValue *float64 = nil
		var LiabilitiesMarketValueAsReported *float64 = nil
		var TobinsQ *float64 = nil
		var TobinsQAsReported *float64 = nil
		var ReceivablesTurnoverRatio *float64 = nil
		var ReceivablesTurnoverRatioAsReported *float64 = nil
		var AverageCollectionPeriod *float64 = nil
		var AverageCollectionPeriodAsReported *float64 = nil
		var AccountsPayableTurnoverRatio *float64 = nil
		var AccountsPayableTurnoverRatioAsReported *float64 = nil
		var AverageAccountsPayablePaymentPeriod *float64 = nil
		var AverageAccountsPayablePaymentPeriodAsReported *float64 = nil
		var InventoryToWorkingCapitalRatio *float64 = nil
		var InventoryToWorkingCapitalRatioAsReported *float64 = nil
		var DaysSalesOutstanding *float64 = nil
		var DaysSalesOutstandingAsReported *float64 = nil
		var DaysPayablesOutstanding *float64 = nil
		var DaysPayablesOutstandingAsReported *float64 = nil
		var CashConversionCycle *float64 = nil
		var CashConversionCycleAsReported *float64 = nil
		var NetWorkingCapital *float64 = nil
		var NetWorkingCapitalAsReported *float64 = nil
		var NOPAT *float64 = nil
		var NOPATAsReported *float64 = nil
		var EconomicValueAdded *float64 = nil
		var EconomicValueAddedAsReported *float64 = nil
		var ReturnOnInvestedCapital *float64 = nil
		var ReturnOnInvestedCapitalAsReported *float64 = nil
		var FreeCashFlowToFirm *float64 = nil
		var FreeCashFlowToFirmAsReported *float64 = nil
		var StockDividendPerShare *float64 = nil
		var StockDividendPerShareAsReported *float64 = nil
		var LeverageRatio *float64 = nil
		var LeverageRatioAsReported *float64 = nil
		var CapitalizationRatio *float64 = nil
		var CapitalizationRatioAsReported *float64 = nil
		var LongTermDebtAsReported *float64 = nil
		var DebtToCapitalRatio *float64 = nil
		var DebtToCapitalRatioAsReported *float64 = nil
		var NetGearingRatio *float64 = nil
		var NetGearingRatioAsReported *float64 = nil
		var TotalDebtToEBITDA *float64 = nil
		var TotalDebtToEBITDAAsReported *float64 = nil
		var DebtToEquityRatio *float64 = nil
		var DebtToEquityRatioAsReported *float64 = nil
		var EquityMultiplierRatio *float64 = nil
		var EquityMultiplierRatioAsReported *float64 = nil
		var DuPontAnalysis *float64 = nil
		var DuPontAnalysisAsReported *float64 = nil

		if TotalAssetsAsReported != nil && InventoryAsReported != nil {
			IntangibleAssetsAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *InventoryAsReported)
		}

		if TotalLiabilitiesAsReported != nil && InventoryAsReported != nil {
			NetDebtAsReported = utils.InterfaceToFloat64Ptr(*CurrentLongTermDebtAsReported + *NonCurrentLongTermDebtAsReported)
		}

		if AccountsReceivableAsReported != nil && NonTradeReceivablesAsReported != nil {
			NetReceivablesAsReported = utils.InterfaceToFloat64Ptr(*AccountsReceivableAsReported + *NonTradeReceivablesAsReported)
		}

		if CommonStock != nil && PricePerShare != 0 {
			SharesOutstanding = utils.InterfaceToFloat64Ptr(*CommonStock / PricePerShare)
		}

		if CashAndCashEquivalents != nil && NetReceivables != nil {
			HighQualityLiquidAssets = utils.InterfaceToFloat64Ptr(*CashAndCashEquivalents + *NetReceivables + curBalanceSheet.CashAndShortTermInvestments - curBalanceSheet.ShortTermInvestments)
		}

		if CashAndCashEquivalentsAsReported != nil && NetReceivablesAsReported != nil {
			HighQualityLiquidAssetsAsReported = utils.InterfaceToFloat64Ptr(*CashAndCashEquivalentsAsReported + *NetReceivablesAsReported)
		}

		if TotalAssets != nil && TotalLiabilities != nil {
			WorkingCapital = utils.InterfaceToFloat64Ptr(*TotalAssets - *TotalLiabilities)
		}

		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			WorkingCapitalAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *TotalLiabilitiesAsReported)
		}

		if TotalAssets != nil && IntangibleAssets != nil && TotalLiabilities != nil {
			TangibleNetWorth = utils.InterfaceToFloat64Ptr(*TotalAssets - *IntangibleAssets - *TotalLiabilities)
		}

		if TotalAssetsAsReported != nil && IntangibleAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			TangibleNetWorthAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *IntangibleAssetsAsReported - *TotalLiabilitiesAsReported)
		}

		if ShareholderEquity != nil && Inventory != nil {
			BookValueOfEquity = utils.InterfaceToFloat64Ptr(*ShareholderEquity - *Inventory)
		}

		if ShareholderEquityAsReported != nil && InventoryAsReported != nil {
			BookValueOfEquityAsReported = utils.InterfaceToFloat64Ptr(*ShareholderEquityAsReported - *InventoryAsReported)
		}

		if TotalLiabilities != nil && Inventory != nil {
			BookValueOfDebt = utils.InterfaceToFloat64Ptr(*TotalLiabilities - *Inventory)
		}

		if TotalLiabilitiesAsReported != nil && InventoryAsReported != nil {
			BookValueOfDebtAsReported = utils.InterfaceToFloat64Ptr(*TotalLiabilitiesAsReported - *InventoryAsReported)
		}

		if BookValueOfEquity != nil && BookValueOfDebt != nil {
			EquityBookValue = utils.InterfaceToFloat64Ptr(*BookValueOfEquity - *BookValueOfDebt)
		}

		if BookValueOfEquityAsReported != nil && BookValueOfDebtAsReported != nil {
			EquityBookValueAsReported = utils.InterfaceToFloat64Ptr(*BookValueOfEquityAsReported - *BookValueOfDebtAsReported)
		}

		if BookValueOfEquity != nil && BookValueOfDebt != nil {
			LiabilitiesBookValue = utils.InterfaceToFloat64Ptr(*BookValueOfEquity - *BookValueOfDebt)
		}

		if BookValueOfEquityAsReported != nil && BookValueOfDebtAsReported != nil {
			LiabilitiesBookValueAsReported = utils.InterfaceToFloat64Ptr(*BookValueOfEquityAsReported - *BookValueOfDebtAsReported)
		}

		if TotalAssets != nil && TotalLiabilities != nil && TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			TotalAccrualsToTotalAssets = utils.InterfaceToFloat64Ptr(*TotalAssets - *TotalLiabilities)
		}

		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			TotalAccrualsToTotalAssetsAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *TotalLiabilitiesAsReported)
		}

		if CurrentMarketableSecuritiesAsReported != nil && NonCurrentMarketableSecuritiesAsReported != nil {
			TotalMarketableSecuritiesAsReported = utils.InterfaceToFloat64Ptr(*CurrentMarketableSecuritiesAsReported + *NonCurrentMarketableSecuritiesAsReported)
		}

		if AssetsNonCurrentAsReported != nil && NetFixedAssetsAsReported != nil && NonCurrentMarketableSecuritiesAsReported != nil && OtherAssetsNonCurrentAsReported != nil {
			LongTermInvestmentsAsReported = utils.InterfaceToFloat64Ptr(*AssetsNonCurrentAsReported - (*NetFixedAssetsAsReported + *NonCurrentMarketableSecuritiesAsReported + *OtherAssetsNonCurrentAsReported))
		}

		if ShortTermInvestments != nil && LongTermInvestments != nil {
			TotalMarketableSecurities = utils.InterfaceToFloat64Ptr(*ShortTermInvestments + *LongTermInvestments)
		}

		if TotalMarketableSecuritiesAsReported != nil {
			TotalInvestmentsAsReported = TotalMarketableSecuritiesAsReported
		}

		if NetRevenueAsReported != nil && NetIncomeAsReported != nil && *NetRevenueAsReported != 0 {
			NetProfitMarginAsReported = utils.InterfaceToFloat64Ptr(*NetIncomeAsReported / *NetRevenueAsReported)
		}

		if NetRevenue != nil && NetIncome != nil && *NetRevenue != 0 {
			NetProfitMargin = utils.InterfaceToFloat64Ptr(*NetIncome / *NetRevenue)
		}

		if GrossProfitAsReported != nil && OperatingExpensesAsReported != nil {
			OperatingIncomeAsReported = utils.InterfaceToFloat64Ptr(*GrossProfitAsReported - *OperatingExpensesAsReported)
		}

		if NetCashOperatingActivitiesAsReported != nil && NetCashInvestingActivitiesAsReported != nil && NetCashFinancingActivitiesAsReported != nil {
			OperatingCashFlowAsReported = utils.InterfaceToFloat64Ptr(*NetCashOperatingActivitiesAsReported + *NetCashInvestingActivitiesAsReported + *NetCashFinancingActivitiesAsReported)
		}

		if NetCashOperatingActivitiesAsReported != nil && CapitalExpendituresAsReported != nil {
			FreeCashFlowAsReported = utils.InterfaceToFloat64Ptr(*NetCashOperatingActivitiesAsReported + *CapitalExpendituresAsReported)
		}

		if OperatingIncomeAsReported != nil && DepreciationAndAmortizationAsReported != nil && TotalInterestPaymentsAsReported != nil && TotalTaxesPaidAsReported != nil {
			EBITDAAsReported = utils.InterfaceToFloat64Ptr(*OperatingIncomeAsReported + *DepreciationAndAmortizationAsReported + *TotalInterestPaymentsAsReported + *TotalTaxesPaidAsReported)
		}

		if EBITDA != nil && NetIncome != nil && TotalTaxesPaid != nil && DepreciationAndAmortization != nil {
			TotalInterestPayments = utils.InterfaceToFloat64Ptr(*EBITDA - *NetIncome - *TotalTaxesPaid - *DepreciationAndAmortization)
		}

		if EBITDA != nil && DepreciationAndAmortization != nil {
			EBIT = utils.InterfaceToFloat64Ptr(*EBITDA - *DepreciationAndAmortization)
		}

		if EBITDAAsReported != nil && DepreciationAndAmortizationAsReported != nil {
			EBITAsReported = utils.InterfaceToFloat64Ptr(*EBITDAAsReported - *DepreciationAndAmortizationAsReported)
		}

		if SharesOutstanding != nil && PricePerShare != 0 {
			MarketValueOfEquity = utils.InterfaceToFloat64Ptr(*SharesOutstanding * PricePerShare)
		}

		if SharesOutstandingAsReported != nil && PricePerShare != 0 {
			MarketValueOfEquityAsReported = utils.InterfaceToFloat64Ptr(*SharesOutstandingAsReported * PricePerShare)
		}

		if ShortTermDebt != nil && LongTermDebt != nil {
			TotalDebt = utils.InterfaceToFloat64Ptr(*ShortTermDebt + *LongTermDebt)
		}

		if ShortTermDebtAsReported != nil && CurrentLongTermDebtAsReported != nil && NonCurrentLongTermDebtAsReported != nil {
			TotalDebtAsReported = utils.InterfaceToFloat64Ptr(*ShortTermDebtAsReported + *CurrentLongTermDebtAsReported + *NonCurrentLongTermDebtAsReported)
		}

		if TotalDebt != nil && TotalInterestPayments != nil && *TotalDebt != 0 {
			CostOfDebt = utils.InterfaceToFloat64Ptr(*TotalInterestPayments / *TotalDebt)
		}

		if TotalDebtAsReported != nil && TotalInterestPaymentsAsReported != nil && *TotalDebtAsReported != 0 {
			CostOfDebtAsReported = utils.InterfaceToFloat64Ptr(*TotalInterestPaymentsAsReported / *TotalDebtAsReported)
		}

		if EBIT != nil && DepreciationAndAmortization != nil {
			UnleveredFirmValue = utils.InterfaceToFloat64Ptr((*EBIT * (1 - *EffectiveTaxRate)) + *DepreciationAndAmortization)
		}

		if EBITAsReported != nil && DepreciationAndAmortizationAsReported != nil {
			UnleveredFirmValueAsReported = utils.InterfaceToFloat64Ptr((*EBITAsReported * (1 - *EffectiveTaxRate)) + *DepreciationAndAmortizationAsReported)
		}

		if TotalInterestPayments != nil {
			TaxShieldBenefits = utils.InterfaceToFloat64Ptr(*TotalInterestPayments * *EffectiveTaxRate)
		}

		if TotalInterestPaymentsAsReported != nil {
			TaxShieldBenefitsAsReported = utils.InterfaceToFloat64Ptr(*TotalInterestPaymentsAsReported * *EffectiveTaxRate)
		}

		if TaxShieldBenefits != nil && TotalInterestPayments != nil {
			NetEffectOfDebt = utils.InterfaceToFloat64Ptr(*TaxShieldBenefits - *TotalInterestPayments)
		}

		if TaxShieldBenefitsAsReported != nil && TotalInterestPaymentsAsReported != nil {
			NetEffectOfDebtAsReported = utils.InterfaceToFloat64Ptr(*TaxShieldBenefitsAsReported - *TotalInterestPaymentsAsReported)
		}

		if OperatingExpenses != nil && TotalInterestPayments != nil {
			NonInterestExpenses = utils.InterfaceToFloat64Ptr(*OperatingExpenses - *TotalInterestPayments)
		}

		if OperatingExpensesAsReported != nil && TotalInterestPaymentsAsReported != nil {
			NonInterestExpensesAsReported = utils.InterfaceToFloat64Ptr(*OperatingExpensesAsReported - *TotalInterestPaymentsAsReported)
		}

		if SharesOutstanding != nil && PricePerShare != 0 {
			MarketCapitalization = utils.InterfaceToFloat64Ptr(*SharesOutstanding * PricePerShare)
		}

		if SharesOutstandingAsReported != nil && PricePerShare != 0 {
			MarketCapitalizationAsReported = utils.InterfaceToFloat64Ptr(*SharesOutstandingAsReported * PricePerShare)
		}

		if MarketCapitalization != nil && TotalDebt != nil && CashAndCashEquivalents != nil {
			EnterpriseValue = utils.InterfaceToFloat64Ptr(*MarketCapitalization + *TotalDebt - *CashAndCashEquivalents)
		}

		if MarketCapitalizationAsReported != nil && TotalDebtAsReported != nil && CashAndCashEquivalentsAsReported != nil {
			EnterpriseValueAsReported = utils.InterfaceToFloat64Ptr(*MarketCapitalizationAsReported + *TotalDebtAsReported - *CashAndCashEquivalentsAsReported)
		}

		if TotalDebt != nil && TotalInterestPayments != nil {
			DebtOutstanding = utils.InterfaceToFloat64Ptr(*TotalDebt - *TotalInterestPayments)
		}

		if TotalDebtAsReported != nil && TotalInterestPaymentsAsReported != nil {
			DebtOutstandingAsReported = utils.InterfaceToFloat64Ptr(*TotalDebtAsReported - *TotalInterestPaymentsAsReported)
		}

		if NetRevenue != nil && TotalAssets != nil && *TotalAssets != 0 {
			AssetTurnoverRatio = utils.InterfaceToFloat64Ptr(*NetRevenue / *TotalAssets)
		}

		if NetRevenueAsReported != nil && TotalAssetsAsReported != nil && *TotalAssetsAsReported != 0 {
			AssetTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(*NetRevenueAsReported / *TotalAssetsAsReported)
		}

		if NetIncome != nil && DividendsPaid != nil && *NetIncome != 0 {
			RetentionRatio = utils.InterfaceToFloat64Ptr((*NetIncome - *DividendsPaid) / *NetIncome)
		}

		if NetIncomeAsReported != nil && DividendsPaidAsReported != nil && *NetIncomeAsReported != 0 {
			RetentionRatioAsReported = utils.InterfaceToFloat64Ptr((*NetIncomeAsReported - *DividendsPaidAsReported) / *NetIncomeAsReported)
		}

		if NetIncome != nil && ShareholderEquity != nil && *ShareholderEquity != 0 {
			ReturnOnEquity = utils.InterfaceToFloat64Ptr(*NetIncome / *ShareholderEquity)
		}

		if NetIncomeAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 {
			ReturnOnEquityAsReported = utils.InterfaceToFloat64Ptr(*NetIncomeAsReported / *ShareholderEquityAsReported)
		}

		if CostAndExpenses != nil && CostOfRevenue != nil && OperatingExpenses != nil && SellingGeneralAndAdministrativeExpenses != nil {
			ExplicitCosts = utils.InterfaceToFloat64Ptr(*CostAndExpenses + *CostOfRevenue + *OperatingExpenses + *SellingGeneralAndAdministrativeExpenses)
		}

		if CostOfGoodsSoldAsReported != nil && OperatingExpensesAsReported != nil && SellingGeneralAndAdministrativeExpensesAsReported != nil {
			ExplicitCostsAsReported = utils.InterfaceToFloat64Ptr(*CostOfGoodsSoldAsReported + *OperatingExpensesAsReported + *SellingGeneralAndAdministrativeExpensesAsReported)
		}

		if Inventory != nil && CostOfRevenue != nil && DaysInPeriod != nil && *DaysInPeriod != 0 {
			DaysInventoryOutstanding = utils.InterfaceToFloat64Ptr((*Inventory / *CostOfRevenue) * *DaysInPeriod)
		}

		if InventoryAsReported != nil && CostOfGoodsSoldAsReported != nil && DaysInPeriod != nil && *DaysInPeriod != 0 {
			DaysInventoryOutstandingAsReported = utils.InterfaceToFloat64Ptr((*InventoryAsReported / *CostOfGoodsSoldAsReported) * *DaysInPeriod)
		}

		if LongTermDebt != nil && ShortTermDebt != nil && ShareholderEquity != nil {
			TotalCapital = utils.InterfaceToFloat64Ptr(*LongTermDebt + *ShortTermDebt + *ShareholderEquity)
		}

		if TotalDebtAsReported != nil && ShareholderEquityAsReported != nil {
			TotalCapitalAsReported = utils.InterfaceToFloat64Ptr(*TotalDebtAsReported + *ShareholderEquityAsReported)
		}

		if NetRevenue != nil && CostOfRevenue != nil && *NetRevenue != 0 {
			NetMargin = utils.InterfaceToFloat64Ptr((*NetRevenue - *CostOfRevenue) / *NetRevenue)
		}

		if NetRevenueAsReported != nil && CostOfGoodsSoldAsReported != nil && *NetRevenueAsReported != 0 {
			NetMarginAsReported = utils.InterfaceToFloat64Ptr((*NetRevenueAsReported - *CostOfGoodsSoldAsReported) / *NetRevenueAsReported)
		}

		if EBITDA != nil && DepreciationAndAmortization != nil && TotalInterestPayments != nil && TotalTaxesPaid != nil && ChangeInWorkingCapital != nil && CapitalExpenditures != nil && NetDebt != nil {
			FreeCashFlowToEquity = utils.InterfaceToFloat64Ptr(Calculations.FreeCashFlowToEquity(*EBITDA, *DepreciationAndAmortization, *TotalInterestPayments, *TotalTaxesPaid, *ChangeInWorkingCapital, *CapitalExpenditures, *NetDebt))
		}

		if UnleveredFirmValue != nil && NetEffectOfDebt != nil {
			AdjustedPresentValue = utils.InterfaceToFloat64Ptr(Calculations.AdjustedPresentValue(*UnleveredFirmValue, *NetEffectOfDebt))
		}

		if UnleveredFirmValueAsReported != nil && NetEffectOfDebtAsReported != nil {
			AdjustedPresentValueAsReported = utils.InterfaceToFloat64Ptr(Calculations.AdjustedPresentValue(*UnleveredFirmValueAsReported, *NetEffectOfDebtAsReported))
		}

		if EBIT != nil && TotalInterestPayments != nil && *TotalInterestPayments != 0 {
			InterestCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.InterestCoverageRatio(*EBIT, *TotalInterestPayments))
		}

		if EBITAsReported != nil && TotalInterestPaymentsAsReported != nil && *TotalInterestPaymentsAsReported != 0 {
			InterestCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.InterestCoverageRatio(*EBITAsReported, *TotalInterestPaymentsAsReported))
		}

		if EBIT != nil && NetFixedAssets != nil && TotalInterestPayments != nil && *TotalInterestPayments != 0 && *NetFixedAssets != 0 {
			FixedChargeCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.FixedChargeCoverageRatio(*EBIT, *NetFixedAssets, *TotalInterestPayments))
		}

		if EBITAsReported != nil && NetFixedAssetsAsReported != nil && TotalInterestPaymentsAsReported != nil && *TotalInterestPaymentsAsReported != 0 && *NetFixedAssetsAsReported != 0 {
			FixedChargeCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.FixedChargeCoverageRatio(*EBITAsReported, *NetFixedAssetsAsReported, *TotalInterestPaymentsAsReported))
		}

		if OperatingIncome != nil && DebtService != nil && *DebtService != 0 {
			DebtServiceCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.DebtServiceCoverageRatio(*OperatingIncome, *DebtService))
		}

		if OperatingIncomeAsReported != nil && DebtServiceAsReported != nil && *DebtServiceAsReported != 0 {
			DebtServiceCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DebtServiceCoverageRatio(*OperatingIncomeAsReported, *DebtServiceAsReported))
		}

		if TotalAssets != nil && ShortTermDebt != nil && TotalDebt != nil && *TotalDebt != 0 {
			AssetCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.AssetCoverageRatio(*TotalAssets, *ShortTermDebt, *TotalDebt))
		}

		if TotalAssetsAsReported != nil && ShortTermDebtAsReported != nil && TotalDebtAsReported != nil && *TotalDebtAsReported != 0 {
			AssetCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.AssetCoverageRatio(*TotalAssetsAsReported, *ShortTermDebtAsReported, *TotalDebtAsReported))
		}

		if EBITDA == nil && TotalInterestPayments != nil && *TotalInterestPayments != 0 {
			EBITDAToInterestCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToInterestCoverageRatio(*EBITDA, *TotalInterestPayments))
		}

		if EBITDAAsReported != nil && TotalInterestPaymentsAsReported != nil && *TotalInterestPaymentsAsReported != 0 {
			EBITDAToInterestCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToInterestCoverageRatio(*EBITDAAsReported, *TotalInterestPaymentsAsReported))
		}

		if NetIncome != nil && DividendsPaid != nil && *DividendsPaid != 0 {
			PreferredDividendCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.PreferredDividendCoverageRatio(*NetIncome, *DividendsPaid))
		}

		if NetIncomeAsReported != nil && DividendsPaidAsReported != nil && *DividendsPaidAsReported != 0 {
			PreferredDividendCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PreferredDividendCoverageRatio(*NetIncomeAsReported, *DividendsPaidAsReported))
		}

		if HighQualityLiquidAssets != nil && OperatingCashflow != nil && *OperatingCashflow != 0 {
			LiquidityCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.LiquidityCoverageRatio(*HighQualityLiquidAssets, *OperatingCashflow))
		}

		if HighQualityLiquidAssetsAsReported != nil && OperatingCashFlowAsReported != nil && *OperatingCashFlowAsReported != 0 {
			LiquidityCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.LiquidityCoverageRatio(*HighQualityLiquidAssetsAsReported, *OperatingCashFlowAsReported))
		}

		if CostOfRevenue != nil && Inventory != nil && *Inventory != 0 {
			InventoryTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.InventoryTurnoverRatio(*CostOfRevenue, *Inventory))
		}

		if CostOfGoodsSoldAsReported != nil && InventoryAsReported != nil && *InventoryAsReported != 0 {
			InventoryTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.InventoryTurnoverRatio(*CostOfGoodsSoldAsReported, *InventoryAsReported))
		}

		if EBIT != nil && TotalAssets != nil && TotalLiabilities != nil && *TotalAssets != 0 && *TotalLiabilities != 0 {
			ReturnOnCapitalEmployed = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnCapitalEmployed(*EBIT, *TotalAssets, *TotalLiabilities))
		}

		if EBITAsReported != nil && TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil && *TotalAssetsAsReported != 0 && *TotalLiabilitiesAsReported != 0 {
			ReturnOnCapitalEmployedAsReported = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnCapitalEmployed(*EBITAsReported, *TotalAssetsAsReported, *TotalLiabilitiesAsReported))
		}

		if NonInterestExpenses != nil && NetRevenue != nil && *NetRevenue != 0 {
			EfficiencyRatio = utils.InterfaceToFloat64Ptr(Calculations.EfficiencyRatio(*NonInterestExpenses, *NetRevenue))
		}

		if NonInterestExpensesAsReported != nil && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			EfficiencyRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EfficiencyRatio(*NonInterestExpensesAsReported, *NetRevenueAsReported))
		}

		if NetRevenue != nil && NumEmployees != 0 {
			RevenuePerEmployee = utils.InterfaceToFloat64Ptr(Calculations.RevenuePerEmployee(*NetRevenue, NumEmployees))
		}

		if NetRevenueAsReported != nil && NumEmployees != 0 {
			RevenuePerEmployeeAsReported = utils.InterfaceToFloat64Ptr(Calculations.RevenuePerEmployee(*NetRevenueAsReported, NumEmployees))
		}

		if CapitalExpenditures != nil && OperatingCashflow != nil && *OperatingCashflow != 0 {
			CapitalExpenditureRatio = utils.InterfaceToFloat64Ptr(Calculations.CapitalExpenditureRatio(*CapitalExpenditures, *OperatingCashflow))
		}

		if CapitalExpendituresAsReported != nil && OperatingCashFlowAsReported != nil && *OperatingCashFlowAsReported != 0 {
			CapitalExpenditureRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CapitalExpenditureRatio(*CapitalExpendituresAsReported, *OperatingCashFlowAsReported))
		}

		if OperatingCashflow != nil && NetRevenue != nil && *NetRevenue != 0 {
			OperatingCashFlowRatio = utils.InterfaceToFloat64Ptr(Calculations.OperatingCashFlowRatio(*OperatingCashflow, *NetRevenue))
		}

		if OperatingCashFlowAsReported != nil && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			OperatingCashFlowRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingCashFlowRatio(*OperatingCashFlowAsReported, *NetRevenueAsReported))
		}

		if EBITDA != nil && EnterpriseValue != nil && *EnterpriseValue != 0 {
			EBITDAToEVRatio = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToEVRatio(*EBITDA, *EnterpriseValue))
		}

		if EBITDAAsReported != nil && EnterpriseValueAsReported != nil && *EnterpriseValueAsReported != 0 {
			EBITDAToEVRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToEVRatio(*EBITDAAsReported, *EnterpriseValueAsReported))
		}

		if TangibleNetWorth != nil && TotalAssets != nil && *TotalAssets != 0 {
			TangibleNetWorthRatio = utils.InterfaceToFloat64Ptr(Calculations.TangibleNetWorthRatio(*TangibleNetWorth, *TotalAssets))
		}

		if TangibleNetWorthAsReported != nil && TotalAssetsAsReported != nil && *TotalAssetsAsReported != 0 {
			TangibleNetWorthRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.TangibleNetWorthRatio(*TangibleNetWorthAsReported, *TotalAssetsAsReported))
		}

		if DeferredTaxLiabilityToEquityRatio != nil && ShareholderEquity != nil && *ShareholderEquity != 0 {
			DeferredTaxLiabilityToEquityRatio = utils.InterfaceToFloat64Ptr(Calculations.DeferredTaxLiabilityToEquityRatio(*DeferredTaxLiabilityToEquityRatio, *ShareholderEquity))
		}

		if DeferredTaxLiabilityToEquityRatioAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 {
			DeferredTaxLiabilityToEquityRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DeferredTaxLiabilityToEquityRatio(*DeferredTaxLiabilityToEquityRatioAsReported, *ShareholderEquityAsReported))
		}

		if ShareholderEquity != nil && IntangibleAssets != nil && TotalAssets != nil && *TotalAssets != 0 {
			TangibleEquityRatio = utils.InterfaceToFloat64Ptr(Calculations.TangibleEquityRatio(*ShareholderEquity, *IntangibleAssets, *TotalAssets))
		}

		if ShareholderEquityAsReported != nil && IntangibleAssetsAsReported != nil && TotalAssetsAsReported != nil && *TotalAssetsAsReported != 0 {
			TangibleEquityRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.TangibleEquityRatio(*ShareholderEquityAsReported, *IntangibleAssetsAsReported, *TotalAssetsAsReported))
		}

		if MarketValueOfEquity != nil && TotalDebt != nil && CostOfDebt != nil && *MarketValueOfEquity != 0 && *TotalDebt != 0 {
			WACC = utils.InterfaceToFloat64Ptr(Calculations.WeightedAverageCostOfCapital(*MarketValueOfEquity, *TotalDebt, CostOfEquity, *CostOfDebt, *EffectiveTaxRate))
		}

		if MarketValueOfEquityAsReported != nil && TotalDebtAsReported != nil && CostOfDebtAsReported != nil && *MarketValueOfEquityAsReported != 0 && *TotalDebtAsReported != 0 {
			WACCAsReported = utils.InterfaceToFloat64Ptr(Calculations.WeightedAverageCostOfCapital(*MarketValueOfEquityAsReported, *TotalDebtAsReported, CostOfEquity, *CostOfDebtAsReported, *EffectiveTaxRate))
		}

		if NetRevenue != nil && NetFixedAssets != nil && *NetFixedAssets != 0 {
			FixedAssetTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.FixedAssetTurnoverRatio(*NetRevenue, *NetFixedAssets))
		}

		if NetRevenueAsReported != nil && NetFixedAssetsAsReported != nil && *NetFixedAssetsAsReported != 0 {
			FixedAssetTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.FixedAssetTurnoverRatio(*NetRevenueAsReported, *NetFixedAssetsAsReported))
		}

		if NetRevenue != nil && NetFixedAssets != nil && DepreciationAndAmortization != nil && *NetFixedAssets != 0 {
			PPETurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.PPETurnoverRatio(*NetRevenue, *NetFixedAssets, *DepreciationAndAmortization))
		}

		if NetRevenueAsReported != nil && NetFixedAssetsAsReported != nil && DepreciationAndAmortizationAsReported != nil && *NetFixedAssetsAsReported != 0 {
			PPETurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PPETurnoverRatio(*NetRevenueAsReported, *NetFixedAssetsAsReported, *DepreciationAndAmortizationAsReported))
		}

		if NetRevenue != nil && TotalInvestments != nil && *TotalInvestments != 0 {
			InvestmentTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.InvestmentTurnoverRatio(*NetRevenue, *TotalInvestments))
		}

		if NetRevenueAsReported != nil && TotalInvestmentsAsReported != nil && *TotalInvestmentsAsReported != 0 {
			InvestmentTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.InvestmentTurnoverRatio(*NetRevenueAsReported, *TotalInvestmentsAsReported))
		}

		if NetRevenue != nil && WorkingCapital != nil && *WorkingCapital != 0 {
			WorkingCapitalTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.WorkingCapitalTurnoverRatio(*NetRevenue, *WorkingCapital))
		}

		if NetRevenueAsReported != nil && WorkingCapitalAsReported != nil && *WorkingCapitalAsReported != 0 {
			WorkingCapitalTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.WorkingCapitalTurnoverRatio(*NetRevenueAsReported, *WorkingCapitalAsReported))
		}

		if NetIncome != nil && TotalAssets != nil && *TotalAssets != 0 {
			ReturnOnAssetRatio = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnAssetRatio(*NetIncome, *TotalAssets))
		}

		if NetIncomeAsReported != nil && TotalAssetsAsReported != nil && *TotalAssetsAsReported != 0 {
			ReturnOnAssetRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnAssetRatio(*NetIncomeAsReported, *TotalAssetsAsReported))
		}

		if GrossProfitMargin == nil {
			if GrossProfit != nil && NetRevenue != nil && *NetRevenue != 0 {
				GrossProfitMargin = utils.InterfaceToFloat64Ptr(Calculations.GrossProfitMargin(*GrossProfit, *NetRevenue))
			}
		}

		if GrossProfitAsReported != nil && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			GrossProfitMarginAsReported = utils.InterfaceToFloat64Ptr(Calculations.GrossProfitMargin(*GrossProfitAsReported, *NetRevenueAsReported))
		}

		if OperatingIncome != nil && NetRevenue != nil && *NetRevenue != 0 {
			OperatingProfitMargin = utils.InterfaceToFloat64Ptr(Calculations.OperatingProfitMargin(*OperatingIncome, *NetRevenue))
		}

		if OperatingIncomeAsReported != nil && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			OperatingProfitMarginAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingProfitMargin(*OperatingIncomeAsReported, *NetRevenueAsReported))
		}

		if EBITDA != nil && NetRevenue != nil && *NetRevenue != 0 {
			EBITDAMarginRatio = utils.InterfaceToFloat64Ptr(Calculations.EBITDAMarginRatio(*EBITDA, *NetRevenue))
		}

		if EBITDAAsReported != nil && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			EBITDAMarginRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAMarginRatio(*EBITDAAsReported, *NetRevenueAsReported))
		}

		if DividendsPaid != nil && NetIncome != nil && *NetIncome != 0 {
			DividendPayoutRatio = utils.InterfaceToFloat64Ptr(Calculations.DividendPayoutRatio(*DividendsPaid, *NetIncome))
		}

		if DividendsPaidAsReported != nil && NetIncomeAsReported != nil && *NetIncomeAsReported != 0 {
			DividendPayoutRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DividendPayoutRatio(*DividendsPaidAsReported, *NetIncomeAsReported))
		}

		if DividendsPaid != nil && NetIncome != nil && *NetIncome != 0 {
			RetentionRate = utils.InterfaceToFloat64Ptr(Calculations.RetentionRate(*DividendsPaid, *NetIncome))
		}

		if DividendsPaidAsReported != nil && NetIncomeAsReported != nil && *NetIncomeAsReported != 0 {
			RetentionRateAsReported = utils.InterfaceToFloat64Ptr(Calculations.RetentionRate(*DividendsPaidAsReported, *NetIncomeAsReported))
		}

		if RetentionRate != nil && ReturnOnEquity != nil {
			SustainableGrowthRate = utils.InterfaceToFloat64Ptr(Calculations.SustainableGrowthRate(*RetentionRate, *ReturnOnEquity))
		}

		if RetentionRateAsReported != nil && ReturnOnEquityAsReported != nil {
			SustainableGrowthRateAsReported = utils.InterfaceToFloat64Ptr(Calculations.SustainableGrowthRate(*RetentionRateAsReported, *ReturnOnEquityAsReported))
		}

		if GrossProfit != nil && Inventory != nil && *Inventory != 0 {
			GrossMarginOnInventory = utils.InterfaceToFloat64Ptr(Calculations.GrossMarginOnInventory(*GrossProfit, *Inventory))
		}

		if GrossProfitAsReported != nil && InventoryAsReported != nil && *InventoryAsReported != 0 {
			GrossMarginOnInventoryAsReported = utils.InterfaceToFloat64Ptr(Calculations.GrossMarginOnInventory(*GrossProfitAsReported, *InventoryAsReported))
		}

		if OperatingCashflow != nil && ShareholderEquity != nil && *ShareholderEquity != 0 {
			CashFlowReturnOnEquity = utils.InterfaceToFloat64Ptr(Calculations.CashFlowReturnOnEquity(*OperatingCashflow, *ShareholderEquity))
		}

		if OperatingCashFlowAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 {
			CashFlowReturnOnEquityAsReported = utils.InterfaceToFloat64Ptr(Calculations.CashFlowReturnOnEquity(*OperatingCashFlowAsReported, *ShareholderEquityAsReported))
		}

		if NetRevenue != nil && CostOfRevenue != nil && *CostOfRevenue != 0 {
			OperatingMargin = utils.InterfaceToFloat64Ptr(Calculations.OperatingMargin(*NetRevenue, *CostOfRevenue))
		}

		if NetRevenueAsReported != nil && CostOfGoodsSoldAsReported != nil && *CostOfGoodsSoldAsReported != 0 {
			OperatingMarginAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingMargin(*NetRevenueAsReported, *CostOfGoodsSoldAsReported))
		}

		if OperatingExpenses != nil && DepreciationAndAmortization != nil && OperatingIncome != nil && *OperatingIncome != 0 {
			OperatingExpenseRatio = utils.InterfaceToFloat64Ptr(Calculations.OperatingExpenseRatio(*OperatingExpenses, *DepreciationAndAmortization, *OperatingIncome))
		}

		if OperatingExpensesAsReported != nil && DepreciationAndAmortizationAsReported != nil && OperatingIncomeAsReported != nil && *OperatingIncomeAsReported != 0 {
			OperatingExpenseRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingExpenseRatio(*OperatingExpensesAsReported, *DepreciationAndAmortizationAsReported, *OperatingIncomeAsReported))
		}

		if TotalAssets != nil && TotalLiabilities != nil && *TotalLiabilities != 0 {
			CurrentRatio = utils.InterfaceToFloat64Ptr(Calculations.CurrentRatio(*TotalAssets, *TotalLiabilities))
		}

		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil && *TotalLiabilitiesAsReported != 0 {
			CurrentRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CurrentRatio(*TotalAssetsAsReported, *TotalLiabilitiesAsReported))
		}

		if TotalAssets != nil && TotalLiabilities != nil && Inventory != nil && *TotalLiabilities != 0 {
			AcidTestRatio = utils.InterfaceToFloat64Ptr(Calculations.AcidTestRatio(*TotalAssets, *Inventory, *TotalLiabilities))
		}

		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil && InventoryAsReported != nil && *TotalLiabilitiesAsReported != 0 {
			AcidTestRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.AcidTestRatio(*TotalAssetsAsReported, *InventoryAsReported, *TotalLiabilitiesAsReported))
		}

		if CashAndCashEquivalents != nil && TotalLiabilities != nil && *TotalLiabilities != 0 {
			CashRatio = utils.InterfaceToFloat64Ptr(Calculations.CashRatio(*CashAndCashEquivalents, *TotalLiabilities))
		}

		if CashAndCashEquivalentsAsReported != nil && TotalLiabilitiesAsReported != nil && *TotalLiabilitiesAsReported != 0 {
			CashRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CashRatio(*CashAndCashEquivalentsAsReported, *TotalLiabilitiesAsReported))
		}

		if CashAndCashEquivalents != nil && NetReceivables != nil && TotalMarketableSecurities != nil && OperatingExpenses != nil && NonCashCharges != nil && DaysInPeriod != nil && *OperatingExpenses != 0 && *NonCashCharges != 0 && *DaysInPeriod != 0 {
			DefensiveIntervalRatio = utils.InterfaceToFloat64Ptr(Calculations.DefensiveIntervalRatio(*CashAndCashEquivalents, *NetReceivables, *TotalMarketableSecurities, *OperatingExpenses, *NonCashCharges, *DaysInPeriod))
		}

		if CashAndCashEquivalentsAsReported != nil && NetReceivablesAsReported != nil && TotalMarketableSecuritiesAsReported != nil && OperatingExpensesAsReported != nil && NonCashChargesAsReported != nil && DaysInPeriod != nil && *OperatingExpensesAsReported != 0 && *NonCashChargesAsReported != 0 && *DaysInPeriod != 0 {
			DefensiveIntervalRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DefensiveIntervalRatio(*CashAndCashEquivalentsAsReported, *NetReceivablesAsReported, *TotalMarketableSecuritiesAsReported, *OperatingExpensesAsReported, *NonCashChargesAsReported, *DaysInPeriod))
		}

		if NetReceivables != nil && NetRevenue != nil && *NetRevenue != 0 {
			DrySalesRatio = utils.InterfaceToFloat64Ptr(Calculations.DrySalesRatio(*NetReceivables, *NetRevenue))
		}

		if NetReceivablesAsReported != nil && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			DrySalesRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DrySalesRatio(*NetReceivablesAsReported, *NetRevenueAsReported))
		}

		if MarketCapitalization != nil && BookValueOfEquity != nil && *BookValueOfEquity != 0 {
			PriceToBookValueRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookValueRatio(*MarketCapitalization, *BookValueOfEquity))
		}

		if MarketCapitalizationAsReported != nil && BookValueOfEquityAsReported != nil && *BookValueOfEquityAsReported != 0 {
			PriceToBookValueRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookValueRatio(*MarketCapitalizationAsReported, *BookValueOfEquityAsReported))
		}

		if NetIncome != nil && DividendsPaid != nil && SharesOutstanding != nil && *SharesOutstanding != 0 {
			EarningsPerShare = utils.InterfaceToFloat64Ptr(Calculations.EarningsPerShare(*NetIncome, *DividendsPaid, *SharesOutstanding))
		}

		if NetIncomeAsReported != nil && DividendsPaidAsReported != nil && SharesOutstandingAsReported != nil && *SharesOutstandingAsReported != 0 {
			EarningsPerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.EarningsPerShare(*NetIncomeAsReported, *DividendsPaidAsReported, *SharesOutstandingAsReported))
		}

		if EBITDA != nil && SharesOutstanding != nil && *SharesOutstanding != 0 {
			EBITDAPerShare = utils.InterfaceToFloat64Ptr(Calculations.EBITDAPerShare(*EBITDA, *SharesOutstanding))
		}

		if EBITDAAsReported != nil && SharesOutstandingAsReported != nil && *SharesOutstandingAsReported != 0 {
			EBITDAPerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAPerShare(*EBITDAAsReported, *SharesOutstandingAsReported))
		}

		if ShareholderEquity != nil && SharesOutstanding != nil && *SharesOutstanding != 0 {
			BookValuePerShare = utils.InterfaceToFloat64Ptr(Calculations.BookValuePerShare(*ShareholderEquity, *SharesOutstanding))
		}

		if ShareholderEquityAsReported != nil && SharesOutstandingAsReported != nil && *SharesOutstandingAsReported != 0 {
			BookValuePerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.BookValuePerShare(*ShareholderEquityAsReported, *SharesOutstandingAsReported))
		}

		if TangibleNetWorth != nil && SharesOutstanding != nil && *SharesOutstanding != 0 {
			NetTangibleAssetsPerShare = utils.InterfaceToFloat64Ptr(Calculations.NetTangibleAssetsPerShare(*TangibleNetWorth, *SharesOutstanding))
		}

		if TangibleNetWorthAsReported != nil && SharesOutstandingAsReported != nil && *SharesOutstandingAsReported != 0 {
			NetTangibleAssetsPerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.NetTangibleAssetsPerShare(*TangibleNetWorthAsReported, *SharesOutstandingAsReported))
		}

		if PricePerShare != 0 && SharesOutstanding != nil && BookValueOfDebt != nil {
			MarketValueOfDebt = utils.InterfaceToFloat64Ptr(Calculations.MarketValueOfDebt(PricePerShare, *SharesOutstanding, *BookValueOfDebt))
		}

		if PricePerShare != 0 && SharesOutstandingAsReported != nil && BookValueOfDebtAsReported != nil {
			MarketValueOfDebtAsReported = utils.InterfaceToFloat64Ptr(Calculations.MarketValueOfDebt(PricePerShare, *SharesOutstandingAsReported, *BookValueOfDebtAsReported))
		}

		if MarketCapitalization != nil && BookValueOfEquity != nil && *BookValueOfEquity != 0 {
			MarketToBookRatio = utils.InterfaceToFloat64Ptr(Calculations.MarketToBookRatio(*MarketCapitalization, *BookValueOfEquity))
		}

		if MarketCapitalizationAsReported != nil && BookValueOfEquityAsReported != nil && *BookValueOfEquityAsReported != 0 {
			MarketToBookRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.MarketToBookRatio(*MarketCapitalizationAsReported, *BookValueOfEquityAsReported))
		}

		if IntangibleAssets != nil && TotalAssets != nil && *TotalAssets != 0 {
			IntangiblesRatio = utils.InterfaceToFloat64Ptr(Calculations.IntangiblesRatio(*IntangibleAssets, *TotalAssets))
		}

		if IntangibleAssetsAsReported != nil && TotalAssetsAsReported != nil && *TotalAssetsAsReported != 0 {
			IntangiblesRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.IntangiblesRatio(*IntangibleAssetsAsReported, *TotalAssetsAsReported))
		}

		if PricePerShare != 0 && NetRevenue != nil && *NetRevenue != 0 {
			PriceToSalesRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToSalesRatio(PricePerShare, *NetRevenue))
		}

		if PricePerShare != 0 && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			PriceToSalesRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToSalesRatio(PricePerShare, *NetRevenueAsReported))
		}

		if PricePerShare != 0 && BookValueOfEquity != nil && *BookValueOfEquity != 0 {
			PriceToBookRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookRatio(PricePerShare, *BookValueOfEquity))
		}

		if PricePerShare != 0 && BookValueOfEquityAsReported != nil && *BookValueOfEquityAsReported != 0 {
			PriceToBookRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookRatio(PricePerShare, *BookValueOfEquityAsReported))
		}

		if MarketCapitalization != nil && NetRevenue != nil && *NetRevenue != 0 {
			PricetoSalesValue = utils.InterfaceToFloat64Ptr(Calculations.PricetoSalesValue(*MarketCapitalization, *NetRevenue))
		}

		if MarketCapitalizationAsReported != nil && NetRevenueAsReported != nil && *NetRevenueAsReported != 0 {
			PricetoSalesValueAsReported = utils.InterfaceToFloat64Ptr(Calculations.PricetoSalesValue(*MarketCapitalizationAsReported, *NetRevenueAsReported))
		}

		if OperatingCashflow != nil && SharesOutstanding != nil && *SharesOutstanding != 0 {
			OperatingCashFlowPerShare = utils.InterfaceToFloat64Ptr(*OperatingCashflow / *SharesOutstanding)
		}

		if OperatingCashFlowAsReported != nil && SharesOutstandingAsReported != nil && *SharesOutstandingAsReported != 0 {
			OperatingCashFlowPerShareAsReported = utils.InterfaceToFloat64Ptr(*OperatingCashFlowAsReported / *SharesOutstandingAsReported)
		}

		if PricePerShare != 0 && OperatingCashflow != nil && *OperatingCashflow != 0 {
			PriceToCashFlowRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowRatio(PricePerShare, *OperatingCashflow))
		}

		if PricePerShare != 0 && OperatingCashFlowAsReported != nil && *OperatingCashFlowAsReported != 0 {
			PriceToCashFlowRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowRatio(PricePerShare, *OperatingCashFlowAsReported))
		}

		if FreeCashFlow != nil && SharesOutstanding != nil && *SharesOutstanding != 0 {
			FreeCashFlowPerShare = utils.InterfaceToFloat64Ptr(*FreeCashFlow / *SharesOutstanding)
		}

		if FreeCashFlowAsReported != nil && SharesOutstandingAsReported != nil && *SharesOutstandingAsReported != 0 {
			FreeCashFlowPerShareAsReported = utils.InterfaceToFloat64Ptr(*FreeCashFlowAsReported / *SharesOutstandingAsReported)
		}

		if PricePerShare != 0 && FreeCashFlowPerShare != nil && *FreeCashFlowPerShare != 0 {
			PriceToFreeCashFlowRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowRatio(PricePerShare, *FreeCashFlowPerShare))
		}

		if PricePerShare != 0 && FreeCashFlowPerShareAsReported != nil && *FreeCashFlowPerShareAsReported != 0 {
			PriceToFreeCashFlowRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowRatio(PricePerShare, *FreeCashFlowPerShareAsReported))
		}

		if MarketCapitalization != nil && OperatingCashflow != nil && *OperatingCashflow != 0 {
			PriceToCashFlowValuation = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowValuation(*MarketCapitalization, *OperatingCashflow))
		}

		if MarketCapitalizationAsReported != nil && OperatingCashFlowAsReported != nil && *OperatingCashFlowAsReported != 0 {
			PriceToCashFlowValuationAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowValuation(*MarketCapitalizationAsReported, *OperatingCashFlowAsReported))
		}

		if MarketCapitalization != nil && FreeCashFlow != nil && *FreeCashFlow != 0 {
			PriceToFreeCashFlowValuation = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowValuation(*MarketCapitalization, *FreeCashFlow))
		}

		if MarketCapitalizationAsReported != nil && FreeCashFlowAsReported != nil && *FreeCashFlowAsReported != 0 {
			PriceToFreeCashFlowValuationAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowValuation(*MarketCapitalizationAsReported, *FreeCashFlowAsReported))
		}

		if MarketCapitalization != nil && NetIncome != nil && *NetIncome != 0 {
			PriceToEarningsValuation = utils.InterfaceToFloat64Ptr(Calculations.PriceToEarningsValuation(*MarketCapitalization, *NetIncome))
		}

		if MarketCapitalizationAsReported != nil && NetIncomeAsReported != nil && *NetIncomeAsReported != 0 {
			PriceToEarningsValuationAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToEarningsValuation(*MarketCapitalizationAsReported, *NetIncomeAsReported))
		}

		if PricePerShare != 0 && SharesOutstanding != nil && BookValueOfDebt != nil {
			LiabilitiesMarketValue = utils.InterfaceToFloat64Ptr(Calculations.LiabilitiesMarketValue(PricePerShare, *SharesOutstanding, *BookValueOfDebt))
		}

		if PricePerShare != 0 && SharesOutstandingAsReported != nil && BookValueOfDebtAsReported != nil {
			LiabilitiesMarketValueAsReported = utils.InterfaceToFloat64Ptr(Calculations.LiabilitiesMarketValue(PricePerShare, *SharesOutstandingAsReported, *BookValueOfDebtAsReported))
		}

		if MarketValueOfEquity != nil && MarketValueOfDebt != nil && BookValueOfEquity != nil && BookValueOfDebt != nil && (*BookValueOfEquity != 0 || *BookValueOfDebt != 0) {
			TobinsQ = utils.InterfaceToFloat64Ptr(Calculations.TobinsQ(*MarketValueOfEquity, *MarketValueOfDebt, *BookValueOfEquity, *BookValueOfDebt))
		}

		if MarketValueOfEquityAsReported != nil && MarketValueOfDebtAsReported != nil && BookValueOfEquityAsReported != nil && BookValueOfDebtAsReported != nil && (*BookValueOfEquityAsReported != 0 || *BookValueOfDebtAsReported != 0) {
			TobinsQAsReported = utils.InterfaceToFloat64Ptr(Calculations.TobinsQ(*MarketValueOfEquityAsReported, *MarketValueOfDebtAsReported, *BookValueOfEquityAsReported, *BookValueOfDebtAsReported))
		}

		if NetRevenue != nil && NetReceivables != nil && *NetReceivables != 0 {
			ReceivablesTurnoverRatio = utils.InterfaceToFloat64Ptr(*NetRevenue / *NetReceivables)
		}

		if NetRevenueAsReported != nil && NetReceivablesAsReported != nil && *NetReceivablesAsReported != 0 {
			ReceivablesTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(*NetRevenueAsReported / *NetReceivablesAsReported)
		}

		if ReceivablesTurnoverRatio != nil && *ReceivablesTurnoverRatio != 0 {
			AverageCollectionPeriod = utils.InterfaceToFloat64Ptr(Calculations.AverageCollectionPeriod(*ReceivablesTurnoverRatio))
		}

		if ReceivablesTurnoverRatioAsReported != nil && *ReceivablesTurnoverRatioAsReported != 0 {
			AverageCollectionPeriodAsReported = utils.InterfaceToFloat64Ptr(Calculations.AverageCollectionPeriod(*ReceivablesTurnoverRatioAsReported))
		}

		if AccountsPayable != nil && NetRevenue != nil && *AccountsPayable != 0 {
			AccountsPayableTurnoverRatio = utils.InterfaceToFloat64Ptr(*NetRevenue / *AccountsPayable)
		}

		if AccountsPayableAsReported != nil && NetRevenueAsReported != nil && *AccountsPayableAsReported != 0 {
			AccountsPayableTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(*NetRevenueAsReported / *AccountsPayableAsReported)
		}

		if AccountsPayableTurnoverRatio != nil && *AccountsPayableTurnoverRatio != 0 {
			AverageAccountsPayablePaymentPeriod = utils.InterfaceToFloat64Ptr(Calculations.AverageAccountsPayablePaymentPeriod(*AccountsPayableTurnoverRatio))
		}

		if AccountsPayableTurnoverRatioAsReported != nil && *AccountsPayableTurnoverRatioAsReported != 0 {
			AverageAccountsPayablePaymentPeriodAsReported = utils.InterfaceToFloat64Ptr(Calculations.AverageAccountsPayablePaymentPeriod(*AccountsPayableTurnoverRatioAsReported))
		}

		if Inventory != nil && WorkingCapital != nil && *WorkingCapital != 0 {
			InventoryToWorkingCapitalRatio = utils.InterfaceToFloat64Ptr(*Inventory / *WorkingCapital)
		}

		if InventoryAsReported != nil && WorkingCapitalAsReported != nil && *WorkingCapitalAsReported != 0 {
			InventoryToWorkingCapitalRatioAsReported = utils.InterfaceToFloat64Ptr(*InventoryAsReported / *WorkingCapitalAsReported)
		}

		if NetRevenue != nil && CostOfRevenue != nil && DaysInPeriod != nil && *CostOfRevenue != 0 {
			DaysSalesOutstanding = utils.InterfaceToFloat64Ptr((*NetRevenue / *CostOfRevenue) * *DaysInPeriod)
		}

		if NetRevenueAsReported != nil && CostOfGoodsSoldAsReported != nil && DaysInPeriod != nil && *CostOfGoodsSoldAsReported != 0 {
			DaysSalesOutstandingAsReported = utils.InterfaceToFloat64Ptr((*NetRevenueAsReported / *CostOfGoodsSoldAsReported) * *DaysInPeriod)
		}

		if AccountsPayable != nil && CostOfRevenue != nil && DaysInPeriod != nil && *CostOfRevenue != 0 {
			DaysPayablesOutstanding = utils.InterfaceToFloat64Ptr((*AccountsPayable / *CostOfRevenue) * *DaysInPeriod)
		}

		if AccountsPayableAsReported != nil && CostOfGoodsSoldAsReported != nil && DaysInPeriod != nil && *CostOfGoodsSoldAsReported != 0 {
			DaysPayablesOutstandingAsReported = utils.InterfaceToFloat64Ptr((*AccountsPayableAsReported / *CostOfGoodsSoldAsReported) * *DaysInPeriod)
		}

		if DaysInventoryOutstanding != nil && DaysSalesOutstanding != nil && DaysPayablesOutstanding != nil {
			CashConversionCycle = utils.InterfaceToFloat64Ptr(Calculations.CashConversionCycle(*DaysInventoryOutstanding, *DaysSalesOutstanding, *DaysPayablesOutstanding))
		}

		if NetReceivables != nil && Inventory != nil && AccountsPayable != nil {
			NetWorkingCapital = utils.InterfaceToFloat64Ptr(Calculations.NetWorkingCapital(*NetReceivables, *Inventory, *AccountsPayable))
		}

		if NetReceivablesAsReported != nil && InventoryAsReported != nil && AccountsPayableAsReported != nil {
			NetWorkingCapitalAsReported = utils.InterfaceToFloat64Ptr(Calculations.NetWorkingCapital(*NetReceivablesAsReported, *InventoryAsReported, *AccountsPayableAsReported))
		}

		if OperatingIncome != nil && EffectiveTaxRate != nil {
			NOPAT = utils.InterfaceToFloat64Ptr(*OperatingIncome * (1 - *EffectiveTaxRate))
		}

		if OperatingIncomeAsReported != nil && EffectiveTaxRate != nil {
			NOPATAsReported = utils.InterfaceToFloat64Ptr(*OperatingIncomeAsReported * (1 - *EffectiveTaxRate))
		}

		if NOPAT != nil && WACC != nil && TotalCapital != nil {
			EconomicValueAdded = utils.InterfaceToFloat64Ptr(Calculations.EconomicValueAdded(*NOPAT, *WACC, *TotalCapital))
		}

		if NOPATAsReported != nil && WACCAsReported != nil && TotalCapitalAsReported != nil {
			EconomicValueAddedAsReported = utils.InterfaceToFloat64Ptr(Calculations.EconomicValueAdded(*NOPATAsReported, *WACCAsReported, *TotalCapitalAsReported))
		}

		if NOPAT != nil && TotalInvestments != nil && *TotalInvestments != 0 {
			ReturnOnInvestedCapital = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnInvestedCapital(*NOPAT, *TotalInvestments))
		}

		if NOPATAsReported != nil && TotalInvestmentsAsReported != nil && *TotalInvestmentsAsReported != 0 {
			ReturnOnInvestedCapitalAsReported = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnInvestedCapital(*NOPATAsReported, *TotalInvestmentsAsReported))
		}

		if NetIncome != nil && NonCashCharges != nil && TotalInterestPayments != nil && EffectiveTaxRate != nil && LongTermInvestments != nil && NetWorkingCapital != nil {
			FreeCashFlowToFirm = utils.InterfaceToFloat64Ptr(Calculations.FreeCashFlowToFirm(*NetIncome, *NonCashCharges, *TotalInterestPayments, *EffectiveTaxRate, *LongTermInvestments, *NetWorkingCapital))
		}

		if NetIncomeAsReported != nil && NonCashChargesAsReported != nil && TotalInterestPaymentsAsReported != nil && EffectiveTaxRate != nil && LongTermInvestmentsAsReported != nil && NetWorkingCapitalAsReported != nil {
			FreeCashFlowToFirmAsReported = utils.InterfaceToFloat64Ptr(Calculations.FreeCashFlowToFirm(*NetIncomeAsReported, *NonCashChargesAsReported, *TotalInterestPaymentsAsReported, *EffectiveTaxRate, *LongTermInvestmentsAsReported, *NetWorkingCapitalAsReported))
		}

		if DividendsPaid != nil && SharesOutstanding != nil && *SharesOutstanding != 0 {
			StockDividendPerShare = utils.InterfaceToFloat64Ptr(*DividendsPaid / *SharesOutstanding)
		}

		if DividendsPaidAsReported != nil && SharesOutstandingAsReported != nil && *SharesOutstandingAsReported != 0 {
			StockDividendPerShareAsReported = utils.InterfaceToFloat64Ptr(*DividendsPaidAsReported / *SharesOutstandingAsReported)
		}

		if TotalDebt != nil && EBITDA != nil && *EBITDA != 0 {
			LeverageRatio = utils.InterfaceToFloat64Ptr(Calculations.LeverageRatio(*TotalDebt, *EBITDA))
		}

		if TotalDebtAsReported != nil && EBITDAAsReported != nil && *EBITDAAsReported != 0 {
			LeverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.LeverageRatio(*TotalDebtAsReported, *EBITDAAsReported))
		}

		if TotalDebt != nil && ShareholderEquity != nil && *ShareholderEquity != 0 {
			CapitalizationRatio = utils.InterfaceToFloat64Ptr(Calculations.CapitalizationRatio(*TotalDebt, *ShareholderEquity))
		}

		if TotalDebtAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 {
			CapitalizationRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CapitalizationRatio(*TotalDebtAsReported, *ShareholderEquityAsReported))
		}

		if CurrentLongTermDebtAsReported != nil && NonCurrentLongTermDebtAsReported != nil {
			LongTermDebtAsReported = utils.InterfaceToFloat64Ptr(*CurrentLongTermDebtAsReported + *NonCurrentLongTermDebtAsReported)
		} else if CurrentLongTermDebtAsReported != nil {
			LongTermDebtAsReported = utils.InterfaceToFloat64Ptr(*CurrentLongTermDebtAsReported)
		} else if NonCurrentLongTermDebtAsReported != nil {
			LongTermDebtAsReported = utils.InterfaceToFloat64Ptr(*NonCurrentLongTermDebtAsReported)
		}

		if ShortTermDebt != nil && LongTermDebt != nil && ShareholderEquity != nil && *ShareholderEquity != 0 && *ShortTermDebt != 0 && *LongTermDebt != 0 {
			DebtToCapitalRatio = utils.InterfaceToFloat64Ptr(Calculations.DebtToCapitalRatio(*ShortTermDebt, *LongTermDebt, *ShareholderEquity))
		}

		if ShortTermDebtAsReported != nil && LongTermDebtAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 && *ShortTermDebtAsReported != 0 && *LongTermDebtAsReported != 0 {
			DebtToCapitalRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DebtToCapitalRatio(*ShortTermDebtAsReported, *LongTermDebtAsReported, *ShareholderEquityAsReported))
		}

		if LongTermDebt != nil && ShortTermDebt != nil && TotalLiabilities != nil && ShareholderEquity != nil && *ShareholderEquity != 0 {
			NetGearingRatio = utils.InterfaceToFloat64Ptr(Calculations.NetGearingRatio(*LongTermDebt, *ShortTermDebt, *TotalLiabilities, *ShareholderEquity))
		}

		if LongTermDebtAsReported != nil && ShortTermDebtAsReported != nil && TotalLiabilitiesAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 {
			NetGearingRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.NetGearingRatio(*LongTermDebtAsReported, *ShortTermDebtAsReported, *TotalLiabilitiesAsReported, *ShareholderEquityAsReported))
		}

		if TotalDebt != nil && EBITDA != nil && *EBITDA != 0 {
			TotalDebtToEBITDA = utils.InterfaceToFloat64Ptr(Calculations.TotalDebtToEBITDA(*TotalDebt, *EBITDA))
		}

		if TotalDebtAsReported != nil && EBITDAAsReported != nil && *EBITDAAsReported != 0 {
			TotalDebtToEBITDAAsReported = utils.InterfaceToFloat64Ptr(Calculations.TotalDebtToEBITDA(*TotalDebtAsReported, *EBITDAAsReported))
		}

		if TotalLiabilities != nil && ShareholderEquity != nil && *ShareholderEquity != 0 {
			DebtToEquityRatio = utils.InterfaceToFloat64Ptr(Calculations.DebtToEquityRatio(*TotalLiabilities, *ShareholderEquity))
		}

		if TotalLiabilitiesAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 {
			DebtToEquityRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DebtToEquityRatio(*TotalLiabilitiesAsReported, *ShareholderEquityAsReported))
		}

		if TotalAssets != nil && ShareholderEquity != nil && *ShareholderEquity != 0 {
			EquityMultiplierRatio = utils.InterfaceToFloat64Ptr(Calculations.EquityMultiplierRatio(*TotalAssets, *ShareholderEquity))
		}

		if TotalAssetsAsReported != nil && ShareholderEquityAsReported != nil && *ShareholderEquityAsReported != 0 {
			EquityMultiplierRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EquityMultiplierRatio(*TotalAssetsAsReported, *ShareholderEquityAsReported))
		}

		if NetProfitMargin != nil && AssetTurnoverRatio != nil && EquityMultiplierRatio != nil {
			DuPontAnalysis = utils.InterfaceToFloat64Ptr(Calculations.DuPontAnalysis(*NetProfitMargin, *AssetTurnoverRatio, *EquityMultiplierRatio))
		}

		if NetProfitMarginAsReported != nil && AssetTurnoverRatioAsReported != nil && EquityMultiplierRatioAsReported != nil {
			DuPontAnalysisAsReported = utils.InterfaceToFloat64Ptr(Calculations.DuPontAnalysis(*NetProfitMarginAsReported, *AssetTurnoverRatioAsReported, *EquityMultiplierRatioAsReported))
		}

		checkAndSet := func(value *float64) *float64 {
			if value == nil {
				return nil
			}
			// TODO: why are we getting infinties and NaNs????
			if math.IsInf(*value, 0) || math.IsNaN(*value) || math.IsInf(*value, -1) || math.IsInf(*value, 1) {
				return nil
			}
			return value
		}

		FullCalcResults["EffectiveTaxRate"] = checkAndSet(EffectiveTaxRate)
		FullCalcResults["DaysInPeriod"] = checkAndSet(DaysInPeriod)
		FullCalcResults["TotalAssets"] = checkAndSet(TotalAssets)
		FullCalcResults["TotalLiabilities"] = checkAndSet(TotalLiabilities)
		FullCalcResults["Inventory"] = checkAndSet(Inventory)
		FullCalcResults["IntangibleAssets"] = checkAndSet(IntangibleAssets)
		FullCalcResults["NetDebt"] = checkAndSet(NetDebt)
		FullCalcResults["CashAndCashEquivalents"] = checkAndSet(CashAndCashEquivalents)
		FullCalcResults["NetReceivables"] = checkAndSet(NetReceivables)
		FullCalcResults["NetFixedAssets"] = checkAndSet(NetFixedAssets)
		FullCalcResults["DeferredTaxLiabilities"] = checkAndSet(DeferredTaxLiabilities)
		FullCalcResults["ShareholderEquity"] = checkAndSet(ShareholderEquity)
		FullCalcResults["AccountsPayable"] = checkAndSet(AccountsPayable)
		FullCalcResults["CommonStock"] = checkAndSet(CommonStock)
		FullCalcResults["SharesOutstanding"] = checkAndSet(SharesOutstanding)
		FullCalcResults["HighQualityLiquidAssets"] = checkAndSet(HighQualityLiquidAssets)
		FullCalcResults["WorkingCapital"] = checkAndSet(WorkingCapital)
		FullCalcResults["TangibleNetWorth"] = checkAndSet(TangibleNetWorth)
		FullCalcResults["BookValueOfEquity"] = checkAndSet(BookValueOfEquity)
		FullCalcResults["BookValueOfDebt"] = checkAndSet(BookValueOfDebt)
		FullCalcResults["EquityBookValue"] = checkAndSet(EquityBookValue)
		FullCalcResults["LiabilitiesBookValue"] = checkAndSet(LiabilitiesBookValue)
		FullCalcResults["TotalAccrualsToTotalAssets"] = checkAndSet(TotalAccrualsToTotalAssets)
		FullCalcResults["ShortTermInvestments"] = checkAndSet(ShortTermInvestments)
		FullCalcResults["LongTermInvestments"] = checkAndSet(LongTermInvestments)
		FullCalcResults["TotalMarketableSecurities"] = checkAndSet(TotalMarketableSecurities)
		FullCalcResults["TotalInvestments"] = checkAndSet(TotalInvestments)
		FullCalcResults["NetIncome"] = checkAndSet(NetIncome)
		FullCalcResults["GrossProfit"] = checkAndSet(GrossProfit)
		FullCalcResults["NetRevenue"] = checkAndSet(NetRevenue)
		FullCalcResults["NetProfitMargin"] = checkAndSet(NetProfitMargin)
		FullCalcResults["OperatingExpenses"] = checkAndSet(OperatingExpenses)
		FullCalcResults["OperatingIncome"] = checkAndSet(OperatingIncome)
		FullCalcResults["DepreciationAndAmortization"] = checkAndSet(DepreciationAndAmortization)
		FullCalcResults["TotalTaxesPaid"] = checkAndSet(TotalTaxesPaid)
		FullCalcResults["ChangeInWorkingCapital"] = checkAndSet(ChangeInWorkingCapital)
		FullCalcResults["CapitalExpenditures"] = checkAndSet(CapitalExpenditures)
		FullCalcResults["OperatingCashflow"] = checkAndSet(OperatingCashflow)
		FullCalcResults["FreeCashFlow"] = checkAndSet(FreeCashFlow)
		FullCalcResults["EBITDA"] = checkAndSet(EBITDA)
		FullCalcResults["TotalInterestPayments"] = checkAndSet(TotalInterestPayments)
		FullCalcResults["EBIT"] = checkAndSet(EBIT)
		FullCalcResults["NonCashCharges"] = checkAndSet(NonCashCharges)
		FullCalcResults["MarketValueOfEquity"] = checkAndSet(MarketValueOfEquity)
		FullCalcResults["ShortTermDebt"] = checkAndSet(ShortTermDebt)
		FullCalcResults["LongTermDebt"] = checkAndSet(LongTermDebt)
		FullCalcResults["TotalDebt"] = checkAndSet(TotalDebt)
		FullCalcResults["CostOfDebt"] = checkAndSet(CostOfDebt)
		FullCalcResults["UnleveredFirmValue"] = checkAndSet(UnleveredFirmValue)
		FullCalcResults["TaxShieldBenefits"] = checkAndSet(TaxShieldBenefits)
		FullCalcResults["NetEffectOfDebt"] = checkAndSet(NetEffectOfDebt)
		FullCalcResults["DebtService"] = checkAndSet(DebtService)
		FullCalcResults["NonInterestExpenses"] = checkAndSet(NonInterestExpenses)
		FullCalcResults["MarketCapitalization"] = checkAndSet(MarketCapitalization)
		FullCalcResults["EnterpriseValue"] = checkAndSet(EnterpriseValue)
		FullCalcResults["DebtOutstanding"] = checkAndSet(DebtOutstanding)
		FullCalcResults["AssetTurnoverRatio"] = checkAndSet(AssetTurnoverRatio)
		FullCalcResults["DividendsPaid"] = checkAndSet(DividendsPaid)
		FullCalcResults["RetentionRatio"] = checkAndSet(RetentionRatio)
		FullCalcResults["ReturnOnEquity"] = checkAndSet(ReturnOnEquity)
		FullCalcResults["CostAndExpenses"] = checkAndSet(CostAndExpenses)
		FullCalcResults["CostOfRevenue"] = checkAndSet(CostOfRevenue)
		FullCalcResults["SellingGeneralAndAdministrativeExpenses"] = checkAndSet(SellingGeneralAndAdministrativeExpenses)
		FullCalcResults["ExplicitCosts"] = checkAndSet(ExplicitCosts)
		FullCalcResults["DaysInventoryOutstanding"] = checkAndSet(DaysInventoryOutstanding)
		FullCalcResults["TotalCapital"] = checkAndSet(TotalCapital)
		FullCalcResults["NetMargin"] = checkAndSet(NetMargin)
		FullCalcResults["FreeCashFlowToEquity"] = checkAndSet(FreeCashFlowToEquity)
		FullCalcResults["AdjustedPresentValue"] = checkAndSet(AdjustedPresentValue)
		FullCalcResults["InterestCoverageRatio"] = checkAndSet(InterestCoverageRatio)
		FullCalcResults["FixedChargeCoverageRatio"] = checkAndSet(FixedChargeCoverageRatio)
		FullCalcResults["DebtServiceCoverageRatio"] = checkAndSet(DebtServiceCoverageRatio)
		FullCalcResults["AssetCoverageRatio"] = checkAndSet(AssetCoverageRatio)
		FullCalcResults["EBITDAToInterestCoverageRatio"] = checkAndSet(EBITDAToInterestCoverageRatio)
		FullCalcResults["PreferredDividendCoverageRatio"] = checkAndSet(PreferredDividendCoverageRatio)
		FullCalcResults["LiquidityCoverageRatio"] = checkAndSet(LiquidityCoverageRatio)
		FullCalcResults["InventoryTurnoverRatio"] = checkAndSet(InventoryTurnoverRatio)
		FullCalcResults["ReturnOnCapitalEmployed"] = checkAndSet(ReturnOnCapitalEmployed)
		FullCalcResults["EfficiencyRatio"] = checkAndSet(EfficiencyRatio)
		FullCalcResults["RevenuePerEmployee"] = checkAndSet(RevenuePerEmployee)
		FullCalcResults["CapitalExpenditureRatio"] = checkAndSet(CapitalExpenditureRatio)
		FullCalcResults["OperatingCashFlowRatio"] = checkAndSet(OperatingCashFlowRatio)
		FullCalcResults["EBITDAToEVRatio"] = checkAndSet(EBITDAToEVRatio)
		FullCalcResults["TangibleNetWorthRatio"] = checkAndSet(TangibleNetWorthRatio)
		FullCalcResults["DeferredTaxLiabilityToEquityRatio"] = checkAndSet(DeferredTaxLiabilityToEquityRatio)
		FullCalcResults["TangibleEquityRatio"] = checkAndSet(TangibleEquityRatio)
		FullCalcResults["WACC"] = checkAndSet(WACC)
		FullCalcResults["FixedAssetTurnoverRatio"] = checkAndSet(FixedAssetTurnoverRatio)
		FullCalcResults["PPETurnoverRatio"] = checkAndSet(PPETurnoverRatio)
		FullCalcResults["InvestmentTurnoverRatio"] = checkAndSet(InvestmentTurnoverRatio)
		FullCalcResults["WorkingCapitalTurnoverRatio"] = checkAndSet(WorkingCapitalTurnoverRatio)
		FullCalcResults["ReturnOnAssetRatio"] = checkAndSet(ReturnOnAssetRatio)
		FullCalcResults["GrossProfitMargin"] = checkAndSet(GrossProfitMargin)
		FullCalcResults["OperatingProfitMargin"] = checkAndSet(OperatingProfitMargin)
		FullCalcResults["EBITDAMarginRatio"] = checkAndSet(EBITDAMarginRatio)
		FullCalcResults["DividendPayoutRatio"] = checkAndSet(DividendPayoutRatio)
		FullCalcResults["RetentionRate"] = checkAndSet(RetentionRate)
		FullCalcResults["SustainableGrowthRate"] = checkAndSet(SustainableGrowthRate)
		FullCalcResults["GrossMarginOnInventory"] = checkAndSet(GrossMarginOnInventory)
		FullCalcResults["CashFlowReturnOnEquity"] = checkAndSet(CashFlowReturnOnEquity)
		FullCalcResults["OperatingMargin"] = checkAndSet(OperatingMargin)
		FullCalcResults["OperatingExpenseRatio"] = checkAndSet(OperatingExpenseRatio)
		FullCalcResults["CurrentRatio"] = checkAndSet(CurrentRatio)
		FullCalcResults["AcidTestRatio"] = checkAndSet(AcidTestRatio)
		FullCalcResults["CashRatio"] = checkAndSet(CashRatio)
		FullCalcResults["DefensiveIntervalRatio"] = checkAndSet(DefensiveIntervalRatio)
		FullCalcResults["DrySalesRatio"] = checkAndSet(DrySalesRatio)
		FullCalcResults["PriceToBookValueRatio"] = checkAndSet(PriceToBookValueRatio)
		FullCalcResults["EarningsPerShare"] = checkAndSet(EarningsPerShare)
		FullCalcResults["EBITDAPerShare"] = checkAndSet(EBITDAPerShare)
		FullCalcResults["BookValuePerShare"] = checkAndSet(BookValuePerShare)
		FullCalcResults["NetTangibleAssetsPerShare"] = checkAndSet(NetTangibleAssetsPerShare)
		FullCalcResults["MarketValueOfDebt"] = checkAndSet(MarketValueOfDebt)
		FullCalcResults["MarketToBookRatio"] = checkAndSet(MarketToBookRatio)
		FullCalcResults["IntangiblesRatio"] = checkAndSet(IntangiblesRatio)
		FullCalcResults["PriceToSalesRatio"] = checkAndSet(PriceToSalesRatio)
		FullCalcResults["PriceToBookRatio"] = checkAndSet(PriceToBookRatio)
		FullCalcResults["PricetoSalesValue"] = checkAndSet(PricetoSalesValue)
		FullCalcResults["OperatingCashFlowPerShare"] = checkAndSet(OperatingCashFlowPerShare)
		FullCalcResults["PriceToCashFlowRatio"] = checkAndSet(PriceToCashFlowRatio)
		FullCalcResults["FreeCashFlowPerShare"] = checkAndSet(FreeCashFlowPerShare)
		FullCalcResults["PriceToFreeCashFlowRatio"] = checkAndSet(PriceToFreeCashFlowRatio)
		FullCalcResults["PriceToCashFlowValuation"] = checkAndSet(PriceToCashFlowValuation)
		FullCalcResults["PriceToFreeCashFlowValuation"] = checkAndSet(PriceToFreeCashFlowValuation)
		FullCalcResults["PriceToEarningsValuation"] = checkAndSet(PriceToEarningsValuation)
		FullCalcResults["LiabilitiesMarketValue"] = checkAndSet(LiabilitiesMarketValue)
		FullCalcResults["TobinsQ"] = checkAndSet(TobinsQ)
		FullCalcResults["ReceivablesTurnoverRatio"] = checkAndSet(ReceivablesTurnoverRatio)
		FullCalcResults["AverageCollectionPeriod"] = checkAndSet(AverageCollectionPeriod)
		FullCalcResults["AccountsPayableTurnoverRatio"] = checkAndSet(AccountsPayableTurnoverRatio)
		FullCalcResults["AverageAccountsPayablePaymentPeriod"] = checkAndSet(AverageAccountsPayablePaymentPeriod)
		FullCalcResults["InventoryToWorkingCapitalRatio"] = checkAndSet(InventoryToWorkingCapitalRatio)
		FullCalcResults["DaysSalesOutstanding"] = checkAndSet(DaysSalesOutstanding)
		FullCalcResults["DaysPayablesOutstanding"] = checkAndSet(DaysPayablesOutstanding)
		FullCalcResults["CashConversionCycle"] = checkAndSet(CashConversionCycle)
		FullCalcResults["NetWorkingCapital"] = checkAndSet(NetWorkingCapital)
		FullCalcResults["NOPAT"] = checkAndSet(NOPAT)
		FullCalcResults["EconomicValueAdded"] = checkAndSet(EconomicValueAdded)
		FullCalcResults["ReturnOnInvestedCapital"] = checkAndSet(ReturnOnInvestedCapital)
		FullCalcResults["FreeCashFlowToFirm"] = checkAndSet(FreeCashFlowToFirm)
		FullCalcResults["StockDividendPerShare"] = checkAndSet(StockDividendPerShare)
		FullCalcResults["LeverageRatio"] = checkAndSet(LeverageRatio)
		FullCalcResults["CapitalizationRatio"] = checkAndSet(CapitalizationRatio)
		FullCalcResults["DebtToCapitalRatio"] = checkAndSet(DebtToCapitalRatio)
		FullCalcResults["NetGearingRatio"] = checkAndSet(NetGearingRatio)
		FullCalcResults["TotalDebtToEBITDA"] = checkAndSet(TotalDebtToEBITDA)
		FullCalcResults["DebtToEquityRatio"] = checkAndSet(DebtToEquityRatio)
		FullCalcResults["EquityMultiplierRatio"] = checkAndSet(EquityMultiplierRatio)
		FullCalcResults["DuPontAnalysis"] = checkAndSet(DuPontAnalysis)

		FullCalcResultsAsReported["EffectiveTaxRate"] = checkAndSet(EffectiveTaxRate)
		FullCalcResultsAsReported["DaysInPeriod"] = checkAndSet(DaysInPeriod)
		FullCalcResultsAsReported["TotalAssets"] = checkAndSet(TotalAssetsAsReported)
		FullCalcResultsAsReported["AssetsNonCurrent"] = checkAndSet(AssetsNonCurrentAsReported)
		FullCalcResultsAsReported["OtherAssetsNonCurrent"] = checkAndSet(OtherAssetsNonCurrentAsReported)
		FullCalcResultsAsReported["TotalLiabilities"] = checkAndSet(TotalLiabilitiesAsReported)
		FullCalcResultsAsReported["Inventory"] = checkAndSet(InventoryAsReported)
		FullCalcResultsAsReported["IntangibleAssets"] = checkAndSet(IntangibleAssetsAsReported)
		FullCalcResultsAsReported["CurrentLongTermDebt"] = checkAndSet(CurrentLongTermDebtAsReported)
		FullCalcResultsAsReported["NonCurrentLongTermDebt"] = checkAndSet(NonCurrentLongTermDebtAsReported)
		FullCalcResultsAsReported["NetDebt"] = checkAndSet(NetDebtAsReported)
		FullCalcResultsAsReported["CashAndCashEquivalents"] = checkAndSet(CashAndCashEquivalentsAsReported)
		FullCalcResultsAsReported["AccountsReceivable"] = checkAndSet(AccountsReceivableAsReported)
		FullCalcResultsAsReported["NonTradeReceivables"] = checkAndSet(NonTradeReceivablesAsReported)
		FullCalcResultsAsReported["NetReceivables"] = checkAndSet(NetReceivablesAsReported)
		FullCalcResultsAsReported["NetFixedAssets"] = checkAndSet(NetFixedAssetsAsReported)
		FullCalcResultsAsReported["ShareholderEquity"] = checkAndSet(ShareholderEquityAsReported)
		FullCalcResultsAsReported["AccountsPayable"] = checkAndSet(AccountsPayableAsReported)
		FullCalcResultsAsReported["SharesOutstanding"] = checkAndSet(SharesOutstandingAsReported)
		FullCalcResultsAsReported["HighQualityLiquidAssets"] = checkAndSet(HighQualityLiquidAssetsAsReported)
		FullCalcResultsAsReported["WorkingCapital"] = checkAndSet(WorkingCapitalAsReported)
		FullCalcResultsAsReported["TangibleNetWorth"] = checkAndSet(TangibleNetWorthAsReported)
		FullCalcResultsAsReported["BookValueOfEquity"] = checkAndSet(BookValueOfEquityAsReported)
		FullCalcResultsAsReported["BookValueOfDebt"] = checkAndSet(BookValueOfDebtAsReported)
		FullCalcResultsAsReported["EquityBookValue"] = checkAndSet(EquityBookValueAsReported)
		FullCalcResultsAsReported["LiabilitiesBookValue"] = checkAndSet(LiabilitiesBookValueAsReported)
		FullCalcResultsAsReported["TotalAccrualsToTotalAssets"] = checkAndSet(TotalAccrualsToTotalAssetsAsReported)
		FullCalcResultsAsReported["TotalMarketableSecurities"] = checkAndSet(TotalMarketableSecuritiesAsReported)
		FullCalcResultsAsReported["CurrentMarketableSecurities"] = checkAndSet(CurrentMarketableSecuritiesAsReported)
		FullCalcResultsAsReported["NonCurrentMarketableSecurities"] = checkAndSet(NonCurrentMarketableSecuritiesAsReported)
		FullCalcResultsAsReported["LongTermInvestments"] = checkAndSet(LongTermInvestmentsAsReported)
		FullCalcResultsAsReported["TotalInvestments"] = checkAndSet(TotalInvestmentsAsReported)
		FullCalcResultsAsReported["NetIncome"] = checkAndSet(NetIncomeAsReported)
		FullCalcResultsAsReported["GrossProfit"] = checkAndSet(GrossProfitAsReported)
		FullCalcResultsAsReported["NetRevenue"] = checkAndSet(NetRevenueAsReported)
		FullCalcResultsAsReported["NetProfitMargin"] = checkAndSet(NetProfitMarginAsReported)
		FullCalcResultsAsReported["OperatingExpenses"] = checkAndSet(OperatingExpensesAsReported)
		FullCalcResultsAsReported["OperatingIncome"] = checkAndSet(OperatingIncomeAsReported)
		FullCalcResultsAsReported["DepreciationAndAmortization"] = checkAndSet(DepreciationAndAmortizationAsReported)
		FullCalcResultsAsReported["TotalInterestPayments"] = checkAndSet(TotalInterestPaymentsAsReported)
		FullCalcResultsAsReported["TotalTaxesPaid"] = checkAndSet(TotalTaxesPaidAsReported)
		FullCalcResultsAsReported["CapitalExpenditures"] = checkAndSet(CapitalExpendituresAsReported)
		FullCalcResultsAsReported["NetCashOperatingActivities"] = checkAndSet(NetCashOperatingActivitiesAsReported)
		FullCalcResultsAsReported["NetCashInvestingActivities"] = checkAndSet(NetCashInvestingActivitiesAsReported)
		FullCalcResultsAsReported["NetCashFinancingActivities"] = checkAndSet(NetCashFinancingActivitiesAsReported)
		FullCalcResultsAsReported["OperatingCashFlow"] = checkAndSet(OperatingCashFlowAsReported)
		FullCalcResultsAsReported["FreeCashFlow"] = checkAndSet(FreeCashFlowAsReported)
		FullCalcResultsAsReported["EBITDA"] = checkAndSet(EBITDAAsReported)
		FullCalcResultsAsReported["EBIT"] = checkAndSet(EBITAsReported)
		FullCalcResultsAsReported["NonCashCharges"] = checkAndSet(NonCashChargesAsReported)
		FullCalcResultsAsReported["MarketValueOfEquity"] = checkAndSet(MarketValueOfEquityAsReported)
		FullCalcResultsAsReported["ShortTermDebt"] = checkAndSet(ShortTermDebtAsReported)
		FullCalcResultsAsReported["TotalDebt"] = checkAndSet(TotalDebtAsReported)
		FullCalcResultsAsReported["CostOfDebt"] = checkAndSet(CostOfDebtAsReported)
		FullCalcResultsAsReported["UnleveredFirmValue"] = checkAndSet(UnleveredFirmValueAsReported)
		FullCalcResultsAsReported["TaxShieldBenefits"] = checkAndSet(TaxShieldBenefitsAsReported)
		FullCalcResultsAsReported["NetEffectOfDebt"] = checkAndSet(NetEffectOfDebtAsReported)
		FullCalcResultsAsReported["DebtService"] = checkAndSet(DebtServiceAsReported)
		FullCalcResultsAsReported["NonInterestExpenses"] = checkAndSet(NonInterestExpensesAsReported)
		FullCalcResultsAsReported["MarketCapitalization"] = checkAndSet(MarketCapitalizationAsReported)
		FullCalcResultsAsReported["EnterpriseValue"] = checkAndSet(EnterpriseValueAsReported)
		FullCalcResultsAsReported["DebtOutstanding"] = checkAndSet(DebtOutstandingAsReported)
		FullCalcResultsAsReported["AssetTurnoverRatio"] = checkAndSet(AssetTurnoverRatioAsReported)
		FullCalcResultsAsReported["DividendsPaid"] = checkAndSet(DividendsPaidAsReported)
		FullCalcResultsAsReported["RetentionRatio"] = checkAndSet(RetentionRatioAsReported)
		FullCalcResultsAsReported["ReturnOnEquity"] = checkAndSet(ReturnOnEquityAsReported)
		FullCalcResultsAsReported["CostOfGoodsSold"] = checkAndSet(CostOfGoodsSoldAsReported)
		FullCalcResultsAsReported["SellingGeneralAndAdministrativeExpenses"] = checkAndSet(SellingGeneralAndAdministrativeExpensesAsReported)
		FullCalcResultsAsReported["ExplicitCosts"] = checkAndSet(ExplicitCostsAsReported)
		FullCalcResultsAsReported["DaysInventoryOutstanding"] = checkAndSet(DaysInventoryOutstandingAsReported)
		FullCalcResultsAsReported["TotalCapital"] = checkAndSet(TotalCapitalAsReported)
		FullCalcResultsAsReported["NetMargin"] = checkAndSet(NetMarginAsReported)
		FullCalcResultsAsReported["AdjustedPresentValue"] = checkAndSet(AdjustedPresentValueAsReported)
		FullCalcResultsAsReported["InterestCoverageRatio"] = checkAndSet(InterestCoverageRatioAsReported)
		FullCalcResultsAsReported["FixedChargeCoverageRatio"] = checkAndSet(FixedChargeCoverageRatioAsReported)
		FullCalcResultsAsReported["DebtServiceCoverageRatio"] = checkAndSet(DebtServiceCoverageRatioAsReported)
		FullCalcResultsAsReported["AssetCoverageRatio"] = checkAndSet(AssetCoverageRatioAsReported)
		FullCalcResultsAsReported["EBITDAToInterestCoverageRatio"] = checkAndSet(EBITDAToInterestCoverageRatioAsReported)
		FullCalcResultsAsReported["PreferredDividendCoverageRatio"] = checkAndSet(PreferredDividendCoverageRatioAsReported)
		FullCalcResultsAsReported["LiquidityCoverageRatio"] = checkAndSet(LiquidityCoverageRatioAsReported)
		FullCalcResultsAsReported["InventoryTurnoverRatio"] = checkAndSet(InventoryTurnoverRatioAsReported)
		FullCalcResultsAsReported["ReturnOnCapitalEmployed"] = checkAndSet(ReturnOnCapitalEmployedAsReported)
		FullCalcResultsAsReported["EfficiencyRatio"] = checkAndSet(EfficiencyRatioAsReported)
		FullCalcResultsAsReported["RevenuePerEmployee"] = checkAndSet(RevenuePerEmployeeAsReported)
		FullCalcResultsAsReported["CapitalExpenditureRatio"] = checkAndSet(CapitalExpenditureRatioAsReported)
		FullCalcResultsAsReported["OperatingCashFlowRatio"] = checkAndSet(OperatingCashFlowRatioAsReported)
		FullCalcResultsAsReported["EBITDAToEVRatio"] = checkAndSet(EBITDAToEVRatioAsReported)
		FullCalcResultsAsReported["TangibleNetWorthRatio"] = checkAndSet(TangibleNetWorthRatioAsReported)
		FullCalcResultsAsReported["DeferredTaxLiabilityToEquityRatio"] = checkAndSet(DeferredTaxLiabilityToEquityRatioAsReported)
		FullCalcResultsAsReported["TangibleEquityRatio"] = checkAndSet(TangibleEquityRatioAsReported)
		FullCalcResultsAsReported["WACC"] = checkAndSet(WACCAsReported)
		FullCalcResultsAsReported["FixedAssetTurnoverRatio"] = checkAndSet(FixedAssetTurnoverRatioAsReported)
		FullCalcResultsAsReported["PPETurnoverRatio"] = checkAndSet(PPETurnoverRatioAsReported)
		FullCalcResultsAsReported["InvestmentTurnoverRatio"] = checkAndSet(InvestmentTurnoverRatioAsReported)
		FullCalcResultsAsReported["WorkingCapitalTurnoverRatio"] = checkAndSet(WorkingCapitalTurnoverRatioAsReported)
		FullCalcResultsAsReported["ReturnOnAssetRatio"] = checkAndSet(ReturnOnAssetRatioAsReported)
		FullCalcResultsAsReported["GrossProfitMargin"] = checkAndSet(GrossProfitMarginAsReported)
		FullCalcResultsAsReported["OperatingProfitMargin"] = checkAndSet(OperatingProfitMarginAsReported)
		FullCalcResultsAsReported["EBITDAMarginRatio"] = checkAndSet(EBITDAMarginRatioAsReported)
		FullCalcResultsAsReported["DividendPayoutRatio"] = checkAndSet(DividendPayoutRatioAsReported)
		FullCalcResultsAsReported["RetentionRate"] = checkAndSet(RetentionRateAsReported)
		FullCalcResultsAsReported["SustainableGrowthRate"] = checkAndSet(SustainableGrowthRateAsReported)
		FullCalcResultsAsReported["GrossMarginOnInventory"] = checkAndSet(GrossMarginOnInventoryAsReported)
		FullCalcResultsAsReported["CashFlowReturnOnEquity"] = checkAndSet(CashFlowReturnOnEquityAsReported)
		FullCalcResultsAsReported["OperatingMargin"] = checkAndSet(OperatingMarginAsReported)
		FullCalcResultsAsReported["OperatingExpenseRatio"] = checkAndSet(OperatingExpenseRatioAsReported)
		FullCalcResultsAsReported["CurrentRatio"] = checkAndSet(CurrentRatioAsReported)
		FullCalcResultsAsReported["AcidTestRatio"] = checkAndSet(AcidTestRatioAsReported)
		FullCalcResultsAsReported["CashRatio"] = checkAndSet(CashRatioAsReported)
		FullCalcResultsAsReported["DefensiveIntervalRatio"] = checkAndSet(DefensiveIntervalRatioAsReported)
		FullCalcResultsAsReported["DrySalesRatio"] = checkAndSet(DrySalesRatioAsReported)
		FullCalcResultsAsReported["PriceToBookValueRatio"] = checkAndSet(PriceToBookValueRatioAsReported)
		FullCalcResultsAsReported["EarningsPerShare"] = checkAndSet(EarningsPerShareAsReported)
		FullCalcResultsAsReported["EBITDAPerShare"] = checkAndSet(EBITDAPerShareAsReported)
		FullCalcResultsAsReported["BookValuePerShare"] = checkAndSet(BookValuePerShareAsReported)
		FullCalcResultsAsReported["NetTangibleAssetsPerShare"] = checkAndSet(NetTangibleAssetsPerShareAsReported)
		FullCalcResultsAsReported["MarketValueOfDebt"] = checkAndSet(MarketValueOfDebtAsReported)
		FullCalcResultsAsReported["MarketToBookRatio"] = checkAndSet(MarketToBookRatioAsReported)
		FullCalcResultsAsReported["IntangiblesRatio"] = checkAndSet(IntangiblesRatioAsReported)
		FullCalcResultsAsReported["PriceToSalesRatio"] = checkAndSet(PriceToSalesRatioAsReported)
		FullCalcResultsAsReported["PriceToBookRatio"] = checkAndSet(PriceToBookRatioAsReported)
		FullCalcResultsAsReported["PricetoSalesValue"] = checkAndSet(PricetoSalesValueAsReported)
		FullCalcResultsAsReported["OperatingCashFlowPerShare"] = checkAndSet(OperatingCashFlowPerShareAsReported)
		FullCalcResultsAsReported["PriceToCashFlowRatio"] = checkAndSet(PriceToCashFlowRatioAsReported)
		FullCalcResultsAsReported["FreeCashFlowPerShare"] = checkAndSet(FreeCashFlowPerShareAsReported)
		FullCalcResultsAsReported["PriceToFreeCashFlowRatio"] = checkAndSet(PriceToFreeCashFlowRatioAsReported)
		FullCalcResultsAsReported["PriceToCashFlowValuation"] = checkAndSet(PriceToCashFlowValuationAsReported)
		FullCalcResultsAsReported["PriceToFreeCashFlowValuation"] = checkAndSet(PriceToFreeCashFlowValuationAsReported)
		FullCalcResultsAsReported["PriceToEarningsValuation"] = checkAndSet(PriceToEarningsValuationAsReported)
		FullCalcResultsAsReported["LiabilitiesMarketValue"] = checkAndSet(LiabilitiesMarketValueAsReported)
		FullCalcResultsAsReported["TobinsQ"] = checkAndSet(TobinsQAsReported)
		FullCalcResultsAsReported["ReceivablesTurnoverRatio"] = checkAndSet(ReceivablesTurnoverRatioAsReported)
		FullCalcResultsAsReported["AverageCollectionPeriod"] = checkAndSet(AverageCollectionPeriodAsReported)
		FullCalcResultsAsReported["AccountsPayableTurnoverRatio"] = checkAndSet(AccountsPayableTurnoverRatioAsReported)
		FullCalcResultsAsReported["AverageAccountsPayablePaymentPeriod"] = checkAndSet(AverageAccountsPayablePaymentPeriodAsReported)
		FullCalcResultsAsReported["InventoryToWorkingCapitalRatio"] = checkAndSet(InventoryToWorkingCapitalRatioAsReported)
		FullCalcResultsAsReported["DaysSalesOutstanding"] = checkAndSet(DaysSalesOutstandingAsReported)
		FullCalcResultsAsReported["DaysPayablesOutstanding"] = checkAndSet(DaysPayablesOutstandingAsReported)
		FullCalcResultsAsReported["CashConversionCycle"] = checkAndSet(CashConversionCycleAsReported)
		FullCalcResultsAsReported["NetWorkingCapital"] = checkAndSet(NetWorkingCapitalAsReported)
		FullCalcResultsAsReported["NOPAT"] = checkAndSet(NOPATAsReported)
		FullCalcResultsAsReported["EconomicValueAdded"] = checkAndSet(EconomicValueAddedAsReported)
		FullCalcResultsAsReported["ReturnOnInvestedCapital"] = checkAndSet(ReturnOnInvestedCapitalAsReported)
		FullCalcResultsAsReported["FreeCashFlowToFirm"] = checkAndSet(FreeCashFlowToFirmAsReported)
		FullCalcResultsAsReported["StockDividendPerShare"] = checkAndSet(StockDividendPerShareAsReported)
		FullCalcResultsAsReported["LeverageRatio"] = checkAndSet(LeverageRatioAsReported)
		FullCalcResultsAsReported["CapitalizationRatio"] = checkAndSet(CapitalizationRatioAsReported)
		FullCalcResultsAsReported["LongTermDebt"] = checkAndSet(LongTermDebtAsReported)
		FullCalcResultsAsReported["DebtToCapitalRatio"] = checkAndSet(DebtToCapitalRatioAsReported)
		FullCalcResultsAsReported["NetGearingRatio"] = checkAndSet(NetGearingRatioAsReported)
		FullCalcResultsAsReported["TotalDebtToEBITDA"] = checkAndSet(TotalDebtToEBITDAAsReported)
		FullCalcResultsAsReported["DebtToEquityRatio"] = checkAndSet(DebtToEquityRatioAsReported)
		FullCalcResultsAsReported["EquityMultiplierRatio"] = checkAndSet(EquityMultiplierRatioAsReported)
		FullCalcResultsAsReported["DuPontAnalysis"] = checkAndSet(DuPontAnalysisAsReported)

		FinalCalcResults = append(FinalCalcResults, FullCalcResults)
		FinalCalcResultsAsReported = append(FinalCalcResultsAsReported, FullCalcResultsAsReported)
	}

	return FinalCalcResults, FinalCalcResultsAsReported
}
