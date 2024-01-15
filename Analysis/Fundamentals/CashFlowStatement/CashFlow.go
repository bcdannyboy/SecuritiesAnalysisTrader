package Fundamentals

import (
	"fmt"
	fundamentals "github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis/Fundamentals"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
	"reflect"
)

func AnalyzeCashFlow(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) ([]objects.CashFlowStatement, []objects.CashFlowStatementGrowth, []objects.CashFlowStatementAsReported, []*fundamentals.CashFlowStatementAsReportedGrowth, []*fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported, error) {
	var CF_STMT []objects.CashFlowStatement
	var CF_STMT_GROWTH []objects.CashFlowStatementGrowth
	var CF_STMT_AS_REPORTED []objects.CashFlowStatementAsReported
	var CF_STMT_AS_REPORTED_GROWTH []*fundamentals.CashFlowStatementAsReportedGrowth
	var CF_DISCREPANCIES []*fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported

	CF_STMT, err := APIClient.CompanyValuation.CashFlowStatement(objects.RequestCashFlowStatement{
		Symbol: Symbol,
		Period: Period,
	})
	if err != nil {
		return CF_STMT, CF_STMT_GROWTH, CF_STMT_AS_REPORTED, CF_STMT_AS_REPORTED_GROWTH, CF_DISCREPANCIES, fmt.Errorf("Failed to get cash flow statement: %s", err.Error())
	}

	CF_STMT_GROWTH, err = APIClient.CompanyValuation.CashFlowStatementGrowth(objects.RequestCashFlowStatementGrowth{
		Symbol: Symbol,
		Period: "quarter",
	})
	if err != nil {
		return CF_STMT, CF_STMT_GROWTH, CF_STMT_AS_REPORTED, CF_STMT_AS_REPORTED_GROWTH, CF_DISCREPANCIES, fmt.Errorf("Failed to get cash flow statement growth: %s", err.Error())
	}

	CF_STMT_AS_REPORTED, err = APIClient.CompanyValuation.CashFlowStatementAsReported(objects.RequestCashFlowStatementAsReported{
		Symbol: Symbol,
		Period: "quarter",
	})
	if err != nil {
		return CF_STMT, CF_STMT_GROWTH, CF_STMT_AS_REPORTED, CF_STMT_AS_REPORTED_GROWTH, CF_DISCREPANCIES, fmt.Errorf("Failed to get cash flow statement as reported: %s", err.Error())
	}

	CF_DISCREPANCIES = IdentifyDiscrepanciesBetweenCashFlowStatementAndCashFlowStatementAsReported(CF_STMT, CF_STMT_AS_REPORTED)
	CF_STMT_AS_REPORTED_GROWTH = GetGrowthOfCashFlowStatementAsReported(CF_STMT_AS_REPORTED)

	return CF_STMT, CF_STMT_GROWTH, CF_STMT_AS_REPORTED, CF_STMT_AS_REPORTED_GROWTH, CF_DISCREPANCIES, nil
}

func IdentifyDiscrepanciesBetweenCashFlowStatementAndCashFlowStatementAsReported(CF_STMT []objects.CashFlowStatement, CF_STMT_AS_REPORTED []objects.CashFlowStatementAsReported) []*fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported {

	calculateDiscrepancyPercentage := func(value1, value2 float64) float64 {
		if value1 == 0 && value2 == 0 {
			return 0
		}
		absoluteDifference := math.Abs(value1 - value2)
		averageValue := (math.Abs(value1) + math.Abs(value2)) / 2
		return absoluteDifference / averageValue
	}

	discrepancies := make([]*fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported, 0)

	for _, cf := range CF_STMT {
		for _, cfar := range CF_STMT_AS_REPORTED {
			if cf.Date == cfar.Date && cf.Symbol == cfar.Symbol && cf.Period == cfar.Period {
				discrepancy := &fundamentals.DiscrepancyCashFlowStatementAndCashFlowStatementAsReported{
					Date:         cf.Date,
					Symbol:       cf.Symbol,
					Period:       cf.Period,
					FillingDate:  cf.FillingDate,
					AcceptedDate: cf.AcceptedDate,
				}

				discrepancy.NetIncomeDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.NetIncome, cfar.Netincomeloss)
				discrepancy.DepreciationAndAmortizationDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.DepreciationAndAmortization, cfar.Depreciationdepletionandamortization)
				discrepancy.DeferredIncomeTaxDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.DeferredIncomeTax, cfar.Deferredincometaxexpensebenefit)
				discrepancy.StockBasedCompensationDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.StockBasedCompensation, cfar.Sharebasedcompensation)
				discrepancy.ChangeInWorkingCapitalDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.ChangeInWorkingCapital, cfar.Increasedecreaseinaccountsreceivable+cfar.Increasedecreaseininventories-cfar.Increasedecreaseinaccountspayable)
				discrepancy.AccountsReceivablesDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.AccountsReceivables, cfar.Increasedecreaseinaccountsreceivable)
				discrepancy.InventoryDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.Inventory, cfar.Increasedecreaseininventories)
				discrepancy.AccountsPayablesDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.AccountsPayables, cfar.Increasedecreaseinaccountspayable)
				discrepancy.OtherNonCashItemsDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.OtherNonCashItems, cfar.Othernoncashincomeexpense)
				discrepancy.NetCashProvidedByOperatingActivitiesDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.NetCashProvidedByOperatingActivities, cfar.Netcashprovidedbyusedinoperatingactivities)
				discrepancy.InvestmentsInPropertyPlantAndEquipmentDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.InvestmentsInPropertyPlantAndEquipment, cfar.Paymentstoacquirepropertyplantandequipment)
				discrepancy.AcquisitionsNetDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.AcquisitionsNet, cfar.Paymentstoacquirebusinessesnetofcashacquired)
				discrepancy.PurchasesOfInvestmentsDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.PurchasesOfInvestments, cfar.Paymentstoacquireotherinvestments)
				discrepancy.SalesMaturitiesOfInvestmentsDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.SalesMaturitiesOfInvestments, cfar.Proceedsfromsaleandmaturityofotherinvestments)
				discrepancy.OtherInvestingActivitesDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.OtherInvestingActivites, cfar.Paymentsforproceedsfromotherinvestingactivities)
				discrepancy.NetCashUsedForInvestingActivitesDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.NetCashUsedForInvestingActivites, 0) // Assuming no direct match
				discrepancy.DebtRepaymentDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.DebtRepayment, cfar.Repaymentsoflongtermdebt)
				discrepancy.CommonStockIssuedDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.CommonStockIssued, cfar.Proceedsfromissuanceofcommonstock)
				discrepancy.CommonStockRepurchasedDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.CommonStockRepurchased, cfar.Paymentsforrepurchaseofcommonstock)
				discrepancy.DividendsPaidDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.DividendsPaid, cfar.Paymentsofdividends)
				discrepancy.OtherFinancingActivitesDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.OtherFinancingActivites, cfar.Proceedsfrompaymentsforotherfinancingactivities)
				discrepancy.NetCashUsedProvidedByFinancingActivitiesDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.NetCashUsedProvidedByFinancingActivities, 0) // Assuming no direct match
				discrepancy.EffectOfForexChangesOnCashDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.EffectOfForexChangesOnCash, 0)                             // Assuming no direct match
				discrepancy.NetChangeInCashDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.NetChangeInCash, cfar.Cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect)
				discrepancy.CashAtEndOfPeriodDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.CashAtEndOfPeriod, 0)             // Assuming no direct match
				discrepancy.CashAtBeginningOfPeriodDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.CashAtBeginningOfPeriod, 0) // Assuming no direct match
				discrepancy.OperatingCashFlowDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.OperatingCashFlow, cfar.Netcashprovidedbyusedinoperatingactivities)
				discrepancy.CapitalExpenditureDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.CapitalExpenditure, cfar.Paymentstoacquirepropertyplantandequipment)
				discrepancy.FreeCashFlowDiscrepancyPercentage = calculateDiscrepancyPercentage(cf.FreeCashFlow, 0) // Assuming no direct match

				discrepancies = append(discrepancies, discrepancy)
			}
		}
	}

	return discrepancies
}

func GetGrowthOfCashFlowStatementAsReported(CFS_STMT_AS_REPORTED []objects.CashFlowStatementAsReported) []*fundamentals.CashFlowStatementAsReportedGrowth {
	Growth := []*fundamentals.CashFlowStatementAsReportedGrowth{}
	LastStatement := objects.CashFlowStatementAsReported{}

	for i, cfs_stmt_as_reported := range CFS_STMT_AS_REPORTED {
		NewGrowthObj := &fundamentals.CashFlowStatementAsReportedGrowth{
			Date:   cfs_stmt_as_reported.Date,
			Symbol: cfs_stmt_as_reported.Symbol,
			Period: cfs_stmt_as_reported.Period,
		}

		if i > 0 {
			// Here, reflect is used to iterate over the fields of the struct
			valCFS := reflect.ValueOf(cfs_stmt_as_reported)
			valLast := reflect.ValueOf(LastStatement)
			valGrowth := reflect.ValueOf(NewGrowthObj).Elem()

			for j := 0; j < valCFS.NumField(); j++ {
				fieldCFS := valCFS.Field(j)
				fieldLast := valLast.Field(j)
				fieldGrowth := valGrowth.Field(j)

				// Check if the field is of type float64
				if fieldCFS.Kind() == reflect.Float64 {
					growthValue := fieldCFS.Float() - fieldLast.Float()
					fieldGrowth.SetFloat(growthValue)
				}

				// For interface{} types, use type assertion
				if fieldCFS.Kind() == reflect.Interface && !fieldCFS.IsNil() {
					curVal, okCur := fieldCFS.Interface().(float64)
					lastVal, okLast := fieldLast.Interface().(float64)
					if okCur && okLast {
						fieldGrowth.SetFloat(curVal - lastVal)
					}
				}
			}
		}

		// Append the new growth object to the slice
		Growth = append(Growth, NewGrowthObj)

		// Update the LastStatement for the next iteration
		LastStatement = cfs_stmt_as_reported
	}

	return Growth
}
