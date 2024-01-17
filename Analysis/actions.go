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

func PerformCustomCalculations(Fundamentals *CompanyFundamentals, Period objects.CompanyValuationPeriod, PricePerShare float64, EffectiveTaxRate float64, NumEmployees int) (map[string]float64, map[string]float64) {
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
			LongTermInvestmentsAsReported = utils.GetFloat64PtrIfNotEmpty(AssetsNonCurrentAsReported - (NetFixedAssetsAsReported + NonCurrentMarketableSecuritiesAsReported + OtherAssetsNonCurrentAsReported))
		}

		var TotalMarketableSecurities *float64 = nil
		if ShortTermInvestments != nil && LongTermInvestments != nil {
			TotalMarketableSecurities = utils.InterfaceToFloat64Ptr(*ShortTermInvestments + *LongTermInvestments)
		}

		var TotalInvestments = utils.GetFloat64PtrIfNotEmpty(curBalanceSheet, "TotalInvestments")
		var TotalInvestmentsAsReported *float64 = nil
		if TotalMarketableSecuritiesAsReported != nil {
			var TotalInvestmentsAsReported = TotalMarketableSecuritiesAsReported
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
		if NetCashOperatingActivitiesAsReported != nil CapitalExpendituresAsReported != nil {
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
		var CurrentLongTermDebtAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Longtermdebtcurrent")
		var NonCurrentLongTermDebtAsReported = utils.GetFloat64PtrIfNotEmpty(curBalanceSheetAsReported, "Longtermdebtnoncurrent")
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
			UnleveredFirmValue = utils.InterfaceToFloat64Ptr((*EBIT * (1 - EffectiveTaxRate)) + *DepreciationAndAmortization)
		}
		if EBITAsReported != nil && DepreciationAndAmortizationAsReported != nil {
			UnleveredFirmValueAsReported = utils.InterfaceToFloat64Ptr((*EBITAsReported * (1 - EffectiveTaxRate)) + *DepreciationAndAmortizationAsReported)&
		}

		var TaxShieldBenefits *float64 = nil
		var TaxShieldBenefitsAsReported *float64 = nil
		if TotalInterestPayments != nil {
			TaxShieldBenefits = utils.InterfaceToFloat64Ptr(*TotalInterestPayments * EffectiveTaxRate)
		}
		if TotalInterestPaymentsAsReported != nil {
			TaxShieldBenefitsAsReported = utils.InterfaceToFloat64Ptr(*TotalInterestPaymentsAsReported * EffectiveTaxRate)
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
		if CostOfGoodsSoldAsReported != nil && OperatingExpensesAsReported != nil  && SellingGeneralAndAdministrativeExpensesAsReported != nil {
            ExplicitCostsAsReported = utils.InterfaceToFloat64Ptr(*CostOfGoodsSoldAsReported + *OperatingExpensesAsReported + *SellingGeneralAndAdministrativeExpensesAsReported)
        }

		var DaysInvewntoryOutstanding *float64 = nil
		var DaysInvewntoryOutstandingAsReported *float64 = nil
		if Inventory != nil && CostOfRevenue != nil {
            DaysInvewntoryOutstanding = utils.InterfaceToFloat64Ptr((*Inventory / *CostOfRevenue) * DaysInPeriod)
        }
		if InventoryAsReported != nil && CostOfGoodsSoldAsReported != nil {
            DaysInvewntoryOutstandingAsReported = utils.InterfaceToFloat64Ptr((*InventoryAsReported / *CostOfGoodsSoldAsReported) * DaysInPeriod)
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
	}

	return nil, nil
}
