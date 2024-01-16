package Analysis

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Calculations"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/spacecodewor/fmpcloud-go/objects"
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

func PerformCustomCalculations(Fundamentals *CompanyFundamentals, Period objects.CompanyValuationPeriod) (map[string]float64, map[string]float64) {
	FullCalcResults := map[string]float64{}
	FullCalcResultsGrowth := map[string]float64{}
	FullCalcResultsAsReported := map[string]float64{}
	FullCalcResultsAsReportedGrowth := map[string]float64{}
	ZippedFullCaclResultsMeanSTD := map[string]float64{}
	ZippedFullCaclResultsGrowthMeanSTD := map[string]float64{}

	LenBalanceSheetStatements := len(Fundamentals.BalanceSheetStatements)
	LenBalanceSheetStatementAsReported := len(Fundamentals.BalanceSheetStatementAsReported)
	LenIncomeStatement := len(Fundamentals.IncomeStatement)
	LenIncomeStatementAsReported := len(Fundamentals.IncomeStatementAsReported)
	LenCashFlowStatement := len(Fundamentals.CashFlowStatement)
	LenCashFlowStatementAsReported := len(Fundamentals.CashFlowStatementAsReported)

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

		/* Balance Sheet */

		// Net Inventory
		var Inventory *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheet) {
			Inventory = utils.InterfaceToFloat64Ptr(curBalanceSheet.Inventory)
		}

		var InventoryAsReported *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheetAsReported) {
			InventoryAsReported = utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Inventorynet)
		}

		// Total Assets
		var TotalAssets *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheet) {
			TotalAssets = utils.InterfaceToFloat64Ptr(curBalanceSheet.TotalAssets)
		}

		var TotalAssetsAsReported *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheetAsReported) {
			TotalAssetsAsReported = utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Assets)
		}

		// Total Liabilities
		var TotalLiabilities *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheet) {
			TotalLiabilities = utils.InterfaceToFloat64Ptr(curBalanceSheet.TotalLiabilities)
		}

		var TotalLiabilitiesAsReported *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheetAsReported) {
			TotalLiabilitiesAsReported = utils.InterfaceToFloat64Ptr(curBalanceSheetAsReported.Liabilities)
		}
		// Intangible Assets
		var IntangibleAssets *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheet) {
			IntangibleAssets = utils.InterfaceToFloat64Ptr(curBalanceSheet.IntangibleAssets)
		}

		var IntangibleAssetsAsReported *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheetAsReported) {
			if TotalAssetsAsReported != nil && InventoryAsReported != nil {
				IntangibleAssetsAsReported = utils.InterfaceToFloat64Ptr(*TotalAssetsAsReported - *InventoryAsReported)
			}
		}

		// Net Debt
		var NetDebt *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheet) {
			NetDebt = utils.InterfaceToFloat64Ptr(curBalanceSheet.NetDebt)
		}

		var NetDebtAsReported *float64 = nil
		if !utils.IsStructEmpty(curBalanceSheetAsReported) {
			if TotalLiabilitiesAsReported != nil && InventoryAsReported != nil {
				NetDebtAsReported = utils.InterfaceToFloat64Ptr(*TotalLiabilitiesAsReported - *InventoryAsReported)
			}
		}
		// Long-Term Investments
		// Short-Term Liabilities
		// High Quality Liquid Assets
		// COGS
		// Average Inventory Total
		// Net Fixed Assets
		// Total Investments
		// Working Capital
		// Tangible Net Worth
		// Deferred Tax Liabilities
		// Common Shareholder Equity
		// Total Shareholder Equity
		// Total Inventory
		// Cash and Cash Equivalents
		// Accounts Receivable
		// Marketable Securities
		// Book Value of Equity
		// Shares Outstanding
		// Book Value of Debt
		// Equity Book Value
		// Liabilities Book Value
		// Total Accruals to Total Assets

		/* Income Statement */

		// EBITDA
		// Net Income
		// Gross Profit
		// Operating Income
		// Net Revenue
		// Net Profit Margin
		// Operating Expenses
		// Return on Assets
		// NOPAT

		/* Cash Flow Statement */

		// Depreciation & Amortization
		// Total Interest Payments
		// Total Taxes Paid
		// Change in Working Capital
		// Capital Expenditures
		// Operating Cash Flow
		// Funds From Operations
		// Free Cashflow
		// Operating Cashflow Per Share
		// Free Cashflow Per Share

		/* Calculated or Derived */

		// EBIT
		// Non-Cash Charges
		// Tax Rate
		// Market Value of Equity
		// Market Value of Debt
		// Cost of Equity
		// Cost of Debt
		// Risk Free Rate
		// Market Return
		// Beta
		// Percent change In Quantity Demanded
		// Percent change in Cost of Goods Sold
		// Percent change in Total Expenses
		// Percent change in Quantity of Units Produced
		// Preferred Stock Dividend Per Share
		// Market Value of Preferred Stock
		// Market Value of Stock
		// Upcoming Dividend Yield
		// Expected Growth Rate
		// Unlevered Firm Value
		// Net Effect of Debt
		// Lease Payments
		// Net Operating Income
		// Total Debt Service
		// Net Present Value of CashFlow
		// Total Preferred Dividend Payments
		// Net Credit Sales
		// Average Accounts Receivable
		// Carrying Cost Per Unit
		// Ordering Cost Per Order
		// Annual Demand
		// Non-Interest Expenses
		// Number of Employees
		// Variable Costs
		// Enterprise Value
		// Percent Change in Income
		// Total Loans Outstanding
		// Total Deposits
		// Non-Performing Assets
		// Short Term Debt
		// Long Term Debt
		// Asset Turnover Ratio
		// Equity Multiplier Ratio
		// Percent Change in EPS
		// Percent Change in EBIT
		// Depreciation Expenses Alone
		// Amortization Expenses Alone
		// Exploration Expenses Alone
		// Retention Ratio
		// Return on Equity
		// Explicit Costs
		// Implicit Costs
		// Period In Days
		// Market Capitalization
		// Equity Market Value
		// Liabilities Market Value
		// Quality of Earnings
		// Accounts Receivable Turnover Ratio
		// Supplier Purchases
		// Average Accounts Payable
		// Accounts Payable Turnover Ratio
		// Days Inventory Outstanding
		// Days Sales Outstanding
		// WACC
		// Total Capital
		// Total Invested Capital
		// Days Sales in Receivables Index
		// Gross Margin Index
		// Sales Growth Index
		// Sales General and Administrative Expenses Index

	}

	return nil, nil
}
