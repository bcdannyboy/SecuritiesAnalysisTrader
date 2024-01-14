package financials

import (
	"encoding/json"
	"fmt"

	"github.com/bcdannyboy/SecuritiesAnalysisTrader/utils"
)

func (FA *FinancialsAPI) BalanceSheet(Symbol string, Type string, DateFrom string, DateTo string) (*BalanceSheetResponse, error) {
	/*
	 * Gets the balance sheet for the specified symbol.
	 *
	 * Parameters:
	 * 	Symbol (string): The symbol to get the balance sheet for.
	 * 	Type (string): The type of balance sheet to get.
	 * 	DateFrom (string): The date to get the balance sheet from.
	 * 	DateTo (string): The date to get the balance sheet to.
	 *
	 * Returns:
	 * 	*BalanceSheetResponse: The balance sheet response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/balancesheet?KEY=%s&ticker=%s", FA.APIURL, FA.APIKey, Symbol)

	PossibleTypes := []string{"year", "quarter", "ttm"}

	if Type != "" && utils.StringSliceContains(PossibleTypes, Type) {
		URL = fmt.Sprintf("%s&type=%s", URL, Type)
	} else {
		URL = fmt.Sprintf("%s&type=annual", URL)
	}

	if DateFrom != "" {
		URL = fmt.Sprintf("%s&date_from=%s", URL, DateFrom)
	}

	if DateTo != "" {
		URL = fmt.Sprintf("%s&date_to=%s", URL, DateTo)
	}

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve balance sheet for %s: %s", Symbol, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving balance sheet: %s", Symbol)
	}

	var BalanceSheetResponseObject *BalanceSheetResponse
	err = json.Unmarshal(ResponseObject, &BalanceSheetResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance sheet response: %s", err.Error())
	}

	return BalanceSheetResponseObject, nil
}
