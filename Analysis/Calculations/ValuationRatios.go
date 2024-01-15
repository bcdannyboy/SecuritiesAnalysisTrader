package Calculations

import "math"

func PriceToBookValueRatio(MarketCapitalization float64, BookValueOfEquity float64) float64 {
	// The price-to-book ratio compares a company's market value to its book value. The market value of a company is its share price multiplied by the number of outstanding shares. The book value is the net assets of a company.

	return MarketCapitalization / BookValueOfEquity
}

func EarningsPerShare(NetIncome, PreferredDividends, SharesOutstanding float64) float64 {
	// Earnings per share (EPS) is calculated as a company's profit divided by the outstanding shares of its common stock. The resulting number serves as an indicator of a company's profitability.

	return (NetIncome - PreferredDividends) / SharesOutstanding
}

func EBITDAPerShare(EBITDA, SharesOutstanding float64) float64 {
	// EBITDA per share is a ratio used to measure a company's return on investment. It is calculated by dividing EBITDA by the number of outstanding shares.

	return EBITDA / SharesOutstanding
}

func BookValuePerShare(ShareHolderEquity float64, SharesOutstanding float64) float64 {
	// Book value per share (BVPS) takes the ratio of a firm's common equity divided by its number of shares outstanding. Book value of equity per share effectively indicates a firm's net asset value (total assets - total liabilities) on a per-share basis.

	return ShareHolderEquity / SharesOutstanding
}

func NetTangibleAssetsPerShare(TangibleNetWorth float64, SharesOutstanding float64) float64 {
	// Net tangible assets per share is a company's total tangible assets minus total liabilities, divided by its number of shares of common stock outstanding.

	return TangibleNetWorth / SharesOutstanding
}

func MarketValueOfEquity(SharePrice, SharesOutstanding float64) float64 {
	// Market value of equity is the total dollar value of a company's equity calculated by multiplying the current stock price by total outstanding shares.

	return SharePrice * SharesOutstanding
}

func MarketValueOfDebt(SharePrice, SharesOutstanding, BookValueOfDebt float64) float64 {
	// Market value of debt is the total dollar value of a company's debt, calculated by multiplying the current market price of a company's debt by its total outstanding debt.

	return (SharePrice * SharesOutstanding) + BookValueOfDebt
}

func MarketToBookRatio(MarketCapitalization, BookValueOfEquity float64) float64 {
	// The market-to-book (M/B) ratio is calculated by dividing the market price per share by book value per share.

	return MarketCapitalization / BookValueOfEquity
}

func IntangiblesRatio(IntangibleAssets, TotalAssets float64) float64 {
	// The intangible ratio is a financial ratio that measures the percentage of intangible assets in comparison to a company's total assets. The intangible ratio is calculated by dividing intangible assets by total assets.

	return IntangibleAssets / TotalAssets
}

func PriceToSalesRatio(SharePrice, NetSales float64) float64 {
	// The price-to-sales (P/S) ratio is a valuation ratio that compares a company’s stock price to its revenues. It is an indicator of the value placed on each dollar of a company’s sales or revenues.

	return SharePrice / NetSales
}

func PriceToBookRatio(SharePrice, BookValueOfEquity float64) float64 {
	// The price-to-book ratio compares a company's market value to its book value. The market value of a company is its share price multiplied by the number of outstanding shares. The book value is the net assets of a company.

	return SharePrice / BookValueOfEquity
}

func PricetoSalesValue(MarketCapitalization, NetSales float64) float64 {
	// The price-to-sales value (P/S) ratio is a valuation ratio that compares a company’s stock price to its revenues. It is an indicator of the value placed on each dollar of a company’s sales or revenues.

	return MarketCapitalization / NetSales
}

func PriceToCashFlowRatio(SharePrice, OperatingCashFlowPerShare float64) float64 {
	// The price-to-cash flow (P/CF) ratio is a stock valuation indicator or multiple that measures the value of a stock’s price relative to its operating cash flow per share.

	return SharePrice / OperatingCashFlowPerShare
}

func PriceToFreeCashFlowRatio(SharePrice, FreeCashFlowPerShare float64) float64 {
	// The price-to-free cash flow (P/FCF) ratio is a valuation method used to compare a company’s current share price to its per-share free cash flow.

	return SharePrice / FreeCashFlowPerShare
}

func PriceToCashFlowValuation(MarketCapitlization, OperatingCashFlow float64) float64 {
	// The price-to-cash flow (P/CF) ratio is a stock valuation indicator or multiple that measures the value of a stock’s price relative to its operating cash flow per share.

	return MarketCapitlization / OperatingCashFlow
}

func PriceToFreeCashFlowValuation(MarketCapitlization, FreeCashFlow float64) float64 {
	// The price-to-free cash flow (P/FCF) ratio is a valuation method used to compare a company’s current share price to its per-share free cash flow.

	return MarketCapitlization / FreeCashFlow
}

func PriceToEarningsGrowth(SharePrice, EarningsPerShare, GrowthRate float64) float64 {
	// The price/earnings to growth (PEG) ratio is a stock's price-to-earnings (P/E) ratio divided by the growth rate of its earnings for a specified time period.

	return SharePrice / (EarningsPerShare * GrowthRate)
}

func PriceToEarningsValuation(MarketCapitalization, NetIncome float64) float64 {
	// The price-to-earnings ratio (P/E ratio) relates a company's share price to its earnings per share. A high P/E ratio could mean that a company's stock is over-valued, or else that investors are expecting high growth rates in the future.

	return MarketCapitalization / NetIncome
}

func EquityMarketValue(SharePrice, SharesOutstanding float64) float64 {
	// Equity market value is the total dollar value of a company's equity calculated by multiplying the current stock price by total outstanding shares.

	return SharePrice * SharesOutstanding
}

func LiabilitiesMarketValue(SharePrice, SharesOutstanding, BookValueOfDebt float64) float64 {
	// Liabilities market value is the total dollar value of a company's debt, calculated by multiplying the current market price of a company's debt by its total outstanding debt.

	return (SharePrice * SharesOutstanding) + BookValueOfDebt
}

func TobinsQ(EquityMarketValue, LiabilitiesMarketValue, EquityBookValue, LiabilitiesBookValue float64) float64 {
	// Tobin's Q ratio is defined as the market value of a company divided by its assets' replacement cost.

	return (EquityMarketValue + LiabilitiesMarketValue) / (EquityBookValue + LiabilitiesBookValue)
}

func CalculateDCF(CashFlows []float64, DiscountRate float64) float64 {
	// The discounted cash flow (DCF) formula is equal to the sum of the cash flow in each period divided by one plus the discount rate (WACC) raised to the power of the period number.
	calculatePresentValue := func(futureCashFlow float64, discountRate float64, year int) float64 {
		return futureCashFlow / math.Pow(1+discountRate, float64(year))
	}

	totalValue := 0.0
	for i, cashFlow := range CashFlows {
		totalValue += calculatePresentValue(cashFlow, DiscountRate, i+1)
	}
	return totalValue
}

func PiotroskiFScore(netIncome, operatingCashFlow, returnOnAssets, qualityOfEarnings, longTermDebt, currentRatio, sharesOutstanding, grossMargin, assetTurnover float64) int {
	// The Piotroski F-Score is a scoring system that uses nine criteria to determine the financial strength of a company. The score is used to determine the best value stocks, nine being the best score and one being the worst score. A score of 8 or 9 is considered strong and a score of 3 or less is considered weak.
	score := 0

	// Profitability
	if netIncome > 0 {
		score++
	}
	if operatingCashFlow > 0 {
		score++
	}
	if returnOnAssets > 0 {
		score++
	}
	if qualityOfEarnings > 0 {
		score++
	}

	// Leverage/Liquidity/Solvency
	if longTermDebt < 0 {
		score++
	}
	if currentRatio > 0 {
		score++
	}
	if sharesOutstanding <= 0 {
		score++
	}

	// Operating Efficiency
	if grossMargin > 0 {
		score++
	}
	if assetTurnover > 0 {
		score++
	}

	return score
}
