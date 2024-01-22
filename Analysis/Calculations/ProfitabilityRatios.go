package Calculations

import "fmt"

func ReturnOnAssetRatio(NetIncome float64, TotalAssets float64) float64 {
	// The return on assets ratio, often called the return on total assets, is a profitability ratio that measures the net income produced by total assets during a period by comparing net income to the average total assets.
	if TotalAssets == 0 {
		fmt.Printf("Got 0 for TotalAssets with NetIncome: %f and TotalAssets: %f\n", NetIncome, TotalAssets)
		return 0
	}
	return NetIncome / TotalAssets
}

func GrossProfitMargin(GrossProfit float64, NetSales float64) float64 {
	// Gross profit margin is a ratio that indicates the performance of a company's sales and production. It measures how much out of every dollar of sales a company actually keeps in earnings.
	if NetSales == 0 {
		fmt.Printf("Got 0 for NetSales with GrossProfit: %f and NetSales: %f\n", GrossProfit, NetSales)
		return 0
	}
	return GrossProfit / NetSales
}

func OperatingProfitMargin(OperatingIncome float64, NetSales float64) float64 {
	// Operating margin is a profitability ratio measuring revenue after covering operating and non-operating expenses of a business. It is calculated by dividing operating profit by revenue.
	if NetSales == 0 {
		fmt.Printf("Got 0 for NetSales with OperatingIncome: %f and NetSales: %f\n", OperatingIncome, NetSales)
		return 0
	}
	return OperatingIncome / NetSales
}

func EBITDAMarginRatio(EBITDA, NetSales float64) float64 {
	// EBITDA margin is a profitability ratio that measures how much earnings a company is generating before interest, taxes, depreciation, and amortization, as a percentage of revenue.
	if NetSales == 0 {
		fmt.Printf("Got 0 for NetSales with EBITDA: %f and NetSales: %f\n", EBITDA, NetSales)
		return 0
	}
	return EBITDA / NetSales
}

func DividendPayoutRatio(Dividends, NetIncome float64) float64 {
	// The dividend payout ratio is the ratio of the total amount of dividends paid out to shareholders relative to the net income of the company. It is the percentage of earnings paid to shareholders in dividends.
	if NetIncome == 0 {
		fmt.Printf("Got 0 for NetIncome with Dividends: %f and NetIncome: %f\n", Dividends, NetIncome)
		return 0
	}
	return Dividends / NetIncome
}

func RetentionRate(Dividends, NetIncome float64) float64 {
	// The retention ratio is the proportion of earnings kept back in the business as retained earnings. Retention ratio refers to the percentage of net income that is retained to grow the business, rather than being paid out as dividends.
	if NetIncome == 0 {
		fmt.Printf("Got 0 for NetIncome with Dividends: %f and NetIncome: %f\n", Dividends, NetIncome)
		return 0
	}
	return (NetIncome - Dividends) / NetIncome
}

func SustainableGrowthRate(RetentionRatio, ReturnOnEquity float64) float64 {
	// The sustainable growth rate (SGR) is the maximum rate of growth that a company can sustain without raising additional equity or taking on new debt. The sustainable growth rate can be calculated by multiplying a company's earnings retention ratio by its return on equity (ROE).

	return RetentionRatio * ReturnOnEquity
}

func GrossMarginOnInventory(GrossProfit, Inventory float64) float64 {
	// Gross margin on inventory is a ratio that measures how profitable a company is relative to its inventory. It is calculated by dividing gross profit by inventory.
	if Inventory == 0 {
		fmt.Printf("Got 0 for Inventory with GrossProfit: %f and Inventory: %f\n", GrossProfit, Inventory)
		return 0
	}
	return GrossProfit / Inventory
}

func Economicprofit(TotalRevenue, ExplicitCost, ImplicitCost float64) float64 {
	// Economic profit is the difference between the revenue received from the sale of an output and the costs of all inputs, including opportunity costs.

	return TotalRevenue - (ExplicitCost + ImplicitCost)
}

func CashFlowReturnOnEquity(OperatingCashFlow, ShareHolderEquity float64) float64 {
	// Cash flow return on equity (CFROE) is a measure of a company's ability to generate cash flow from equity. It is calculated by dividing cash flow from operations by shareholders' equity.
	if ShareHolderEquity == 0 {
		fmt.Printf("Got 0 for ShareHolderEquity with OperatingCashFlow: %f and ShareHolderEquity: %f\n", OperatingCashFlow, ShareHolderEquity)
		return 0
	}
	return OperatingCashFlow / ShareHolderEquity
}

func OperatingMargin(Revenue, COGS float64) float64 {
	// operating margin is a profitability ratio that measures how much of every dollar of revenue is left over after paying the cost of goods sold (COGS).
	if COGS == 0 {
		fmt.Printf("Got 0 for COGS with Revenue: %f and COGS: %f\n", Revenue, COGS)
		return 0
	}
	return Revenue / COGS
}

func OperatingExpenseRatio(OperatingExpenses, DepreciationAmortization, OperatingIncome float64) float64 {
	// The operating expense ratio (OER) is a measure of what it costs to operate a piece of property compared to the income that the property brings in. The operating expense ratio is calculated by dividing a property's operating expense by its gross operating income.
	if OperatingIncome == 0 {
		fmt.Printf("Got 0 for OperatingIncome with OperatingExpenses: %f and OperatingIncome: %f\n", OperatingExpenses, OperatingIncome)
		return 0
	}
	return (OperatingExpenses + DepreciationAmortization) / OperatingIncome
}
