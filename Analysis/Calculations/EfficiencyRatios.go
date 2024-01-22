package Calculations

import (
	"fmt"
	"math"
)

func InventoryTurnoverRatio(CostOfGoodsSold float64, AverageInventory float64) float64 {
	// The inventory turnover ratio is a key measure for evaluating how effective a company's management is at managing inventory levels and generating sales from it. The formula for calculating inventory turnover equals the cost of goods sold divided by the average inventory.
	if AverageInventory == 0 {
		fmt.Printf("got 0 for AverageInventory with CostOfGoodsSold: %f, AverageInventory: %f\n", CostOfGoodsSold, AverageInventory)
		return 0
	}
	return CostOfGoodsSold / AverageInventory
}

func AccountsReceivableTurnoverRatio(NetCreditSales float64, AverageAccountsReceivable float64) float64 {
	// The accounts receivable turnover ratio measures a companies effectiveness in terms of qualifying their credit borrowers and collecting monies owed from them. The A/R turnover ratio is an indication to how many times the accounts receivables are "turned over" throughout the year
	if AverageAccountsReceivable == 0 {
		fmt.Printf("got 0 for AverageAccountsReceivable with NetCreditSales: %f, AverageAccountsReceivable: %f\n", NetCreditSales, AverageAccountsReceivable)
		return 0
	}
	return NetCreditSales / AverageAccountsReceivable
}

func FixedAssetTurnoverRatio(NetSales float64, NetFixedAssets float64) float64 {
	// The fixed-asset turnover ratio measures a company’s ability to generate net sales from fixed-asset investments – specifically property, plant and equipment (PP&E) – net of depreciation. The fixed-asset turnover ratio is calculated by dividing net sales by net fixed assets.

	if NetFixedAssets == 0 {
		fmt.Printf("got 0 for NetFixedAssets with NetSales: %f, NetFixedAssets: %f\n", NetSales, NetFixedAssets)
		return 0
	}
	return NetSales / NetFixedAssets
}

func PPETurnoverRatio(NetSales float64, NetFixedAssets float64, AccumulatedDepreciation float64) float64 {
	// The fixed-asset turnover ratio measures a company’s ability to generate net sales from fixed-asset investments – specifically property, plant and equipment (PP&E) – net of depreciation. The fixed-asset turnover ratio is calculated by dividing net sales by net fixed assets.

	if (NetFixedAssets - AccumulatedDepreciation) == 0 {
		fmt.Printf("got 0 for NetFixedAssets - AccumulatedDepreciation with NetSales: %f, NetFixedAssets: %f, AccumulatedDepreciation: %f\n", NetSales, NetFixedAssets, AccumulatedDepreciation)
		return 0
	}
	return NetSales / (NetFixedAssets - AccumulatedDepreciation)
}

func InvestmentTurnoverRatio(NetSales, TotalInvestments float64) float64 {
	// The investment turnover ratio measures a company's efficiency in using its assets to generate revenue or sales. The ratio formula is to divide net sales by the average total assets.

	if TotalInvestments == 0 {
		fmt.Printf("got 0 for TotalInvestments with NetSales: %f, TotalInvestments: %f\n", NetSales, TotalInvestments)
		return 0
	}
	return NetSales / TotalInvestments
}

func WorkingCapitalTurnoverRatio(NetSales, WorkingCapital float64) float64 {
	// The working capital turnover ratio measures how efficiently a company is using its working capital to support a given level of sales. It is calculated by dividing net sales by the average working capital.

	if WorkingCapital == 0 {
		fmt.Printf("got 0 for WorkingCapital with NetSales: %f, WorkingCapital: %f\n", NetSales, WorkingCapital)
		return 0
	}
	return NetSales / WorkingCapital
}

func EconomicOrderQuantity(CarryingCostPerUnit, OrderingCostPerOrder, AnnualDemand float64) float64 {
	// Economic order quantity (EOQ) is an equation for inventory that determines the ideal order quantity a company should purchase for its inventory given a set cost of production, demand rate, and other variables.

	if CarryingCostPerUnit == 0 {
		fmt.Printf("got 0 for CarryingCostPerUnit with CarryingCostPerUnit: %f, OrderingCostPerOrder: %f, AnnualDemand: %f\n", CarryingCostPerUnit, OrderingCostPerOrder, AnnualDemand)
		return 0
	}
	return math.Sqrt((2 * AnnualDemand * OrderingCostPerOrder) / CarryingCostPerUnit)
}

func ReturnOnCapitalEmployed(EBIT, TotalAssets, TotalLiabilities float64) float64 {
	// Return on capital employed (ROCE) is a financial ratio that measures a company's profitability and the efficiency with which its capital is employed. ROCE is calculated as: ROCE = EBIT / Capital Employed.

	if (TotalAssets - TotalLiabilities) == 0 {
		fmt.Printf("got 0 for TotalAssets - TotalLiabilities with EBIT: %f, TotalAssets: %f, TotalLiabilities: %f\n", EBIT, TotalAssets, TotalLiabilities)
		return 0
	}
	return EBIT / (TotalAssets - TotalLiabilities)
}

func EfficiencyRatio(NonInterestExpense, NetRevenue float64) float64 {
	// The efficiency ratio is typically used to analyze how well a company uses its assets and liabilities internally. An efficiency ratio can calculate the turnover of receivables, the repayment of liabilities, the quantity and usage of equity, and the general use of inventory and machinery.

	if NetRevenue == 0 {
		fmt.Printf("got 0 for NetRevenue with NonInterestExpense: %f, NetRevenue: %f\n", NonInterestExpense, NetRevenue)
		return 0
	}
	return NonInterestExpense / NetRevenue
}

func RevenuePerEmployee(NetRevenue, NumberOfEmployees float64) float64 {
	// Revenue per employee is an important ratio that looks at a company's revenue in relation to the number of employees it has. The ratio is generally used as a way of measuring productivity.

	if NumberOfEmployees == 0 {
		fmt.Printf("got 0 for NumberOfEmployees with NetRevenue: %f, NumberOfEmployees: %f\n", NetRevenue, NumberOfEmployees)
		return 0
	}
	return NetRevenue / NumberOfEmployees
}

func VariableCostRatio(VariableCosts, NetRevenue float64) float64 {
	// The variable cost ratio is the proportion of variable costs to sales. It is computed by dividing the variable costs by the net revenue.

	if NetRevenue == 0 {
		fmt.Printf("got 0 for NetRevenue with VariableCosts: %f, NetRevenue: %f\n", VariableCosts, NetRevenue)
		return 0
	}
	return VariableCosts / NetRevenue
}

func CapitalExpenditureRatio(CapitalExpenditures, OperatingCashFlow float64) float64 {
	// The capital expenditure ratio is a ratio that measures the amount of a company's capital expenditures to the amount of its operating cash flow.

	if OperatingCashFlow == 0 {
		fmt.Printf("got 0 for OperatingCashFlow with CapitalExpenditures: %f, OperatingCashFlow: %f\n", CapitalExpenditures, OperatingCashFlow)
		return 0
	}
	return CapitalExpenditures / OperatingCashFlow
}

func OperatingCashFlowRatio(OperatingCashFlow, NetSales float64) float64 {
	// The operating cash flow ratio is a measure of how well current liabilities are covered by the cash flows generated from a company's operations.

	if NetSales == 0 {
		fmt.Printf("got 0 for NetSales with OperatingCashFlow: %f, NetSales: %f\n", OperatingCashFlow, NetSales)
		return 0
	}
	return OperatingCashFlow / NetSales
}

func EBITDAToEVRatio(EBITDA, EnterpriseValue float64) float64 {
	// The EBITDA-to-enterprise value (EV) ratio is a measure of a company's return on investment (ROI). The EV/EBITDA ratio compares a company's enterprise value (EV) to its earnings before interest, taxes, depreciation, and amortization (EBITDA).

	if EnterpriseValue == 0 {
		fmt.Printf("got 0 for EnterpriseValue with EBITDA: %f, EnterpriseValue: %f\n", EBITDA, EnterpriseValue)
		return 0
	}
	return EBITDA / EnterpriseValue
}

func IncomeElasticityOfDemand(PercentChangeInQuantityDemanded, PercentChangeInIncome float64) float64 {
	// Income elasticity of demand is a measure of the relationship between a change in the quantity demanded for a particular good and a change in real income.

	if PercentChangeInIncome == 0 {
		fmt.Printf("got 0 for PercentChangeInIncome with PercentChangeInQuantityDemanded: %f, PercentChangeInIncome: %f\n", PercentChangeInQuantityDemanded, PercentChangeInIncome)
		return 0
	}
	return PercentChangeInQuantityDemanded / PercentChangeInIncome
}
