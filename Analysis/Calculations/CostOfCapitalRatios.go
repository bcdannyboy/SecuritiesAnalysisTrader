package Calculations

import "fmt"

func WeightedAverageCostOfCapital(MarketValueOfEquity, MarketValueOfDebt, CostOfEquity, CostOfDebt, CorporateTaxRate float64) float64 {
	// Weighted average cost of capital (WACC) represents a company's cost of capital, with each category of capital (debt and equity) proportionately weighted.
	E := MarketValueOfEquity
	D := MarketValueOfDebt
	V := E + D

	if V == 0 {
		fmt.Printf("got a zero value for V in WACC with MarketValueOfEquity: %f, MarketValueOfDebt: %f, CostOfEquity: %f, CostOfDebt: %f, CorporateTaxRate: %f\n", MarketValueOfEquity, MarketValueOfDebt, CostOfEquity, CostOfDebt, CorporateTaxRate)
		return (0 * CostOfEquity) + (0 * CostOfDebt * (1 - CorporateTaxRate))
	}

	return ((E / V) * CostOfEquity) + ((D / V) * CostOfDebt * (1 - CorporateTaxRate))
}

func PriceElasticityOfDemand(PercentageChangeInQuantityDemanded, PercentageChangeInCostOfGoodsSold float64) float64 {
	// Price elasticity of demand is a measure of the relationship between a change in the quantity demanded of a particular good and a change in its price.
	if PercentageChangeInCostOfGoodsSold == 0 {
		fmt.Printf("got a zero value for PercentageChangeInCostOfGoodsSold in PriceElasticityOfDemand with PercentageChangeInQuantityDemanded: %f, PercentageChangeInCostOfGoodsSold: %f\n", PercentageChangeInQuantityDemanded, PercentageChangeInCostOfGoodsSold)
		return 0
	}
	return PercentageChangeInQuantityDemanded / PercentageChangeInCostOfGoodsSold
}

func MarginalCostOfCapital(PercentChangeInTotalExpenses, PercentChangeInQuantityOfUnitsProduced float64) float64 {
	// Marginal cost of capital is the weighted average cost of the last dollar of capital raised by a company.

	if PercentChangeInQuantityOfUnitsProduced == 0 {
		fmt.Printf("got a zero value for PercentChangeInQuantityOfUnitsProduced in MarginalCostOfCapital with PercentChangeInTotalExpenses: %f, PercentChangeInQuantityOfUnitsProduced: %f\n", PercentChangeInTotalExpenses, PercentChangeInQuantityOfUnitsProduced)
		return 0
	}
	return PercentChangeInTotalExpenses / PercentChangeInQuantityOfUnitsProduced
}

func CostOfPreferredStock(PrefferedStockDividendPerShare, MarketValueOfPreferredStock float64) float64 {
	// The cost of preferred stock is the rate of return required by holders of a company's preferred stock.

	if MarketValueOfPreferredStock == 0 {
		fmt.Printf("got a zero value for MarketValueOfPreferredStock in CostOfPreferredStock with PrefferedStockDividendPerShare: %f, MarketValueOfPreferredStock: %f\n", PrefferedStockDividendPerShare, MarketValueOfPreferredStock)
		return 0
	}
	return PrefferedStockDividendPerShare / MarketValueOfPreferredStock
}

func CostOfRetainedEarnings(MarketValueOfStock, UpcomingDividendYield, ExpectedGrowthRate float64) float64 {
	// The cost of retained earnings is the opportunity cost associated with the use of retained earnings as a source of funding.

	if MarketValueOfStock == 0 {
		fmt.Printf("got a zero value for MarketValueOfStock in CostOfRetainedEarnings with MarketValueOfStock: %f, UpcomingDividendYield: %f, ExpectedGrowthRate: %f\n", MarketValueOfStock, UpcomingDividendYield, ExpectedGrowthRate)
		return 0 + ExpectedGrowthRate
	}
	return (UpcomingDividendYield / MarketValueOfStock) + ExpectedGrowthRate
}

func AdjustedPresentValue(UnleveredFirmValue, NetEffectOfDebt float64) float64 {
	// Adjusted present value (APV) is the net present value of a project if financed solely by ownership equity plus the present value of all the benefits of financing.

	return UnleveredFirmValue + NetEffectOfDebt
}
