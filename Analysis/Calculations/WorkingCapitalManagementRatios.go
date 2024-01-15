package Calculations

func AverageCollectionPeriod(AccountsReceivableTurnoverRatio float64) float64 {
	// The average collection period is the average number of days in a year that a company takes to collect payments owed, after a sale of its goods or services has been made.

	return 365 / AccountsReceivableTurnoverRatio
}

func AccountsPayableTurnoverRatio(SupplierPurchases float64, AverageAccountsPayable float64) float64 {
	// The accounts payable turnover ratio is a short-term liquidity measure used to quantify the rate at which a company pays off its suppliers. Accounts payable turnover ratio is calculated by taking the total purchases made from suppliers and dividing it by the average accounts payable amount during the same period.

	return SupplierPurchases / AverageAccountsPayable
}

func AverageAccountsPayablePaymentPeriod(AccountsPayableTurnoverRatio float64) float64 {
	// The average payable period is the average number of days it takes a company to pay off credit accounts payable.

	return 365 / AccountsPayableTurnoverRatio
}

func InventoryToWorkingCapitalRatio(Inventory float64, WorkingCapital float64) float64 {
	// The inventory to working capital ratio is a liquidity ratio that measures a company's ability to pay off its current liabilities with inventory.

	return Inventory / WorkingCapital
}

func CashConversionCycle(DaysInventoryOutstanding float64, DaysSalesOutstanding float64, DaysPayablesOutstanding float64) float64 {
	// The cash conversion cycle (CCC) is a metric that expresses the length of time, in days, that it takes for a company to convert resource inputs into cash flows. The cash conversion cycle attempts to measure the amount of time each net input dollar is tied up in the production and sales process before it is converted into cash through sales to customers.

	return DaysInventoryOutstanding + DaysSalesOutstanding - DaysPayablesOutstanding
}

func NetWorkingCapital(AccountsReceivable float64, Inventory float64, AccountsPayable float64) float64 {
	// Net working capital is calculated as current assets minus current liabilities. It is a derivation of working capital, that is commonly used in valuation techniques such as discounted cash flows (DCF) that look at cash flows over the next five to ten years.

	return AccountsReceivable + Inventory - AccountsPayable
}

func EconomicValueAdded(NOPAT, WACC, TotalCapital float64) float64 {
	// Economic value added (EVA) is a measure of a company's financial performance based on the residual wealth calculated by deducting its cost of capital from its operating profit, adjusted for taxes on a cash basis.

	return NOPAT - (WACC * TotalCapital)
}

func ReturnOnInvestedCapital(NOPAT, TotalInvestedCapital float64) float64 {
	// Return on invested capital (ROIC) is a calculation used to assess a company's efficiency at allocating the capital under its control to profitable investments.

	return NOPAT / TotalInvestedCapital
}

func BeneishMScore(daysSalesInReceivablesIndex, grossMarginIndex, assetQualityIndex, salesGrowthIndex, depreciationIndex, salesGeneralAndAdministrativeExpensesIndex, totalAccrualsToTotalAssets, leverageIndex float64) float64 {
	// The Beneish M-score is a mathematical model that uses financial ratios and eight variables to identify whether a company has manipulated its earnings. The variables are constructed from the data in the company's financial statements and, once calculated, create an M-score to describe the degree to which the earnings have been manipulated.

	return -4.84 + 0.92*daysSalesInReceivablesIndex + 0.528*grossMarginIndex + 0.404*assetQualityIndex + 0.892*salesGrowthIndex + 0.115*depreciationIndex - 0.172*salesGeneralAndAdministrativeExpensesIndex + 4.679*totalAccrualsToTotalAssets - 0.327*leverageIndex
}
