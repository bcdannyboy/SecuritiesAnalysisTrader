package FullFinancialStatement

import (
	"fmt"
	"github.com/spacecodewor/fmpcloud-go"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

func AnalyzeFinancialStatement(APIClient *fmpcloud.APIClient, Symbol string, Period objects.CompanyValuationPeriod) ([]objects.FullFinancialStatementAsReported, []objects.FinancialStatementsGrowth, error) {
	var F_STMT []objects.FullFinancialStatementAsReported
	var F_STMT_GROWTH []objects.FinancialStatementsGrowth

	F_STMT, err := APIClient.CompanyValuation.FullFinancialStatementAsReported(objects.RequestFullFinancialStatementAsReported{
		Symbol: Symbol,
		Period: Period,
	})
	if err != nil {
		return F_STMT, F_STMT_GROWTH, fmt.Errorf("Failed to get full financial statement as reported: %s", err.Error())
	}

	F_STMT_GROWTH, err = APIClient.CompanyValuation.FinancialStatementsGrowth(objects.RequestFinancialStatementsGrowth{
		Symbol: Symbol,
		Period: Period,
	})
	if err != nil {
		return F_STMT, F_STMT_GROWTH, fmt.Errorf("Failed to get financial statement growth: %s", err.Error())
	}

	return F_STMT, F_STMT_GROWTH, nil
}
