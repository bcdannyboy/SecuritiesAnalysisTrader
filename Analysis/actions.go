package Analysis

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Calculations"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func PerformFundamentalsCalculations(Fundamentals *CompanyFundamentals, Period string, PricePerShare float64, EffectiveTaxRate float64, NumEmployees float64, RiskFreeRate float64, MarketReturn float64, Beta float64) *FundamentalsCalculationsResults {
	CalculationResults := &FundamentalsCalculationsResults{
		Symbol:       Fundamentals.Symbol,
		Fundamentals: Fundamentals,
	}

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

	_, _, BalanceSheetStatementMissingPeriods, BalanceSheetStatementConsecutivePeriods, BalanceSheetStatementGapPeriods := Calculations.ProcessReportDates(BalanceSheetStatementReportDates, Period)
	_, _, BalanceSheetStatementAsReportedMissingPeriods, BalanceSheetStatementAsReportedConsecutivePeriods, BalanceSheetStatementAsReportedGapPeriods := Calculations.ProcessReportDates(BalanceSheetAsReportedReportDates, Period)

	_, _, IncomeStatementMissingPeriods, IncomeStatementConsecutivePeriods, IncomeStatementGapPeriods := Calculations.ProcessReportDates(IncomeStatementReportDates, Period)
	_, _, IncomeStatementAsReportedMissingPeriods, IncomeStatementAsReportedConsecutivePeriods, IncomeStatementAsReportedGapPeriods := Calculations.ProcessReportDates(IncomeStatementAsReportedReportDates, Period)

	_, _, CashFlowStatementMissingPeriods, CashFlowStatementConsecutivePeriods, CashFlowStatementGapPeriods := Calculations.ProcessReportDates(CashFlowReportDates, Period)
	_, _, CashFlowStatementAsReportedMissingPeriods, CashFlowStatementAsReportedConsecutivePeriods, CashFlowStatementAsReportedGapPeriods := Calculations.ProcessReportDates(CashFlowAsReportedReportDates, Period)

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
	CustomCalculationResults, CustomCalculationAsReportedResults := PerformCustomCalculations(Fundamentals, cvp, PricePerShare, EffectiveTaxRate, NumEmployees, CostOfEquity)

	CalculationResults.PeriodLength = cvp
	CalculationResults.CostOfEquity = CostOfEquity
	CalculationResults.Beta = Beta
	CalculationResults.EffectiveTaxRate = EffectiveTaxRate
	CalculationResults.CustomCalculations = CustomCalculationResults
	CalculationResults.CustomCalculationsAsReported = CustomCalculationAsReportedResults

	CalculationResults.CustomCalculationsGrowth = calculateGrowth(CustomCalculationResults)
	CalculationResults.CustomCalculationsAsReportedGrowth = calculateGrowth(CustomCalculationAsReportedResults)

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

	return CalculationResults
}

func PerformCustomCalculations(Fundamentals *CompanyFundamentals, Period objects.CompanyValuationPeriod, PricePerShare float64, ETR float64, NumEmployees float64, CostOfEquity float64) ([]map[string]*float64, []map[string]*float64) {
	FinalCalcResults := []map[string]*float64{}
	FinalCalcResultsAsReported := []map[string]*float64{}

	LenBalanceSheetStatements := len(Fundamentals.BalanceSheetStatements)
	LenBalanceSheetStatementAsReported := len(Fundamentals.BalanceSheetStatementAsReported)
	LenIncomeStatement := len(Fundamentals.IncomeStatement)
	LenIncomeStatementAsReported := len(Fundamentals.IncomeStatementAsReported)
	LenCashFlowStatement := len(Fundamentals.CashFlowStatement)
	LenCashFlowStatementAsReported := len(Fundamentals.CashFlowStatementAsReported)

	var EffectiveTaxRate = &ETR
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
		if current_iteration < LenBalanceSheetStatements {
			curBalanceSheet = Fundamentals.BalanceSheetStatements[current_iteration]
		}

		curBalanceSheetAsReported := objects.BalanceSheetStatementAsReported{}
		if current_iteration < LenBalanceSheetStatementAsReported {
			curBalanceSheetAsReported = Fundamentals.BalanceSheetStatementAsReported[current_iteration]
		}

		curIncomeStatement := objects.IncomeStatement{}
		if current_iteration < LenIncomeStatement {
			curIncomeStatement = Fundamentals.IncomeStatement[current_iteration]
		}

		curIncomeStatementAsReported := objects.IncomeStatementAsReported{}
		if current_iteration < LenIncomeStatementAsReported {
			curIncomeStatementAsReported = Fundamentals.IncomeStatementAsReported[current_iteration]
		}

		curCashFlowStatement := objects.CashFlowStatement{}
		if current_iteration < LenCashFlowStatement {
			curCashFlowStatement = Fundamentals.CashFlowStatement[current_iteration]
		}

		curCashFlowStatementAsReported := objects.CashFlowStatementAsReported{}
		if current_iteration < LenCashFlowStatementAsReported {
			curCashFlowStatementAsReported = Fundamentals.CashFlowStatementAsReported[current_iteration]
		}

		var TotalAssets = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "TotalAssets")
		var TotalAssetsAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Assets")
		var AssetsNonCurrentAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Assetsnoncurrent")
		var OtherAssetsNonCurrentAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Otherassetsnoncurrent")

		var TotalLiabilities = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "TotalLiabilities")
		var TotalLiabilitiesAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Liabilities")

		var Inventory = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "Inventory")
		var InventoryAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Inventorynet")

		var IntangibleAssets = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "IntangibleAssets")
		var IntangibleAssetsAsReported *float64 = nil
		if TotalAssetsAsReported != nil && InventoryAsReported != nil {
			IntangibleAssetsAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *InventoryAsReported)
		}

		var NetDebt = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "NetDebt")
		var CurrentLongTermDebtAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Longtermdebtcurrent")
		var NonCurrentLongTermDebtAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Longtermdebtnoncurrent")
		var NetDebtAsReported *float64 = nil
		if TotalLiabilitiesAsReported != nil && InventoryAsReported != nil {
			NetDebtAsReported = utils.InterfaceToFloat64Ptr(*CurrentLongTermDebtAsReported + *NonCurrentLongTermDebtAsReported)
		}

		var CashAndCashEquivalents = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "CashAndCashEquivalents")
		var CashAndCashEquivalentsAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Cashandcashequivalentsatcarryingvalue")

		var NetReceivables = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "NetReceivables")
		var AccountsReceivableAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Accountsreceivablenetcurrent")
		var NonTradeReceivablesAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Nontradereceivablescurrent")
		var NetReceivablesAsReported *float64 = nil
		if AccountsReceivableAsReported != nil && NonTradeReceivablesAsReported != nil {
			NetReceivablesAsReported = utils.InterfaceToFloat64Ptr(*AccountsReceivableAsReported + *NonTradeReceivablesAsReported)
		}

		var NetFixedAssets = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "PropertyPlantEquipmentNet")
		var NetFixedAssetsAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Propertyplantandequipmentnet")

		var DeferredTaxLiabilities = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "DeferredTaxLiabilitiesNonCurrent")
		var DeferredTaxLiabilitiesAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Deferredtaxliabilitiesnoncurrent")

		var ShareholderEquity = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "TotalStockholdersEquity")
		var ShareholderEquityAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Stockholdersequity")

		var AccountsPayable = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "AccountsPayable")
		var AccountsPayableAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Accountspayablecurrent")

		var CommonStock = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "CommonStock")
		var SharesOutstanding *float64 = nil
		if CommonStock != nil && PricePerShare != 0 {
			SharesOutstanding = utils.InterfaceToFloat64Ptr(*CommonStock / PricePerShare)
		}
		var SharesOutstandingAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Commonstocksharesoutstanding")

		var HighQualityLiquidAssets *float64 = nil
		var HighQualityLiquidAssetsAsReported *float64 = nil
		if CashAndCashEquivalents != nil && NetReceivables != nil {
			HighQualityLiquidAssets = utils.InterfaceToFloat64Ptr(*CashAndCashEquivalents + *NetReceivables + curBalanceSheet.CashAndShortTermInvestments - curBalanceSheet.ShortTermInvestments)
		}
		if CashAndCashEquivalentsAsReported != nil && NetReceivablesAsReported != nil {
			HighQualityLiquidAssetsAsReported = utils.InterfaceToFloat64Ptr(*CashAndCashEquivalentsAsReported + *NetReceivablesAsReported)
		}

		var WorkingCapital *float64 = nil
		var WorkingCapitalAsReported *float64 = nil
		if TotalAssets != nil && TotalLiabilities != nil {
			WorkingCapital = utils.InterfaceToFloat64Ptr(*TotalAssets - *TotalLiabilities)
		}
		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			WorkingCapitalAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *TotalLiabilitiesAsReported)
		}

		var TangibleNetWorth *float64 = nil
		var TangibleNetWorthAsReported *float64 = nil
		if TotalAssets != nil && IntangibleAssets != nil && TotalLiabilities != nil {
			TangibleNetWorth = utils.InterfaceToFloat64Ptr(*TotalAssets - *IntangibleAssets - *TotalLiabilities)
		}
		if TotalAssetsAsReported != nil && IntangibleAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			TangibleNetWorthAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *IntangibleAssetsAsReported - *TotalLiabilitiesAsReported)
		}

		var BookValueOfEquity *float64 = nil
		var BookValueOfEquityAsReported *float64 = nil
		if ShareholderEquity != nil && Inventory != nil {
			BookValueOfEquity = utils.InterfaceToFloat64Ptr(*ShareholderEquity - *Inventory)
		}
		if ShareholderEquityAsReported != nil && InventoryAsReported != nil {
			BookValueOfEquityAsReported = utils.InterfaceToFloat64Ptr(*ShareholderEquityAsReported - *InventoryAsReported)
		}

		var BookValueOfDebt *float64 = nil
		var BookValueOfDebtAsReported *float64 = nil
		if TotalLiabilities != nil && Inventory != nil {
			BookValueOfDebt = utils.InterfaceToFloat64Ptr(*TotalLiabilities - *Inventory)
		}
		if TotalLiabilitiesAsReported != nil && InventoryAsReported != nil {
			BookValueOfDebtAsReported = utils.InterfaceToFloat64Ptr(*TotalLiabilitiesAsReported - *InventoryAsReported)
		}

		var EquityBookValue *float64 = nil
		var EquityBookValueAsReported *float64 = nil
		if BookValueOfEquity != nil && BookValueOfDebt != nil {
			EquityBookValue = utils.InterfaceToFloat64Ptr(*BookValueOfEquity - *BookValueOfDebt)
		}
		if BookValueOfEquityAsReported != nil && BookValueOfDebtAsReported != nil {
			EquityBookValueAsReported = utils.InterfaceToFloat64Ptr(*BookValueOfEquityAsReported - *BookValueOfDebtAsReported)
		}

		var LiabilitiesBookValue *float64 = nil
		var LiabilitiesBookValueAsReported *float64 = nil
		if BookValueOfEquity != nil && BookValueOfDebt != nil {
			LiabilitiesBookValue = utils.InterfaceToFloat64Ptr(*BookValueOfEquity - *BookValueOfDebt)
		}
		if BookValueOfEquityAsReported != nil && BookValueOfDebtAsReported != nil {
			LiabilitiesBookValueAsReported = utils.InterfaceToFloat64Ptr(*BookValueOfEquityAsReported - *BookValueOfDebtAsReported)
		}

		var TotalAccrualsToTotalAssets *float64 = nil
		var TotalAccrualsToTotalAssetsAsReported *float64 = nil
		if TotalAssets != nil && TotalLiabilities != nil && TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			TotalAccrualsToTotalAssets = utils.InterfaceToFloat64Ptr(*TotalAssets - *TotalLiabilities)
		}
		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			TotalAccrualsToTotalAssetsAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *TotalLiabilitiesAsReported)
		}

		var TotalMarketableSecuritiesAsReported *float64 = nil
		var CurrentMarketableSecuritiesAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Marketablesecuritiescurrent")
		var NonCurrentMarketableSecuritiesAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Marketablesecuritiesnoncurrent")
		if CurrentMarketableSecuritiesAsReported != nil && NonCurrentMarketableSecuritiesAsReported != nil {
			TotalMarketableSecuritiesAsReported = utils.InterfaceToFloat64Ptr(*CurrentMarketableSecuritiesAsReported + *NonCurrentMarketableSecuritiesAsReported)
		}

		var ShortTermInvestments = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "ShortTermInvestments")
		var LongTermInvestments = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "LongTermInvestments")
		var LongTermInvestmentsAsReported *float64 = nil
		if AssetsNonCurrentAsReported != nil && NetFixedAssetsAsReported != nil && NonCurrentMarketableSecuritiesAsReported != nil && OtherAssetsNonCurrentAsReported != nil {
			LongTermInvestmentsAsReported = utils.InterfaceToFloat64Ptr(*AssetsNonCurrentAsReported - (*NetFixedAssetsAsReported + *NonCurrentMarketableSecuritiesAsReported + *OtherAssetsNonCurrentAsReported))
		}

		var TotalMarketableSecurities *float64 = nil
		if ShortTermInvestments != nil && LongTermInvestments != nil {
			TotalMarketableSecurities = utils.InterfaceToFloat64Ptr(*ShortTermInvestments + *LongTermInvestments)
		}

		var TotalInvestments = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "TotalInvestments")
		var TotalInvestmentsAsReported *float64 = nil
		if TotalMarketableSecuritiesAsReported != nil {
			TotalInvestmentsAsReported = TotalMarketableSecuritiesAsReported
		}

		var NetIncome = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "NetIncome")
		var NetIncomeAsReported = utils.GetFloat64PtrIfNotEmpty(curIncomeStatementAsReported, "Comprehensiveincomenetoftax")

		var GrossProfit = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "GrossProfit")
		var GrossProfitAsReported = utils.GetFloat64PtrIfNotEmpty(curIncomeStatementAsReported, "Grossprofit")

		var NetRevenue = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "Revenue")
		var NetRevenueAsReported = utils.GetFloat64PtrIfNotEmpty(curIncomeStatementAsReported, "Revenuefromcontractwithcustomerexcludingassessedtax")

		var NetProfitMargin = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "GrossProfitRatio")
		var NetProfitMarginAsReported *float64 = nil
		if NetRevenueAsReported != nil && NetIncomeAsReported != nil {
			NetProfitMarginAsReported = utils.InterfaceToFloat64Ptr(*NetIncomeAsReported / *NetRevenueAsReported)
		}

		var OperatingExpenses = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "OperatingExpenses")
		var OperatingExpensesAsReported = utils.GetFloat64PtrIfNotEmpty(curIncomeStatementAsReported, "Operatingexpenses")

		var OperatingIncome = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "OperatingIncome")
		var OperatingIncomeAsReported *float64 = nil
		if GrossProfitAsReported != nil && OperatingExpensesAsReported != nil {
			OperatingIncomeAsReported = utils.InterfaceToFloat64Ptr(*GrossProfitAsReported - *OperatingExpensesAsReported)
		}

		var DepreciationAndAmortization = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "DepreciationAndAmortization")
		var DepreciationAndAmortizationAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Depreciationdepletionandamortization")

		var TotalInterestPaymentsAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Interestpaidnet")

		var TotalTaxesPaid = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "DeferredIncomeTax")
		var TotalTaxesPaidAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Incometaxespaidnet")

		var ChangeInWorkingCapital = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "ChangeInWorkingCapital")

		var CapitalExpenditures = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "CapitalExpenditure")
		var CapitalExpendituresAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Paymentstoacquirepropertyplantandequipment")

		var NetCashOperatingActivitiesAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Netcashprovidedbyusedinoperatingactivities")
		var NetCashInvestingActivitiesAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "Netcashprovidedbyusedininvestingactivities")
		var NetCashFinancingActivitiesAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "Netcashprovidedbyusedinfinancingactivities")
		var OperatingCashflow = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "OperatingCashFlow")
		var OperatingCashFlowAsReported *float64 = nil
		if NetCashOperatingActivitiesAsReported != nil && NetCashInvestingActivitiesAsReported != nil && NetCashFinancingActivitiesAsReported != nil {
			OperatingCashFlowAsReported = utils.InterfaceToFloat64Ptr(*NetCashOperatingActivitiesAsReported + *NetCashInvestingActivitiesAsReported + *NetCashFinancingActivitiesAsReported)
		}

		var FreeCashFlow = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "FreeCashFlow")
		var FreeCashFlowAsReported *float64 = nil
		if NetCashOperatingActivitiesAsReported != nil && CapitalExpendituresAsReported != nil {
			FreeCashFlowAsReported = utils.InterfaceToFloat64Ptr(*NetCashOperatingActivitiesAsReported + *CapitalExpendituresAsReported)
		}

		var EBITDA = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "Ebitda")
		var EBITDAAsReported *float64 = nil
		if OperatingIncomeAsReported != nil && DepreciationAndAmortizationAsReported != nil && TotalInterestPaymentsAsReported != nil && TotalTaxesPaidAsReported != nil {
			EBITDAAsReported = utils.InterfaceToFloat64Ptr(*OperatingIncomeAsReported + *DepreciationAndAmortizationAsReported + *TotalInterestPaymentsAsReported + *TotalTaxesPaidAsReported)
		}

		var TotalInterestPayments *float64 = nil
		if EBITDA != nil && NetIncome != nil && TotalTaxesPaid != nil && DepreciationAndAmortization != nil {
			TotalInterestPayments = utils.InterfaceToFloat64Ptr(*EBITDA - *NetIncome - *TotalTaxesPaid - *DepreciationAndAmortization)
		}

		var EBIT *float64 = nil
		var EBITAsReported *float64 = nil
		if EBITDA != nil && DepreciationAndAmortization != nil {
			EBIT = utils.InterfaceToFloat64Ptr(*EBITDA - *DepreciationAndAmortization)
		}
		if EBITDAAsReported != nil && DepreciationAndAmortizationAsReported != nil {
			EBITAsReported = utils.InterfaceToFloat64Ptr(*EBITDAAsReported - *DepreciationAndAmortizationAsReported)
		}

		var NonCashCharges = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "OtherNonCashItems")
		var NonCashChargesAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Othernoncashincomeexpense")

		var MarketValueOfEquity *float64 = nil
		var MarketValueOfEquityAsReported *float64 = nil
		if SharesOutstanding != nil && PricePerShare != 0 {
			MarketValueOfEquity = utils.InterfaceToFloat64Ptr(*SharesOutstanding * PricePerShare)
		}
		if SharesOutstandingAsReported != nil && PricePerShare != 0 {
			MarketValueOfEquityAsReported = utils.InterfaceToFloat64Ptr(*SharesOutstandingAsReported * PricePerShare)
		}

		var ShortTermDebt = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "ShortTermDebt")
		var LongTermDebt = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "LongTermDebt")
		var TotalDebt *float64 = nil
		if ShortTermDebt != nil && LongTermDebt != nil {
			TotalDebt = utils.InterfaceToFloat64Ptr(*ShortTermDebt + *LongTermDebt)
		}
		var ShortTermDebtAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Othershorttermborrowings")
		var TotalDebtAsReported *float64 = nil
		if ShortTermDebtAsReported != nil && CurrentLongTermDebtAsReported != nil && NonCurrentLongTermDebtAsReported != nil {
			TotalDebtAsReported = utils.InterfaceToFloat64Ptr(*ShortTermDebtAsReported + *CurrentLongTermDebtAsReported + *NonCurrentLongTermDebtAsReported)
		}

		var CostOfDebt *float64 = nil
		var CostOfDebtAsReported *float64 = nil
		if TotalDebt != nil && TotalInterestPayments != nil {
			CostOfDebt = utils.InterfaceToFloat64Ptr(*TotalInterestPayments / *TotalDebt)
		}
		if TotalDebtAsReported != nil && TotalInterestPaymentsAsReported != nil {
			CostOfDebtAsReported = utils.InterfaceToFloat64Ptr(*TotalInterestPaymentsAsReported / *TotalDebtAsReported)
		}

		var UnleveredFirmValue *float64 = nil
		var UnleveredFirmValueAsReported *float64 = nil
		if EBIT != nil && DepreciationAndAmortization != nil {
			UnleveredFirmValue = utils.InterfaceToFloat64Ptr((*EBIT * (1 - *EffectiveTaxRate)) + *DepreciationAndAmortization)
		}
		if EBITAsReported != nil && DepreciationAndAmortizationAsReported != nil {
			UnleveredFirmValueAsReported = utils.InterfaceToFloat64Ptr((*EBITAsReported * (1 - *EffectiveTaxRate)) + *DepreciationAndAmortizationAsReported)
		}

		var TaxShieldBenefits *float64 = nil
		var TaxShieldBenefitsAsReported *float64 = nil
		if TotalInterestPayments != nil {
			TaxShieldBenefits = utils.InterfaceToFloat64Ptr(*TotalInterestPayments * *EffectiveTaxRate)
		}
		if TotalInterestPaymentsAsReported != nil {
			TaxShieldBenefitsAsReported = utils.InterfaceToFloat64Ptr(*TotalInterestPaymentsAsReported * *EffectiveTaxRate)
		}

		var NetEffectOfDebt *float64 = nil
		var NetEffectOfDebtAsReported *float64 = nil
		if TaxShieldBenefits != nil && TotalInterestPayments != nil {
			NetEffectOfDebt = utils.InterfaceToFloat64Ptr(*TaxShieldBenefits - *TotalInterestPayments)
		}
		if TaxShieldBenefitsAsReported != nil && TotalInterestPaymentsAsReported != nil {
			NetEffectOfDebtAsReported = utils.InterfaceToFloat64Ptr(*TaxShieldBenefitsAsReported - *TotalInterestPaymentsAsReported)
		}

		var DebtService = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "DebtService")
		var DebtServiceAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Debtservice")

		var NonInterestExpenses *float64 = nil
		var NonInterestExpensesAsReported *float64 = nil
		if OperatingExpenses != nil && TotalInterestPayments != nil {
			NonInterestExpenses = utils.InterfaceToFloat64Ptr(*OperatingExpenses - *TotalInterestPayments)
		}
		if OperatingExpensesAsReported != nil && TotalInterestPaymentsAsReported != nil {
			NonInterestExpensesAsReported = utils.InterfaceToFloat64Ptr(*OperatingExpensesAsReported - *TotalInterestPaymentsAsReported)
		}

		var MarketCapitalization *float64 = nil
		var MarketCapitalizationAsReported *float64 = nil
		if SharesOutstanding != nil && PricePerShare != 0 {
			MarketCapitalization = utils.InterfaceToFloat64Ptr(*SharesOutstanding * PricePerShare)
		}
		if SharesOutstandingAsReported != nil && PricePerShare != 0 {
			MarketCapitalizationAsReported = utils.InterfaceToFloat64Ptr(*SharesOutstandingAsReported * PricePerShare)
		}

		var EnterpriseValue *float64 = nil
		var EnterpriseValueAsReported *float64 = nil
		if MarketCapitalization != nil && TotalDebt != nil && CashAndCashEquivalents != nil {
			EnterpriseValue = utils.InterfaceToFloat64Ptr(*MarketCapitalization + *TotalDebt - *CashAndCashEquivalents)
		}
		if MarketCapitalizationAsReported != nil && TotalDebtAsReported != nil && CashAndCashEquivalentsAsReported != nil {
			EnterpriseValueAsReported = utils.InterfaceToFloat64Ptr(*MarketCapitalizationAsReported + *TotalDebtAsReported - *CashAndCashEquivalentsAsReported)
		}

		var DebtOutstanding *float64 = nil
		var DebtOutstandingAsReported *float64 = nil
		if TotalDebt != nil && TotalInterestPayments != nil {
			DebtOutstanding = utils.InterfaceToFloat64Ptr(*TotalDebt - *TotalInterestPayments)
		}
		if TotalDebtAsReported != nil && TotalInterestPaymentsAsReported != nil {
			DebtOutstandingAsReported = utils.InterfaceToFloat64Ptr(*TotalDebtAsReported - *TotalInterestPaymentsAsReported)
		}

		var AssetTurnoverRatio *float64 = nil
		var AssetTurnoverRatioAsReported *float64 = nil
		if NetRevenue != nil && TotalAssets != nil {
			AssetTurnoverRatio = utils.InterfaceToFloat64Ptr(*NetRevenue / *TotalAssets)
		}
		if NetRevenueAsReported != nil && TotalAssetsAsReported != nil {
			AssetTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(*NetRevenueAsReported / *TotalAssetsAsReported)
		}

		var DividendsPaid = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatement, "DividendsPaid")
		var DividendsPaidAsReported = utils.GetFloat64PtrIfNotEmpty(curCashFlowStatementAsReported, "Paymentsofdividends")

		var RetentionRatio *float64 = nil
		var RetentionRatioAsReported *float64 = nil
		if NetIncome != nil && DividendsPaid != nil {
			RetentionRatio = utils.InterfaceToFloat64Ptr((*NetIncome - *DividendsPaid) / *NetIncome)
		}
		if NetIncomeAsReported != nil && DividendsPaidAsReported != nil {
			RetentionRatioAsReported = utils.InterfaceToFloat64Ptr((*NetIncomeAsReported - *DividendsPaidAsReported) / *NetIncomeAsReported)
		}

		var ReturnOnEquity *float64 = nil
		var ReturnOnEquityAsReported *float64 = nil
		if NetIncome != nil && ShareholderEquity != nil {
			ReturnOnEquity = utils.InterfaceToFloat64Ptr(*NetIncome / *ShareholderEquity)
		}
		if NetIncomeAsReported != nil && ShareholderEquityAsReported != nil {
			ReturnOnEquityAsReported = utils.InterfaceToFloat64Ptr(*NetIncomeAsReported / *ShareholderEquityAsReported)
		}

		var CostAndExpenses = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "CostAndExpenses")
		var CostOfRevenue = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "CostOfRevenue")
		var CostOfGoodsSoldAsReported = utils.GetFloat64PtrIfNotEmpty(curIncomeStatementAsReported, "Costofgoodsandservicessold")

		var SellingGeneralAndAdministrativeExpenses = utils.GetFloat64PtrIfNotEmpty(curIncomeStatement, "SellingGeneralAndAdministrativeExpenses")
		var SellingGeneralAndAdministrativeExpensesAsReported = utils.GetFloat64PtrIfNotEmpty(curIncomeStatementAsReported, "Sellinggeneralandadministrativeexpense")

		var ExplicitCosts *float64 = nil
		var ExplicitCostsAsReported *float64 = nil
		if CostAndExpenses != nil && CostOfRevenue != nil && OperatingExpenses != nil && SellingGeneralAndAdministrativeExpenses != nil {
			ExplicitCosts = utils.InterfaceToFloat64Ptr(*CostAndExpenses + *CostOfRevenue + *OperatingExpenses + *SellingGeneralAndAdministrativeExpenses)
		}
		if CostOfGoodsSoldAsReported != nil && OperatingExpensesAsReported != nil && SellingGeneralAndAdministrativeExpensesAsReported != nil {
			ExplicitCostsAsReported = utils.InterfaceToFloat64Ptr(*CostOfGoodsSoldAsReported + *OperatingExpensesAsReported + *SellingGeneralAndAdministrativeExpensesAsReported)
		}

		var DaysInventoryOutstanding *float64 = nil
		var DaysInventoryOutstandingAsReported *float64 = nil
		if Inventory != nil && CostOfRevenue != nil && DaysInPeriod != nil {
			DaysInventoryOutstanding = utils.InterfaceToFloat64Ptr((*Inventory / *CostOfRevenue) * *DaysInPeriod)
		}
		if InventoryAsReported != nil && CostOfGoodsSoldAsReported != nil && DaysInPeriod != nil {
			DaysInventoryOutstandingAsReported = utils.InterfaceToFloat64Ptr((*InventoryAsReported / *CostOfGoodsSoldAsReported) * *DaysInPeriod)
		}

		var TotalCapital *float64 = nil
		var TotalCapitalAsReported *float64 = nil
		if LongTermDebt != nil && ShortTermDebt != nil && ShareholderEquity != nil {
			TotalCapital = utils.InterfaceToFloat64Ptr(*LongTermDebt + *ShortTermDebt + *ShareholderEquity)
		}
		if TotalDebtAsReported != nil && ShareholderEquityAsReported != nil {
			TotalCapitalAsReported = utils.InterfaceToFloat64Ptr(*TotalDebtAsReported + *ShareholderEquityAsReported)
		}

		var NetMargin *float64 = nil
		var NetMarginAsReported *float64 = nil
		if NetRevenue != nil && CostOfRevenue != nil {
			NetMargin = utils.InterfaceToFloat64Ptr((*NetRevenue - *CostOfRevenue) / *NetRevenue)
		}
		if NetRevenueAsReported != nil && CostOfGoodsSoldAsReported != nil {
			NetMarginAsReported = utils.InterfaceToFloat64Ptr((*NetRevenueAsReported - *CostOfGoodsSoldAsReported) / *NetRevenueAsReported)
		}

		var FreeCashFlowToEquity *float64 = nil
		if EBITDA != nil && DepreciationAndAmortization != nil && TotalInterestPayments != nil && TotalTaxesPaid != nil && ChangeInWorkingCapital != nil && CapitalExpenditures != nil && NetDebt != nil {
			FreeCashFlowToEquity = utils.InterfaceToFloat64Ptr(Calculations.FreeCashFlowToEquity(*EBITDA, *DepreciationAndAmortization, *TotalInterestPayments, *TotalTaxesPaid, *ChangeInWorkingCapital, *CapitalExpenditures, *NetDebt))
		}

		var AdjustedPresentValue *float64 = nil
		var AdjustedPresentValueAsReported *float64 = nil
		if UnleveredFirmValue != nil && NetEffectOfDebt != nil {
			AdjustedPresentValue = utils.InterfaceToFloat64Ptr(Calculations.AdjustedPresentValue(*UnleveredFirmValue, *NetEffectOfDebt))
		}
		if UnleveredFirmValueAsReported != nil && NetEffectOfDebtAsReported != nil {
			AdjustedPresentValueAsReported = utils.InterfaceToFloat64Ptr(Calculations.AdjustedPresentValue(*UnleveredFirmValueAsReported, *NetEffectOfDebtAsReported))
		}

		var InterestCoverageRatio *float64 = nil
		var InterestCoverageRatioAsReported *float64 = nil
		if EBIT != nil && TotalInterestPayments != nil {
			InterestCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.InterestCoverageRatio(*EBIT, *TotalInterestPayments))
		}
		if EBITAsReported != nil && TotalInterestPaymentsAsReported != nil {
			InterestCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.InterestCoverageRatio(*EBITAsReported, *TotalInterestPaymentsAsReported))
		}

		var FixedChargeCoverageRatio *float64 = nil
		var FixedChargeCoverageRatioAsReported *float64 = nil
		if EBIT != nil && NetFixedAssets != nil && TotalInterestPayments != nil {
			FixedChargeCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.FixedChargeCoverageRatio(*EBIT, *NetFixedAssets, *TotalInterestPayments))
		}
		if EBITAsReported != nil && NetFixedAssetsAsReported != nil && TotalInterestPaymentsAsReported != nil {
			FixedChargeCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.FixedChargeCoverageRatio(*EBITAsReported, *NetFixedAssetsAsReported, *TotalInterestPaymentsAsReported))
		}

		var DebtServiceCoverageRatio *float64 = nil
		var DebtServiceCoverageRatioAsReported *float64 = nil
		if OperatingIncome != nil && DebtService != nil {
			DebtServiceCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.DebtServiceCoverageRatio(*OperatingIncome, *DebtService))
		}
		if OperatingIncomeAsReported != nil && DebtServiceAsReported != nil {
			DebtServiceCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DebtServiceCoverageRatio(*OperatingIncomeAsReported, *DebtServiceAsReported))
		}

		var AssetCoverageRatio *float64 = nil
		var AssetCoverageRatioAsReported *float64 = nil
		if TotalAssets != nil && ShortTermDebt != nil && TotalDebt != nil {
			AssetCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.AssetCoverageRatio(*TotalAssets, *ShortTermDebt, *TotalDebt))
		}
		if TotalAssetsAsReported != nil && ShortTermDebtAsReported != nil && TotalDebtAsReported != nil {
			AssetCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.AssetCoverageRatio(*TotalAssetsAsReported, *ShortTermDebtAsReported, *TotalDebtAsReported))
		}

		var EBITDAToInterestCoverageRatio *float64 = nil
		var EBITDAToInterestCoverageRatioAsReported *float64 = nil
		if EBITDA == nil && TotalInterestPayments != nil {
			EBITDAToInterestCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToInterestCoverageRatio(*EBITDA, *TotalInterestPayments))
		}
		if EBITDAAsReported != nil && TotalInterestPaymentsAsReported != nil {
			EBITDAToInterestCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToInterestCoverageRatio(*EBITDAAsReported, *TotalInterestPaymentsAsReported))
		}

		var PreferredDividendCoverageRatio *float64 = nil
		var PreferredDividendCoverageRatioAsReported *float64 = nil
		if NetIncome != nil && DividendsPaid != nil {
			PreferredDividendCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.PreferredDividendCoverageRatio(*NetIncome, *DividendsPaid))
		}
		if NetIncomeAsReported != nil && DividendsPaidAsReported != nil {
			PreferredDividendCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PreferredDividendCoverageRatio(*NetIncomeAsReported, *DividendsPaidAsReported))
		}

		var LiquidityCoverageRatio *float64 = nil
		var LiquidityCoverageRatioAsReported *float64 = nil
		if HighQualityLiquidAssets != nil && OperatingCashflow != nil {
			LiquidityCoverageRatio = utils.InterfaceToFloat64Ptr(Calculations.LiquidityCoverageRatio(*HighQualityLiquidAssets, *OperatingCashflow))
		}
		if HighQualityLiquidAssetsAsReported != nil && OperatingCashFlowAsReported != nil {
			LiquidityCoverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.LiquidityCoverageRatio(*HighQualityLiquidAssetsAsReported, *OperatingCashFlowAsReported))
		}

		var InventoryTurnoverRatio *float64 = nil
		var InventoryTurnoverRatioAsReported *float64 = nil
		if CostOfRevenue != nil && Inventory != nil {
			InventoryTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.InventoryTurnoverRatio(*CostOfRevenue, *Inventory))
		}
		if CostOfGoodsSoldAsReported != nil && InventoryAsReported != nil {
			InventoryTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.InventoryTurnoverRatio(*CostOfGoodsSoldAsReported, *InventoryAsReported))
		}

		var ReturnOnCapitalEmployed *float64 = nil
		var ReturnOnCapitalEmployedAsReported *float64 = nil
		if EBIT != nil && TotalAssets != nil && TotalLiabilities != nil {
			ReturnOnCapitalEmployed = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnCapitalEmployed(*EBIT, *TotalAssets, *TotalLiabilities))
		}
		if EBITAsReported != nil && TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			ReturnOnCapitalEmployedAsReported = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnCapitalEmployed(*EBITAsReported, *TotalAssetsAsReported, *TotalLiabilitiesAsReported))
		}

		var EfficiencyRatio *float64 = nil
		var EfficiencyRatioAsReported *float64 = nil
		if NonInterestExpenses != nil && NetRevenue != nil {
			EfficiencyRatio = utils.InterfaceToFloat64Ptr(Calculations.EfficiencyRatio(*NonInterestExpenses, *NetRevenue))
		}
		if NonInterestExpensesAsReported != nil && NetRevenueAsReported != nil {
			EfficiencyRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EfficiencyRatio(*NonInterestExpensesAsReported, *NetRevenueAsReported))
		}

		var RevenuePerEmployee *float64 = nil
		var RevenuePerEmployeeAsReported *float64 = nil
		if NetRevenue != nil && NumEmployees != 0 {
			RevenuePerEmployee = utils.InterfaceToFloat64Ptr(Calculations.RevenuePerEmployee(*NetRevenue, NumEmployees))
		}
		if NetRevenueAsReported != nil && NumEmployees != 0 {
			RevenuePerEmployeeAsReported = utils.InterfaceToFloat64Ptr(Calculations.RevenuePerEmployee(*NetRevenueAsReported, NumEmployees))
		}

		var CapitalExpenditureRatio *float64 = nil
		var CapitalExpenditureRatioAsReported *float64 = nil
		if CapitalExpenditures != nil && OperatingCashflow != nil {
			CapitalExpenditureRatio = utils.InterfaceToFloat64Ptr(Calculations.CapitalExpenditureRatio(*CapitalExpenditures, *OperatingCashflow))
		}
		if CapitalExpendituresAsReported != nil && OperatingCashFlowAsReported != nil {
			CapitalExpenditureRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CapitalExpenditureRatio(*CapitalExpendituresAsReported, *OperatingCashFlowAsReported))
		}

		var OperatingCashFlowRatio *float64 = nil
		var OperatingCashFlowRatioAsReported *float64 = nil
		if OperatingCashflow != nil && NetRevenue != nil {
			OperatingCashFlowRatio = utils.InterfaceToFloat64Ptr(Calculations.OperatingCashFlowRatio(*OperatingCashflow, *NetRevenue))
		}
		if OperatingCashFlowAsReported != nil && NetRevenueAsReported != nil {
			OperatingCashFlowRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingCashFlowRatio(*OperatingCashFlowAsReported, *NetRevenueAsReported))
		}

		var EBITDAToEVRatio *float64 = nil
		var EBITDAToEVRatioAsReported *float64 = nil
		if EBITDA != nil && EnterpriseValue != nil {
			EBITDAToEVRatio = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToEVRatio(*EBITDA, *EnterpriseValue))
		}
		if EBITDAAsReported != nil && EnterpriseValueAsReported != nil {
			EBITDAToEVRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAToEVRatio(*EBITDAAsReported, *EnterpriseValueAsReported))
		}

		var TangibleNetWorthRatio *float64 = nil
		var TangibleNetWorthRatioAsReported *float64 = nil
		if TangibleNetWorth != nil && TotalAssets != nil {
			TangibleNetWorthRatio = utils.InterfaceToFloat64Ptr(Calculations.TangibleNetWorthRatio(*TangibleNetWorth, *TotalAssets))
		}
		if TangibleNetWorthAsReported != nil && TotalAssetsAsReported != nil {
			TangibleNetWorthRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.TangibleNetWorthRatio(*TangibleNetWorthAsReported, *TotalAssetsAsReported))
		}

		var DeferredTaxLiabilityToEquityRatio *float64 = nil
		var DeferredTaxLiabilityToEquityRatioAsReported *float64 = nil
		if DeferredTaxLiabilityToEquityRatio != nil && ShareholderEquity != nil {
			DeferredTaxLiabilityToEquityRatio = utils.InterfaceToFloat64Ptr(Calculations.DeferredTaxLiabilityToEquityRatio(*DeferredTaxLiabilityToEquityRatio, *ShareholderEquity))
		}
		if DeferredTaxLiabilityToEquityRatioAsReported != nil && ShareholderEquityAsReported != nil {
			DeferredTaxLiabilityToEquityRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DeferredTaxLiabilityToEquityRatio(*DeferredTaxLiabilityToEquityRatioAsReported, *ShareholderEquityAsReported))
		}

		var TangibleEquityRatio *float64 = nil
		var TangibleEquityRatioAsReported *float64 = nil
		if ShareholderEquity != nil && IntangibleAssets != nil && TotalAssets != nil {
			TangibleEquityRatio = utils.InterfaceToFloat64Ptr(Calculations.TangibleEquityRatio(*ShareholderEquity, *IntangibleAssets, *TotalAssets))
		}
		if ShareholderEquityAsReported != nil && IntangibleAssetsAsReported != nil && TotalAssetsAsReported != nil {
			TangibleEquityRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.TangibleEquityRatio(*ShareholderEquityAsReported, *IntangibleAssetsAsReported, *TotalAssetsAsReported))
		}

		var WACC *float64 = nil
		var WACCAsReported *float64 = nil
		if MarketValueOfEquity != nil && TotalDebt != nil && CostOfDebt != nil {
			WACC = utils.InterfaceToFloat64Ptr(Calculations.WeightedAverageCostOfCapital(*MarketValueOfEquity, *TotalDebt, CostOfEquity, *CostOfDebt, *EffectiveTaxRate))
		}
		if MarketValueOfEquityAsReported != nil && TotalDebtAsReported != nil && CostOfDebtAsReported != nil {
			WACCAsReported = utils.InterfaceToFloat64Ptr(Calculations.WeightedAverageCostOfCapital(*MarketValueOfEquityAsReported, *TotalDebtAsReported, CostOfEquity, *CostOfDebtAsReported, *EffectiveTaxRate))
		}

		var FixedAssetTurnoverRatio *float64 = nil
		var FixedAssetTurnoverRatioAsReported *float64 = nil
		if NetRevenue != nil && NetFixedAssets != nil {
			FixedAssetTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.FixedAssetTurnoverRatio(*NetRevenue, *NetFixedAssets))
		}
		if NetRevenueAsReported != nil && NetFixedAssetsAsReported != nil {
			FixedAssetTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.FixedAssetTurnoverRatio(*NetRevenueAsReported, *NetFixedAssetsAsReported))
		}

		var PPETurnoverRatio *float64 = nil
		var PPETurnoverRatioAsReported *float64 = nil
		if NetRevenue != nil && NetFixedAssets != nil && DepreciationAndAmortization != nil {
			PPETurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.PPETurnoverRatio(*NetRevenue, *NetFixedAssets, *DepreciationAndAmortization))
		}
		if NetRevenueAsReported != nil && NetFixedAssetsAsReported != nil && DepreciationAndAmortizationAsReported != nil {
			PPETurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PPETurnoverRatio(*NetRevenueAsReported, *NetFixedAssetsAsReported, *DepreciationAndAmortizationAsReported))
		}

		var InvestmentTurnoverRatio *float64 = nil
		var InvestmentTurnoverRatioAsReported *float64 = nil
		if NetRevenue != nil && TotalInvestments != nil {
			InvestmentTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.InvestmentTurnoverRatio(*NetRevenue, *TotalInvestments))
		}
		if NetRevenueAsReported != nil && TotalInvestmentsAsReported != nil {
			InvestmentTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.InvestmentTurnoverRatio(*NetRevenueAsReported, *TotalInvestmentsAsReported))
		}

		var WorkingCapitalTurnoverRatio *float64 = nil
		var WorkingCapitalTurnoverRatioAsReported *float64 = nil
		if NetRevenue != nil && WorkingCapital != nil {
			WorkingCapitalTurnoverRatio = utils.InterfaceToFloat64Ptr(Calculations.WorkingCapitalTurnoverRatio(*NetRevenue, *WorkingCapital))
		}
		if NetRevenueAsReported != nil && WorkingCapitalAsReported != nil {
			WorkingCapitalTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.WorkingCapitalTurnoverRatio(*NetRevenueAsReported, *WorkingCapitalAsReported))
		}

		var ReturnOnAssetRatio *float64 = nil
		var ReturnOnAssetRatioAsReported *float64 = nil
		if NetIncome != nil && TotalAssets != nil {
			ReturnOnAssetRatio = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnAssetRatio(*NetIncome, *TotalAssets))
		}
		if NetIncomeAsReported != nil && TotalAssetsAsReported != nil {
			ReturnOnAssetRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnAssetRatio(*NetIncomeAsReported, *TotalAssetsAsReported))
		}

		var GrossProfitMargin *float64 = nil
		var GrossProfitMarginAsReported *float64 = nil
		if GrossProfit != nil && NetRevenue != nil {
			GrossProfitMargin = utils.InterfaceToFloat64Ptr(Calculations.GrossProfitMargin(*GrossProfit, *NetRevenue))
		}
		if GrossProfitAsReported != nil && NetRevenueAsReported != nil {
			GrossProfitMarginAsReported = utils.InterfaceToFloat64Ptr(Calculations.GrossProfitMargin(*GrossProfitAsReported, *NetRevenueAsReported))
		}

		var OperatingProfitMargin *float64 = nil
		var OperatingProfitMarginAsReported *float64 = nil
		if OperatingIncome != nil && NetRevenue != nil {
			OperatingProfitMargin = utils.InterfaceToFloat64Ptr(Calculations.OperatingProfitMargin(*OperatingIncome, *NetRevenue))
		}
		if OperatingIncomeAsReported != nil && NetRevenueAsReported != nil {
			OperatingProfitMarginAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingProfitMargin(*OperatingIncomeAsReported, *NetRevenueAsReported))
		}

		var EBITDAMarginRatio *float64 = nil
		var EBITDAMarginRatioAsReported *float64 = nil
		if EBITDA != nil && NetRevenue != nil {
			EBITDAMarginRatio = utils.InterfaceToFloat64Ptr(Calculations.EBITDAMarginRatio(*EBITDA, *NetRevenue))
		}
		if EBITDAAsReported != nil && NetRevenueAsReported != nil {
			EBITDAMarginRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAMarginRatio(*EBITDAAsReported, *NetRevenueAsReported))
		}

		var DividendPayoutRatio *float64 = nil
		var DividendPayoutRatioAsReported *float64 = nil
		if DividendsPaid != nil && NetIncome != nil {
			DividendPayoutRatio = utils.InterfaceToFloat64Ptr(Calculations.DividendPayoutRatio(*DividendsPaid, *NetIncome))
		}
		if DividendsPaidAsReported != nil && NetIncomeAsReported != nil {
			DividendPayoutRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DividendPayoutRatio(*DividendsPaidAsReported, *NetIncomeAsReported))
		}

		var RetentionRate *float64 = nil
		var RetentionRateAsReported *float64 = nil
		if DividendsPaid != nil && NetIncome != nil {
			RetentionRate = utils.InterfaceToFloat64Ptr(Calculations.RetentionRate(*DividendsPaid, *NetIncome))
		}
		if DividendsPaidAsReported != nil && NetIncomeAsReported != nil {
			RetentionRateAsReported = utils.InterfaceToFloat64Ptr(Calculations.RetentionRate(*DividendsPaidAsReported, *NetIncomeAsReported))
		}

		var SustainableGrowthRate *float64 = nil
		var SustainableGrowthRateAsReported *float64 = nil
		if RetentionRate != nil && ReturnOnEquity != nil {
			SustainableGrowthRate = utils.InterfaceToFloat64Ptr(Calculations.SustainableGrowthRate(*RetentionRate, *ReturnOnEquity))
		}
		if RetentionRateAsReported != nil && ReturnOnEquityAsReported != nil {
			SustainableGrowthRateAsReported = utils.InterfaceToFloat64Ptr(Calculations.SustainableGrowthRate(*RetentionRateAsReported, *ReturnOnEquityAsReported))
		}

		var GrossMarginOnInventory *float64 = nil
		var GrossMarginOnInventoryAsReported *float64 = nil
		if GrossProfit != nil && Inventory != nil {
			GrossMarginOnInventory = utils.InterfaceToFloat64Ptr(Calculations.GrossMarginOnInventory(*GrossProfit, *Inventory))
		}
		if GrossProfitAsReported != nil && InventoryAsReported != nil {
			GrossMarginOnInventoryAsReported = utils.InterfaceToFloat64Ptr(Calculations.GrossMarginOnInventory(*GrossProfitAsReported, *InventoryAsReported))
		}

		var CashFlowReturnOnEquity *float64 = nil
		var CashFlowReturnOnEquityAsReported *float64 = nil
		if OperatingCashflow != nil && ShareholderEquity != nil {
			CashFlowReturnOnEquity = utils.InterfaceToFloat64Ptr(Calculations.CashFlowReturnOnEquity(*OperatingCashflow, *ShareholderEquity))
		}
		if OperatingCashFlowAsReported != nil && ShareholderEquityAsReported != nil {
			CashFlowReturnOnEquityAsReported = utils.InterfaceToFloat64Ptr(Calculations.CashFlowReturnOnEquity(*OperatingCashFlowAsReported, *ShareholderEquityAsReported))
		}

		var OperatingMargin *float64 = nil
		var OperatingMarginAsReported *float64 = nil
		if NetRevenue != nil && CostOfRevenue != nil {
			OperatingMargin = utils.InterfaceToFloat64Ptr(Calculations.OperatingMargin(*NetRevenue, *CostOfRevenue))
		}
		if NetRevenueAsReported != nil && CostOfGoodsSoldAsReported != nil {
			OperatingMarginAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingMargin(*NetRevenueAsReported, *CostOfGoodsSoldAsReported))
		}

		var OperatingExpenseRatio *float64 = nil
		var OperatingExpenseRatioAsReported *float64 = nil
		if OperatingExpenses != nil && DepreciationAndAmortization != nil && OperatingIncome != nil {
			OperatingExpenseRatio = utils.InterfaceToFloat64Ptr(Calculations.OperatingExpenseRatio(*OperatingExpenses, *DepreciationAndAmortization, *OperatingIncome))
		}
		if OperatingExpensesAsReported != nil && DepreciationAndAmortizationAsReported != nil && OperatingIncomeAsReported != nil {
			OperatingExpenseRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.OperatingExpenseRatio(*OperatingExpensesAsReported, *DepreciationAndAmortizationAsReported, *OperatingIncomeAsReported))
		}

		var CurrentRatio *float64 = nil
		var CurrentRatioAsReported *float64 = nil
		if TotalAssets != nil && TotalLiabilities != nil {
			CurrentRatio = utils.InterfaceToFloat64Ptr(Calculations.CurrentRatio(*TotalAssets, *TotalLiabilities))
		}
		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil {
			CurrentRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CurrentRatio(*TotalAssetsAsReported, *TotalLiabilitiesAsReported))
		}

		var AcidTestRatio *float64 = nil
		var AcidTestRatioAsReported *float64 = nil
		if TotalAssets != nil && TotalLiabilities != nil && Inventory != nil {
			AcidTestRatio = utils.InterfaceToFloat64Ptr(Calculations.AcidTestRatio(*TotalAssets, *Inventory, *TotalLiabilities))
		}
		if TotalAssetsAsReported != nil && TotalLiabilitiesAsReported != nil && InventoryAsReported != nil {
			AcidTestRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.AcidTestRatio(*TotalAssetsAsReported, *InventoryAsReported, *TotalLiabilitiesAsReported))
		}

		var CashRatio *float64 = nil
		var CashRatioAsReported *float64 = nil
		if CashAndCashEquivalents != nil && TotalLiabilities != nil {
			CashRatio = utils.InterfaceToFloat64Ptr(Calculations.CashRatio(*CashAndCashEquivalents, *TotalLiabilities))
		}
		if CashAndCashEquivalentsAsReported != nil && TotalLiabilitiesAsReported != nil {
			CashRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CashRatio(*CashAndCashEquivalentsAsReported, *TotalLiabilitiesAsReported))
		}

		var DefensiveIntervalRatio *float64 = nil
		var DefensiveIntervalRatioAsReported *float64 = nil
		if CashAndCashEquivalents != nil && NetReceivables != nil && TotalMarketableSecurities != nil && OperatingExpenses != nil && NonCashCharges != nil && DaysInPeriod != nil {
			DefensiveIntervalRatio = utils.InterfaceToFloat64Ptr(Calculations.DefensiveIntervalRatio(*CashAndCashEquivalents, *NetReceivables, *TotalMarketableSecurities, *OperatingExpenses, *NonCashCharges, *DaysInPeriod))
		}
		if CashAndCashEquivalentsAsReported != nil && NetReceivablesAsReported != nil && TotalMarketableSecuritiesAsReported != nil && OperatingExpensesAsReported != nil && NonCashChargesAsReported != nil && DaysInPeriod != nil {
			DefensiveIntervalRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DefensiveIntervalRatio(*CashAndCashEquivalentsAsReported, *NetReceivablesAsReported, *TotalMarketableSecuritiesAsReported, *OperatingExpensesAsReported, *NonCashChargesAsReported, *DaysInPeriod))
		}

		var DrySalesRatio *float64 = nil
		var DrySalesRatioAsReported *float64 = nil
		if NetReceivables != nil && NetRevenue != nil {
			DrySalesRatio = utils.InterfaceToFloat64Ptr(Calculations.DrySalesRatio(*NetReceivables, *NetRevenue))
		}
		if NetReceivablesAsReported != nil && NetRevenueAsReported != nil {
			DrySalesRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DrySalesRatio(*NetReceivablesAsReported, *NetRevenueAsReported))
		}

		var PriceToBookValueRatio *float64 = nil
		var PriceToBookValueRatioAsReported *float64 = nil
		if MarketCapitalization != nil && BookValueOfEquity != nil {
			PriceToBookValueRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookValueRatio(*MarketCapitalization, *BookValueOfEquity))
		}
		if MarketCapitalizationAsReported != nil && BookValueOfEquityAsReported != nil {
			PriceToBookValueRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookValueRatio(*MarketCapitalizationAsReported, *BookValueOfEquityAsReported))
		}

		var EarningsPerShare *float64 = nil
		var EarningsPerShareAsReported *float64 = nil
		if NetIncome != nil && DividendsPaid != nil && SharesOutstanding != nil {
			EarningsPerShare = utils.InterfaceToFloat64Ptr(Calculations.EarningsPerShare(*NetIncome, *DividendsPaid, *SharesOutstanding))
		}
		if NetIncomeAsReported != nil && DividendsPaidAsReported != nil && SharesOutstandingAsReported != nil {
			EarningsPerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.EarningsPerShare(*NetIncomeAsReported, *DividendsPaidAsReported, *SharesOutstandingAsReported))
		}

		var EBITDAPerShare *float64 = nil
		var EBITDAPerShareAsReported *float64 = nil
		if EBITDA != nil && SharesOutstanding != nil {
			EBITDAPerShare = utils.InterfaceToFloat64Ptr(Calculations.EBITDAPerShare(*EBITDA, *SharesOutstanding))
		}
		if EBITDAAsReported != nil && SharesOutstandingAsReported != nil {
			EBITDAPerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.EBITDAPerShare(*EBITDAAsReported, *SharesOutstandingAsReported))
		}

		var BookValuePerShare *float64 = nil
		var BookValuePerShareAsReported *float64 = nil
		if ShareholderEquity != nil && SharesOutstanding != nil {
			BookValuePerShare = utils.InterfaceToFloat64Ptr(Calculations.BookValuePerShare(*ShareholderEquity, *SharesOutstanding))
		}
		if ShareholderEquityAsReported != nil && SharesOutstandingAsReported != nil {
			BookValuePerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.BookValuePerShare(*ShareholderEquityAsReported, *SharesOutstandingAsReported))
		}

		var NetTangibleAssetsPerShare *float64 = nil
		var NetTangibleAssetsPerShareAsReported *float64 = nil
		if TangibleNetWorth != nil && SharesOutstanding != nil {
			NetTangibleAssetsPerShare = utils.InterfaceToFloat64Ptr(Calculations.NetTangibleAssetsPerShare(*TangibleNetWorth, *SharesOutstanding))
		}
		if TangibleNetWorthAsReported != nil && SharesOutstandingAsReported != nil {
			NetTangibleAssetsPerShareAsReported = utils.InterfaceToFloat64Ptr(Calculations.NetTangibleAssetsPerShare(*TangibleNetWorthAsReported, *SharesOutstandingAsReported))
		}

		var MarketValueOfDebt *float64 = nil
		var MarketValueOfDebtAsReported *float64 = nil
		if PricePerShare != 0 && SharesOutstanding != nil && BookValueOfDebt != nil {
			MarketValueOfDebt = utils.InterfaceToFloat64Ptr(Calculations.MarketValueOfDebt(PricePerShare, *SharesOutstanding, *BookValueOfDebt))
		}
		if PricePerShare != 0 && SharesOutstandingAsReported != nil && BookValueOfDebtAsReported != nil {
			MarketValueOfDebtAsReported = utils.InterfaceToFloat64Ptr(Calculations.MarketValueOfDebt(PricePerShare, *SharesOutstandingAsReported, *BookValueOfDebtAsReported))
		}

		var MarketToBookRatio *float64 = nil
		var MarketToBookRatioAsReported *float64 = nil
		if MarketCapitalization != nil && BookValueOfEquity != nil {
			MarketToBookRatio = utils.InterfaceToFloat64Ptr(Calculations.MarketToBookRatio(*MarketCapitalization, *BookValueOfEquity))
		}
		if MarketCapitalizationAsReported != nil && BookValueOfEquityAsReported != nil {
			MarketToBookRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.MarketToBookRatio(*MarketCapitalizationAsReported, *BookValueOfEquityAsReported))
		}

		var IntangiblesRatio *float64 = nil
		var IntangiblesRatioAsReported *float64 = nil
		if IntangibleAssets != nil && TotalAssets != nil {
			IntangiblesRatio = utils.InterfaceToFloat64Ptr(Calculations.IntangiblesRatio(*IntangibleAssets, *TotalAssets))
		}
		if IntangibleAssetsAsReported != nil && TotalAssetsAsReported != nil {
			IntangiblesRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.IntangiblesRatio(*IntangibleAssetsAsReported, *TotalAssetsAsReported))
		}

		var PriceToSalesRatio *float64 = nil
		var PriceToSalesRatioAsReported *float64 = nil
		if PricePerShare != 0 && NetRevenue != nil {
			PriceToSalesRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToSalesRatio(PricePerShare, *NetRevenue))
		}
		if PricePerShare != 0 && NetRevenueAsReported != nil {
			PriceToSalesRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToSalesRatio(PricePerShare, *NetRevenueAsReported))
		}

		var PriceToBookRatio *float64 = nil
		var PriceToBookRatioAsReported *float64 = nil
		if PricePerShare != 0 && BookValueOfEquity != nil {
			PriceToBookRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookRatio(PricePerShare, *BookValueOfEquity))
		}
		if PricePerShare != 0 && BookValueOfEquityAsReported != nil {
			PriceToBookRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToBookRatio(PricePerShare, *BookValueOfEquityAsReported))
		}

		var PricetoSalesValue *float64 = nil
		var PricetoSalesValueAsReported *float64 = nil
		if MarketCapitalization != nil && NetRevenue != nil {
			PricetoSalesValue = utils.InterfaceToFloat64Ptr(Calculations.PricetoSalesValue(*MarketCapitalization, *NetRevenue))
		}
		if MarketCapitalizationAsReported != nil && NetRevenueAsReported != nil {
			PricetoSalesValueAsReported = utils.InterfaceToFloat64Ptr(Calculations.PricetoSalesValue(*MarketCapitalizationAsReported, *NetRevenueAsReported))
		}

		var OperatingCashFlowPerShare *float64 = nil
		var OperatingCashFlowPerShareAsReported *float64 = nil
		if OperatingCashflow != nil && SharesOutstanding != nil {
			OperatingCashFlowPerShare = utils.InterfaceToFloat64Ptr(*OperatingCashflow / *SharesOutstanding)
		}
		if OperatingCashFlowAsReported != nil && SharesOutstandingAsReported != nil {
			OperatingCashFlowPerShareAsReported = utils.InterfaceToFloat64Ptr(*OperatingCashFlowAsReported / *SharesOutstandingAsReported)
		}

		var PriceToCashFlowRatio *float64 = nil
		var PriceToCashFlowRatioAsReported *float64 = nil
		if PricePerShare != 0 && OperatingCashflow != nil {
			PriceToCashFlowRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowRatio(PricePerShare, *OperatingCashflow))
		}
		if PricePerShare != 0 && OperatingCashFlowAsReported != nil {
			PriceToCashFlowRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowRatio(PricePerShare, *OperatingCashFlowAsReported))
		}

		var FreeCashFlowPerShare *float64 = nil
		var FreeCashFlowPerShareAsReported *float64 = nil
		if FreeCashFlow != nil && SharesOutstanding != nil {
			FreeCashFlowPerShare = utils.InterfaceToFloat64Ptr(*FreeCashFlow / *SharesOutstanding)
		}
		if FreeCashFlowAsReported != nil && SharesOutstandingAsReported != nil {
			FreeCashFlowPerShareAsReported = utils.InterfaceToFloat64Ptr(*FreeCashFlowAsReported / *SharesOutstandingAsReported)
		}

		var PriceToFreeCashFlowRatio *float64 = nil
		var PriceToFreeCashFlowRatioAsReported *float64 = nil
		if PricePerShare != 0 && FreeCashFlowPerShare != nil {
			PriceToFreeCashFlowRatio = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowRatio(PricePerShare, *FreeCashFlowPerShare))
		}
		if PricePerShare != 0 && FreeCashFlowPerShareAsReported != nil {
			PriceToFreeCashFlowRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowRatio(PricePerShare, *FreeCashFlowPerShareAsReported))
		}

		var PriceToCashFlowValuation *float64 = nil
		var PriceToCashFlowValuationAsReported *float64 = nil
		if MarketCapitalization != nil && OperatingCashflow != nil {
			PriceToCashFlowValuation = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowValuation(*MarketCapitalization, *OperatingCashflow))
		}
		if MarketCapitalizationAsReported != nil && OperatingCashFlowAsReported != nil {
			PriceToCashFlowValuationAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToCashFlowValuation(*MarketCapitalizationAsReported, *OperatingCashFlowAsReported))
		}

		var PriceToFreeCashFlowValuation *float64 = nil
		var PriceToFreeCashFlowValuationAsReported *float64 = nil
		if MarketCapitalization != nil && FreeCashFlow != nil {
			PriceToFreeCashFlowValuation = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowValuation(*MarketCapitalization, *FreeCashFlow))
		}
		if MarketCapitalizationAsReported != nil && FreeCashFlowAsReported != nil {
			PriceToFreeCashFlowValuationAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToFreeCashFlowValuation(*MarketCapitalizationAsReported, *FreeCashFlowAsReported))
		}

		var PriceToEarningsValuation *float64 = nil
		var PriceToEarningsValuationAsReported *float64 = nil
		if MarketCapitalization != nil && NetIncome != nil {
			PriceToEarningsValuation = utils.InterfaceToFloat64Ptr(Calculations.PriceToEarningsValuation(*MarketCapitalization, *NetIncome))
		}
		if MarketCapitalizationAsReported != nil && NetIncomeAsReported != nil {
			PriceToEarningsValuationAsReported = utils.InterfaceToFloat64Ptr(Calculations.PriceToEarningsValuation(*MarketCapitalizationAsReported, *NetIncomeAsReported))
		}

		var LiabilitiesMarketValue *float64 = nil
		var LiabilitiesMarketValueAsReported *float64 = nil
		if PricePerShare != 0 && SharesOutstanding != nil && BookValueOfDebt != nil {
			LiabilitiesMarketValue = utils.InterfaceToFloat64Ptr(Calculations.LiabilitiesMarketValue(PricePerShare, *SharesOutstanding, *BookValueOfDebt))
		}
		if PricePerShare != 0 && SharesOutstandingAsReported != nil && BookValueOfDebtAsReported != nil {
			LiabilitiesMarketValueAsReported = utils.InterfaceToFloat64Ptr(Calculations.LiabilitiesMarketValue(PricePerShare, *SharesOutstandingAsReported, *BookValueOfDebtAsReported))
		}

		var TobinsQ *float64 = nil
		var TobinsQAsReported *float64 = nil
		if MarketValueOfEquity != nil && MarketValueOfDebt != nil && BookValueOfEquity != nil && BookValueOfDebt != nil {
			TobinsQ = utils.InterfaceToFloat64Ptr(Calculations.TobinsQ(*MarketValueOfEquity, *MarketValueOfDebt, *BookValueOfEquity, *BookValueOfDebt))
		}
		if MarketValueOfEquityAsReported != nil && MarketValueOfDebtAsReported != nil && BookValueOfEquityAsReported != nil && BookValueOfDebtAsReported != nil {
			TobinsQAsReported = utils.InterfaceToFloat64Ptr(Calculations.TobinsQ(*MarketValueOfEquityAsReported, *MarketValueOfDebtAsReported, *BookValueOfEquityAsReported, *BookValueOfDebtAsReported))
		}

		var ReceivablesTurnoverRatio *float64 = nil
		var ReceivablesTurnoverRatioAsReported *float64 = nil
		if NetRevenue != nil && NetReceivables != nil {
			ReceivablesTurnoverRatio = utils.InterfaceToFloat64Ptr(*NetRevenue / *NetReceivables)
		}
		if NetRevenueAsReported != nil && NetReceivablesAsReported != nil {
			ReceivablesTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(*NetRevenueAsReported / *NetReceivablesAsReported)
		}

		var AverageCollectionPeriod *float64 = nil
		var AverageCollectionPeriodAsReported *float64 = nil
		if ReceivablesTurnoverRatio != nil {
			AverageCollectionPeriod = utils.InterfaceToFloat64Ptr(Calculations.AverageCollectionPeriod(*ReceivablesTurnoverRatio))
		}

		var AccountsPayableTurnoverRatio *float64 = nil
		var AccountsPayableTurnoverRatioAsReported *float64 = nil
		if AccountsPayable != nil && NetRevenue != nil {
			AccountsPayableTurnoverRatio = utils.InterfaceToFloat64Ptr(*NetRevenue / *AccountsPayable)
		}
		if AccountsPayableAsReported != nil && NetRevenueAsReported != nil {
			AccountsPayableTurnoverRatioAsReported = utils.InterfaceToFloat64Ptr(*NetRevenueAsReported / *AccountsPayableAsReported)
		}

		var AverageAccountsPayablePaymentPeriod *float64 = nil
		var AverageAccountsPayablePaymentPeriodAsReported *float64 = nil
		if AccountsPayableTurnoverRatio != nil {
			AverageAccountsPayablePaymentPeriod = utils.InterfaceToFloat64Ptr(Calculations.AverageAccountsPayablePaymentPeriod(*AccountsPayableTurnoverRatio))
		}
		if AccountsPayableTurnoverRatioAsReported != nil {
			AverageAccountsPayablePaymentPeriodAsReported = utils.InterfaceToFloat64Ptr(Calculations.AverageAccountsPayablePaymentPeriod(*AccountsPayableTurnoverRatioAsReported))
		}

		var InventoryToWorkingCapitalRatio *float64 = nil
		var InventoryToWorkingCapitalRatioAsReported *float64 = nil
		if Inventory != nil && WorkingCapital != nil {
			InventoryToWorkingCapitalRatio = utils.InterfaceToFloat64Ptr(*Inventory / *WorkingCapital)
		}
		if InventoryAsReported != nil && WorkingCapitalAsReported != nil {
			InventoryToWorkingCapitalRatioAsReported = utils.InterfaceToFloat64Ptr(*InventoryAsReported / *WorkingCapitalAsReported)
		}

		var DaysSalesOutstanding *float64 = nil
		var DaysSalesOutstandingAsReported *float64 = nil
		if NetRevenue != nil && CostOfRevenue != nil && DaysInPeriod != nil {
			DaysSalesOutstanding = utils.InterfaceToFloat64Ptr((*NetRevenue / *CostOfRevenue) * *DaysInPeriod)
		}
		if NetRevenueAsReported != nil && CostOfGoodsSoldAsReported != nil && DaysInPeriod != nil {
			DaysSalesOutstandingAsReported = utils.InterfaceToFloat64Ptr((*NetRevenueAsReported / *CostOfGoodsSoldAsReported) * *DaysInPeriod)
		}

		var DaysPayablesOutstanding *float64 = nil
		var DaysPayablesOutstandingAsReported *float64 = nil
		if AccountsPayable != nil && CostOfRevenue != nil && DaysInPeriod != nil {
			DaysPayablesOutstanding = utils.InterfaceToFloat64Ptr((*AccountsPayable / *CostOfRevenue) * *DaysInPeriod)
		}
		if AccountsPayableAsReported != nil && CostOfGoodsSoldAsReported != nil && DaysInPeriod != nil {
			DaysPayablesOutstandingAsReported = utils.InterfaceToFloat64Ptr((*AccountsPayableAsReported / *CostOfGoodsSoldAsReported) * *DaysInPeriod)
		}

		var CashConversionCycle *float64 = nil
		var CashConversionCycleAsReported *float64 = nil
		if DaysInventoryOutstanding != nil && DaysSalesOutstanding != nil && DaysPayablesOutstanding != nil {
			CashConversionCycle = utils.InterfaceToFloat64Ptr(Calculations.CashConversionCycle(*DaysInventoryOutstanding, *DaysSalesOutstanding, *DaysPayablesOutstanding))
		}

		var NetWorkingCapital *float64 = nil
		var NetWorkingCapitalAsReported *float64 = nil
		if NetReceivables != nil && Inventory != nil && AccountsPayable != nil {
			NetWorkingCapital = utils.InterfaceToFloat64Ptr(Calculations.NetWorkingCapital(*NetReceivables, *Inventory, *AccountsPayable))
		}
		if NetReceivablesAsReported != nil && InventoryAsReported != nil && AccountsPayableAsReported != nil {
			NetWorkingCapitalAsReported = utils.InterfaceToFloat64Ptr(Calculations.NetWorkingCapital(*NetReceivablesAsReported, *InventoryAsReported, *AccountsPayableAsReported))
		}

		var NOPAT *float64 = nil
		var NOPATAsReported *float64 = nil
		if OperatingIncome != nil && EffectiveTaxRate != nil {
			NOPAT = utils.InterfaceToFloat64Ptr(*OperatingIncome * (1 - *EffectiveTaxRate))
		}
		if OperatingIncomeAsReported != nil && EffectiveTaxRate != nil {
			NOPATAsReported = utils.InterfaceToFloat64Ptr(*OperatingIncomeAsReported * (1 - *EffectiveTaxRate))
		}

		var EconomicValueAdded *float64 = nil
		var EconomicValueAddedAsReported *float64 = nil
		if NOPAT != nil && WACC != nil && TotalCapital != nil {
			EconomicValueAdded = utils.InterfaceToFloat64Ptr(Calculations.EconomicValueAdded(*NOPAT, *WACC, *TotalCapital))
		}
		if NOPATAsReported != nil && WACCAsReported != nil && TotalCapitalAsReported != nil {
			EconomicValueAddedAsReported = utils.InterfaceToFloat64Ptr(Calculations.EconomicValueAdded(*NOPATAsReported, *WACCAsReported, *TotalCapitalAsReported))
		}

		var ReturnOnInvestedCapital *float64 = nil
		var ReturnOnInvestedCapitalAsReported *float64 = nil
		if NOPAT != nil && TotalInvestments != nil {
			ReturnOnInvestedCapital = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnInvestedCapital(*NOPAT, *TotalInvestments))
		}
		if NOPATAsReported != nil && TotalInvestmentsAsReported != nil {
			ReturnOnInvestedCapitalAsReported = utils.InterfaceToFloat64Ptr(Calculations.ReturnOnInvestedCapital(*NOPATAsReported, *TotalInvestmentsAsReported))
		}

		var FreeCashFlowToFirm *float64 = nil
		var FreeCashFlowToFirmAsReported *float64 = nil
		if NetIncome != nil && NonCashCharges != nil && TotalInterestPayments != nil && EffectiveTaxRate != nil && LongTermInvestments != nil && NetWorkingCapital != nil {
			FreeCashFlowToFirm = utils.InterfaceToFloat64Ptr(Calculations.FreeCashFlowToFirm(*NetIncome, *NonCashCharges, *TotalInterestPayments, *EffectiveTaxRate, *LongTermInvestments, *NetWorkingCapital))
		}
		if NetIncomeAsReported != nil && NonCashChargesAsReported != nil && TotalInterestPaymentsAsReported != nil && EffectiveTaxRate != nil && LongTermInvestmentsAsReported != nil && NetWorkingCapitalAsReported != nil {
			FreeCashFlowToFirmAsReported = utils.InterfaceToFloat64Ptr(Calculations.FreeCashFlowToFirm(*NetIncomeAsReported, *NonCashChargesAsReported, *TotalInterestPaymentsAsReported, *EffectiveTaxRate, *LongTermInvestmentsAsReported, *NetWorkingCapitalAsReported))
		}

		var StockDividendPerShare *float64 = nil
		var StockDividendPerShareAsReported *float64 = nil
		if DividendsPaid != nil && SharesOutstanding != nil {
			StockDividendPerShare = utils.InterfaceToFloat64Ptr(*DividendsPaid / *SharesOutstanding)
		}
		if DividendsPaidAsReported != nil && SharesOutstandingAsReported != nil {
			StockDividendPerShareAsReported = utils.InterfaceToFloat64Ptr(*DividendsPaidAsReported / *SharesOutstandingAsReported)
		}

		var LeverageRatio *float64 = nil
		var LeverageRatioAsReported *float64 = nil
		if TotalDebt != nil && EBITDA != nil {
			LeverageRatio = utils.InterfaceToFloat64Ptr(Calculations.LeverageRatio(*TotalDebt, *EBITDA))
		}
		if TotalDebtAsReported != nil && EBITDAAsReported != nil {
			LeverageRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.LeverageRatio(*TotalDebtAsReported, *EBITDAAsReported))
		}

		var CapitalizationRatio *float64 = nil
		var CapitalizationRatioAsReported *float64 = nil
		if TotalDebt != nil && ShareholderEquity != nil {
			CapitalizationRatio = utils.InterfaceToFloat64Ptr(Calculations.CapitalizationRatio(*TotalDebt, *ShareholderEquity))
		}
		if TotalDebtAsReported != nil && ShareholderEquityAsReported != nil {
			CapitalizationRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.CapitalizationRatio(*TotalDebtAsReported, *ShareholderEquityAsReported))
		}

		var LongTermDebtAsReported *float64 = nil
		if CurrentLongTermDebtAsReported != nil && NonCurrentLongTermDebtAsReported != nil {
			LongTermDebtAsReported = utils.InterfaceToFloat64Ptr(*CurrentLongTermDebtAsReported + *NonCurrentLongTermDebtAsReported)
		} else if CurrentLongTermDebtAsReported != nil {
			LongTermDebtAsReported = utils.InterfaceToFloat64Ptr(*CurrentLongTermDebtAsReported)
		} else if NonCurrentLongTermDebtAsReported != nil {
			LongTermDebtAsReported = utils.InterfaceToFloat64Ptr(*NonCurrentLongTermDebtAsReported)
		}

		var DebtToCapitalRatio *float64 = nil
		var DebtToCapitalRatioAsReported *float64 = nil
		if ShortTermDebt != nil && LongTermDebt != nil && ShareholderEquity != nil {
			DebtToCapitalRatio = utils.InterfaceToFloat64Ptr(Calculations.DebtToCapitalRatio(*ShortTermDebt, *LongTermDebt, *ShareholderEquity))
		}
		if ShortTermDebtAsReported != nil && LongTermDebtAsReported != nil && ShareholderEquityAsReported != nil {
			DebtToCapitalRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DebtToCapitalRatio(*ShortTermDebtAsReported, *LongTermDebtAsReported, *ShareholderEquityAsReported))
		}

		var NetGearingRatio *float64 = nil
		var NetGearingRatioAsReported *float64 = nil
		if LongTermDebt != nil && ShortTermDebt != nil && TotalLiabilities != nil && ShareholderEquity != nil {
			NetGearingRatio = utils.InterfaceToFloat64Ptr(Calculations.NetGearingRatio(*LongTermDebt, *ShortTermDebt, *TotalLiabilities, *ShareholderEquity))
		}
		if LongTermDebtAsReported != nil && ShortTermDebtAsReported != nil && TotalLiabilitiesAsReported != nil && ShareholderEquityAsReported != nil {
			NetGearingRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.NetGearingRatio(*LongTermDebtAsReported, *ShortTermDebtAsReported, *TotalLiabilitiesAsReported, *ShareholderEquityAsReported))
		}

		var TotalDebtToEBITDA *float64 = nil
		var TotalDebtToEBITDAAsReported *float64 = nil
		if TotalDebt != nil && EBITDA != nil {
			TotalDebtToEBITDA = utils.InterfaceToFloat64Ptr(Calculations.TotalDebtToEBITDA(*TotalDebt, *EBITDA))
		}
		if TotalDebtAsReported != nil && EBITDAAsReported != nil {
			TotalDebtToEBITDAAsReported = utils.InterfaceToFloat64Ptr(Calculations.TotalDebtToEBITDA(*TotalDebtAsReported, *EBITDAAsReported))
		}

		var DebtToEquityRatio *float64 = nil
		var DebtToEquityRatioAsReported *float64 = nil
		if TotalLiabilities != nil && ShareholderEquity != nil {
			DebtToEquityRatio = utils.InterfaceToFloat64Ptr(Calculations.DebtToEquityRatio(*TotalLiabilities, *ShareholderEquity))
		}
		if TotalLiabilitiesAsReported != nil && ShareholderEquityAsReported != nil {
			DebtToEquityRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.DebtToEquityRatio(*TotalLiabilitiesAsReported, *ShareholderEquityAsReported))
		}

		var EquityMultiplierRatio *float64 = nil
		var EquityMultiplierRatioAsReported *float64 = nil
		if TotalAssets != nil && ShareholderEquity != nil {
			EquityMultiplierRatio = utils.InterfaceToFloat64Ptr(Calculations.EquityMultiplierRatio(*TotalAssets, *ShareholderEquity))
		}
		if TotalAssetsAsReported != nil && ShareholderEquityAsReported != nil {
			EquityMultiplierRatioAsReported = utils.InterfaceToFloat64Ptr(Calculations.EquityMultiplierRatio(*TotalAssetsAsReported, *ShareholderEquityAsReported))
		}

		var DuPontAnalysis *float64 = nil
		var DuPontAnalysisAsReported *float64 = nil
		if NetProfitMargin != nil && AssetTurnoverRatio != nil && EquityMultiplierRatio != nil {
			DuPontAnalysis = utils.InterfaceToFloat64Ptr(Calculations.DuPontAnalysis(*NetProfitMargin, *AssetTurnoverRatio, *EquityMultiplierRatio))
		}
		if NetProfitMarginAsReported != nil && AssetTurnoverRatioAsReported != nil && EquityMultiplierRatioAsReported != nil {
			DuPontAnalysisAsReported = utils.InterfaceToFloat64Ptr(Calculations.DuPontAnalysis(*NetProfitMarginAsReported, *AssetTurnoverRatioAsReported, *EquityMultiplierRatioAsReported))
		}
		FullCalcResults["EffectiveTaxRate"] = EffectiveTaxRate
		FullCalcResults["DaysInPeriod"] = DaysInPeriod
		FullCalcResults["TotalAssets"] = TotalAssets
		FullCalcResults["TotalLiabilities"] = TotalLiabilities
		FullCalcResults["Inventory"] = Inventory
		FullCalcResults["IntangibleAssets"] = IntangibleAssets
		FullCalcResults["NetDebt"] = NetDebt
		FullCalcResults["CashAndCashEquivalents"] = CashAndCashEquivalents
		FullCalcResults["NetReceivables"] = NetReceivables
		FullCalcResults["NetFixedAssets"] = NetFixedAssets
		FullCalcResults["DeferredTaxLiabilities"] = DeferredTaxLiabilities
		FullCalcResults["ShareholderEquity"] = ShareholderEquity
		FullCalcResults["AccountsPayable"] = AccountsPayable
		FullCalcResults["CommonStock"] = CommonStock
		FullCalcResults["SharesOutstanding"] = SharesOutstanding
		FullCalcResults["HighQualityLiquidAssets"] = HighQualityLiquidAssets
		FullCalcResults["WorkingCapital"] = WorkingCapital
		FullCalcResults["TangibleNetWorth"] = TangibleNetWorth
		FullCalcResults["BookValueOfEquity"] = BookValueOfEquity
		FullCalcResults["BookValueOfDebt"] = BookValueOfDebt
		FullCalcResults["EquityBookValue"] = EquityBookValue
		FullCalcResults["LiabilitiesBookValue"] = LiabilitiesBookValue
		FullCalcResults["TotalAccrualsToTotalAssets"] = TotalAccrualsToTotalAssets
		FullCalcResults["ShortTermInvestments"] = ShortTermInvestments
		FullCalcResults["LongTermInvestments"] = LongTermInvestments
		FullCalcResults["TotalMarketableSecurities"] = TotalMarketableSecurities
		FullCalcResults["TotalInvestments"] = TotalInvestments
		FullCalcResults["NetIncome"] = NetIncome
		FullCalcResults["GrossProfit"] = GrossProfit
		FullCalcResults["NetRevenue"] = NetRevenue
		FullCalcResults["NetProfitMargin"] = NetProfitMargin
		FullCalcResults["OperatingExpenses"] = OperatingExpenses
		FullCalcResults["OperatingIncome"] = OperatingIncome
		FullCalcResults["DepreciationAndAmortization"] = DepreciationAndAmortization
		FullCalcResults["TotalTaxesPaid"] = TotalTaxesPaid
		FullCalcResults["ChangeInWorkingCapital"] = ChangeInWorkingCapital
		FullCalcResults["CapitalExpenditures"] = CapitalExpenditures
		FullCalcResults["OperatingCashflow"] = OperatingCashflow
		FullCalcResults["FreeCashFlow"] = FreeCashFlow
		FullCalcResults["EBITDA"] = EBITDA
		FullCalcResults["TotalInterestPayments"] = TotalInterestPayments
		FullCalcResults["EBIT"] = EBIT
		FullCalcResults["NonCashCharges"] = NonCashCharges
		FullCalcResults["MarketValueOfEquity"] = MarketValueOfEquity
		FullCalcResults["ShortTermDebt"] = ShortTermDebt
		FullCalcResults["LongTermDebt"] = LongTermDebt
		FullCalcResults["TotalDebt"] = TotalDebt
		FullCalcResults["CostOfDebt"] = CostOfDebt
		FullCalcResults["UnleveredFirmValue"] = UnleveredFirmValue
		FullCalcResults["TaxShieldBenefits"] = TaxShieldBenefits
		FullCalcResults["NetEffectOfDebt"] = NetEffectOfDebt
		FullCalcResults["DebtService"] = DebtService
		FullCalcResults["NonInterestExpenses"] = NonInterestExpenses
		FullCalcResults["MarketCapitalization"] = MarketCapitalization
		FullCalcResults["EnterpriseValue"] = EnterpriseValue
		FullCalcResults["DebtOutstanding"] = DebtOutstanding
		FullCalcResults["AssetTurnoverRatio"] = AssetTurnoverRatio
		FullCalcResults["DividendsPaid"] = DividendsPaid
		FullCalcResults["RetentionRatio"] = RetentionRatio
		FullCalcResults["ReturnOnEquity"] = ReturnOnEquity
		FullCalcResults["CostAndExpenses"] = CostAndExpenses
		FullCalcResults["CostOfRevenue"] = CostOfRevenue
		FullCalcResults["SellingGeneralAndAdministrativeExpenses"] = SellingGeneralAndAdministrativeExpenses
		FullCalcResults["ExplicitCosts"] = ExplicitCosts
		FullCalcResults["DaysInventoryOutstanding"] = DaysInventoryOutstanding
		FullCalcResults["TotalCapital"] = TotalCapital
		FullCalcResults["NetMargin"] = NetMargin
		FullCalcResults["FreeCashFlowToEquity"] = FreeCashFlowToEquity
		FullCalcResults["AdjustedPresentValue"] = AdjustedPresentValue
		FullCalcResults["InterestCoverageRatio"] = InterestCoverageRatio
		FullCalcResults["FixedChargeCoverageRatio"] = FixedChargeCoverageRatio
		FullCalcResults["DebtServiceCoverageRatio"] = DebtServiceCoverageRatio
		FullCalcResults["AssetCoverageRatio"] = AssetCoverageRatio
		FullCalcResults["EBITDAToInterestCoverageRatio"] = EBITDAToInterestCoverageRatio
		FullCalcResults["PreferredDividendCoverageRatio"] = PreferredDividendCoverageRatio
		FullCalcResults["LiquidityCoverageRatio"] = LiquidityCoverageRatio
		FullCalcResults["InventoryTurnoverRatio"] = InventoryTurnoverRatio
		FullCalcResults["ReturnOnCapitalEmployed"] = ReturnOnCapitalEmployed
		FullCalcResults["EfficiencyRatio"] = EfficiencyRatio
		FullCalcResults["RevenuePerEmployee"] = RevenuePerEmployee
		FullCalcResults["CapitalExpenditureRatio"] = CapitalExpenditureRatio
		FullCalcResults["OperatingCashFlowRatio"] = OperatingCashFlowRatio
		FullCalcResults["EBITDAToEVRatio"] = EBITDAToEVRatio
		FullCalcResults["TangibleNetWorthRatio"] = TangibleNetWorthRatio
		FullCalcResults["DeferredTaxLiabilityToEquityRatio"] = DeferredTaxLiabilityToEquityRatio
		FullCalcResults["TangibleEquityRatio"] = TangibleEquityRatio
		FullCalcResults["WACC"] = WACC
		FullCalcResults["FixedAssetTurnoverRatio"] = FixedAssetTurnoverRatio
		FullCalcResults["PPETurnoverRatio"] = PPETurnoverRatio
		FullCalcResults["InvestmentTurnoverRatio"] = InvestmentTurnoverRatio
		FullCalcResults["WorkingCapitalTurnoverRatio"] = WorkingCapitalTurnoverRatio
		FullCalcResults["ReturnOnAssetRatio"] = ReturnOnAssetRatio
		FullCalcResults["GrossProfitMargin"] = GrossProfitMargin
		FullCalcResults["OperatingProfitMargin"] = OperatingProfitMargin
		FullCalcResults["EBITDAMarginRatio"] = EBITDAMarginRatio
		FullCalcResults["DividendPayoutRatio"] = DividendPayoutRatio
		FullCalcResults["RetentionRate"] = RetentionRate
		FullCalcResults["SustainableGrowthRate"] = SustainableGrowthRate
		FullCalcResults["GrossMarginOnInventory"] = GrossMarginOnInventory
		FullCalcResults["CashFlowReturnOnEquity"] = CashFlowReturnOnEquity
		FullCalcResults["OperatingMargin"] = OperatingMargin
		FullCalcResults["OperatingExpenseRatio"] = OperatingExpenseRatio
		FullCalcResults["CurrentRatio"] = CurrentRatio
		FullCalcResults["AcidTestRatio"] = AcidTestRatio
		FullCalcResults["CashRatio"] = CashRatio
		FullCalcResults["DefensiveIntervalRatio"] = DefensiveIntervalRatio
		FullCalcResults["DrySalesRatio"] = DrySalesRatio
		FullCalcResults["PriceToBookValueRatio"] = PriceToBookValueRatio
		FullCalcResults["EarningsPerShare"] = EarningsPerShare
		FullCalcResults["EBITDAPerShare"] = EBITDAPerShare
		FullCalcResults["BookValuePerShare"] = BookValuePerShare
		FullCalcResults["NetTangibleAssetsPerShare"] = NetTangibleAssetsPerShare
		FullCalcResults["MarketValueOfDebt"] = MarketValueOfDebt
		FullCalcResults["MarketToBookRatio"] = MarketToBookRatio
		FullCalcResults["IntangiblesRatio"] = IntangiblesRatio
		FullCalcResults["PriceToSalesRatio"] = PriceToSalesRatio
		FullCalcResults["PriceToBookRatio"] = PriceToBookRatio
		FullCalcResults["PricetoSalesValue"] = PricetoSalesValue
		FullCalcResults["OperatingCashFlowPerShare"] = OperatingCashFlowPerShare
		FullCalcResults["PriceToCashFlowRatio"] = PriceToCashFlowRatio
		FullCalcResults["FreeCashFlowPerShare"] = FreeCashFlowPerShare
		FullCalcResults["PriceToFreeCashFlowRatio"] = PriceToFreeCashFlowRatio
		FullCalcResults["PriceToCashFlowValuation"] = PriceToCashFlowValuation
		FullCalcResults["PriceToFreeCashFlowValuation"] = PriceToFreeCashFlowValuation
		FullCalcResults["PriceToEarningsValuation"] = PriceToEarningsValuation
		FullCalcResults["LiabilitiesMarketValue"] = LiabilitiesMarketValue
		FullCalcResults["TobinsQ"] = TobinsQ
		FullCalcResults["ReceivablesTurnoverRatio"] = ReceivablesTurnoverRatio
		FullCalcResults["AverageCollectionPeriod"] = AverageCollectionPeriod
		FullCalcResults["AccountsPayableTurnoverRatio"] = AccountsPayableTurnoverRatio
		FullCalcResults["AverageAccountsPayablePaymentPeriod"] = AverageAccountsPayablePaymentPeriod
		FullCalcResults["InventoryToWorkingCapitalRatio"] = InventoryToWorkingCapitalRatio
		FullCalcResults["DaysSalesOutstanding"] = DaysSalesOutstanding
		FullCalcResults["DaysPayablesOutstanding"] = DaysPayablesOutstanding
		FullCalcResults["CashConversionCycle"] = CashConversionCycle
		FullCalcResults["NetWorkingCapital"] = NetWorkingCapital
		FullCalcResults["NOPAT"] = NOPAT
		FullCalcResults["EconomicValueAdded"] = EconomicValueAdded
		FullCalcResults["ReturnOnInvestedCapital"] = ReturnOnInvestedCapital
		FullCalcResults["FreeCashFlowToFirm"] = FreeCashFlowToFirm
		FullCalcResults["StockDividendPerShare"] = StockDividendPerShare
		FullCalcResults["LeverageRatio"] = LeverageRatio
		FullCalcResults["CapitalizationRatio"] = CapitalizationRatio
		FullCalcResults["DebtToCapitalRatio"] = DebtToCapitalRatio
		FullCalcResults["NetGearingRatio"] = NetGearingRatio
		FullCalcResults["TotalDebtToEBITDA"] = TotalDebtToEBITDA
		FullCalcResults["DebtToEquityRatio"] = DebtToEquityRatio
		FullCalcResults["EquityMultiplierRatio"] = EquityMultiplierRatio
		FullCalcResults["DuPontAnalysis"] = DuPontAnalysis

		FullCalcResultsAsReported["EffectiveTaxRate"] = EffectiveTaxRate
		FullCalcResultsAsReported["DaysInPeriod"] = DaysInPeriod
		FullCalcResultsAsReported["TotalAssets"] = TotalAssetsAsReported
		FullCalcResultsAsReported["AssetsNonCurrent"] = AssetsNonCurrentAsReported
		FullCalcResultsAsReported["OtherAssetsNonCurrent"] = OtherAssetsNonCurrentAsReported
		FullCalcResultsAsReported["TotalLiabilities"] = TotalLiabilitiesAsReported
		FullCalcResultsAsReported["Inventory"] = InventoryAsReported
		FullCalcResultsAsReported["IntangibleAssets"] = IntangibleAssetsAsReported
		FullCalcResultsAsReported["CurrentLongTermDebt"] = CurrentLongTermDebtAsReported
		FullCalcResultsAsReported["NonCurrentLongTermDebt"] = NonCurrentLongTermDebtAsReported
		FullCalcResultsAsReported["NetDebt"] = NetDebtAsReported
		FullCalcResultsAsReported["CashAndCashEquivalents"] = CashAndCashEquivalentsAsReported
		FullCalcResultsAsReported["AccountsReceivable"] = AccountsReceivableAsReported
		FullCalcResultsAsReported["NonTradeReceivables"] = NonTradeReceivablesAsReported
		FullCalcResultsAsReported["NetReceivables"] = NetReceivablesAsReported
		FullCalcResultsAsReported["NetFixedAssets"] = NetFixedAssetsAsReported
		FullCalcResultsAsReported["DeferredTaxLiabilities"] = DeferredTaxLiabilitiesAsReported
		FullCalcResultsAsReported["ShareholderEquity"] = ShareholderEquityAsReported
		FullCalcResultsAsReported["AccountsPayable"] = AccountsPayableAsReported
		FullCalcResultsAsReported["SharesOutstanding"] = SharesOutstandingAsReported
		FullCalcResultsAsReported["HighQualityLiquidAssets"] = HighQualityLiquidAssetsAsReported
		FullCalcResultsAsReported["WorkingCapital"] = WorkingCapitalAsReported
		FullCalcResultsAsReported["TangibleNetWorth"] = TangibleNetWorthAsReported
		FullCalcResultsAsReported["BookValueOfEquity"] = BookValueOfEquityAsReported
		FullCalcResultsAsReported["BookValueOfDebt"] = BookValueOfDebtAsReported
		FullCalcResultsAsReported["EquityBookValue"] = EquityBookValueAsReported
		FullCalcResultsAsReported["LiabilitiesBookValue"] = LiabilitiesBookValueAsReported
		FullCalcResultsAsReported["TotalAccrualsToTotalAssets"] = TotalAccrualsToTotalAssetsAsReported
		FullCalcResultsAsReported["TotalMarketableSecurities"] = TotalMarketableSecuritiesAsReported
		FullCalcResultsAsReported["CurrentMarketableSecurities"] = CurrentMarketableSecuritiesAsReported
		FullCalcResultsAsReported["NonCurrentMarketableSecurities"] = NonCurrentMarketableSecuritiesAsReported
		FullCalcResultsAsReported["LongTermInvestments"] = LongTermInvestmentsAsReported
		FullCalcResultsAsReported["TotalInvestments"] = TotalInvestmentsAsReported
		FullCalcResultsAsReported["NetIncome"] = NetIncomeAsReported
		FullCalcResultsAsReported["GrossProfit"] = GrossProfitAsReported
		FullCalcResultsAsReported["NetRevenue"] = NetRevenueAsReported
		FullCalcResultsAsReported["NetProfitMargin"] = NetProfitMarginAsReported
		FullCalcResultsAsReported["OperatingExpenses"] = OperatingExpensesAsReported
		FullCalcResultsAsReported["OperatingIncome"] = OperatingIncomeAsReported
		FullCalcResultsAsReported["DepreciationAndAmortization"] = DepreciationAndAmortizationAsReported
		FullCalcResultsAsReported["TotalInterestPayments"] = TotalInterestPaymentsAsReported
		FullCalcResultsAsReported["TotalTaxesPaid"] = TotalTaxesPaidAsReported
		FullCalcResultsAsReported["CapitalExpenditures"] = CapitalExpendituresAsReported
		FullCalcResultsAsReported["NetCashOperatingActivities"] = NetCashOperatingActivitiesAsReported
		FullCalcResultsAsReported["NetCashInvestingActivities"] = NetCashInvestingActivitiesAsReported
		FullCalcResultsAsReported["NetCashFinancingActivities"] = NetCashFinancingActivitiesAsReported
		FullCalcResultsAsReported["OperatingCashFlow"] = OperatingCashFlowAsReported
		FullCalcResultsAsReported["FreeCashFlow"] = FreeCashFlowAsReported
		FullCalcResultsAsReported["EBITDA"] = EBITDAAsReported
		FullCalcResultsAsReported["EBIT"] = EBITAsReported
		FullCalcResultsAsReported["NonCashCharges"] = NonCashChargesAsReported
		FullCalcResultsAsReported["MarketValueOfEquity"] = MarketValueOfEquityAsReported
		FullCalcResultsAsReported["ShortTermDebt"] = ShortTermDebtAsReported
		FullCalcResultsAsReported["TotalDebt"] = TotalDebtAsReported
		FullCalcResultsAsReported["CostOfDebt"] = CostOfDebtAsReported
		FullCalcResultsAsReported["UnleveredFirmValue"] = UnleveredFirmValueAsReported
		FullCalcResultsAsReported["TaxShieldBenefits"] = TaxShieldBenefitsAsReported
		FullCalcResultsAsReported["NetEffectOfDebt"] = NetEffectOfDebtAsReported
		FullCalcResultsAsReported["DebtService"] = DebtServiceAsReported
		FullCalcResultsAsReported["NonInterestExpenses"] = NonInterestExpensesAsReported
		FullCalcResultsAsReported["MarketCapitalization"] = MarketCapitalizationAsReported
		FullCalcResultsAsReported["EnterpriseValue"] = EnterpriseValueAsReported
		FullCalcResultsAsReported["DebtOutstanding"] = DebtOutstandingAsReported
		FullCalcResultsAsReported["AssetTurnoverRatio"] = AssetTurnoverRatioAsReported
		FullCalcResultsAsReported["DividendsPaid"] = DividendsPaidAsReported
		FullCalcResultsAsReported["RetentionRatio"] = RetentionRatioAsReported
		FullCalcResultsAsReported["ReturnOnEquity"] = ReturnOnEquityAsReported
		FullCalcResultsAsReported["CostOfGoodsSold"] = CostOfGoodsSoldAsReported
		FullCalcResultsAsReported["SellingGeneralAndAdministrativeExpenses"] = SellingGeneralAndAdministrativeExpensesAsReported
		FullCalcResultsAsReported["ExplicitCosts"] = ExplicitCostsAsReported
		FullCalcResultsAsReported["DaysInventoryOutstanding"] = DaysInventoryOutstandingAsReported
		FullCalcResultsAsReported["TotalCapital"] = TotalCapitalAsReported
		FullCalcResultsAsReported["NetMargin"] = NetMarginAsReported
		FullCalcResultsAsReported["AdjustedPresentValue"] = AdjustedPresentValueAsReported
		FullCalcResultsAsReported["InterestCoverageRatio"] = InterestCoverageRatioAsReported
		FullCalcResultsAsReported["FixedChargeCoverageRatio"] = FixedChargeCoverageRatioAsReported
		FullCalcResultsAsReported["DebtServiceCoverageRatio"] = DebtServiceCoverageRatioAsReported
		FullCalcResultsAsReported["AssetCoverageRatio"] = AssetCoverageRatioAsReported
		FullCalcResultsAsReported["EBITDAToInterestCoverageRatio"] = EBITDAToInterestCoverageRatioAsReported
		FullCalcResultsAsReported["PreferredDividendCoverageRatio"] = PreferredDividendCoverageRatioAsReported
		FullCalcResultsAsReported["LiquidityCoverageRatio"] = LiquidityCoverageRatioAsReported
		FullCalcResultsAsReported["InventoryTurnoverRatio"] = InventoryTurnoverRatioAsReported
		FullCalcResultsAsReported["ReturnOnCapitalEmployed"] = ReturnOnCapitalEmployedAsReported
		FullCalcResultsAsReported["EfficiencyRatio"] = EfficiencyRatioAsReported
		FullCalcResultsAsReported["RevenuePerEmployee"] = RevenuePerEmployeeAsReported
		FullCalcResultsAsReported["CapitalExpenditureRatio"] = CapitalExpenditureRatioAsReported
		FullCalcResultsAsReported["OperatingCashFlowRatio"] = OperatingCashFlowRatioAsReported
		FullCalcResultsAsReported["EBITDAToEVRatio"] = EBITDAToEVRatioAsReported
		FullCalcResultsAsReported["TangibleNetWorthRatio"] = TangibleNetWorthRatioAsReported
		FullCalcResultsAsReported["DeferredTaxLiabilityToEquityRatio"] = DeferredTaxLiabilityToEquityRatioAsReported
		FullCalcResultsAsReported["TangibleEquityRatio"] = TangibleEquityRatioAsReported
		FullCalcResultsAsReported["WACC"] = WACCAsReported
		FullCalcResultsAsReported["FixedAssetTurnoverRatio"] = FixedAssetTurnoverRatioAsReported
		FullCalcResultsAsReported["PPETurnoverRatio"] = PPETurnoverRatioAsReported
		FullCalcResultsAsReported["InvestmentTurnoverRatio"] = InvestmentTurnoverRatioAsReported
		FullCalcResultsAsReported["WorkingCapitalTurnoverRatio"] = WorkingCapitalTurnoverRatioAsReported
		FullCalcResultsAsReported["ReturnOnAssetRatio"] = ReturnOnAssetRatioAsReported
		FullCalcResultsAsReported["GrossProfitMargin"] = GrossProfitMarginAsReported
		FullCalcResultsAsReported["OperatingProfitMargin"] = OperatingProfitMarginAsReported
		FullCalcResultsAsReported["EBITDAMarginRatio"] = EBITDAMarginRatioAsReported
		FullCalcResultsAsReported["DividendPayoutRatio"] = DividendPayoutRatioAsReported
		FullCalcResultsAsReported["RetentionRate"] = RetentionRateAsReported
		FullCalcResultsAsReported["SustainableGrowthRate"] = SustainableGrowthRateAsReported
		FullCalcResultsAsReported["GrossMarginOnInventory"] = GrossMarginOnInventoryAsReported
		FullCalcResultsAsReported["CashFlowReturnOnEquity"] = CashFlowReturnOnEquityAsReported
		FullCalcResultsAsReported["OperatingMargin"] = OperatingMarginAsReported
		FullCalcResultsAsReported["OperatingExpenseRatio"] = OperatingExpenseRatioAsReported
		FullCalcResultsAsReported["CurrentRatio"] = CurrentRatioAsReported
		FullCalcResultsAsReported["AcidTestRatio"] = AcidTestRatioAsReported
		FullCalcResultsAsReported["CashRatio"] = CashRatioAsReported
		FullCalcResultsAsReported["DefensiveIntervalRatio"] = DefensiveIntervalRatioAsReported
		FullCalcResultsAsReported["DrySalesRatio"] = DrySalesRatioAsReported
		FullCalcResultsAsReported["PriceToBookValueRatio"] = PriceToBookValueRatioAsReported
		FullCalcResultsAsReported["EarningsPerShare"] = EarningsPerShareAsReported
		FullCalcResultsAsReported["EBITDAPerShare"] = EBITDAPerShareAsReported
		FullCalcResultsAsReported["BookValuePerShare"] = BookValuePerShareAsReported
		FullCalcResultsAsReported["NetTangibleAssetsPerShare"] = NetTangibleAssetsPerShareAsReported
		FullCalcResultsAsReported["MarketValueOfDebt"] = MarketValueOfDebtAsReported
		FullCalcResultsAsReported["MarketToBookRatio"] = MarketToBookRatioAsReported
		FullCalcResultsAsReported["IntangiblesRatio"] = IntangiblesRatioAsReported
		FullCalcResultsAsReported["PriceToSalesRatio"] = PriceToSalesRatioAsReported
		FullCalcResultsAsReported["PriceToBookRatio"] = PriceToBookRatioAsReported
		FullCalcResultsAsReported["PricetoSalesValue"] = PricetoSalesValueAsReported
		FullCalcResultsAsReported["OperatingCashFlowPerShare"] = OperatingCashFlowPerShareAsReported
		FullCalcResultsAsReported["PriceToCashFlowRatio"] = PriceToCashFlowRatioAsReported
		FullCalcResultsAsReported["FreeCashFlowPerShare"] = FreeCashFlowPerShareAsReported
		FullCalcResultsAsReported["PriceToFreeCashFlowRatio"] = PriceToFreeCashFlowRatioAsReported
		FullCalcResultsAsReported["PriceToCashFlowValuation"] = PriceToCashFlowValuationAsReported
		FullCalcResultsAsReported["PriceToFreeCashFlowValuation"] = PriceToFreeCashFlowValuationAsReported
		FullCalcResultsAsReported["PriceToEarningsValuation"] = PriceToEarningsValuationAsReported
		FullCalcResultsAsReported["LiabilitiesMarketValue"] = LiabilitiesMarketValueAsReported
		FullCalcResultsAsReported["TobinsQ"] = TobinsQAsReported
		FullCalcResultsAsReported["ReceivablesTurnoverRatio"] = ReceivablesTurnoverRatioAsReported
		FullCalcResultsAsReported["AverageCollectionPeriod"] = AverageCollectionPeriodAsReported
		FullCalcResultsAsReported["AccountsPayableTurnoverRatio"] = AccountsPayableTurnoverRatioAsReported
		FullCalcResultsAsReported["AverageAccountsPayablePaymentPeriod"] = AverageAccountsPayablePaymentPeriodAsReported
		FullCalcResultsAsReported["InventoryToWorkingCapitalRatio"] = InventoryToWorkingCapitalRatioAsReported
		FullCalcResultsAsReported["DaysSalesOutstanding"] = DaysSalesOutstandingAsReported
		FullCalcResultsAsReported["DaysPayablesOutstanding"] = DaysPayablesOutstandingAsReported
		FullCalcResultsAsReported["CashConversionCycle"] = CashConversionCycleAsReported
		FullCalcResultsAsReported["NetWorkingCapital"] = NetWorkingCapitalAsReported
		FullCalcResultsAsReported["NOPAT"] = NOPATAsReported
		FullCalcResultsAsReported["EconomicValueAdded"] = EconomicValueAddedAsReported
		FullCalcResultsAsReported["ReturnOnInvestedCapital"] = ReturnOnInvestedCapitalAsReported
		FullCalcResultsAsReported["FreeCashFlowToFirm"] = FreeCashFlowToFirmAsReported
		FullCalcResultsAsReported["StockDividendPerShare"] = StockDividendPerShareAsReported
		FullCalcResultsAsReported["LeverageRatio"] = LeverageRatioAsReported
		FullCalcResultsAsReported["CapitalizationRatio"] = CapitalizationRatioAsReported
		FullCalcResultsAsReported["LongTermDebt"] = LongTermDebtAsReported
		FullCalcResultsAsReported["DebtToCapitalRatio"] = DebtToCapitalRatioAsReported
		FullCalcResultsAsReported["NetGearingRatio"] = NetGearingRatioAsReported
		FullCalcResultsAsReported["TotalDebtToEBITDA"] = TotalDebtToEBITDAAsReported
		FullCalcResultsAsReported["DebtToEquityRatio"] = DebtToEquityRatioAsReported
		FullCalcResultsAsReported["EquityMultiplierRatio"] = EquityMultiplierRatioAsReported
		FullCalcResultsAsReported["DuPontAnalysis"] = DuPontAnalysisAsReported

		FinalCalcResults = append(FinalCalcResults, FullCalcResults)
		FinalCalcResultsAsReported = append(FinalCalcResultsAsReported, FullCalcResultsAsReported)
	}

	return FinalCalcResults, FinalCalcResultsAsReported
}

func calculateGrowth(data []map[string]*float64) map[string][]float64 {
	growthMap := make(map[string][]float64)

	// Previous values for each key to calculate growth
	prevValues := make(map[string]float64)

	for _, entry := range data {
		for key, valuePtr := range entry {
			if valuePtr != nil {
				value := *valuePtr

				// Check if we have a previous value
				if prevValue, exists := prevValues[key]; exists && prevValue != 0 {
					// Calculate growth
					growth := (value - prevValue) / prevValue
					growthMap[key] = append(growthMap[key], growth)
				} else {
					// For the first non-nil value, growth is 0
					growthMap[key] = append(growthMap[key], 0)
				}

				// Update previous value
				prevValues[key] = value
			}
			// If value is nil, do nothing for this iteration
		}
	}

	// Remove keys that only had nil values
	for key, growths := range growthMap {
		if len(growths) == 0 {
			delete(growthMap, key)
		}
	}

	return growthMap
}
