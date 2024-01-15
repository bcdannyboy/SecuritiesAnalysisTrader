package Fundamentals

import (
	"fmt"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
)

func AnalyzeBalanceSheet(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) ([]objects.BalanceSheetStatement, []objects.BalanceSheetStatementGrowth, []objects.BalanceSheetStatementAsReported, []*GrowthBalanceSheetStatementAsReported, []*DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported, error) {

	var BS_STMT []objects.BalanceSheetStatement
	var BS_STMT_GROWTH []objects.BalanceSheetStatementGrowth
	var BS_STMT_AS_REPORTED []objects.BalanceSheetStatementAsReported
	var BS_STMT_AS_REPORTED_GROWTH []*GrowthBalanceSheetStatementAsReported
	var BS_DISCREPANCIES []*DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported

	BS_STMT, err := APIClient.CompanyValuation.BalanceSheetStatement(
		objects.RequestBalanceSheetStatement{
			Symbol: Symbol,
			Period: Period,
		})
	if err != nil {
		return BS_STMT, BS_STMT_GROWTH, BS_STMT_AS_REPORTED, BS_STMT_AS_REPORTED_GROWTH, BS_DISCREPANCIES, fmt.Errorf("failed to get balance sheet statement: %s", err.Error())
	}

	BS_STMT_GROWTH, err = APIClient.CompanyValuation.BalanceSheetStatementGrowth(
		objects.RequestBalanceSheetStatementGrowth{
			Symbol: Symbol,
			Period: Period,
		})
	if err != nil {
		return BS_STMT, BS_STMT_GROWTH, BS_STMT_AS_REPORTED, BS_STMT_AS_REPORTED_GROWTH, BS_DISCREPANCIES, fmt.Errorf("failed to get balance sheet statement growth: %s", err.Error())
	}

	BS_STMT_AS_REPORTED, err = APIClient.CompanyValuation.BalanceSheetStatementAsReported(
		objects.RequestBalanceSheetStatementAsReported{
			Symbol: Symbol,
			Period: Period,
		})
	if err != nil {
		return BS_STMT, BS_STMT_GROWTH, BS_STMT_AS_REPORTED, BS_STMT_AS_REPORTED_GROWTH, BS_DISCREPANCIES, fmt.Errorf("failed to get balance sheet statement as reported: %s", err.Error())
	}

	BS_STMT_AS_REPORTED_GROWTH = GetGrowthOfBalanceSheetStatementAsReported(BS_STMT_AS_REPORTED)

	BS_DISCREPANCIES = IdentifyDiscrepanciesBetweenBalanceSheetStatementAndBalanceSheetStatementAsReported(BS_STMT, BS_STMT_AS_REPORTED)

	return BS_STMT, BS_STMT_GROWTH, BS_STMT_AS_REPORTED, BS_STMT_AS_REPORTED_GROWTH, BS_DISCREPANCIES, nil
}

func GetGrowthOfBalanceSheetStatementAsReported(BS_STMT_AS_REPORTED []objects.BalanceSheetStatementAsReported) []*GrowthBalanceSheetStatementAsReported {
	Growth := []*GrowthBalanceSheetStatementAsReported{}

	LastStatement := objects.BalanceSheetStatementAsReported{}
	for i, bs_stmt_as_reported := range BS_STMT_AS_REPORTED {
		if i == 0 {

			NewGrowthObj := &GrowthBalanceSheetStatementAsReported{
				Date:   bs_stmt_as_reported.Date,
				Symbol: bs_stmt_as_reported.Symbol,
				Period: bs_stmt_as_reported.Period,
			}

			// all of the specifically called out ones are defined as interface{} and need to be checked
			Cur_Liabilitiesandstockholdersequity := bs_stmt_as_reported.Liabilitiesandstockholdersequity
			Old_Liabilitiesandstockholdersequity := LastStatement.Liabilitiesandstockholdersequity

			if Cur_Liabilitiesandstockholdersequity != nil && Old_Liabilitiesandstockholdersequity != nil {
				NewGrowthObj.GrowthLiabilitiesandstockholdersequity = Cur_Liabilitiesandstockholdersequity.(float64) - Old_Liabilitiesandstockholdersequity.(float64)
			}

			Cur_Liabilities := bs_stmt_as_reported.Liabilities
			Old_Liabilities := LastStatement.Liabilities

			if Cur_Liabilities != nil && Old_Liabilities != nil {
				NewGrowthObj.GrowthLiabilities = Cur_Liabilities.(float64) - Old_Liabilities.(float64)
			}

			Cur_Liabilitiescurrent := bs_stmt_as_reported.Liabilitiescurrent
			Old_Liabilitiescurrent := LastStatement.Liabilitiescurrent

			if Cur_Liabilitiescurrent != nil && Old_Liabilitiescurrent != nil {
				NewGrowthObj.GrowthLiabilitiescurrent = Cur_Liabilitiescurrent.(float64) - Old_Liabilitiescurrent.(float64)
			}

			Cur_Cashandcashequivalentsatcarryingvalue := bs_stmt_as_reported.Cashandcashequivalentsatcarryingvalue
			Old_Retainedearningsaccumulateddeficit := LastStatement.Retainedearningsaccumulateddeficit

			if Cur_Cashandcashequivalentsatcarryingvalue != nil && Old_Retainedearningsaccumulateddeficit != nil {
				NewGrowthObj.GrowthCashandcashequivalentsatcarryingvalue = Cur_Cashandcashequivalentsatcarryingvalue.(float64) - Old_Retainedearningsaccumulateddeficit.(float64)
			}

			Cur_Liabilitiesnoncurrent := bs_stmt_as_reported.Liabilitiesnoncurrent
			Old_Liabilitiesnoncurrent := LastStatement.Liabilitiesnoncurrent

			if Cur_Liabilitiesnoncurrent != nil && Old_Liabilitiesnoncurrent != nil {
				NewGrowthObj.GrowthLiabilitiesnoncurrent = Cur_Liabilitiesnoncurrent.(float64) - Old_Liabilitiesnoncurrent.(float64)
			}

			Cur_Commonstocksincludingadditionalpaidincapital := bs_stmt_as_reported.Commonstocksincludingadditionalpaidincapital
			Old_Commonstocksincludingadditionalpaidincapital := LastStatement.Commonstocksincludingadditionalpaidincapital

			if Cur_Commonstocksincludingadditionalpaidincapital != nil && Old_Commonstocksincludingadditionalpaidincapital != nil {
				NewGrowthObj.GrowthCommonstocksincludingadditionalpaidincapital = Cur_Commonstocksincludingadditionalpaidincapital.(float64) - Old_Commonstocksincludingadditionalpaidincapital.(float64)
			}

			Cur_Assetscurrent := bs_stmt_as_reported.Assetscurrent
			Old_Assetscurrent := LastStatement.Assetscurrent

			if Cur_Assetscurrent != nil && Old_Assetscurrent != nil {
				NewGrowthObj.GrowthAssetscurrent = Cur_Assetscurrent.(float64) - Old_Assetscurrent.(float64)
			}

			Cur_Stockholdersequity := bs_stmt_as_reported.Stockholdersequity
			Old_Stockholdersequity := LastStatement.Stockholdersequity

			if Cur_Stockholdersequity != nil && Old_Stockholdersequity != nil {
				NewGrowthObj.GrowthStockholdersequity = Cur_Stockholdersequity.(float64) - Old_Stockholdersequity.(float64)
			}

			Cur_Assets := bs_stmt_as_reported.Assets
			Old_Assets := LastStatement.Assets

			if Cur_Assets != nil && Old_Assets != nil {
				NewGrowthObj.GrowthAssets = Cur_Assets.(float64) - Old_Assets.(float64)
			}

			Cur_Assetsnoncurrent := bs_stmt_as_reported.Assetsnoncurrent
			Old_Assetsnoncurrent := LastStatement.Assetsnoncurrent

			if Cur_Assetsnoncurrent != nil && Old_Assetsnoncurrent != nil {
				NewGrowthObj.GrowthAssetsnoncurrent = Cur_Assetsnoncurrent.(float64) - Old_Assetsnoncurrent.(float64)
			}

			Cur_Marketablesecuritiesnoncurrent := bs_stmt_as_reported.Marketablesecuritiesnoncurrent
			Old_Marketablesecuritiesnoncurrent := LastStatement.Marketablesecuritiesnoncurrent

			if Cur_Marketablesecuritiesnoncurrent != nil && Old_Marketablesecuritiesnoncurrent != nil {
				NewGrowthObj.GrowthMarketablesecuritiesnoncurrent = Cur_Marketablesecuritiesnoncurrent.(float64) - Old_Marketablesecuritiesnoncurrent.(float64)
			}

			NewGrowthObj.GrowthCommonstocksharesauthorized = (bs_stmt_as_reported.Commonstocksharesauthorized - LastStatement.Commonstocksharesauthorized)
			NewGrowthObj.GrowthPropertyplantandequipmentnet = (bs_stmt_as_reported.Propertyplantandequipmentnet - LastStatement.Propertyplantandequipmentnet)
			NewGrowthObj.GrowthCommercialpaper = (bs_stmt_as_reported.Commercialpaper - LastStatement.Commercialpaper)
			NewGrowthObj.GrowthLongtermdebtcurrent = (bs_stmt_as_reported.Longtermdebtcurrent - LastStatement.Longtermdebtcurrent)
			NewGrowthObj.GrowthCommonstocksharesoutstanding = (bs_stmt_as_reported.Commonstocksharesoutstanding - LastStatement.Commonstocksharesoutstanding)
			NewGrowthObj.GrowthOtherliabilitiesnoncurrent = (bs_stmt_as_reported.Otherliabilitiesnoncurrent - LastStatement.Otherliabilitiesnoncurrent)
			NewGrowthObj.GrowthMarketablesecuritiescurrent = (bs_stmt_as_reported.Marketablesecuritiescurrent - LastStatement.Marketablesecuritiescurrent)
			NewGrowthObj.GrowthOtherliabilitiescurrent = (bs_stmt_as_reported.Otherliabilitiescurrent - LastStatement.Otherliabilitiescurrent)
			NewGrowthObj.GrowthLongtermdebtnoncurrent = (bs_stmt_as_reported.Longtermdebtnoncurrent - LastStatement.Longtermdebtnoncurrent)
			NewGrowthObj.GrowthContractwithcustomerliabilitycurrent = (bs_stmt_as_reported.Contractwithcustomerliabilitycurrent - LastStatement.Contractwithcustomerliabilitycurrent)
			NewGrowthObj.GrowthNontradereceivablescurrent = (bs_stmt_as_reported.Nontradereceivablescurrent - LastStatement.Nontradereceivablescurrent)
			NewGrowthObj.GrowthCommonstocksharesissued = (bs_stmt_as_reported.Commonstocksharesissued - LastStatement.Commonstocksharesissued)
			NewGrowthObj.GrowthAccountsreceivablenetcurrent = (bs_stmt_as_reported.Accountsreceivablenetcurrent - LastStatement.Accountsreceivablenetcurrent)
			NewGrowthObj.GrowthAccountspayablecurrent = (bs_stmt_as_reported.Accountspayablecurrent - LastStatement.Accountspayablecurrent)
			NewGrowthObj.GrowthOtherassetscurrent = (bs_stmt_as_reported.Otherassetscurrent - LastStatement.Otherassetscurrent)
			NewGrowthObj.GrowthOtherassetsnoncurrent = (bs_stmt_as_reported.Otherassetsnoncurrent - LastStatement.Otherassetsnoncurrent)
			NewGrowthObj.GrowthInventorynet = (bs_stmt_as_reported.Inventorynet - LastStatement.Inventorynet)
			NewGrowthObj.GrowthAccumulatedothercomprehensiveincomelossnetoftax = (bs_stmt_as_reported.Accumulatedothercomprehensiveincomelossnetoftax - LastStatement.Accumulatedothercomprehensiveincomelossnetoftax)
			NewGrowthObj.GrowthOthershorttermborrowings = (bs_stmt_as_reported.Othershorttermborrowings - LastStatement.Othershorttermborrowings)

			// first element has no growth
			Growth = append(Growth, NewGrowthObj)
			LastStatement = bs_stmt_as_reported
			continue
		}
	}

	return Growth
}

func IdentifyDiscrepanciesBetweenBalanceSheetStatementAndBalanceSheetStatementAsReported(BS_STMT []objects.BalanceSheetStatement, BS_STMT_AS_REPORTED []objects.BalanceSheetStatementAsReported) []*DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported {

	calculateDiscrepancyPercentage := func(value1, value2 float64) float64 {
		if value1 == 0 && value2 == 0 {
			return 0
		}
		absoluteDifference := math.Abs(value1 - value2)
		averageValue := (math.Abs(value1) + math.Abs(value2)) / 2
		return absoluteDifference / averageValue
	}

	discrepancies := make([]*DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported, 0)

	for _, bs := range BS_STMT {
		for _, bsar := range BS_STMT_AS_REPORTED {
			if bs.Date == bsar.Date && bs.Symbol == bsar.Symbol && bs.Period == bsar.Period {
				discrepancy := &DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported{
					Date:   bs.Date,
					Symbol: bs.Symbol,
					Period: bs.Period,
				}

				if bsar.Cashandcashequivalentsatcarryingvalue != nil {
					discrepancy.CashAndCashEquivalentsMatch = calculateDiscrepancyPercentage(bs.CashAndCashEquivalents, bsar.Cashandcashequivalentsatcarryingvalue.(float64))
				} else {
					discrepancy.CashAndCashEquivalentsMatch = calculateDiscrepancyPercentage(bs.CashAndCashEquivalents, 0)
				}

				discrepancy.PropertyPlantEquipmentNetMatch = calculateDiscrepancyPercentage(bs.PropertyPlantEquipmentNet, bsar.Propertyplantandequipmentnet)

				if bsar.Commonstocksincludingadditionalpaidincapital != nil {
					discrepancy.CommonStockMatch = calculateDiscrepancyPercentage(bs.CommonStock, bsar.Commonstocksincludingadditionalpaidincapital.(float64))
				} else {
					discrepancy.CommonStockMatch = calculateDiscrepancyPercentage(bs.CommonStock, 0)
				}

				if bsar.Retainedearningsaccumulateddeficit != nil {
					discrepancy.RetainedEarningsMatch = calculateDiscrepancyPercentage(bs.RetainedEarnings, bsar.Retainedearningsaccumulateddeficit.(float64))
				} else {
					discrepancy.RetainedEarningsMatch = calculateDiscrepancyPercentage(bs.RetainedEarnings, 0)
				}

				if bsar.Retainedearningsaccumulateddeficit != nil {
					discrepancy.RetainedEarningsMatch = calculateDiscrepancyPercentage(bs.RetainedEarnings, bsar.Retainedearningsaccumulateddeficit.(float64))
				} else {
					discrepancy.RetainedEarningsMatch = calculateDiscrepancyPercentage(bs.RetainedEarnings, 0)
				}

				if bsar.Assetscurrent != nil {
					discrepancy.TotalCurrentAssetsMatch = calculateDiscrepancyPercentage(bs.TotalCurrentAssets, bsar.Assetscurrent.(float64))
				} else {
					discrepancy.TotalCurrentAssetsMatch = calculateDiscrepancyPercentage(bs.TotalCurrentAssets, 0)
				}

				if bsar.Assets != nil {
					discrepancy.TotalAssetsMatch = calculateDiscrepancyPercentage(bs.TotalAssets, bsar.Assets.(float64))
				} else {
					discrepancy.TotalAssetsMatch = calculateDiscrepancyPercentage(bs.TotalAssets, 0)
				}

				if bsar.Liabilitiescurrent != nil {
					discrepancy.TotalCurrentLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalCurrentLiabilities, bsar.Liabilitiescurrent.(float64))
				} else {
					discrepancy.TotalCurrentLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalCurrentLiabilities, 0)
				}

				if bsar.Liabilities != nil {
					discrepancy.TotalLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalLiabilities, bsar.Liabilities.(float64))
				} else {
					discrepancy.TotalLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalLiabilities, 0)
				}

				if bsar.Stockholdersequity != nil {
					discrepancy.TotalStockholdersEquityMatch = calculateDiscrepancyPercentage(bs.TotalStockholdersEquity, bsar.Stockholdersequity.(float64))
				} else {
					discrepancy.TotalStockholdersEquityMatch = calculateDiscrepancyPercentage(bs.TotalStockholdersEquity, 0)
				}

				if bsar.Liabilitiesandstockholdersequity != nil {
					discrepancy.TotalLiabilitiesAndStockholdersEquityMatch = calculateDiscrepancyPercentage(bs.TotalLiabilitiesAndStockholdersEquity, bsar.Liabilitiesandstockholdersequity.(float64))
				} else {
					discrepancy.TotalLiabilitiesAndStockholdersEquityMatch = calculateDiscrepancyPercentage(bs.TotalLiabilitiesAndStockholdersEquity, 0)
				}
			}
		}
	}

	return discrepancies
}
