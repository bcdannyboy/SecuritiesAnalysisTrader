package Calculations

import (
	"fmt"
	"math"
)

func LoanToDepositRatio(Loans, Deposits float64) float64 {
	// The loan-to-deposit ratio (LDR) is used to assess a bank's liquidity by comparing a bank's total loans to its total deposits for the same period. The LDR is expressed as a percentage.
	if Deposits == 0 {
		fmt.Printf("Got 0 for Deposits, with Loans %f and Deposits %f\n", Loans, Deposits)
		return 0
	}
	return Loans / Deposits
}

func TangibleNetWorthRatio(TNW float64, TotalAssets float64) float64 {
	// Tangible net worth is most commonly a calculation of the net worth of a company that excludes any value derived from intangible assets such as copyrights.
	if TotalAssets == 0 {
		fmt.Printf("Got 0 for TotalAssets, with TNW %f and TotalAssets %f\n", TNW, TotalAssets)
		return 0
	}
	return TNW / TotalAssets
}

func NonPerformingAssetRatio(NonPerformingAssets float64, TotalAssets float64) float64 {
	// The nonperforming asset ratio is a measurement of the percentage of nonperforming assets to the total assets of a bank or company. A nonperforming asset refers to loans or advances that are in jeopardy of default.
	if TotalAssets == 0 {
		fmt.Printf("Got 0 for TotalAssets, with NonPerformingAssets %f and TotalAssets %f\n", NonPerformingAssets, TotalAssets)
		return 0
	}

	return NonPerformingAssets / TotalAssets
}

func DeferredTaxLiabilityToEquityRatio(DeferredTaxLiabilities float64, ShareHolderEquity float64) float64 {
	// Deferred tax liability is a tax that is assessed or is due for the current period but has not yet been paid. The deferral arises because of timing differences between the accrual of the tax and payment of the tax.
	if ShareHolderEquity == 0 {
		fmt.Printf("Got 0 for ShareHolderEquity, with DeferredTaxLiabilities %f and ShareHolderEquity %f\n", DeferredTaxLiabilities, ShareHolderEquity)
		return 0
	}
	return DeferredTaxLiabilities / ShareHolderEquity
}

func TangibleEquityRatio(CommonShareHolderEquity, IntangibleAssets, TotalAssets float64) float64 {
	// The tangible common equity (TCE) ratio measures the percentage of a company’s common stock that is tangible common equity. The ratio is used to calculate a bank's ability to deal with potential losses. The higher the ratio, the more likely it is that the bank will be able to absorb the losses it incurs.
	if TotalAssets == 0 {
		fmt.Printf("Got 0 for TotalAssets, with CommonShareHolderEquity %f, IntangibleAssets %f and TotalAssets %f\n", CommonShareHolderEquity, IntangibleAssets, TotalAssets)
		return 0
	}
	return (CommonShareHolderEquity - IntangibleAssets) / (TotalAssets - IntangibleAssets)
}

func OhlsonOScore(totalAssets, totalLiabilities, workingCapital, currentLiabilities, currentAssets, netIncome, fundsFromOperations float64, liabilitiesExceedAssets, negativeNetIncomeLastTwoYears int) float64 {
	// The Ohlson O-Score is a statistical model that predicts the probability of a firm going bankrupt. It is a useful tool for investors to assess the financial health of a company and to determine whether it is a good investment.
	if totalAssets == 0 {
		fmt.Printf("Got 0 for totalAssets, with totalAssets %f, totalLiabilities %f, workingCapital %f, currentLiabilities %f, currentAssets %f, netIncome %f, fundsFromOperations %f, liabilitiesExceedAssets %d, negativeNetIncomeLastTwoYears %d\n", totalAssets, totalLiabilities, workingCapital, currentLiabilities, currentAssets, netIncome, fundsFromOperations, liabilitiesExceedAssets, negativeNetIncomeLastTwoYears)
		return 0
	}
	if currentAssets == 0 {
		fmt.Printf("Got 0 for currentAssets, with totalAssets %f, totalLiabilities %f, workingCapital %f, currentLiabilities %f, currentAssets %f, netIncome %f, fundsFromOperations %f, liabilitiesExceedAssets %d, negativeNetIncomeLastTwoYears %d\n", totalAssets, totalLiabilities, workingCapital, currentLiabilities, currentAssets, netIncome, fundsFromOperations, liabilitiesExceedAssets, negativeNetIncomeLastTwoYears)
		return 0
	}
	if totalLiabilities == 0 {
		fmt.Printf("Got 0 for totalLiabilities, with totalAssets %f, totalLiabilities %f, workingCapital %f, currentLiabilities %f, currentAssets %f, netIncome %f, fundsFromOperations %f, liabilitiesExceedAssets %d, negativeNetIncomeLastTwoYears %d\n", totalAssets, totalLiabilities, workingCapital, currentLiabilities, currentAssets, netIncome, fundsFromOperations, liabilitiesExceedAssets, negativeNetIncomeLastTwoYears)
		return 0
	}
	return -1.32 -
		0.407*math.Log(totalAssets) +
		6.03*(totalLiabilities/totalAssets) -
		1.43*(workingCapital/totalAssets) +
		0.0757*(currentLiabilities/currentAssets) -
		1.72*float64(liabilitiesExceedAssets) -
		2.37*(netIncome/totalAssets) -
		1.83*(fundsFromOperations/totalLiabilities) +
		0.285*float64(negativeNetIncomeLastTwoYears) -
		0.521*(float64(negativeNetIncomeLastTwoYears)*(netIncome/totalAssets))
}

func AltmanZScore(workingCapital, retainedEarnings, earningsBeforeInterestAndTaxes, marketValueOfEquity, sales, totalAssets, totalLiabilities float64) float64 {
	// The Altman Z-score is the output of a credit-strength test that gauges a publicly traded manufacturing company's likelihood of bankruptcy. The Altman Z-score is based on five financial ratios that can be calculated from data found on a company's annual 10K report.
	if totalAssets == 0 {
		fmt.Printf("Got 0 for totalAssets, with workingCapital %f, retainedEarnings %f, earningsBeforeInterestAndTaxes %f, marketValueOfEquity %f, sales %f, totalAssets %f, totalLiabilities %f\n", workingCapital, retainedEarnings, earningsBeforeInterestAndTaxes, marketValueOfEquity, sales, totalAssets, totalLiabilities)
		return 0
	}
	if totalLiabilities == 0 {
		fmt.Printf("Got 0 for totalLiabilities, with workingCapital %f, retainedEarnings %f, earningsBeforeInterestAndTaxes %f, marketValueOfEquity %f, sales %f, totalAssets %f, totalLiabilities %f\n", workingCapital, retainedEarnings, earningsBeforeInterestAndTaxes, marketValueOfEquity, sales, totalAssets, totalLiabilities)
		return 0
	}
	X1 := workingCapital / totalAssets
	X2 := retainedEarnings / totalAssets
	X3 := earningsBeforeInterestAndTaxes / totalAssets
	X4 := marketValueOfEquity / totalLiabilities
	X5 := sales / totalAssets

	return 1.2*X1 + 1.4*X2 + 3.3*X3 + 0.6*X4 + 1.0*X5
}
