package Analysis

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Calculations"
)

func PerformFundamentalsCalculations(Fundamentals *CompanyFundamentals, Period string) *FundamentalsCalculationsResults {
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

	return CalculationResults
}
