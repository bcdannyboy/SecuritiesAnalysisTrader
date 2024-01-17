package Calculations

func WeightedAverageCostOfCapital(MarketValueOfEquity, MarketValueOfDebt, CostOfEquity, CostOfDebt, CorporateTaxRate float64) float64 {
	// Weighted average cost of capital (WACC) represents a company's cost of capital, with each category of capital (debt and equity) proportionately weighted.
	E := MarketValueOfEquity
	D := MarketValueOfDebt
	V := E + D

	return ((E / V) * CostOfEquity) + ((D / V) * CostOfDebt * (1 - CorporateTaxRate))
}

func PriceElasticityOfDemand(PercentageChangeInQuantityDemanded, PercentageChangeInCostOfGoodsSold float64) float64 {
	// Price elasticity of demand is a measure of the relationship between a change in the quantity demanded of a particular good and a change in its price.

	return PercentageChangeInQuantityDemanded / PercentageChangeInCostOfGoodsSold
}

func MarginalCostOfCapital(PercentChangeInTotalExpenses, PercentChangeInQuantityOfUnitsProduced float64) float64 {
	// Marginal cost of capital is the weighted average cost of the last dollar of capital raised by a company.

	return PercentChangeInTotalExpenses / PercentChangeInQuantityOfUnitsProduced
}

func CostOfPreferredStock(PrefferedStockDividendPerShare, MarketValueOfPreferredStock float64) float64 {
	// The cost of preferred stock is the rate of return required by holders of a company's preferred stock.

	return PrefferedStockDividendPerShare / MarketValueOfPreferredStock
}

func CostOfRetainedEarnings(MarketValueOfStock, UpcomingDividendYield, ExpectedGrowthRate float64) float64 {
	// The cost of retained earnings is the opportunity cost associated with the use of retained earnings as a source of funding.

	return (UpcomingDividendYield / MarketValueOfStock) + ExpectedGrowthRate
}

func AdjustedPresentValue(UnleveredFirmValue, NetEffectOfDebt float64) float64 {
	// Adjusted present value (APV) is the net present value of a project if financed solely by ownership equity plus the present value of all the benefits of financing.

	return UnleveredFirmValue + NetEffectOfDebt
}
