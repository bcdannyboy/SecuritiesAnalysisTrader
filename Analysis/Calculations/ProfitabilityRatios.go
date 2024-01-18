package Calculations

func ReturnOnAssetRatio(NetIncome float64, TotalAssets float64) float64 {
	// The return on assets ratio, often called the return on total assets, is a profitability ratio that measures the net income produced by total assets during a period by comparing net income to the average total assets.

	return NetIncome / TotalAssets
}

func GrossProfitMargin(GrossProfit float64, NetSales float64) float64 {
	// Gross profit margin is a ratio that indicates the performance of a company's sales and production. It measures how much out of every dollar of sales a company actually keeps in earnings.

	return GrossProfit / NetSales
}

func OperatingProfitMargin(OperatingIncome float64, NetSales float64) float64 {
	// Operating margin is a profitability ratio measuring revenue after covering operating and non-operating expenses of a business. It is calculated by dividing operating profit by revenue.

	return OperatingIncome / NetSales
}

func EBITDAMarginRatio(EBITDA, NetSales float64) float64 {
	// EBITDA margin is a profitability ratio that measures how much earnings a company is generating before interest, taxes, depreciation, and amortization, as a percentage of revenue.

	return EBITDA / NetSales
}

func DividendPayoutRatio(Dividends, NetIncome float64) float64 {
	// The dividend payout ratio is the ratio of the total amount of dividends paid out to shareholders relative to the net income of the company. It is the percentage of earnings paid to shareholders in dividends.

	return Dividends / NetIncome
}

func RetentionRate(Dividends, NetIncome float64) float64 {
	// The retention ratio is the proportion of earnings kept back in the business as retained earnings. Retention ratio refers to the percentage of net income that is retained to grow the business, rather than being paid out as dividends.

	return (NetIncome - Dividends) / NetIncome
}

func SustainableGrowthRate(RetentionRatio, ReturnOnEquity float64) float64 {
	// The sustainable growth rate (SGR) is the maximum rate of growth that a company can sustain without raising additional equity or taking on new debt. The sustainable growth rate can be calculated by multiplying a company's earnings retention ratio by its return on equity (ROE).

	return RetentionRatio * ReturnOnEquity
}

func GrossMarginOnInventory(GrossProfit, Inventory float64) float64 {
	// Gross margin on inventory is a ratio that measures how profitable a company is relative to its inventory. It is calculated by dividing gross profit by inventory.

	return GrossProfit / Inventory
}

func Economicprofit(TotalRevenue, ExplicitCost, ImplicitCost float64) float64 {
	// Economic profit is the difference between the revenue received from the sale of an output and the costs of all inputs, including opportunity costs.

	return TotalRevenue - (ExplicitCost + ImplicitCost)
}

func CashFlowReturnOnEquity(OperatingCashFlow, ShareHolderEquity float64) float64 {
	// Cash flow return on equity (CFROE) is a measure of a company's ability to generate cash flow from equity. It is calculated by dividing cash flow from operations by shareholders' equity.

	return OperatingCashFlow / ShareHolderEquity
}

func OperatingMargin(Revenue, COGS float64) float64 {
	// operating margin is a profitability ratio that measures how much of every dollar of revenue is left over after paying the cost of goods sold (COGS).

	return Revenue / COGS
}

func OperatingExpenseRatio(OperatingExpenses, DepreciationAmortization, OperatingIncome float64) float64 {
	// The operating expense ratio (OER) is a measure of what it costs to operate a piece of property compared to the income that the property brings in. The operating expense ratio is calculated by dividing a property's operating expense by its gross operating income.

	return (OperatingExpenses + DepreciationAmortization) / OperatingIncome
}
