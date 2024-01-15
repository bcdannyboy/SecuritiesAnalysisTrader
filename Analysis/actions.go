package Analysis

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/BalanceSheet"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/CashFlowStatement"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/FinancialRatios"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/FullFinancialStatement"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/IncomeStatement"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func PullCompanyFundamentals(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) (*CompanyFundamentals, error) {
	// Pull fundamentals / ratios from FMP / calculations for a given ticker
	FundamentalsObj := &CompanyFundamentals{
		Symbol: Symbol,
	}

	BS_STMT, BS_STMT_GROWTH, BS_STMT_AS_REPORTED, BS_STMT_AS_REPORTED_GROWTH, BS_DISCREPANCIES, err := BalanceSheet.AnalyzeBalanceSheet(APIClient, Symbol, Period)
	if err != nil {
		if len(BS_STMT) == 0 {
			fmt.Printf("Failed to get balance sheet statement for %s: %s\n", Symbol, err.Error())
		} else if len(BS_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get balance sheet statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.BalanceSheetStatementGrowth = BS_STMT_GROWTH
		} else if len(BS_STMT_AS_REPORTED) == 0 {
			fmt.Printf("Failed to get balance sheet statement as reported for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.BalanceSheetStatementGrowth = BS_STMT_GROWTH
			FundamentalsObj.BalanceSheetStatementAsReported = BS_STMT_AS_REPORTED
		}
	} else {
		FundamentalsObj.BalanceSheetStatements = BS_STMT
		FundamentalsObj.BalanceSheetStatementGrowth = BS_STMT_GROWTH
		FundamentalsObj.BalanceSheetStatementAsReported = BS_STMT_AS_REPORTED
	}

	FundamentalsObj.GrowthBalanceSheetStatementAsReported = BS_STMT_AS_REPORTED_GROWTH
	FundamentalsObj.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported = BS_DISCREPANCIES

	CF_STMT, CF_STMT_GROWTH, CF_STMT_AS_REPORTED, CF_STMT_AS_REPORTED_GROWTH, CF_DISCREPANCIES, err := CashFlowStatement.AnalyzeCashFlow(APIClient, Symbol, Period)
	if err != nil {
		if len(CF_STMT) == 0 {
			fmt.Printf("Failed to get cash flow statement for %s: %s\n", Symbol, err.Error())
		} else if len(CF_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get cash flow statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.CashFlowStatementGrowth = CF_STMT_GROWTH
		} else if len(CF_STMT_AS_REPORTED) == 0 {
			fmt.Printf("Failed to get cash flow statement as reported for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.CashFlowStatementGrowth = CF_STMT_GROWTH
			FundamentalsObj.CashFlowStatementAsReported = CF_STMT_AS_REPORTED
		}
	} else {
		FundamentalsObj.CashFlowStatement = CF_STMT
		FundamentalsObj.CashFlowStatementGrowth = CF_STMT_GROWTH
		FundamentalsObj.CashFlowStatementAsReported = CF_STMT_AS_REPORTED
	}

	FundamentalsObj.CashFlowStatementAsReportedGrowth = CF_STMT_AS_REPORTED_GROWTH
	FundamentalsObj.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported = CF_DISCREPANCIES

	I_STMT, I_STMT_GROWTH, I_STMT_AS_REPORTED, I_STMT_AS_REPORTED_GROWTH, I_DISCREPANCIES, err := IncomeStatement.AnalyzeIncomeStatement(APIClient, Symbol, Period)
	if err != nil {
		if len(I_STMT) == 0 {
			fmt.Printf("Failed to get income statement for %s: %s\n", Symbol, err.Error())
		} else if len(I_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get income statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.IncomeStatement = I_STMT
		} else if len(I_STMT_AS_REPORTED) == 0 {
			fmt.Printf("Failed to get income statement as reported for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.IncomeStatementGrowth = I_STMT_GROWTH
			FundamentalsObj.IncomeStatementAsReported = I_STMT_AS_REPORTED
		}
	} else {
		FundamentalsObj.IncomeStatement = I_STMT
		FundamentalsObj.IncomeStatementGrowth = I_STMT_GROWTH
		FundamentalsObj.IncomeStatementAsReported = I_STMT_AS_REPORTED
	}

	FundamentalsObj.GrowthIncomeStatementAsReported = I_STMT_AS_REPORTED_GROWTH
	FundamentalsObj.DiscrepancyIncomeStatementAndIncomeStatementAsReported = I_DISCREPANCIES

	F_STMT, F_STMT_GROWTH, err := FullFinancialStatement.AnalyzeFinancialStatement(APIClient, Symbol, Period)
	if err != nil {
		if len(F_STMT) == 0 {
			fmt.Printf("Failed to get full financial statement for %s: %s\n", Symbol, err.Error())
		} else if len(F_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get full financial statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.FullFinancialStatement = F_STMT
		}
	} else {
		FundamentalsObj.FullFinancialStatement = F_STMT
		FundamentalsObj.FullFinancialStatementGrowth = F_STMT_GROWTH
	}

	FIN_RATIOS, FIN_RATIOS_TTM, FR_GROWTH, FR_TTM_GROWTH, err := FinancialRatios.AnalyzeFinancialRatios(APIClient, Symbol, Period)
	if err != nil {
		if len(FIN_RATIOS) == 0 {
			print("Failed to get financial ratios for %s: %s\n", Symbol, err.Error())
		} else if len(FIN_RATIOS_TTM) == 0 {
			print("Failed to get financial ratios TTM for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.FinancialRatios = FIN_RATIOS
		}
	} else {
		FundamentalsObj.FinancialRatios = FIN_RATIOS
		FundamentalsObj.FinancialRatiosTTM = FIN_RATIOS_TTM
	}

	FundamentalsObj.FinancialRatiosGrowth = FR_GROWTH
	FundamentalsObj.FinancialRatiosTTMGrowth = FR_TTM_GROWTH

	return FundamentalsObj, nil
}
