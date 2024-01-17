package Analysis

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Calculations"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/BalanceSheet"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/CashFlowStatement"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/FinancialRatios"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/FullFinancialStatement"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals/IncomeStatement"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func PullCompanyFundamentals(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) (*CompanyFundamentals, error) {
	// Pull fundamentals / ratios from FMP / calculations for a given ticker
	FundamentalsObj := &CompanyFundamentals{
		Symbol: Symbol,
	}

	BS_STMT, BS_STMT_GROWTH, BS_STMT_AS_REPORTED, BS_STMT_AS_REPORTED_GROWTH, BS_DISCREPANCIES, err := BalanceSheet.AnalyzeBalanceSheet(APIClient, Symbol, Period)
	if err != nil {
		if len(BS_STMT) == 0 {
			fmt.Printf("Failed to get balance sheet statement for %s: %s\n", Symbol, err.Error())
		} else if len(BS_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get balance sheet statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.BalanceSheetStatementGrowth = BS_STMT_GROWTH
		} else if len(BS_STMT_AS_REPORTED) == 0 {
			fmt.Printf("Failed to get balance sheet statement as reported for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.BalanceSheetStatementGrowth = BS_STMT_GROWTH
			FundamentalsObj.BalanceSheetStatementAsReported = BS_STMT_AS_REPORTED
		}
	} else {
		FundamentalsObj.BalanceSheetStatements = BS_STMT
		FundamentalsObj.BalanceSheetStatementGrowth = BS_STMT_GROWTH
		FundamentalsObj.BalanceSheetStatementAsReported = BS_STMT_AS_REPORTED
	}

	FundamentalsObj.GrowthBalanceSheetStatementAsReported = BS_STMT_AS_REPORTED_GROWTH
	FundamentalsObj.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported = BS_DISCREPANCIES

	CF_STMT, CF_STMT_GROWTH, CF_STMT_AS_REPORTED, CF_STMT_AS_REPORTED_GROWTH, CF_DISCREPANCIES, err := CashFlowStatement.AnalyzeCashFlow(APIClient, Symbol, Period)
	if err != nil {
		if len(CF_STMT) == 0 {
			fmt.Printf("Failed to get cash flow statement for %s: %s\n", Symbol, err.Error())
		} else if len(CF_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get cash flow statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.CashFlowStatementGrowth = CF_STMT_GROWTH
		} else if len(CF_STMT_AS_REPORTED) == 0 {
			fmt.Printf("Failed to get cash flow statement as reported for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.CashFlowStatementGrowth = CF_STMT_GROWTH
			FundamentalsObj.CashFlowStatementAsReported = CF_STMT_AS_REPORTED
		}
	} else {
		FundamentalsObj.CashFlowStatement = CF_STMT
		FundamentalsObj.CashFlowStatementGrowth = CF_STMT_GROWTH
		FundamentalsObj.CashFlowStatementAsReported = CF_STMT_AS_REPORTED
	}

	FundamentalsObj.CashFlowStatementAsReportedGrowth = CF_STMT_AS_REPORTED_GROWTH
	FundamentalsObj.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported = CF_DISCREPANCIES

	I_STMT, I_STMT_GROWTH, I_STMT_AS_REPORTED, I_STMT_AS_REPORTED_GROWTH, I_DISCREPANCIES, err := IncomeStatement.AnalyzeIncomeStatement(APIClient, Symbol, Period)
	if err != nil {
		if len(I_STMT) == 0 {
			fmt.Printf("Failed to get income statement for %s: %s\n", Symbol, err.Error())
		} else if len(I_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get income statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.IncomeStatement = I_STMT
		} else if len(I_STMT_AS_REPORTED) == 0 {
			fmt.Printf("Failed to get income statement as reported for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.IncomeStatementGrowth = I_STMT_GROWTH
			FundamentalsObj.IncomeStatementAsReported = I_STMT_AS_REPORTED
		}
	} else {
		FundamentalsObj.IncomeStatement = I_STMT
		FundamentalsObj.IncomeStatementGrowth = I_STMT_GROWTH
		FundamentalsObj.IncomeStatementAsReported = I_STMT_AS_REPORTED
	}

	FundamentalsObj.GrowthIncomeStatementAsReported = I_STMT_AS_REPORTED_GROWTH
	FundamentalsObj.DiscrepancyIncomeStatementAndIncomeStatementAsReported = I_DISCREPANCIES

	F_STMT, F_STMT_GROWTH, err := FullFinancialStatement.AnalyzeFinancialStatement(APIClient, Symbol, Period)
	if err != nil {
		if len(F_STMT) == 0 {
			fmt.Printf("Failed to get full financial statement for %s: %s\n", Symbol, err.Error())
		} else if len(F_STMT_GROWTH) == 0 {
			fmt.Printf("Failed to get full financial statement growth for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.FullFinancialStatement = F_STMT
		}
	} else {
		FundamentalsObj.FullFinancialStatement = F_STMT
		FundamentalsObj.FullFinancialStatementGrowth = F_STMT_GROWTH
	}

	FIN_RATIOS, FIN_RATIOS_TTM, FR_GROWTH, FR_TTM_GROWTH, err := FinancialRatios.AnalyzeFinancialRatios(APIClient, Symbol, Period)
	if err != nil {
		if len(FIN_RATIOS) == 0 {
			print("Failed to get financial ratios for %s: %s\n", Symbol, err.Error())
		} else if len(FIN_RATIOS_TTM) == 0 {
			print("Failed to get financial ratios TTM for %s: %s\n", Symbol, err.Error())
			FundamentalsObj.FinancialRatios = FIN_RATIOS
		}
	} else {
		FundamentalsObj.FinancialRatios = FIN_RATIOS
		FundamentalsObj.FinancialRatiosTTM = FIN_RATIOS_TTM
	}

	FundamentalsObj.FinancialRatiosGrowth = FR_GROWTH
	FundamentalsObj.FinancialRatiosTTMGrowth = FR_TTM_GROWTH

	return FundamentalsObj, nil
}

func PullCompanyDCFs(APIClient *fmpcloud.APIClient, Symbol string) ([]objects.DiscountedCashFlow, map[string][]interface{}, error) {
	CompanyDCFs, err := APIClient.CompanyValuation.DiscountedCashFlow(Symbol)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get company DCFs for %s: %s", Symbol, err.Error())
	}
	CompanyDCFMeanSTD, err := Calculations.CalculateMeanSTDObjs([]interface{}{CompanyDCFs})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to calculate mean and std for company DCFs for %s: %s", Symbol, err.Error())
	}
	return CompanyDCFs, CompanyDCFMeanSTD, nil
}

func PullCompanyRatings(APIClient *fmpcloud.APIClient, Symbol string) ([]map[string]*float64, map[string][]float64, map[string][]interface{}, map[string][]interface{}, error) {
	// TODO: should rating dates come with the results too?
	Ratings, err := APIClient.CompanyValuation.Rating(Symbol)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to get company ratings for %s: %s", Symbol, err.Error())
	}
	RatingsMaps := []map[string]*float64{}

	for _, Rating := range Ratings {
		RatingsMap := map[string]*float64{
			"RatingScore":           utils.InterfaceToFloat64Ptr(Rating.RatingScore),
			"RatingDetailsDCFScore": utils.InterfaceToFloat64Ptr(Rating.RatingDetailsDCFScore),
			"RatingDetailsROEScore": utils.InterfaceToFloat64Ptr(Rating.RatingDetailsROEScore),
			"RatingDetailsROAScore": utils.InterfaceToFloat64Ptr(Rating.RatingDetailsROAScore),
			"RatingDetailsDEScore":  utils.InterfaceToFloat64Ptr(Rating.RatingDetailsDEScore),
			"RatingDetailsPEScore":  utils.InterfaceToFloat64Ptr(Rating.RatingDetailsPEScore),
			"RatingDetailsPBScore":  utils.InterfaceToFloat64Ptr(Rating.RatingDetailsPBScore),
		}

		RatingsMaps = append(RatingsMaps, RatingsMap)
	}

	RatingsMapsGrowth := Calculations.CalculateGrowthF64P(RatingsMaps)

	RatingsMapsMeanSTD, err := Calculations.CalculateMeanSTDObjs([]interface{}{RatingsMaps})
	if err != nil {
		return RatingsMaps, RatingsMapsGrowth, nil, nil, fmt.Errorf("failed to calculate mean and std for company ratings for %s: %s", Symbol, err.Error())
	}
	RatingsMapsGrowthMeanSTD, err := Calculations.CalculateMeanSTDObjs([]interface{}{RatingsMapsGrowth})
	if err != nil {
		return RatingsMaps, RatingsMapsGrowth, RatingsMapsMeanSTD, nil, fmt.Errorf("failed to calculate mean and std for company ratings growth for %s: %s", Symbol, err.Error())
	}

	return RatingsMaps, RatingsMapsGrowth, RatingsMapsMeanSTD, RatingsMapsGrowthMeanSTD, nil
}

func PullCompanyOutlook(APIClient *fmpcloud.APIClient, Symbol string) (*CompanyOutlook, error) {
	ValuationInfo, err := APIClient.CompanyValuation.CompanyOutlook(Symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to get company outlook for %s: %s", Symbol, err.Error())
	}

	Beta := ValuationInfo.Profile.Beta
	AvgVolume := ValuationInfo.Profile.VolAvg
	MarketCap := ValuationInfo.Profile.MktCap
	LastDividend := ValuationInfo.Profile.LastDiv
	Changes := ValuationInfo.Profile.Changes
	DcfDiff := ValuationInfo.Profile.DcfDiff
	Dcf := ValuationInfo.Profile.Dcf
	StockPrice := ValuationInfo.Profile.Price
	TotalExecutivePay := float64(0)
	var InsiderTradesMeanSTD map[string][]interface{} = nil
	InsiderTradesGrowth := map[string][]float64{}
	var InsiderTradesGrowthMeanSTD map[string][]interface{} = nil
	InsiderTrades := []map[string]*float64{}
	Splits := []float64{}
	var SplitMeanSTD map[string][]interface{} = nil
	DividendHist := []map[string]*float64{}
	DividendHistGrowth := map[string][]float64{}
	var DividendHistMeanSTD map[string][]interface{} = nil
	var DividendHistGrowthMeanSTD map[string][]interface{} = nil
	DividendYieldTTM := ValuationInfo.Metrics.DividendYielTTM
	CurrentVolume := ValuationInfo.Metrics.Volume
	YearHigh := ValuationInfo.Metrics.YearHigh
	YearLow := ValuationInfo.Metrics.YearLow
	Ratios := []map[string]*float64{}
	RatiosGrowth := map[string][]float64{}
	var RatiosMeanSTD map[string][]interface{} = nil
	var RatiosGrowthMeanSTD map[string][]interface{} = nil

	for _, InsiderTrade := range ValuationInfo.InsideTrades {
		TradeMap := map[string]*float64{
			"SecuritiesOwned":      utils.InterfaceToFloat64Ptr(InsiderTrade.SecuritiesOwned),
			"SecuritiesTransacted": utils.InterfaceToFloat64Ptr(InsiderTrade.SecuritiesTransacted),
			"TotalPrice":           utils.InterfaceToFloat64Ptr(InsiderTrade.SecuritiesTransacted * InsiderTrade.Price),
		}

		InsiderTrades = append(InsiderTrades, TradeMap)
	}

	InsiderTradesGrowth = Calculations.CalculateGrowthF64P(InsiderTrades)

	if len(InsiderTrades) > 0 {
		InsiderTradesMeanSTD, err = Calculations.CalculateMeanSTDObjs([]interface{}{InsiderTrades})
		if err != nil {
			fmt.Printf("Failed to calculate mean and std for insider trades for %s: %s\n", Symbol, err.Error())
		}
	}

	if len(InsiderTradesGrowth) > 0 {
		InsiderTradesGrowthMeanSTD, err = Calculations.CalculateMeanSTDObjs([]interface{}{InsiderTradesGrowth})
		if err != nil {
			fmt.Printf("Failed to calculate mean and std for insider trades growth for %s: %s\n", Symbol, err.Error())
		}
	}

	for _, Executive := range ValuationInfo.KeyExecutives {
		TotalExecutivePay += Executive.Pay
	}

	for _, Split := range ValuationInfo.SplitHistory {
		Splits = append(Splits, Split.Numerator/Split.Denominator)
	}

	if len(Splits) > 0 {
		SplitMeanSTD, err = Calculations.CalculateMeanSTDObjs([]interface{}{Splits})
		if err != nil {
			fmt.Printf("Failed to calculate mean and std for splits for %s: %s\n", Symbol, err.Error())
		}
	}

	for _, Dividend := range ValuationInfo.StockDividend {
		DividendMap := map[string]*float64{
			"Dividend":    utils.InterfaceToFloat64Ptr(Dividend.Dividend),
			"AdjDividend": utils.InterfaceToFloat64Ptr(Dividend.AdjDividend),
		}

		DividendHist = append(DividendHist, DividendMap)
	}

	DividendHistGrowth = Calculations.CalculateGrowthF64P(DividendHist)

	if len(DividendHist) > 0 {
		DividendHistMeanSTD, err = Calculations.CalculateMeanSTDObjs([]interface{}{DividendHist})
		if err != nil {
			fmt.Printf("Failed to calculate mean and std for dividend history for %s: %s\n", Symbol, err.Error())
		}
	}

	if len(DividendHistGrowth) > 0 {
		DividendHistGrowthMeanSTD, err = Calculations.CalculateMeanSTDObjs([]interface{}{DividendHistGrowth})
		if err != nil {
			fmt.Printf("Failed to calculate mean and std for dividend history growth for %s: %s\n", Symbol, err.Error())
		}
	}

	for _, Ratio := range ValuationInfo.Ratios {
		RatioMap := map[string]*float64{
			"DividendYielTTM":                       utils.InterfaceToFloat64Ptr(Ratio.DividendYielTTM),
			"DividendYielPercentageTTM":             utils.InterfaceToFloat64Ptr(Ratio.DividendYielPercentageTTM),
			"PeRatioTTM":                            utils.InterfaceToFloat64Ptr(Ratio.PeRatioTTM),
			"PegRatioTTM":                           utils.InterfaceToFloat64Ptr(Ratio.PegRatioTTM),
			"PayoutRatioTTM":                        utils.InterfaceToFloat64Ptr(Ratio.PayoutRatioTTM),
			"CurrentRatioTTM":                       utils.InterfaceToFloat64Ptr(Ratio.CurrentRatioTTM),
			"QuickRatioTTM":                         utils.InterfaceToFloat64Ptr(Ratio.QuickRatioTTM),
			"CashRatioTTM":                          utils.InterfaceToFloat64Ptr(Ratio.CashRatioTTM),
			"DaysOfSalesOutstandingTTM":             utils.InterfaceToFloat64Ptr(Ratio.DaysOfSalesOutstandingTTM),
			"DaysOfInventoryOutstandingTTM":         utils.InterfaceToFloat64Ptr(Ratio.DaysOfInventoryOutstandingTTM),
			"OperatingCycleTTM":                     utils.InterfaceToFloat64Ptr(Ratio.OperatingCycleTTM),
			"DaysOfPayablesOutstandingTTM":          utils.InterfaceToFloat64Ptr(Ratio.DaysOfPayablesOutstandingTTM),
			"CashConversionCycleTTM":                utils.InterfaceToFloat64Ptr(Ratio.CashConversionCycleTTM),
			"GrossProfitMarginTTM":                  utils.InterfaceToFloat64Ptr(Ratio.GrossProfitMarginTTM),
			"OperatingProfitMarginTTM":              utils.InterfaceToFloat64Ptr(Ratio.OperatingProfitMarginTTM),
			"PretaxProfitMarginTTM":                 utils.InterfaceToFloat64Ptr(Ratio.PretaxProfitMarginTTM),
			"NetProfitMarginTTM":                    utils.InterfaceToFloat64Ptr(Ratio.NetProfitMarginTTM),
			"EffectiveTaxRateTTM":                   utils.InterfaceToFloat64Ptr(Ratio.EffectiveTaxRateTTM),
			"ReturnOnAssetsTTM":                     utils.InterfaceToFloat64Ptr(Ratio.ReturnOnAssetsTTM),
			"ReturnOnEquityTTM":                     utils.InterfaceToFloat64Ptr(Ratio.ReturnOnEquityTTM),
			"ReturnOnCapitalEmployedTTM":            utils.InterfaceToFloat64Ptr(Ratio.ReturnOnCapitalEmployedTTM),
			"NetIncomePerEBTTTM":                    utils.InterfaceToFloat64Ptr(Ratio.NetIncomePerEBTTTM),
			"EbtPerEbitTTM":                         utils.InterfaceToFloat64Ptr(Ratio.EbtPerEbitTTM),
			"EbitPerRevenueTTM":                     utils.InterfaceToFloat64Ptr(Ratio.EbitPerRevenueTTM),
			"DebtRatioTTM":                          utils.InterfaceToFloat64Ptr(Ratio.DebtRatioTTM),
			"DebtEquityRatioTTM":                    utils.InterfaceToFloat64Ptr(Ratio.DebtEquityRatioTTM),
			"LongTermDebtToCapitalizationTTM":       utils.InterfaceToFloat64Ptr(Ratio.LongTermDebtToCapitalizationTTM),
			"TotalDebtToCapitalizationTTM":          utils.InterfaceToFloat64Ptr(Ratio.TotalDebtToCapitalizationTTM),
			"InterestCoverageTTM":                   utils.InterfaceToFloat64Ptr(Ratio.InterestCoverageTTM),
			"CashFlowToDebtRatioTTM":                utils.InterfaceToFloat64Ptr(Ratio.CashFlowToDebtRatioTTM),
			"CompanyEquityMultiplierTTM":            utils.InterfaceToFloat64Ptr(Ratio.CompanyEquityMultiplierTTM),
			"ReceivablesTurnoverTTM":                utils.InterfaceToFloat64Ptr(Ratio.ReceivablesTurnoverTTM),
			"PayablesTurnoverTTM":                   utils.InterfaceToFloat64Ptr(Ratio.PayablesTurnoverTTM),
			"InventoryTurnoverTTM":                  utils.InterfaceToFloat64Ptr(Ratio.InventoryTurnoverTTM),
			"FixedAssetTurnoverTTM":                 utils.InterfaceToFloat64Ptr(Ratio.FixedAssetTurnoverTTM),
			"AssetTurnoverTTM":                      utils.InterfaceToFloat64Ptr(Ratio.AssetTurnoverTTM),
			"OperatingCashFlowPerShareTTM":          utils.InterfaceToFloat64Ptr(Ratio.OperatingCashFlowPerShareTTM),
			"FreeCashFlowPerShareTTM":               utils.InterfaceToFloat64Ptr(Ratio.FreeCashFlowPerShareTTM),
			"CashPerShareTTM":                       utils.InterfaceToFloat64Ptr(Ratio.CashPerShareTTM),
			"OperatingCashFlowSalesRatioTTM":        utils.InterfaceToFloat64Ptr(Ratio.OperatingCashFlowSalesRatioTTM),
			"FreeCashFlowOperatingCashFlowRatioTTM": utils.InterfaceToFloat64Ptr(Ratio.FreeCashFlowOperatingCashFlowRatioTTM),
			"CashFlowCoverageRatiosTTM":             utils.InterfaceToFloat64Ptr(Ratio.CashFlowCoverageRatiosTTM),
			"ShortTermCoverageRatiosTTM":            utils.InterfaceToFloat64Ptr(Ratio.ShortTermCoverageRatiosTTM),
			"CapitalExpenditureCoverageRatioTTM":    utils.InterfaceToFloat64Ptr(Ratio.CapitalExpenditureCoverageRatioTTM),
			"DividendPaidAndCapexCoverageRatioTTM":  utils.InterfaceToFloat64Ptr(Ratio.DividendPaidAndCapexCoverageRatioTTM),
			"PriceBookValueRatioTTM":                utils.InterfaceToFloat64Ptr(Ratio.PriceBookValueRatioTTM),
			"PriceToSalesRatioTTM":                  utils.InterfaceToFloat64Ptr(Ratio.PriceToSalesRatioTTM),
			"PriceEarningsRatioTTM":                 utils.InterfaceToFloat64Ptr(Ratio.PriceEarningsRatioTTM),
			"PriceToFreeCashFlowsRatioTTM":          utils.InterfaceToFloat64Ptr(Ratio.PriceToFreeCashFlowsRatioTTM),
			"PriceToOperatingCashFlowsRatioTTM":     utils.InterfaceToFloat64Ptr(Ratio.PriceToOperatingCashFlowsRatioTTM),
			"PriceCashFlowRatioTTM":                 utils.InterfaceToFloat64Ptr(Ratio.PriceCashFlowRatioTTM),
			"PriceEarningsToGrowthRatioTTM":         utils.InterfaceToFloat64Ptr(Ratio.PriceEarningsToGrowthRatioTTM),
			"PriceSalesRatioTTM":                    utils.InterfaceToFloat64Ptr(Ratio.PriceSalesRatioTTM),
			"DividendYieldTTM":                      utils.InterfaceToFloat64Ptr(Ratio.DividendYieldTTM),
			"EnterpriseValueMultipleTTM":            utils.InterfaceToFloat64Ptr(Ratio.EnterpriseValueMultipleTTM),
			"PriceFairValueTTM":                     utils.InterfaceToFloat64Ptr(Ratio.PriceFairValueTTM),
		}

		Ratios = append(Ratios, RatioMap)
	}

	RatiosGrowth = Calculations.CalculateGrowthF64P(Ratios)

	if len(Ratios) > 0 {
		RatiosMeanSTD, err = Calculations.CalculateMeanSTDObjs([]interface{}{Ratios})
		if err != nil {
			fmt.Printf("Failed to calculate mean and std for ratios for %s: %s\n", Symbol, err.Error())
		}
	}

	if len(RatiosGrowth) > 0 {
		RatiosGrowthMeanSTD, err = Calculations.CalculateMeanSTDObjs([]interface{}{RatiosGrowth})
		if err != nil {
			fmt.Printf("Failed to calculate mean and std for ratios growth for %s: %s\n", Symbol, err.Error())
		}
	}

	return &CompanyOutlook{
		Beta:                       Beta,
		AvgVolume:                  AvgVolume,
		MarketCap:                  MarketCap,
		LastDividend:               LastDividend,
		Changes:                    Changes,
		FMPDcfDiff:                 DcfDiff,
		FMPDcf:                     Dcf,
		StockPrice:                 StockPrice,
		TotalExecutivePay:          TotalExecutivePay,
		InsiderTrades:              InsiderTrades,
		InsiderTradesGrowth:        InsiderTradesGrowth,
		InsiderTradesMeanSTD:       InsiderTradesMeanSTD,
		InsiderTradesGrowthMeanSTD: InsiderTradesGrowthMeanSTD,
		Splits:                     Splits,
		SplitMeanSTD:               SplitMeanSTD,
		DividendHist:               DividendHist,
		DividendHistGrowth:         DividendHistGrowth,
		DividendHistMeanSTD:        DividendHistMeanSTD,
		DividendHistGrowthMeanSTD:  DividendHistGrowthMeanSTD,
		DividendYieldTTM:           DividendYieldTTM,
		CurrentVolume:              CurrentVolume,
		YearHigh:                   YearHigh,
		YearLow:                    YearLow,
		Ratios:                     Ratios,
		RatiosGrowth:               RatiosGrowth,
		RatiosMeanSTD:              RatiosMeanSTD,
		RatiosGrowthMeanSTD:        RatiosGrowthMeanSTD,
	}, nil
}

func PullEmployeeCount(APIClient *fmpcloud.APIClient, Symbol string) (*float64, error) {
	EmployeeCountOBj, err := APIClient.CompanyValuation.EmployeeCount(Symbol)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee count for %s: %s", Symbol, err.Error())
	}

	return utils.InterfaceToFloat64Ptr(EmployeeCountOBj.EmployeeCount), nil
}
