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
