package Calculations

import "fmt"

func DebtToCapitalRatio(ShortTermDebt float64, LongTermDebt float64, TotalShareholderEquity float64) float64 {
	// The debt-to-capital ratio is a measurement of a company's financial leverage. The debt-to-capital ratio is calculated by taking the company's debt, including both short- and long-term liabilities and dividing it by the total capital.
	if (ShortTermDebt + LongTermDebt + TotalShareholderEquity) == 0 {
		fmt.Printf("(ShortTermDebt + LongTermDebt + TotalShareholderEquity) == 0, with ShortTermDebt: %f LongTermDebt: %f TotalShareholderEquity: %f\n", ShortTermDebt, LongTermDebt, TotalShareholderEquity)
		return 0
	}
	return (ShortTermDebt + LongTermDebt) / (ShortTermDebt + LongTermDebt + TotalShareholderEquity)
}

func LeverageRatio(TotalDebt float64, EBITDA float64) float64 {
	// The leverage ratio is a measure of the financial leverage of a company. It is calculated by dividing total debt (the sum of current liabilities and long-term liabilities) by total assets (the sum of current assets, fixed assets, and other assets such as 'goodwill').
	if EBITDA == 0 {
		fmt.Printf("EBITDA == 0, with TotalDebt: %f EBITDA: %f\n", TotalDebt, EBITDA)
		return 0
	}
	return TotalDebt / EBITDA
}

func CapitalizationRatio(TotalDebt float64, ShareHolderEquity float64) float64 {
	// The capitalization ratio, often called the Cap ratio, is a financial metric that measures a company's solvency by calculating the total debt component of the company's capital structure of debt and equity.
	if ShareHolderEquity == 0 {
		fmt.Printf("ShareHolderEquity == 0, with TotalDebt: %f ShareHolderEquity: %f\n", TotalDebt, ShareHolderEquity)
		return 0
	}
	return TotalDebt / ShareHolderEquity
}

func NetGearingRatio(LongTermDebt, ShortTermDebt, AdditionalLiabilities, ShareHoldersEquity float64) float64 {
	// The net gearing ratio is a liquidity ratio that measures a company's ability to repay its debts if they were all due today. It is calculated by dividing a company's net debt by its net assets.
	if ShareHoldersEquity == 0 {
		fmt.Printf("got ShareHoldersEquity == 0, with LongTermDebt: %f ShortTermDebt: %f AdditionalLiabilities: %f ShareHoldersEquity: %f\n", LongTermDebt, ShortTermDebt, AdditionalLiabilities, ShareHoldersEquity)
		return 0
	}
	return (LongTermDebt + ShortTermDebt + AdditionalLiabilities) / ShareHoldersEquity
}

func TotalDebtToEBITDA(TotalDebt float64, EBITDA float64) float64 {
	// The leverage ratio is a measure of the financial leverage of a company. It is calculated by dividing total debt (the sum of current liabilities and long-term liabilities) by total assets (the sum of current assets, fixed assets, and other assets such as 'goodwill').
	if EBITDA == 0 {
		fmt.Printf("got EBITDA == 0, with TotalDebt: %f EBITDA: %f\n", TotalDebt, EBITDA)
		return 0
	}
	return TotalDebt / EBITDA
}

func DebtToEquityRatio(TotalLiabilities float64, TotalShareholderEquity float64) float64 {
	// The debt-to-equity (D/E) ratio is calculated by dividing a company’s total liabilities by its shareholder equity. These numbers are available on the balance sheet of a company’s financial statements. The ratio is used to evaluate a company’s financial leverage.
	if TotalShareholderEquity == 0 {
		fmt.Printf("got TotalShareholderEquity == 0, with TotalLiabilities: %f TotalShareholderEquity: %f\n", TotalLiabilities, TotalShareholderEquity)
		return 0
	}
	return TotalLiabilities / TotalShareholderEquity
}

func EquityMultiplierRatio(TotalAssets float64, ShareHolderEquity float64) float64 {
	// The equity multiplier is a financial leverage ratio that measures the portion of company’s assets that are financed by stockholder's equity. It is calculated by dividing total assets by total equity.
	if ShareHolderEquity == 0 {
		fmt.Printf("got ShareHolderEquity == 0, with TotalAssets: %f ShareHolderEquity: %f\n", TotalAssets, ShareHolderEquity)
		return 0
	}
	return TotalAssets / ShareHolderEquity
}

func DuPontAnalysis(NetProfitMargin, AssetTurnoverRatio, EquityMultiplierRatio float64) float64 {
	// The DuPont analysis is an expression which breaks ROE (return on equity) into three parts: profit margin, total asset turnover, and financial leverage. It is also known as the DuPont identity.

	return NetProfitMargin * AssetTurnoverRatio * EquityMultiplierRatio
}

func DegreeOfFinancialLeverage(PercentageChangeInEPS, PercentageChangeInEBIT float64) float64 {
	// The degree of financial leverage (DFL) is a ratio that measures the sensitivity of a company’s earnings per share (EPS) to fluctuations in its operating income, as a result of changes in its capital structure.
	if PercentageChangeInEBIT == 0 {
		fmt.Printf("got PercentageChangeInEBIT == 0, with PercentageChangeInEPS: %f PercentageChangeInEBIT: %f\n", PercentageChangeInEPS, PercentageChangeInEBIT)
		return 0
	}
	return PercentageChangeInEPS / PercentageChangeInEBIT
}
