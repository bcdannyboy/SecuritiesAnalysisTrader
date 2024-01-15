package Calculations

import "math"

func InventoryTurnoverRatio(CostOfGoodsSold float64, AverageInventory float64) float64 {
	// The inventory turnover ratio is a key measure for evaluating how effective a company's management is at managing inventory levels and generating sales from it. The formula for calculating inventory turnover equals the cost of goods sold divided by the average inventory.

	return CostOfGoodsSold / AverageInventory
}

func AccountsReceivableTurnoverRatio(NetCreditSales float64, AverageAccountsReceivable float64) float64 {
	// The accounts receivable turnover ratio measures a companies effectiveness in terms of qualifying their credit borrowers and collecting monies owed from them. The A/R turnover ratio is an indication to how many times the accounts receivables are "turned over" throughout the year

	return NetCreditSales / AverageAccountsReceivable
}

func FixedAssetTurnoverRatio(NetSales float64, NetFixedAssets float64) float64 {
	// The fixed-asset turnover ratio measures a company’s ability to generate net sales from fixed-asset investments – specifically property, plant and equipment (PP&E) – net of depreciation. The fixed-asset turnover ratio is calculated by dividing net sales by net fixed assets.

	return NetSales / NetFixedAssets
}

func PPETurnoverRatio(NetSales float64, NetFixedAssets float64, AccumulatedDepreciation float64) float64 {
	// The fixed-asset turnover ratio measures a company’s ability to generate net sales from fixed-asset investments – specifically property, plant and equipment (PP&E) – net of depreciation. The fixed-asset turnover ratio is calculated by dividing net sales by net fixed assets.

	return NetSales / (NetFixedAssets - AccumulatedDepreciation)
}

func InvestmentTurnoverRatio(NetSales, TotalInvestments float64) float64 {
	// The investment turnover ratio measures a company's efficiency in using its assets to generate revenue or sales. The ratio formula is to divide net sales by the average total assets.

	return NetSales / TotalInvestments
}

func WorkingCapitalTurnoverRatio(NetSales, WorkingCapital float64) float64 {
	// The working capital turnover ratio measures how efficiently a company is using its working capital to support a given level of sales. It is calculated by dividing net sales by the average working capital.

	return NetSales / WorkingCapital
}

func EconomicOrderQuantity(CarryingCostPerUnit, OrderingCostPerOrder, AnnualDemand float64) float64 {
	// Economic order quantity (EOQ) is an equation for inventory that determines the ideal order quantity a company should purchase for its inventory given a set cost of production, demand rate, and other variables.

	return math.Sqrt((2 * AnnualDemand * OrderingCostPerOrder) / CarryingCostPerUnit)
}

func ReturnOnCapitalEmployed(EBIT, TotalAssets, TotalLiabilities float64) float64 {
	// Return on capital employed (ROCE) is a financial ratio that measures a company's profitability and the efficiency with which its capital is employed. ROCE is calculated as: ROCE = EBIT / Capital Employed.

	return EBIT / (TotalAssets - TotalLiabilities)
}

func EfficiencyRatio(NonInterestExpense, NetRevenue float64) float64 {
	// The efficiency ratio is typically used to analyze how well a company uses its assets and liabilities internally. An efficiency ratio can calculate the turnover of receivables, the repayment of liabilities, the quantity and usage of equity, and the general use of inventory and machinery.

	return NonInterestExpense / NetRevenue
}

func RevenuePerEmployee(NetRevenue, NumberOfEmployees float64) float64 {
	// Revenue per employee is an important ratio that looks at a company's revenue in relation to the number of employees it has. The ratio is generally used as a way of measuring productivity.

	return NetRevenue / NumberOfEmployees
}

func VariableCostRatio(VariableCosts, NetRevenue float64) float64 {
	// The variable cost ratio is the proportion of variable costs to sales. It is computed by dividing the variable costs by the net revenue.

	return VariableCosts / NetRevenue
}

func CapitalExpenditureRatio(CapitalExpenditures, OperatingCashFlow float64) float64 {
	// The capital expenditure ratio is a ratio that measures the amount of a company's capital expenditures to the amount of its operating cash flow.

	return CapitalExpenditures / OperatingCashFlow
}

func OperatingCashFlowRatio(OperatingCashFlow, NetSales float64) float64 {
	// The operating cash flow ratio is a measure of how well current liabilities are covered by the cash flows generated from a company's operations.

	return OperatingCashFlow / NetSales
}

func EBITDAToEVRatio(EBITDA, EnterpriseValue float64) float64 {
	// The EBITDA-to-enterprise value (EV) ratio is a measure of a company's return on investment (ROI). The EV/EBITDA ratio compares a company's enterprise value (EV) to its earnings before interest, taxes, depreciation, and amortization (EBITDA).

	return EBITDA / EnterpriseValue
}

func IncomeElasticityOfDemand(PercentChangeInQuantityDemanded, PercentChangeInIncome float64) float64 {
	// Income elasticity of demand is a measure of the relationship between a change in the quantity demanded for a particular good and a change in real income.

	return PercentChangeInQuantityDemanded / PercentChangeInIncome
}
