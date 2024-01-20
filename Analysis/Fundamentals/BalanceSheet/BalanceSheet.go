package BalanceSheet

import (
	"fmt"
	fundamentals "github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
	"reflect"
)

func AnalyzeBalanceSheet(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) ([]objects.BalanceSheetStatement, []objects.BalanceSheetStatementGrowth, []objects.BalanceSheetStatementAsReported, []fundamentals.GrowthBalanceSheetStatementAsReported, []fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported, error) {

	var BS_STMT []objects.BalanceSheetStatement
	var BS_STMT_GROWTH []objects.BalanceSheetStatementGrowth
	var BS_STMT_AS_REPORTED []objects.BalanceSheetStatementAsReported
	var BS_STMT_AS_REPORTED_GROWTH []fundamentals.GrowthBalanceSheetStatementAsReported
	var BS_DISCREPANCIES []fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported

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

func GetGrowthOfBalanceSheetStatementAsReported(BS_STMT_AS_REPORTED []objects.BalanceSheetStatementAsReported) []fundamentals.GrowthBalanceSheetStatementAsReported {
	Growth := []fundamentals.GrowthBalanceSheetStatementAsReported{}
	LastStatement := objects.BalanceSheetStatementAsReported{}

	for i, bs_stmt_as_reported := range BS_STMT_AS_REPORTED {
		NewGrowthObj := fundamentals.GrowthBalanceSheetStatementAsReported{
			Date:   bs_stmt_as_reported.Date,
			Symbol: bs_stmt_as_reported.Symbol,
			Period: bs_stmt_as_reported.Period,
		}

		if i > 0 {
			// Here, reflect is used to iterate over the fields of the struct
			valBS := reflect.ValueOf(bs_stmt_as_reported)
			valLast := reflect.ValueOf(LastStatement)
			valGrowth := reflect.ValueOf(&NewGrowthObj)
			if valGrowth.Kind() == reflect.Ptr && !valGrowth.IsNil() {
				valGrowth = valGrowth.Elem() // Now it's safe to call Elem()
			}

			for j := 0; j < valBS.NumField(); j++ {
				fieldBS := valBS.Field(j)
				fieldLast := valLast.Field(j)
				fieldGrowth := valGrowth.Field(j)

				// Calculate growthValue as usual
				var growthValue float64
				if fieldBS.Kind() == reflect.Float64 && fieldLast.Kind() == reflect.Float64 {
					growthValue = fieldBS.Float() - fieldLast.Float()
				}

				// Set the value based on the kind of fieldGrowth
				if fieldGrowth.Kind() == reflect.Float64 {
					fieldGrowth.SetFloat(growthValue)
				} else if fieldGrowth.Kind() == reflect.Interface {
					// Convert growthValue to interface{} and set it
					fieldGrowth.Set(reflect.ValueOf(growthValue))
				}
				// Handle other cases as necessary
			}

		}

		// Append the new growth object to the slice
		Growth = append(Growth, NewGrowthObj)

		// Update the LastStatement for the next iteration
		LastStatement = bs_stmt_as_reported
	}

	return Growth
}

func IdentifyDiscrepanciesBetweenBalanceSheetStatementAndBalanceSheetStatementAsReported(BS_STMT []objects.BalanceSheetStatement, BS_STMT_AS_REPORTED []objects.BalanceSheetStatementAsReported) []fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported {

	calculateDiscrepancyPercentage := func(value1, value2 float64) float64 {
		if value1 == 0 && value2 == 0 {
			return 0
		}
		absoluteDifference := math.Abs(value1 - value2)
		averageValue := (math.Abs(value1) + math.Abs(value2)) / 2
		return absoluteDifference / averageValue
	}

	discrepancies := make([]fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported, 0)

	for _, bs := range BS_STMT {
		for _, bsar := range BS_STMT_AS_REPORTED {
			if bs.Date == bsar.Date && bs.Symbol == bsar.Symbol && bs.Period == bsar.Period {
				discrepancy := fundamentals.DiscrepancyBalanceSheetStatementAndBalanceSheetStatementAsReported{
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
					if floatValue, ok := bsar.Assets.(float64); ok {
						discrepancy.TotalAssetsMatch = calculateDiscrepancyPercentage(bs.TotalAssets, floatValue)
					} else {
						discrepancy.TotalAssetsMatch = calculateDiscrepancyPercentage(bs.TotalAssets, 0)
					}
				} else {
					discrepancy.TotalAssetsMatch = calculateDiscrepancyPercentage(bs.TotalAssets, 0)
				}

				if bsar.Liabilitiescurrent != nil {
					discrepancy.TotalCurrentLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalCurrentLiabilities, bsar.Liabilitiescurrent.(float64))
				} else {
					discrepancy.TotalCurrentLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalCurrentLiabilities, 0)
				}

				if bsar.Liabilities != nil {
					if floatValue, ok := bsar.Liabilities.(float64); ok {
						discrepancy.TotalLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalLiabilities, floatValue)
					} else {
						discrepancy.TotalLiabilitiesMatch = calculateDiscrepancyPercentage(bs.TotalLiabilities, 0)
					}
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
