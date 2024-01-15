package Fundamentals

import (
	"fmt"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
	"reflect"
)

func AnalyzeIncomeStatement(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) ([]objects.IncomeStatement, []objects.IncomeStatementGrowth, []objects.IncomeStatementAsReported, []*GrowthIncomeStatementAsReported, []*DiscrepancyIncomeStatementAndIncomeStatementAsReported, error) {
	I_STMT := []objects.IncomeStatement{}
	I_STMT_GROWTH := []objects.IncomeStatementGrowth{}
	I_STMT_AS_REPORTED := []objects.IncomeStatementAsReported{}
	I_STMT_AS_REPORTED_GROWTH := []*GrowthIncomeStatementAsReported{}
	I_DISCREPANCIES := []*DiscrepancyIncomeStatementAndIncomeStatementAsReported{}

	I_STMT, err := APIClient.CompanyValuation.IncomeStatement(
		objects.RequestIncomeStatement{
			Symbol: Symbol,
			Period: Period,
		})
	if err != nil {
		return I_STMT, I_STMT_GROWTH, I_STMT_AS_REPORTED, I_STMT_AS_REPORTED_GROWTH, I_DISCREPANCIES, fmt.Errorf("failed to get income statement: %s", err.Error())
	}

	I_STMT_GROWTH, err = APIClient.CompanyValuation.IncomeStatementGrowth(
		objects.RequestIncomeStatementGrowth{
			Symbol: Symbol,
			Period: Period,
		})
	if err != nil {
		return I_STMT, I_STMT_GROWTH, I_STMT_AS_REPORTED, I_STMT_AS_REPORTED_GROWTH, I_DISCREPANCIES, fmt.Errorf("failed to get income statement growth: %s", err.Error())
	}

	I_STMT_AS_REPORTED, err = APIClient.CompanyValuation.IncomeStatementAsReported(
		objects.RequestIncomeStatementAsReported{
			Symbol: Symbol,
			Period: Period,
		})
	if err != nil {
		return I_STMT, I_STMT_GROWTH, I_STMT_AS_REPORTED, I_STMT_AS_REPORTED_GROWTH, I_DISCREPANCIES, fmt.Errorf("failed to get income statement as reported: %s", err.Error())
	}

	I_DISCREPANCIES = IdentifyDiscrepanciesBetweenIncomeStatementAndIncomeStatementAsReported(I_STMT, I_STMT_AS_REPORTED)
	I_STMT_AS_REPORTED_GROWTH = GetGrowthOfIncomeStatementAsReported(I_STMT_AS_REPORTED)

	return I_STMT, I_STMT_GROWTH, I_STMT_AS_REPORTED, I_STMT_AS_REPORTED_GROWTH, I_DISCREPANCIES, nil
}

func IdentifyDiscrepanciesBetweenIncomeStatementAndIncomeStatementAsReported(IS_STMT []objects.IncomeStatement, IS_STMT_AS_REPORTED []objects.IncomeStatementAsReported) []*DiscrepancyIncomeStatementAndIncomeStatementAsReported {
	calculateDiscrepancyPercentage := func(value1, value2 float64) float64 {
		if value1 == 0 && value2 == 0 {
			return 0
		}
		absoluteDifference := math.Abs(value1 - value2)
		averageValue := (math.Abs(value1) + math.Abs(value2)) / 2
		return absoluteDifference / averageValue
	}

	discrepancies := make([]*DiscrepancyIncomeStatementAndIncomeStatementAsReported, 0)

	for _, is := range IS_STMT {
		for _, isar := range IS_STMT_AS_REPORTED {
			if is.Date == isar.Date && is.Symbol == isar.Symbol && is.Period == isar.Period {
				discrepancy := &DiscrepancyIncomeStatementAndIncomeStatementAsReported{
					Date:   is.Date,
					Symbol: is.Symbol,
					Period: is.Period,
				}

				discrepancy.NetIncomeDiscrepancy = calculateDiscrepancyPercentage(is.NetIncome, isar.Netincomeloss)

				// Assuming GrossProfit and Grossprofit are comparable
				if gp, ok := isar.Grossprofit.(float64); ok {
					discrepancy.GrossProfitDiscrepancy = calculateDiscrepancyPercentage(is.GrossProfit, gp)
				}

				discrepancy.ResearchAndDevelopmentExpensesDiscrepancy = calculateDiscrepancyPercentage(is.ResearchAndDevelopmentExpenses, isar.Researchanddevelopmentexpense)

				// Assuming OperatingIncome and Operatingincomeloss are comparable
				discrepancy.OperatingIncomeDiscrepancy = calculateDiscrepancyPercentage(is.OperatingIncome, isar.Operatingincomeloss)

				discrepancy.EpsDiscrepancy = calculateDiscrepancyPercentage(is.Eps, isar.Earningspersharebasic)
				discrepancy.EpsDilutedDiscrepancy = calculateDiscrepancyPercentage(is.Epsdiluted, isar.Earningspersharediluted)

				discrepancy.WeightedAverageShsOutDiscrepancy = calculateDiscrepancyPercentage(is.WeightedAverageShsOut, isar.Weightedaveragenumberofsharesoutstandingbasic)
				discrepancy.WeightedAverageShsOutDilDiscrepancy = calculateDiscrepancyPercentage(is.WeightedAverageShsOutDil, isar.Weightedaveragenumberofdilutedsharesoutstanding)

				discrepancy.IncomeTaxExpenseDiscrepancy = calculateDiscrepancyPercentage(is.IncomeTaxExpense, isar.Incometaxexpensebenefit)

				discrepancies = append(discrepancies, discrepancy)
			}
		}
	}

	return discrepancies
}

func GetGrowthOfIncomeStatementAsReported(IS_STMT_AS_REPORTED []objects.IncomeStatementAsReported) []*GrowthIncomeStatementAsReported {
	Growth := []*GrowthIncomeStatementAsReported{}
	LastStatement := objects.IncomeStatementAsReported{}

	for i, is_stmt_as_reported := range IS_STMT_AS_REPORTED {
		NewGrowthObj := &GrowthIncomeStatementAsReported{
			Date:   is_stmt_as_reported.Date,
			Symbol: is_stmt_as_reported.Symbol,
			Period: is_stmt_as_reported.Period,
		}

		if i > 0 {
			valIS := reflect.ValueOf(is_stmt_as_reported)
			valLast := reflect.ValueOf(LastStatement)
			valGrowth := reflect.ValueOf(NewGrowthObj).Elem()

			for j := 0; j < valIS.NumField(); j++ {
				fieldIS := valIS.Field(j)
				fieldLast := valLast.Field(j)
				fieldGrowth := valGrowth.Field(j)

				if fieldIS.Kind() == reflect.Float64 {
					growthValue := fieldIS.Float() - fieldLast.Float()
					fieldGrowth.SetFloat(growthValue)
				}

				if fieldIS.Kind() == reflect.Interface && !fieldIS.IsNil() {
					curVal, okCur := fieldIS.Interface().(float64)
					lastVal, okLast := fieldLast.Interface().(float64)
					if okCur && okLast {
						fieldGrowth.SetFloat(curVal - lastVal)
					}
				}
			}
		}

		Growth = append(Growth, NewGrowthObj)
		LastStatement = is_stmt_as_reported
	}

	return Growth
}
