package Fundamentals

import (
	"fmt"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func AnalyzeCashFlow(APIClient *fmpcloud.APIClient, Symbol string) error {
	var CF_STMT []objects.CashFlowStatement
	var CF_STMT_GROWTH []objects.CashFlowStatementGrowth
	var CF_STMT_AS_REPORTED []objects.CashFlowStatementAsReported
	var CF_STMT_AS_REPORTED_GROWTH []*CashFlowStatementAsReportedGrowth
	var CF_DISCREPANCIES []interface{}

	CF_STMT, err := APIClient.CompanyValuation.CashFlowStatement(objects.RequestCashFlowStatement{
		Symbol: Symbol,
		Period: "quarter",
	})
	if err != nil {
		return fmt.Errorf("Failed to get cash flow statement: %s", err.Error())
	}

	return nil
}

// def AnalyzeCashFlowStatementAsReportedGrowth(CF_STMT_AS_REPORTED []objects.CashFlowStatementAsReported)
