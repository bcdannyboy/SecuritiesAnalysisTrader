package Calculations

func FreeCashFlowToEquity(EBITDA, DepreciationAmortization, InterestPayments, Taxes, ChangeInWorkingCapital, CapitalExpenditures, NetDebt float64) float64 {
	// Free cash flow to equity (FCFE) is a measure of how much cash can be paid to the equity shareholders of a company after all expenses, reinvestment, and debt are paid. FCFE is a measure of equity capital usage.

	return EBITDA - (CapitalExpenditures - DepreciationAmortization) - ChangeInWorkingCapital + (NetDebt - InterestPayments) - Taxes
}

func FreeCashFlowToFirm(NetIncome, NonCashCharges, InterestPayments, TaxRate, LongTermInvestments, InvestmentsInWorkingCapital float64) float64 {
	// Free cash flow to the firm (FCFF) represents the amount of cash flow from operations available for distribution after certain expenses are paid. FCF is a measure of a company's profitability after all expenses and reinvestments.

	return NetIncome + NonCashCharges - (InterestPayments * (1 - TaxRate)) - LongTermInvestments - InvestmentsInWorkingCapital
}
