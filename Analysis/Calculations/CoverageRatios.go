package Calculations

import "fmt"

func InterestCoverageRatio(EBIT float64, InterestExpense float64) float64 {
	// The interest coverage ratio is used to determine how easily a company can pay their interest expenses on outstanding debt. The ratio is calculated by dividing a company's earnings before interest and taxes (EBIT) by the company's interest expenses for the same period.

	if InterestExpense == 0 {
		fmt.Printf("Interest Expense cannot be 0, got EBIT: %f, Interest Expense: %f\n", EBIT, InterestExpense)
		return 0
	}
	return EBIT / InterestExpense
}

func FixedChargeCoverageRatio(EBIT, LeasePayments, InterestExpense float64) float64 {
	// The fixed-charge coverage ratio (FCCR) is a measure of a firm's ability to cover its fixed charges, such as debt payments, interest expense, and equipment lease expense. The ratio is calculated by dividing a company's earnings before interest and taxes (EBIT) by its total fixed charges.

	if (LeasePayments + (LeasePayments * InterestExpense)) == 0 {
		fmt.Printf("Lease Payments + (Lease Payments * Interest Expense) cannot be 0, got EBIT: %f, Lease Payments: %f, Interest Expense: %f\n", EBIT, LeasePayments, InterestExpense)
		return 0
	}
	return (EBIT + LeasePayments) / (LeasePayments + (LeasePayments * InterestExpense))
}

func DebtServiceCoverageRatio(NetOperatingIncome float64, TotalDebtService float64) float64 {
	// The debt service coverage ratio (DSCR) measures the ability of a company to use its operating income to repay all its debt obligations, including repayment of principal and interest on both short-term and long-term debt.
	if TotalDebtService == 0 {
		fmt.Printf("Total Debt Service cannot be 0, got Net Operating Income: %f, Total Debt Service: %f\n", NetOperatingIncome, TotalDebtService)
		return 0
	}
	return NetOperatingIncome / TotalDebtService
}

func AssetCoverageRatio(TotalAssets, ShortTermLiabilities, TotalDebt float64) float64 {
	// The asset coverage ratio determines a company's ability to cover debt obligations with its assets after all liabilities have been satisfied. The ratio can be calculated by dividing a company's total assets by its total liabilities.
	if TotalDebt == 0 {
		fmt.Printf("Total Debt cannot be 0, got Total Assets: %f, Short Term Liabilities: %f, Total Debt: %f\n", TotalAssets, ShortTermLiabilities, TotalDebt)
		return 0
	}
	return (TotalAssets - ShortTermLiabilities) / TotalDebt
}

func LoanLifeCoverageRatio(NetPresentValueOfCashFlow float64, TotalDebt float64) float64 {
	// The loan life coverage ratio (LLCR) is a financial ratio used to estimate the ability of the borrowing company to repay an outstanding loan. The loan life coverage ratio is calculated by dividing the net present value (NPV) of the cash flow available for debt repayment by the amount of debt outstanding.
	if TotalDebt == 0 {
		fmt.Printf("Total Debt cannot be 0, got Net Present Value Of Cash Flow: %f, Total Debt: %f\n", NetPresentValueOfCashFlow, TotalDebt)
		return 0
	}
	return NetPresentValueOfCashFlow / TotalDebt
}

func EBITDAToInterestCoverageRatio(EBITDA, InterestExpense float64) float64 {
	// The EBITDA-to-interest coverage ratio is used to assess a company's financial durability by examining whether it is at least profitable enough to pay off its interest expenses. The ratio is calculated by dividing a company's earnings before interest, taxes, depreciation, and amortization (EBITDA) by the company's interest expenses for the same period.
	if InterestExpense == 0 {
		fmt.Printf("Interest Expense cannot be 0, got EBITDA: %f, Interest Expense: %f\n", EBITDA, InterestExpense)
		return 0
	}
	return EBITDA / InterestExpense
}

func PreferredDividendCoverageRatio(NetIncome, PreferredDividends float64) float64 {
	// The preferred dividend coverage ratio is a coverage ratio that measures a company's ability to pay off its required preferred dividend payments. The preferred dividend coverage ratio is calculated by dividing a company's net income by its current preferred dividend payments.
	if PreferredDividends == 0 {
		fmt.Printf("Preferred Dividends cannot be 0, got Net Income: %f, Preferred Dividends: %f\n", NetIncome, PreferredDividends)
		return 0
	}
	return NetIncome / PreferredDividends
}

func LiquidityCoverageRatio(HighQualityLiquidAssets, TotalNetCashOutflows float64) float64 {
	// The liquidity coverage ratio (LCR) refers to the proportion of highly liquid assets held by financial institutions to ensure their ongoing ability to meet short-term obligations. The ratio is calculated by dividing a bank's stock of high-quality liquid assets by its total net cash outflows over the next 30 days.
	if TotalNetCashOutflows == 0 {
		fmt.Printf("Total Net Cash Outflows cannot be 0, got High Quality Liquid Assets: %f, Total Net Cash Outflows: %f\n", HighQualityLiquidAssets, TotalNetCashOutflows)
		return 0
	}
	return HighQualityLiquidAssets / TotalNetCashOutflows
}
