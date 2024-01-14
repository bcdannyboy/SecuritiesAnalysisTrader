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
		return nil, fmt.Errorf("failed to unmarshal balance sheet response for %s: %s", Symbol, err.Error())
	}

	return BalanceSheetResponseObject, nil
}

func (FA *FinancialsAPI) IncomeStatements(Symbol string, Type string, DateFrom string, DateTo string) (*IncomeStatementResponse, error) {
	/*
	 * Gets the income statement for the specified symbol.
	 *
	 * Parameters:
	 * 	Symbol (string): The symbol to get the income statement for.
	 * 	Type (string): The type of income statement to get.
	 * 	DateFrom (string): The date to get the income statement from.
	 * 	DateTo (string): The date to get the income statement to.
	 *
	 * Returns:
	 * 	*IncomeStatementResponse: The income statement response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/incomestatements?KEY=%s&ticker=%s", FA.APIURL, FA.APIKey, Symbol)

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
		return nil, fmt.Errorf("failed to retrieve income statement for %s: %s", Symbol, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving income statement: %s", Symbol)
	}

	var IncomeStatementResponseObject *IncomeStatementResponse
	err = json.Unmarshal(ResponseObject, &IncomeStatementResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal income statement response for %s: %s", Symbol, err.Error())
	}

	return IncomeStatementResponseObject, nil
}

func (FA *FinancialsAPI) CashFlows(Symbol string, Type string, DateFrom string, DateTo string) (*CashFlowResponse, error) {
	/*
	 * Gets the cash flows for the specified symbol.
	 *
	 * Parameters:
	 * 	Symbol (string): The symbol to get the cash flows for.
	 * 	Type (string): The type of cash flows to get.
	 * 	DateFrom (string): The date to get the cash flows from.
	 * 	DateTo (string): The date to get the cash flows to.
	 *
	 * Returns:
	 * 	*CashFlowResponse: The cash flows response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/cashflows?KEY=%s&ticker=%s", FA.APIURL, FA.APIKey, Symbol)

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
		return nil, fmt.Errorf("failed to retrieve cash flows for %s: %s", Symbol, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving cash flows: %s", Symbol)
	}

	var CashFlowResponseObject *CashFlowResponse
	err = json.Unmarshal(ResponseObject, &CashFlowResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal cash flows response for %s: %s", Symbol, err.Error())
	}

	return CashFlowResponseObject, nil
}

func (FA *FinancialsAPI) Dividends(Symbol string, Type string, DateTo string) (*DividendsResponse, error) {
	/*
	 * Gets the dividends for the specified symbol.
	 *
	 * Parameters:
	 * 	Symbol (string): The symbol to get the dividends for.
	 * 	Type (string): The type of dividends to get.
	 * 	DateTo (string): The date to get the dividends to.
	 *
	 * Returns:
	 * 	*DividendsResponse: The dividends response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/dividens?KEY=%s&ticker=%s", FA.APIURL, FA.APIKey, Symbol)

	PossibleTypes := []string{"year", "quarter", "ttm"}

	if Type != "" && utils.StringSliceContains(PossibleTypes, Type) {
		URL = fmt.Sprintf("%s&type=%s", URL, Type)
	} else {
		URL = fmt.Sprintf("%s&type=annual", URL)
	}

	if DateTo != "" {
		URL = fmt.Sprintf("%s&date_to=%s", URL, DateTo)
	}

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve dividends for %s: %s", Symbol, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving dividends: %s", Symbol)
	}

	var DividendsResponseObject *DividendsResponse
	err = json.Unmarshal(ResponseObject, &DividendsResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal dividends response for %s: %s", Symbol, err.Error())
	}

	return DividendsResponseObject, nil
}

func (FA *FinancialsAPI) StockSplits(Symbol string, Type string, DateTo string) (*StockSplitsResponse, error) {
	/*
	 * Gets the economic calendar for the specified symbol.
	 *
	 * Parameters:
	 * 	Symbol (string): The symbol to get the stock splits for.
	 * 	Type (string): The type of stock splits to get.
	 * 	DateTo (string): The date to get the stock splits to.
	 *
	 * Returns:
	 * 	*StockSplitsResponse: The stock split response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/stocksplits?KEY=%s&ticker=%s", FA.APIURL, FA.APIKey, Symbol)

	PossibleTypes := []string{"year", "quarter", "ttm"}

	if Type != "" && utils.StringSliceContains(PossibleTypes, Type) {
		URL = fmt.Sprintf("%s&type=%s", URL, Type)
	} else {
		URL = fmt.Sprintf("%s&type=annual", URL)
	}

	if DateTo != "" {
		URL = fmt.Sprintf("%s&date_to=%s", URL, DateTo)
	}

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve stock splits for %s: %s", Symbol, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving stock splits: %s", Symbol)
	}

	var StockSplitsResponseObject *StockSplitsResponse
	err = json.Unmarshal(ResponseObject, &StockSplitsResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal stock splits response for %s: %s", Symbol, err.Error())
	}

	return StockSplitsResponseObject, nil
}

func (FA *FinancialsAPI) EconomicCalendar(Country string, Date string, ISO_Country_String string) (*CalendarResponse, error) {
	/*
	 * Gets the economic calendar for the specified symbol.
	 *
	 * Parameters:
	 * 	Country (string): The country to get the economic calendar for.
	 * 	Date (string): The date to get the economic calendar for.
	 * 	ISO_Country_String (string): The ISO country string to get the economic calendar for.
	 *
	 * Returns:
	 * 	*CalendarResponse: The economic calendar response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/macrocalendar?KEY=%s", FA.APIURL, FA.APIKey)

	if Country != "" {
		URL = fmt.Sprintf("%s&country=%s", URL, Country)
	} else {
		URL = fmt.Sprintf("%s&country=United_States", URL)
	}

	if Date != "" {
		URL = fmt.Sprintf("%s&date=%s", URL, Date)
	}

	if ISO_Country_String != "" {
		URL = fmt.Sprintf("%s&iso_country_code=%s", URL, ISO_Country_String)
	} else {
		URL = fmt.Sprintf("%s&iso_country_code=US", URL)
	}

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve economic calendar for %s: %s", Country, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving economic calendar: %s", Country)
	}

	var CalendarResponseObject *CalendarResponse
	err = json.Unmarshal(ResponseObject, &CalendarResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal economic calendar response for %s: %s", Country, err.Error())
	}

	return CalendarResponseObject, nil
}

func (FA *FinancialsAPI) SECFilings(Ticker string) (*SECFilingsResponse, error) {
	/*
	 * Gets the SEC Filings for the specified symbol.
	 *
	 * Parameters:
	 * 	Ticker (string): The ticker to get the sec filings for.
	 *
	 * Returns:
	 * 	*SECFilingsResponse: The sec filings response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/secfilings?KEY=%s&stock_ticker_symbol=%s", FA.APIURL, FA.APIKey, Ticker)

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve sec filings for %s: %s", Ticker, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving sec filings: %s", Ticker)
	}

	var SECFilingsResponseObject *SECFilingsResponse
	err = json.Unmarshal(ResponseObject, &SECFilingsResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal sec filings response for %s: %s", Ticker, err.Error())
	}

	return SECFilingsResponseObject, nil
}

func (FA *FinancialsAPI) ETFHoldings(Series_ID string, Series_LEI string, Date_From string, Date_To string) (*ETFHoldingsResponse, error) {
	/*
	 * Gets the ETF Holdings for the specified symbol.
	 *
	 * Parameters:
	 * 	Series_ID (string): The series ID to get the ETF Holdings for.
	 * 	Series_LEI (string): The series LEI to get the ETF Holdings for.
	 * 	Date_From (string): The date to get the ETF Holdings from.
	 * 	Date_To (string): The date to get the ETF Holdings to.
	 *
	 * Returns:
	 * 	*ETFHoldingsResponse: The ETF Holdings response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/etfholdings?KEY=%s", FA.APIURL, FA.APIKey)

	if Series_ID != "" {
		URL = fmt.Sprintf("%s&series_id=%s", URL, Series_ID)
	}

	if Series_LEI != "" {
		URL = fmt.Sprintf("%s&series_lei=%s", URL, Series_LEI)
	}

	if Date_From != "" {
		URL = fmt.Sprintf("%s&date_from=%s", URL, Date_From)
	}

	if Date_To != "" {
		URL = fmt.Sprintf("%s&date_to=%s", URL, Date_To)
	}

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ETF Holdings for %s: %s", Series_ID, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving ETF Holdings: %s", Series_ID)
	}

	var ETFHoldingsResponseObject *ETFHoldingsResponse
	err = json.Unmarshal(ResponseObject, &ETFHoldingsResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal ETF Holdings response for %s: %s", Series_ID, err.Error())
	}

	return ETFHoldingsResponseObject, nil
}

func (FA *FinancialsAPI) ETFInfo(Series_ID string, Series_LEI string, CIK string) (*ETFInfoResponse, error) {
	/*
	 * Gets the ETF Information for the specified symbol.
	 *
	 * Parameters:
	 * 	Series_ID (string): The series ID to get the ETF Information for.
	 * 	Series_LEI (string): The series LEI to get the ETF Information for.
	 * 	CIK (string): The CIK to get the ETF Information for.
	 *
	 * Returns:
	 * 	*ETFInfoResponse: The ETF Information response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/etfinfo?KEY=%s", FA.APIURL, FA.APIKey)

	if CIK != "" {
		URL = fmt.Sprintf("%s&cik=%s", URL, CIK)
	}

	if Series_ID != "" {
		URL = fmt.Sprintf("%s&series_id=%s", URL, Series_ID)
	}

	if Series_LEI != "" {
		URL = fmt.Sprintf("%s&series_lei=%s", URL, Series_LEI)
	}

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ETF Information for %s: %s", Series_ID, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving ETF Information: %s", Series_ID)
	}

	var ETFInfoResponseObject *ETFInfoResponse
	err = json.Unmarshal(ResponseObject, &ETFInfoResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal ETF Information response for %s: %s", Series_ID, err.Error())
	}

	return ETFInfoResponseObject, nil
}

func (FA *FinancialsAPI) BondsYield(Country string, Region string, Type string) (*BondsResponse, error) {
	/*
	 * Gets the Bonds Yield for the specified symbol.
	 *
	 * Parameters:
	 * 	Country (string): The country to get the Bonds Yield for.
	 * 	Region (string): The region to get the Bonds Yield for.
	 * 	Type (string): The type to get the Bonds Yield for.
	 *
	 * Returns:
	 * 	*BondsResponse: The Bonds Yield response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/etfinfo?KEY=%s", FA.APIURL, FA.APIKey)

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Bonds Yield for %s-%s-%s: %s", Country, Region, Type, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving Bonds Yield: %s-%s-%s", Country, Region, Type)
	}

	var BondsResponseObject *BondsResponse
	err = json.Unmarshal(ResponseObject, &BondsResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Bonds Yield response for %s-%s-%s: %s", Country, Region, Type, err.Error())
	}

	return BondsResponseObject, nil
}

func (FA *FinancialsAPI) FinancialRatios(Ticker string) (*FinancialRatiosResponse, error) {
	/*
	 * Gets the Financial Ratios for the specified symbol.
	 *
	 * Parameters:
	 * 	Ticker (string): The ticker to get the Financial Ratios for.
	 *
	 * Returns:
	 * 	*FinancialRatiosResponse: The Financial Ratios response object.
	 * 	error: The error object.
	 */

	URL := fmt.Sprintf("%s/financialratios?KEY=%s&stock_ticker_symbol=%s", FA.APIURL, FA.APIKey, Ticker)

	StatusCode, ResponseObject, err := utils.SendHTTPGETRequest(URL, map[string]string{}, map[string]string{}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Financial Ratios for %s: %s", Ticker, err.Error())
	}

	if StatusCode != 200 {
		return nil, fmt.Errorf("invalid request while retrieving financial ratios: %s", Ticker)
	}

	var FinancialRatiosResponseObject *FinancialRatiosResponse
	err = json.Unmarshal(ResponseObject, &FinancialRatiosResponseObject)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal financial ratios response for %s: %s", Ticker, err.Error())
	}

	return FinancialRatiosResponseObject, nil
}
