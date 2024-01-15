package Fundamentals

import (
	"fmt"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
	"math"
)

func AnalyzeCashFlow(APIClient *fmpcloud.APIClient, Symbol string) ([]objects.CashFlowStatement, []objects.CashFlowStatementGrowth, []objects.CashFlowStatementAsReported, []*CashFlowStatementAsReportedGrowth, []*DiscrepancyCashFlowStatementAndCashFlowStatementAsReported, error) {
	var CF_STMT []objects.CashFlowStatement
	var CF_STMT_GROWTH []objects.CashFlowStatementGrowth
	var CF_STMT_AS_REPORTED []objects.CashFlowStatementAsReported
	var CF_STMT_AS_REPORTED_GROWTH []*CashFlowStatementAsReportedGrowth
	var CF_DISCREPANCIES []*DiscrepancyCashFlowStatementAndCashFlowStatementAsReported

	CF_STMT, err := APIClient.CompanyValuation.CashFlowStatement(objects.RequestCashFlowStatement{
		Symbol: Symbol,
		Period: "quarter",
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

	return CF_STMT, CF_STMT_GROWTH, CF_STMT_AS_REPORTED, CF_STMT_AS_REPORTED_GROWTH, CF_DISCREPANCIES, nil
}

func IdentifyDiscrepanciesBetweenCashFlowStatementAndCashFlowStatementAsReported(CF_STMT []objects.CashFlowStatement, CF_STMT_AS_REPORTED []objects.CashFlowStatementAsReported) []*DiscrepancyCashFlowStatementAndCashFlowStatementAsReported {

	calculateDiscrepancyPercentage := func(value1, value2 float64) float64 {
		if value1 == 0 && value2 == 0 {
			return 0
		}
		absoluteDifference := math.Abs(value1 - value2)
		averageValue := (math.Abs(value1) + math.Abs(value2)) / 2
		return absoluteDifference / averageValue
	}

	discrepancies := make([]*DiscrepancyCashFlowStatementAndCashFlowStatementAsReported, 0)

	for _, cf := range CF_STMT {
		for _, cfar := range CF_STMT_AS_REPORTED {
			if cf.Date == cfar.Date && cf.Symbol == cfar.Symbol && cf.Period == cfar.Period {
				discrepancy := &DiscrepancyCashFlowStatementAndCashFlowStatementAsReported{
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
